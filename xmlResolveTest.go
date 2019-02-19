package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//Recurlyservers ...
type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svrs        []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func parseXML() {
	file, err := os.Open("D:/study/excises/Go/src/haipi/resources/servers.xml")
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	fmt.Println(v)
}

type servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svrs    []svr    `xml:"server"`
}

type svr struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func generateXML() {
	v := servers{Version: "1"}
	v.Svrs = append(v.Svrs, svr{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	v.Svrs = append(v.Svrs, svr{"Beijing_VPN", "127.0.0.2"})
	output, _ := xml.MarshalIndent(v, "  ", "   ")

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
