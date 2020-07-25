/*
Copyright 2020 kubeflow.org.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

// SplitterSpec defines a simple weighted traffic split interface
type SplitterSpec struct {
	// Weights defines the weights of the routes
	Weights []*WeightsSpec `json:"weights,omitempty"`
}

// WeightsSpec defines a simple weighted traffic split interface
type WeightsSpec struct {
	// The name for the route
	Name string `json:"name"`
	// The weight to send traffic to this route.
	Weight int `json:"weight"`
}