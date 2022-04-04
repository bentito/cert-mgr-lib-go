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
package main

import (
	"context"
	"fmt"
	"github.com/openshift/cert-manager-operator/pkg/cmd/operator"
	"github.com/spf13/cobra"
	"os"

	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	ctx := context.TODO()

	// Create a *rest.Config for talking to a Kubernetes apiserver.
	// If you would like to specify the kubeconfig manually, this can be removed.
	_ = ctrl.GetConfigOrDie()

	// TODO: Instantiate the required resources to create new instances of the
	// controller. In order to run the controller successfully, make sure to start
	// the informers first.

	// Start the informers to make sure their caches are in sync and are updated periodically.
	for _, informer := range []interface {
		Start(stopCh <-chan struct{})
	}{
		// TODO: If there are any informers for your controller, make sure to
		// add them here to start the informer.
	} {
		informer.Start(ctx.Done())
	}

	// Start and run the controller
	for _, controllerint := range []interface {
		Run(ctx context.Context, workers int)
	}{
		// TODO: Add the name of controllers which have been instantiated previosuly for the
		// operator.
	} {
		go controllerint.Run(ctx, 1)
	}

	<-ctx.Done()
	command := NewOperatorCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return

}

func NewOperatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cert-manager-operator",
		Short: "OpenShift cluster cert-manager operator",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	cmd.AddCommand(operator.NewOperator())
	return cmd
}
