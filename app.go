//
// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//   ███╗   ██╗ █████╗ ███╗   ███╗██╗
//   ████╗  ██║██╔══██╗████╗ ████║██║
//   ██╔██╗ ██║███████║██╔████╔██║██║
//   ██║╚██╗██║██╔══██║██║╚██╔╝██║██║
//   ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗
//   ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
//

package app

import (
	"context"
	"fmt"

	naml2 "github.com/kris-nova/naml"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Version string

// MySampleApp is the purest and most true form of an application. In this example
// we keep SOME of the field unexported (lowercase) so that they are only accessible by this
// package.
type MySampleApp struct {
	metav1.ObjectMeta
	exampleString string
	exampleInt    int
	description   string
	public *MySampleAppPublic
}

// MySampleAppPublic can be used for any public (or exported) facing mechanisms.
// - Kubernetes Custom Resources
// - Alternative to Values.yaml
// - Exposed over an HTTP API
// - Exposed over a gRPC API
//
// Here is where you could define a large amount of values that another mechanism could "tweak" or "configure"
// just like a Values.yaml.
//
// By making this (and the sub fields) exported we could expose this to other Go packages or even to a Kubernetes
// custom resource.
type MySampleAppPublic struct {

	// In case anyone is wondering this is "the new Values.yaml" as long as you plumb the fields
	// through in the "implementation" below.

	ExampleValue string // See line 84, and line 110
	ExampleNumber int
	ExampleText string
	ExampleToggle bool
	ExampleVerbose int
	ExampleName string
	ExampleAnnotations map[string]string
	ExampleValues map[int]string
	ExampleValue1 string
	ExampleValue2 string
	ExampleValue3 string
}

// New will create a naml compatible application.
//
// Note: There is nothing intrinsically wrong with making all the MySampleApplication fields
// exported (and deleting the MySampleAppPublic struct) if you wanted that functionality.
//
// I just suggested that we have an "exposed" and an "internal" representation of the same object. There are trade offs
// to this approach. Feel free to do whatever you want. 🤷‍♀️
func New(name string, namespace string, description string, publicApp *MySampleAppPublic) *MySampleApp {
	return &MySampleApp{
		ObjectMeta: metav1.ObjectMeta{
			Name:            name,
			Namespace:       namespace,
			ResourceVersion: "v1.0.0",
			Labels: map[string]string{
				"k8s-app":       "mysampleapp",
				"app":           "mysampleapp",
				"example-label": publicApp.ExampleValue, // Here we "plumb" through to our application from line 62
				"description":   description,
			},
			Annotations: publicApp.ExampleAnnotations,
		},
		exampleInt:    publicApp.ExampleNumber, // More plumbing
		exampleString: publicApp.ExampleValue1, // More plumbing
		description:   description,
		public: publicApp,
	}
}

func (v *MySampleApp) Install(client *kubernetes.Clientset) error {
	deployment := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: v.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: naml2.I32p(int32(v.exampleInt)),
			Selector: &metav1.LabelSelector{
				MatchLabels: v.Labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: v.Labels,
					Name: v.public.ExampleValue, // Here we "plumb" through to our application from line 62
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  v.Name,
							Image: "busybox",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	_, err := client.AppsV1().Deployments(v.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("unable to install deployment in Kubernetes: %v", err)
	}
	return nil
}

func (v *MySampleApp) Uninstall(client *kubernetes.Clientset) error {
	return client.AppsV1().Deployments(v.Namespace).Delete(context.TODO(), v.Name, metav1.DeleteOptions{})
}

func (v *MySampleApp) Meta() *metav1.ObjectMeta {
	return &v.ObjectMeta
}

func (v *MySampleApp) Description() string {
	return v.description
}
