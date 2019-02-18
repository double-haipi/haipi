package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func isIP(ip string) bool {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		fmt.Println("ip 格式非法")
		return false
	}
	return true
}

func resolveContentFromWeb() {
	resp, err := http.Get("https://www.qq.com")
	if err != nil {
		fmt.Println("https get error")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)
	//转化为小写
	re, err := regexp.Compile(`\<[\S\s]+?\>`)
	if err != nil {
		fmt.Println("regex compile failed")
	}
	src = re.ReplaceAllStringFunc(src, strings.ToUpper)
	fmt.Println("after resolved：", src)
}

func regexpFindTest() {
	a := "I am learning Go language"
	re, _ := regexp.Compile(`[a-z]{2,4}`)

	target := []byte(a)
	//查找符合正则的第一个
	one := re.Find(target)
	formatPrint("Find", string(one))

	// 查找符合正则的所有slice，n小于0表示返回全部的符合项，否则返回指定长度的
	all := re.FindAll(target, -1)
	formatPrint("FindAll", "")
	for _, v := range all {
		fmt.Println(string(v))
	}

	//查找符合条件的index位置，开始位置和结束位置
	index := re.FindIndex(target)
	formatPrint("FindIndex", index)

	//所有符合条件的位置
	allindex := re.FindAllIndex(target, -1)
	formatPrint("FinaAllIndex", allindex)

	re2, _ := regexp.Compile(`am(.*)lang(.*)`)
	//查找Submatch，返回数组，第一个是匹配的全部原始，第二个是第一个（）里的，依次类推
	submatch := re2.FindSubmatch(target)
	fmt.Println("FindSubmatch", "")
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	submatchIndex := re2.FindSubmatchIndex(target)
	formatPrint("FindSubmatchIndex", "")
	fmt.Println(submatchIndex)

	submatchall := re2.FindAllSubmatch(target, -1)
	formatPrint("FindAllSubmatch", "")
	fmt.Println(submatchall)

	submatchallindex := re2.FindAllSubmatchIndex(target, -1)
	formatPrint("FindAllSubmatchIndex", "")
	fmt.Println(submatchallindex)
}

func formatPrint(prefix string, content interface{}) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println(prefix, content)
}
