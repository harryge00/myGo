package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"net/http"
	"io/ioutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ipMap := make(map[string]bool)
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for _, pod := range pods.Items {
		ipArr := strings.Split(pod.Annotations["ips"],"-")
		if len(ipArr) == 2 && ipArr[1] != "none" && ipArr[1] != "" {
			fmt.Println(ipArr[1])
			ipMap[ipArr[1]] = true
		}
	}
	baseIP := "10.30.99."
	mask := 21
	for i := 1; i < 255; i++ {
		ip := fmt.Sprintf("%s%d", baseIP, i)
		if ipMap[ip] {
			continue
		}
		requestStr := fmt.Sprintf("{\"startip\":\"%s\",\"mask\":%d}", ip, mask)
		fmt.Printf("Releasing IP %s\n", requestStr)
		ipURL := "http://10.30.100.12:7000/resource/delete"
		req, err := http.NewRequest("POST", ipURL, strings.NewReader(requestStr))
		if err != nil {
			fmt.Printf("http.NewRequest() error: %v\n", err)
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		c := &http.Client{}
		resp, err := c.Do(req)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}