/*
Copyright 2016 The Kubernetes Authors.

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
	"flag"
	"fmt"
	"os"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/pkg/api/v1"
	"strings"
)


func main() {
	namespace := 0
	if len(os.Args) == 2 {
		namespace = os.Args[1]
	}
	occupied_map := make(map[string]bool)
	for _, ip := range occupied_ip {
		occupied_map[ip] = false
	}
	kubeconfig := flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	eps, err := clientset.CoreV1().Endpoints(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, ep := range eps.Items {
		if pod.Annotations["ips"] != "" {
			ips := strings.Split(pod.Annotations["ips"], "-")
			occupied_map[ips[1]] = true
		}
	}
	for key, val := range occupied_map {
		if val == false {
			fmt.Println(pod.Name, key)
		}
	}
}
