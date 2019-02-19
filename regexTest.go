package main

import (
	"fmt"
<<<<<<< HEAD
	"regexp"
)

func regexTest() {
	text := `Hello 世界！123 Go.`
	reg := regexp.MustCompile(`[^a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, 2))

	// 查找连续的单词字母
	reg = regexp.MustCompile(`[\w]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "123" "Go"]

	// 查找连续的非单词字母、非空白字符
	reg = regexp.MustCompile(`[^\w\s]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["世界！" "."]

	// 查找连续的大写字母
	reg = regexp.MustCompile(`[[:upper:]]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["H" "G"]

	// 查找连续的非 ASCII 字符
	reg = regexp.MustCompile(`[[:^ascii:]]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["世界！"]

	// 查找连续的标点符号
	reg = regexp.MustCompile(`[\pP]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["！" "."]

	// 查找连续的非标点符号字符
	reg = regexp.MustCompile(`[\PP]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界" "123 Go"]

	// 查找连续的汉字
	reg = regexp.MustCompile(`[\p{Han}]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["世界"]

	// 查找连续的非汉字字符
	reg = regexp.MustCompile(`[\P{Han}]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello " "！123 Go."]

	// 查找 Hello 或 Go
	reg = regexp.MustCompile(`Hello|Go`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找行首以 H 开头，以空格结尾的字符串
	reg = regexp.MustCompile(`^H.*\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界！123 "]

	// 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
	reg = regexp.MustCompile(`(?U)^H.*\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello "]

	// 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
	reg = regexp.MustCompile(`(?i:^hello).*Go`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界！123 Go"]

	// 查找 Go.
	reg = regexp.MustCompile(`\QGo.\E`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Go."]

	// 查找从行首开始，以空格结尾的字符串（非贪婪模式）
	reg = regexp.MustCompile(`(?U)^.* `)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello "]

	// 查找以空格开头，到行尾结束，中间不包含空格字符串
	reg = regexp.MustCompile(` [^ ]*$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// [" Go."]

	// 查找“单词边界”之间的字符串
	reg = regexp.MustCompile(`(?U)\b.+\b`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" " 世界！" "123" " " "Go"]

	// 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
	reg = regexp.MustCompile(`[^ ]{1,4}o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找 Hello 或 Go
	reg = regexp.MustCompile(`(?:Hell|G)o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找 Hello 或 Go，替换为 Hellooo、Gooo
	reg = regexp.MustCompile(`(?PHell|G)o`)
	fmt.Printf("%q\n", reg.ReplaceAllString(text, "${n}ooo"))
	// "Hellooo 世界！123 Gooo."

	// 交换 Hello 和 Go
	// reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
	// fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
	// "Go 世界！123 Hello."

	// 特殊字符的查找
	reg = regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]\|]`)
	fmt.Printf("%q\n", reg.ReplaceAllString("\f\t\n\r\v\123\x7F\U0010FFFF\\^$.*+?{}()[]|", "-"))

}

func zhMatchTest() {
	s := "中文1"
	if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", s); !m {
		fmt.Println("not all are zh-cn")
	} else {
		fmt.Println(" all are zh-cn")
	}
}

func numMatchTest() {
	if m, _ := regexp.MatchString("^[0-9]+$", "12a3"); !m {
		fmt.Println("非数字")
	} else {
		fmt.Println("数字")
	}
=======
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

func regexExpandTest() {
	source := []byte(`
	call hello alice
	hello bob
	call hello eve
	`)
	pattern := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	result := []byte{}
	indexs := pattern.FindAllSubmatchIndex(source, -1)
	fmt.Println("indexes:", indexs)
	for _, s := range indexs {
		result = pattern.Expand(result, []byte("$cmd('$arg')\n"), source, s)
	}
	fmt.Println(string(result))
>>>>>>> 473aa854f40af9689c974fb5c731b54cc5acd3fd
}
