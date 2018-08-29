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
	var ids=[]int{106,
		99,
		99,
		99,
		142,
		99,
		99,
		149,
		149,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		210,
		252,
		106,
		106,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		115,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		106,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		84,
		78,
		162,
		206,
		122,
		206,
		179,
		135,
		72,
		187,
		84,
		84,
		84,
		84,
		84,
		132,
		137,
		84,
		84,
		84,
		84,
		84,
		258,
		137,
		137,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		86,
		85,
		85,
		114,
		114,
		281,
		281,
		281,
		281,
		281,
		278,
		278,
		278,
		278,
		278,
		278,
		278,
		278,
		3,
		258,
		283,
		191,
		210,
		210,
		210,
		210,
		210,
		32,
		210,
		210,
		210,
		210,
		137,
		210,
		258,
		278,
		132,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		180,
		204,
		285,
		285,
		285,
		285,
		285,
		204,
		43,
		43,
		43,
		204,
		132,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		28,
		204,
		204,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		263,
		132,
		114,
		32,
		193,
		32,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		284,
		72,
		72,
		72,
		72,
		72,
		137,
		137,
		278,
		278,
		278,
		278,
		278,
		278,
		278,
		278,
		278,
		46,
		46,
		127,
		193,
		211,
		211,
		211,
		211,
		211,
		28,
	}
	var ips=[]string{
		"172.25.12.174",
		"172.25.12.180",
		"172.25.12.222",
		"172.25.12.223",
		"172.25.13.144",
		"172.25.13.150",
		"172.25.13.151",
		"172.25.12.218",
		"172.25.12.173",
		"172.25.42.6",
		"172.25.42.19",
		"172.25.42.55",
		"172.25.42.57",
		"172.25.42.58",
		"172.25.42.65",
		"172.25.42.66",
		"172.25.42.67",
		"172.25.42.68",
		"172.25.42.69",
		"172.25.42.70",
		"172.25.42.71",
		"172.25.42.103",
		"172.25.42.126",
		"172.25.42.127",
		"172.25.42.129",
		"172.25.42.141",
		"172.25.13.195",
		"172.25.13.198",
		"172.25.13.171",
		"172.25.13.176",
		"172.25.13.197",
		"172.25.42.175",
		"172.25.42.176",
		"172.25.42.177",
		"172.25.42.178",
		"172.25.42.179",
		"172.25.42.180",
		"172.25.42.181",
		"172.25.42.182",
		"172.25.42.183",
		"172.25.12.226",
		"172.25.42.18",
		"172.25.42.73",
		"172.25.13.194",
		"172.25.13.214",
		"172.25.42.2",
		"172.25.42.4",
		"172.25.42.24",
		"172.25.42.26",
		"172.25.42.30",
		"172.25.42.34",
		"172.25.42.37",
		"172.25.42.38",
		"172.25.42.168",
		"172.25.42.171",
		"172.25.42.174",
		"172.25.12.242",
		"172.25.13.222",
		"172.25.13.225",
		"172.25.42.5",
		"172.25.42.7",
		"172.25.42.8",
		"172.25.42.9",
		"172.25.42.45",
		"172.25.42.48",
		"172.25.42.60",
		"172.25.12.131",
		"172.25.12.136",
		"172.25.12.142",
		"172.25.12.143",
		"172.25.42.61",
		"172.25.42.62",
		"172.25.42.79",
		"172.25.42.85",
		"172.25.42.89",
		"172.25.42.90",
		"172.25.42.92",
		"172.25.42.105",
		"172.25.42.106",
		"172.25.42.118",
		"172.25.42.123",
		"172.25.42.140",
		"172.25.42.142",
		"172.25.42.156",
		"172.25.42.157",
		"172.25.42.13",
		"172.25.42.14",
		"172.25.42.15",
		"172.25.42.16",
		"172.25.42.124",
		"172.25.42.158",
		"172.25.42.159",
		"172.25.42.161",
		"172.25.42.163",
		"172.25.42.78",
		"172.25.42.131",
		"172.25.42.132",
		"172.25.42.165",
		"172.25.42.166",
		"172.25.42.167",
		"172.25.42.169",
		"172.25.42.172",
		"172.25.42.173",
		"172.25.42.184",
		"172.25.42.187",
		"172.25.42.188",
		"172.25.42.189",
		"172.25.42.190",
		"172.25.42.191",
		"172.25.42.3",
		"172.25.42.20",
		"172.25.42.12",
		"172.25.12.188",
		"172.25.42.11",
		"172.25.42.41",
		"172.25.13.164",
		"172.25.42.46",
		"172.25.12.153",
		"172.25.13.166",
		"172.25.13.190",
		"172.25.13.207",
		"172.25.13.211",
		"172.25.13.227",
		"172.25.12.132",
		"172.25.12.128",
		"172.25.12.140",
		"172.25.12.141",
		"172.25.12.144",
		"172.25.12.146",
		"172.25.12.147",
		"172.25.12.210",
		"172.25.12.211",
		"172.25.12.215",
		"172.25.12.231",
		"172.25.12.232",
		"172.25.12.239",
		"172.25.12.240",
		"172.25.12.241",
		"172.25.12.243",
		"172.25.12.244",
		"172.25.12.245",
		"172.25.12.251",
		"172.25.12.252",
		"172.25.12.253",
		"172.25.13.159",
		"172.25.13.160",
		"172.25.13.161",
		"172.25.13.162",
		"172.25.13.163",
		"172.25.13.165",
		"172.25.12.254",
		"172.25.13.167",
		"172.25.13.168",
		"172.25.13.169",
		"172.25.13.173",
		"172.25.13.174",
		"172.25.12.230",
		"172.25.13.179",
		"172.25.13.181",
		"172.25.13.192",
		"172.25.13.193",
		"172.25.13.196",
		"172.25.13.199",
		"172.25.13.200",
		"172.25.13.215",
		"172.25.42.10",
		"172.25.42.21",
		"172.25.42.23",
		"172.25.42.27",
		"172.25.42.28",
		"172.25.42.29",
		"172.25.42.31",
		"172.25.42.32",
		"172.25.42.35",
		"172.25.12.156",
		"172.25.42.39",
		"172.25.12.145",
		"172.25.12.187",
		"172.25.12.189",
		"172.25.12.204",
		"172.25.12.212",
		"172.25.13.172",
		"172.25.13.217",
		"172.25.13.221",
		"172.25.42.36",
		"172.25.42.40",
		"172.25.12.214",
		"172.25.13.220",
		"172.25.12.172",
		"172.25.42.22",
		"172.25.42.95",
		"172.25.13.177",
		"172.25.13.178",
		"172.25.13.180",
		"172.25.13.183",
		"172.25.13.184",
		"172.25.13.185",
		"172.25.13.186",
		"172.25.13.187",
		"172.25.13.188",
		"172.25.13.191",
		"172.25.13.201",
		"172.25.13.202",
		"172.25.13.206",
		"172.25.13.208",
		"172.25.13.209",
		"172.25.13.210",
		"172.25.13.212",
		"172.25.13.218",
		"172.25.42.54",
		"172.25.42.72",
		"172.25.42.82",
		"172.25.42.110",
		"172.25.42.107",
		"172.25.42.108",
		"172.25.42.109",
		"172.25.42.113",
		"172.25.42.115",
		"172.25.42.88",
		"172.25.42.117",
		"172.25.12.219",
		"172.25.42.83",
		"172.25.42.112",
		"172.25.42.119",
		"172.25.42.122",
		"172.25.42.133",
		"172.25.42.134",
		"172.25.42.135",
		"172.25.42.136",
		"172.25.42.137",
		"172.25.42.145",
		"172.25.42.146",
		"172.25.42.147",
		"172.25.42.148",
		"172.25.42.44",
		"172.25.42.47",
		"172.25.42.49",
		"172.25.42.50",
		"172.25.42.51",
		"172.25.42.52",
		"172.25.42.53",
		"172.25.42.59",
		"172.25.42.63",
		"172.25.42.64",
		"172.25.42.74",
		"172.25.12.134",
		"172.25.12.167",
		"172.25.12.175",
		"172.25.12.176",
		"172.25.12.181",
		"172.25.12.182",
		"172.25.12.183",
		"172.25.12.190",
		"172.25.12.192",
		"172.25.12.193",
		"172.25.12.209",
		"172.25.12.221",
		"172.25.12.228",
		"172.25.12.229",
		"172.25.13.213",
		"172.25.13.224",
		"172.25.13.228",
		"172.25.13.229",
		"172.25.42.42",
		"172.25.42.43",
		"172.25.42.75",
		"172.25.42.77",
		"172.25.42.80",
		"172.25.42.81",
		"172.25.42.86",
		"172.25.42.87",
		"172.25.13.219",
		"172.25.42.99",
		"172.25.42.100",
		"172.25.42.101",
		"172.25.42.104",
		"172.25.42.116",
		"172.25.42.149",
		"172.25.42.150",
		"172.25.42.151",
		"172.25.42.152",
		"172.25.42.153",
		"172.25.42.154",
		"172.25.42.155",
		"172.25.42.162",
		"172.25.42.170",
		"172.25.42.185",
		"172.25.42.186",
		"172.25.42.192",
		"172.25.42.193",
		"172.25.42.194",
		"172.25.42.195",
		"172.25.42.196",
		"172.25.42.197",
		"172.25.12.139",
		"172.25.12.155",
		"172.25.12.179",
		"172.25.12.205",
		"172.25.12.224",
		"172.25.12.216",
		"172.25.12.217",
		"172.25.42.198",
		"172.25.42.201",
		"172.25.42.202",
		"172.25.42.203",
		"172.25.42.204",
		"172.25.42.212",
		"172.25.42.213",
		"172.25.42.214",
		"172.25.42.215",
		"172.25.13.216",
		"172.25.42.98",
		"172.25.13.170",
		"172.25.42.91",
		"172.25.12.129",
		"172.25.12.137",
		"172.25.12.138",
		"172.25.12.165",
		"172.25.12.166",
		"172.25.42.255",
	}

	users := []string{
		"zhangyuandao",
		"wangzhuzhen",
		"fangying",
		"jianmei1987",
		"xufei",
		"dongshuling",
		"chenmiao",
		"haoyuan",
		"zoufan",
		"senyang",
		"12729",
		"zhangkan",
		"dongsl",
		"dongsl2",
		"zoufan1",
		"pxg_30686",
		"peiqi",
		"dongsl3",
		"wupeiyu",
		"wanghui",
		"guest306",
		"hqw151436",
		"13110",
		"25010",
		"12971",
		"dongsl4",
		"pxg-test1",
		"wanghui3",
		"docker-test",
		"lifangyuan",
		"dingyangzhen",
		"luotingting",
		"zenglihui",
		"peiqi-2",
		"guest308",
		"dong-sl",
		"31852",
		"12381",
		"dhtest",
		"yanhe",
		"xuyan",
		"song-jing",
		"yuanbo",
		"32310",
		"pshi",
		"chaofan",
		"lizhile",
		"33384",
		"dssc8500",
		"xiaowei",
		"37299",
		"chenwangxin",
		"senyang3",
		"36302",
		"zmeng",
		"falcon-app",
		"16945",
		"chenshanwei",
		"abcd",
		"abcd123",
		"abcd1234",
		"28851",
		"cci",
		"33165",
		"huang-zhixin",
		"xuweiping",
		"liangchao-1988",
		"huangzongyin",
		"32528",
		"pilaph",
		"liang-chao",
		"zhangneng",
		"ray110",
		"33914",
		"chebysheph",
		"wuchao",
		"32281",
		"jiangzhe",
		"33402",
		"z35437",
		"helinqiang",
		"kankan",
		"yangli",
		"chenkunning",
		"wang-huan3",
		"leiwangsheng",
		"27590",
		"yegaojiang",
		"huanwang",
		"36928",
		"richard",
		"cuican",
		"gu86266593",
		"25279",
		"jchen",
		"37304",
		"26803",
		"sggh",
		"yaojianbin",
		"panxing123",
		"18962",
		"leon",
		"limingxia",
		"fankaihui",
		"14457",
		"lixiangjing",
		"lynn0616",
		"dingbaozhong",
		"chen-555",
		"test-pao",
		"32043",
		"yangtze",
		"24826",
		"leo",
		"39561",
		"20808",
		"40082",
		"majun",
		"vms",
		"wang-hui3",
		"jhj",
		"liweina",
		"26007",
		"dhcloud",
		"27281",
		"31936",
		"hejinbiao",
		"zyh",
		"myhorchid",
		"13021",
		"yxx916",
		"123",
		"34583",
		"hu-yaolin",
		"tym",
		"33435",
		"25274",
		"sunwei",
		"23245",
		"xqf",
		"wangmeng",
		"toge",
		"wangjun",
		"test",
		"senyang2",
		"36687",
		"32274",
		"yulexin",
		"28775",
		"test1234",
		"36302ydw",
		"lmx",
		"hezhenyu",
		"fdasfa",
		"wangchenfen",
		"wang-chenfen",
		"31777",
		"qq123",
		"suncheng",
		"37305",
		"liqi",
		"35437",
		"23972",
		"zmwtest",
		"jiang",
		"36437",
		"31049",
		"zhulong",
		"20717",
		"chenhaijun",
		"xuechu123",
		"sipeng",
		"wangqq",
		"39564",
		"33344",
		"admin",
		"robot-ai-group",
		"30808",
		"23533",
		"zheng-ji",
		"lx01",
		"20756",
		"liuxuan",
		"hebingting",
		"27388",
		"33476",
		"38051",
		"27905",
		"35657",
		"26013",
		"27706",
		"liuxin",
		"mengfanxu",
		"50518",
		"kb24",
		"liujianhai",
		"24415",
		"guxiawei",
		"32460",
		"haoyuan2",
		"xufei1",
		"dahua12819",
		"15692",
		"40824",
		"32901",
		"26089",
		"23126",
		"42077",
		"50780",
		"32067",
		"36189",
		"29068",
		"lgying",
		"29456",
		"fangzhenyu",
		"jiang-yu",
		"dh39944",
		"27524",
		"dh399442",
		"panhj",
		"liuguiying",
		"pentest",
		"pen1",
		"demozyd",
		"guanbo",
		"mike",
		"jiang-huizhong",
		"lzy",
		"yangchao9",
		"demojhz",
		"40313",
		"40392",
		"41224",
		"wang-zhihao",
		"myc",
		"z42048dahuatech",
		"27461",
		"41080",
		"40918",
		"xiajinming",
		"yukan",
		"33920",
		"37849",
		"panxingtest123",
		"42892",
		"dhmap",
		"yuzhuo",
		"33511",
		"35443",
		"zhang-jie20",
		"dingchao",
		"32789daiyuanpeng",
		"bai-feijian1009",
		"zhang-fan",
		"34049",
		"dongchaohui",
		"qihuidong",
		"zhangweiyi",
		"pippo",
		"hzdmw",
		"40360",
		"hncscwc",
		"32073",
		"21451",
		"16810",
		"30984",
		"43189",
		"spencer",
		"wangyunjin",
		"gaozhiqiang",
		"taoxianhua",
		"41717",
		"cyj",
		"xuechu",
		"yeyuntest",
	}
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
			fmt.Println(string(bytes))
			//_, err = sendReleaseIpReq(bytes, "http://10.30.100.10:8090/api/net/ip/release")
			//if err != nil {
			//	fmt.Println(err)
			//}
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