// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package builder

import (
	"fmt"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gardener/gardener/pkg/client/kubernetes/clientmap"
	"github.com/gardener/gardener/pkg/client/kubernetes/clientmap/internal"
)

// PlantClientMapBuilder can build a ClientMap which can be used to construct a ClientMap for requesting and storing
// clients for Plant clusters.
type PlantClientMapBuilder struct {
	gardenReader client.Reader
}

// NewPlantClientMapBuilder constructs a new PlantClientMapBuilder.
func NewPlantClientMapBuilder() *PlantClientMapBuilder {
	return &PlantClientMapBuilder{}
}

// WithGardenReader sets the reader that should be used for retrieving Plants' kubeconfig secrets.
func (b *PlantClientMapBuilder) WithGardenReader(reader client.Reader) *PlantClientMapBuilder {
	b.gardenReader = reader
	return b
}

// Build builds the PlantClientMap using the provided attributes.
func (b *PlantClientMapBuilder) Build(log logr.Logger) (clientmap.ClientMap, error) {
	if b.gardenReader == nil {
		return nil, fmt.Errorf("garden reader is required but not set")
	}

	return internal.NewPlantClientMap(log, &internal.PlantClientSetFactory{
		GardenReader: b.gardenReader,
	}), nil
}
