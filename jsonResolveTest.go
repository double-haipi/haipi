package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	simplejson "github.com/bitly/go-simplejson"
)

//ServerJSON ...
type ServerJSON struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

//ServerJSONSlice ...
type ServerJSONSlice struct {
	Servers []ServerJSON `json:"servers"`
}

func parseJSONTest() {
	//user relative path
	file, err := os.Open("./resources/seversInfo.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	serverSlice := ServerJSONSlice{}
	json.Unmarshal(data, &serverSlice)
	fmt.Println(serverSlice)
}

//解析不知道其结构的json
func parseJSONTest2() {
	data := []byte(`{"Name":"Wedensday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var result interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("json 解析出错")
		return
	}

	//遍历
	for k, v := range result.(map[string]interface{}) {
		switch valueType := v.(type) {
		case string:
			fmt.Println(k, "is string", v)
		case float64:
			fmt.Println(k, "is int", v)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range valueType {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}

//使用simpleJson包测试
func simpleJSONTest() {
	origianJSON := `{
		"test":{
			"array":[
				1,
				"2",
				3
			],
			"int":10,
			"float":5.15,
			"bignum":91124444445,
			"string":"simpleJson",
			"bool":true
		}
	}`
	data := []byte(origianJSON)
	js, _ := simplejson.NewJson(data)
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	fmt.Println("print parsed value:")
	fmt.Println(arr)
	fmt.Println(i)
	fmt.Println(ms)

}

func generateJSONTest() {
	var s ServerJSONSlice
	s.Servers = append(s.Servers, ServerJSON{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, ServerJSON{"Beijing_VPN", "127.0.0.2"})
	json, err := json.Marshal(s)
	if err != nil {
		fmt.Println("encode to json failed")
		return
	}
	fmt.Println("encoded:", string(json))
}

//Employee ...
type Employee struct {
	ID        int    `json:"-"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PhoneNum  string `josn:"phoneNum,omitempty"`
}

func jsonTagTest() {
	employee := Employee{ID: 2011, FirstName: "haipi", LastName: `huang "1.0"`}
	jsonEmployee, _ := json.Marshal(employee)
	os.Stdout.Write(jsonEmployee)
}
