package main

import (
	"fmt"
	"strings"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"

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
	var ips = []string{
		"10.6.13.31 453",
		"10.6.13.32 455",
		"10.6.13.33 405",
		"10.6.13.35 457",
		"10.6.13.37 459",
		"10.6.13.38 462",
		"10.6.13.39 461",
		"10.6.13.40 460",
		"10.6.13.43 467",
		"10.6.13.44 466",
		"10.6.13.45 469",
		"10.6.13.46 470",
		"10.6.13.47 471",
		"10.6.13.48 468",
		"10.6.13.49 474",
		"10.6.13.50 438",
		"10.6.13.52 90",
		"10.6.13.53 475",
		"10.6.13.56  481",
		"10.6.13.57  480",
		"10.6.13.58  482",
		"10.6.13.59  484",
		"10.6.13.62  486",
		"10.6.13.63  477",
		"10.6.13.64  487",
		"10.6.13.65  488",
		"10.6.13.66  489",
		"10.6.13.68  491",
		"10.6.13.70  352",
		"10.6.13.71  493",
		"10.6.13.73  495",
		"10.6.13.61  440",
		"10.6.13.74  496",
		"10.6.13.76  502",
		"10.6.13.77  503",
		"10.6.13.79  507",
		"10.6.13.80  508",
		"10.6.13.81  398",
		"10.6.13.82  510",
		"10.6.13.83  516",
		"10.6.13.84  473",
		"10.6.13.4     8",
		"10.6.13.86  523",
	}
	for _, x := range ips {
		ipUser := strings.Fields(x)
		err := releaseIP(ipUser)
		if err != nil {
			panic(err)
		}
	}
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

func releaseIP(ipuser []string) error {
	// Currently, no userid 0
	userId := ipuser[1]
	ip := ipuser[0]
	uid, err := strconv.Atoi(userId)
    if err != nil {
    	panic(err)
    }

	req := IpRelease{
		IP:     ip,
		UserId: uid,
		Group: "adefaultservice",
	}
	bytes, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	// _, err = sendReleaseIpReq(bytes, "http://10.30.100.10:8090/api/net/ip/release")
	fmt.Printf("releaseIP %v %v", req, err)
	return err
}