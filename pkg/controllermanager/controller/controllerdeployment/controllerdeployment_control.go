// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllerdeployment

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"github.com/gardener/gardener/pkg/controllerutils"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
)

func (c *Controller) controllerDeploymentAdd(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		c.log.Error(err, "Couldn't get key for object", "object", obj)
		return
	}
	c.controllerDeploymentQueue.Add(key)
}

func (c *Controller) controllerDeploymentUpdate(_, newObj interface{}) {
	c.controllerDeploymentAdd(newObj)
}

// NewReconciler creates a new instance of a reconciler which reconciles ControllerDeployments.
func NewReconciler(gardenClient client.Client) reconcile.Reconciler {
	return &controllerDeploymentReconciler{
		gardenClient: gardenClient,
	}
}

type controllerDeploymentReconciler struct {
	gardenClient client.Client
}

func (c *controllerDeploymentReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	log := logf.FromContext(ctx)

	controllerDeployment := &gardencorev1beta1.ControllerDeployment{}
	if err := c.gardenClient.Get(ctx, kutil.Key(req.Name), controllerDeployment); err != nil {
		if apierrors.IsNotFound(err) {
			log.V(1).Info("Object is gone, stop reconciling")
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, fmt.Errorf("error retrieving object from store: %w", err)
	}

	if controllerDeployment.DeletionTimestamp != nil {
		if !controllerutil.ContainsFinalizer(controllerDeployment, FinalizerName) {
			return reconcile.Result{}, nil
		}

		controllerRegistrationList := &gardencorev1beta1.ControllerRegistrationList{}
		if err := c.gardenClient.List(ctx, controllerRegistrationList); err != nil {
			return reconcile.Result{}, err
		}

		for _, controllerRegistration := range controllerRegistrationList.Items {
			deployment := controllerRegistration.Spec.Deployment
			if deployment == nil {
				continue
			}
			for _, deploymentRef := range deployment.DeploymentRefs {
				if deploymentRef.Name == controllerDeployment.Name {
					return reconcile.Result{}, fmt.Errorf("cannot remove finalizer of ControllerDeployment %q because still found one ControllerRegistration", controllerRegistration.Name)
				}
			}
		}

		if controllerutil.ContainsFinalizer(controllerDeployment, FinalizerName) {
			log.Info("Removing finalizer")
			if err := controllerutils.RemoveFinalizers(ctx, c.gardenClient, controllerDeployment, FinalizerName); err != nil {
				return reconcile.Result{}, fmt.Errorf("failed to remove finalizer: %w", err)
			}
		}

		return reconcile.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(controllerDeployment, FinalizerName) {
		log.Info("Adding finalizer")
		if err := controllerutils.AddFinalizers(ctx, c.gardenClient, controllerDeployment, FinalizerName); err != nil {
			return reconcile.Result{}, fmt.Errorf("failed to add finalizer: %w", err)
		}
	}

	return reconcile.Result{}, nil
}
