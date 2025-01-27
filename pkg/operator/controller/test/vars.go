/*
Copyright 2021.

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

package test

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	v1alpha1 "github.com/openshift/node-observability-operator/api/v1alpha1"
)

const (
	Name              = "test"
	OperandNamespace  = "node-observability-operator"
	TestNamespace     = ""
	OperandName       = "node-observability-operator"
	OperatorNamespace = "node-observability-operator"
	SecretName        = "node-observability-secret"
)

var (
	TrueVar = true
	Scheme  = runtime.NewScheme()
)

func init() {
	if err := clientgoscheme.AddToScheme(Scheme); err != nil {
		panic(err)
	}
	if err := v1alpha1.AddToScheme(Scheme); err != nil {
		panic(err)
	}

}
