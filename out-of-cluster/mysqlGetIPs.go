package main

import (
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"flag"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var IPUser struct {
	IP string
	UserID string
}
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(172.25.3.195:3306)/net_manager")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("select * from alloc_ip where occupied=1;")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	ipToUsers := make(map[string]string)
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value, userID, ip string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			//fmt.Println(columns[i], ": ", value)
			if columns[i] == "user_id" {
				userID = value
			} else if columns[i] == "ip" {
				ip = value
			}
		}
		ipToUsers[ip] = userID
		//fmt.Println("-----------------------------------")
	}
	fmt.Println(ipToUsers)
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
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
	pods, err := clientset.CoreV1().Pods("").List(meta_v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, po := range pods.Items {
		if po.Annotations["ips"] != "" {
			strArr := strings.Split(po.Annotations["ips"], "-")
			if len(strArr) == 2 {
				ipToUsers[strArr[1]] = "found"
			}
		}
	}
	for ip := range ipToUsers {
		if ipToUsers[ip] != "found" {
			fmt.Println(ip, ipToUsers[ip])
		}
	}

}