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
	"fmt"
	"os"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	//"flag"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/tools/clientcmd"
	//"strings"
)

func getOccupiedIPFromMysql(username string) []string {
	return []string{"172.25.42.180"}
}
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(10.30.100.6:3306)/net_manager?charset=utf8")
	db.Ping()
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM alloc_ip where user_id = " + os.Args[1])
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Fetch rows
	for rows.Next() {
		var id int
		var ip string
		var typ int
		var mask int
		var routes string
		var createTime time.Time
		var updateTime time.Time
		var user_id int
		var occupied int
		var group sql.NullString
		var location string
		// get RawBytes from data
		err = rows.Scan(&id, &typ, &ip, &mask, &routes, &createTime, &updateTime, &user_id, &occupied, &group, &location)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		fmt.Println(id, ip, user_id, occupied, group, location)
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//fmt.Println("occupied_ip ", len(occupied_ip))
	//
	//kubeconfig := flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
	//flag.Parse()
	//// uses the current context in kubeconfig
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err.Error())
	//}
	//// creates the clientset
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}
	//pods, err := clientset.CoreV1().Pods(os.Args[1]).List(metav1.ListOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//for _, pod := range pods.Items {
	//	if pod.Annotations["ips"] != "" && strings.Contains(pod.Name, os.Args[2]) {
	//		ips := strings.Split(pod.Annotations["ips"], "-")
	//		fmt.Println(pod.OwnerReferences[0].Name, ips[1])
	//	}
	//}
}


