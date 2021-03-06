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

package controllers

import (
	"context"

	"github.com/openshift/library-go/pkg/controller/factory"
)

//+kubebuilder:rbac:groups=operator.openshift.io.my.domain,resources=certmanagers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=operator.openshift.io.my.domain,resources=certmanagers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=operator.openshift.io.my.domain,resources=certmanagers/finalizers,verbs=update

type CertManagerController struct {
	// TODO: fill in the controller you would like to configure
	// for the resource.
}

// NewCertManagerController() creates a new instance of your controller.
func NewCertManagerController() factory.Controller {
	// TODO: use this to create a new instance of your controller.
	// The controller can further be configured here to be either event or time
	// based.

	return nil
}

// sync contains the logic of the reconciler.
func (c *CertManagerController) sync(ctx context.Context, syncContext factory.SyncContext) error {

	// TODO: implement your reconciler logic here.

	return nil
}
