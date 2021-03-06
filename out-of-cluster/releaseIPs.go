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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
	"bytes"
	"github.com/golang/glog"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"k8s.io/client-go/pkg/api/v1"
)

type IpRelease struct {
	IP     string `json:"ip,omitempty"`
	Group  string `json:"group,omitempty"`
	UserId int    `json:"userId,omitempty"`
}

type IpGroupClear struct {
	Group  string `json:"group,omitempty"`
	UserId int    `json:"userId,omitempty"`
}

type IpReleaseResp struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func main() {
	
	fmt.Println(len(ids), len(ips))
	if len(ids) != len(ips) {
		panic("Len Not EQ!")
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
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for i := range ips {
		ip := ips[i]
		id := ids[i]
		findFlag := false
		for _, pod := range pods.Items {
			if strings.Contains(pod.Annotations["ips"], ip) {
				findFlag = true
				break
			}
		}

		if !findFlag {
			namespace := fmt.Sprintf("%s-%d", users[ids[i] - 1], ids[i])
			fmt.Println(ip, namespace)
			req := IpRelease{
				IP:     ip,
				UserId: id,
			}
			bytes, err := json.Marshal(req)
			if err != nil {
				glog.Errorf("json errorf %v", err)
			}
			_, err = sendReleaseIpReq(bytes, "http://10.30.100.10:8090/api/net/ip/release")
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	//for key, val := range occupied_map {
	//	if val == false {
	//		fmt.Println(pod.Name, key)
	//	}
	//}
}

func sendReleaseIpReq(reqBytes []byte, url string) (code int, err error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var ipResp IpReleaseResp
	err = json.Unmarshal(body, &ipResp)
	if err != nil {
		return
	}
	code = ipResp.Code
	if code != 200 {
		err = fmt.Errorf("%v", ipResp.Message)
	}
	return
}

func releaseIP(userId int, po *v1.Pod) error {
	// Currently, no userid 0
	if userId == 0 {
		return nil
	}
	if ips := po.Annotations["ips"]; ips != "" {
		ipArr := strings.Split(ips, "-")
		if len(ipArr) == 2 && ipArr[1] != "none" && ipArr[1] != "empty" {
			req := IpRelease{
				IP:     ipArr[1],
				UserId: userId,
			}
			if po.Labels["networkgroup"] != "" {
				req.Group = po.Labels["networkgroup"]
			}
			bytes, err := json.Marshal(req)
			if err != nil {
				glog.Errorf("json errorf %v", err)
			}
			_, err = sendReleaseIpReq(bytes, "http://10.30.100.10:8090/api/net/ip/release")
			glog.V(6).Infof("releaseIP %v %v", req, err)
			return err
		}
	}
	return nil
}