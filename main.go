//main comment
package main

import (
	"errors"
	"fmt"
	"haipi/session"
	"net/http"
	"time"
	_ "time"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
}

// entrance
func main() {
	// addTest()
	// typeFuncTest()
	// ArrayTest()
	// getAverageTest()
	// booksTest()
	// sliceTest()
	// rangeTest()
	// mapTest()
	// chanelTest()
	// stringTest()
	// errorTest()
	// constTest()
	// sliceAndArray()
	// mapTest1()
	// switchTest()
	// sumAndProductTest()
	paramsTest(1, 2, 3, 4)
}

//example
func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {

	}
}

func paramsTest(arg ...int) {
	fmt.Println("arg list:")
	for _, v := range arg {
		fmt.Println(v)
	}
	fmt.Println("remove prefix of fmt")
}

func sumAndProductTest() {
	sum, product := SumAndProduct(10, 20)
	fmt.Println("SumAndProduct:", sum, product)
}

//SumAndProduct 计算和和积
func SumAndProduct(A, B int) (sum int, product int) {
	sum = A + B
	product = A * B
	return
}

func switchTest() {
	number := 5
	switch number {
	case 2:
		fmt.Println("number is less than 2")
		fallthrough
	case 3:
		fmt.Println("less than 3")
		fallthrough
	case 4, 5:
		fmt.Println("less than 5")
		fallthrough
	case 6:
		fmt.Println("less than 6")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func mapTest1() {
	numbers := make(map[string]int)
	numbers["haipi"] = 30
	numbers["xiaoying"] = 29
	delete(numbers, "haipi")
	fmt.Printf("map:%v", numbers)

	rating := map[string]float32{"c": 5, "go": 4.5, "python": 4.5, "c++": 2}
	fmt.Printf("rating:%v", rating)
	csharpRating, ok := rating["c#"]
	if ok {
		fmt.Println("c# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("c# is not in the map")
	}
}

func sliceAndArray() {
	arr := [...]int{1, 2, 3, 4, 5}
	a, b := arr[4:], arr[2:4]
	a[0] = 7
	a = append(a, 10)
	fmt.Printf("slice1:%v,slice2:%v,array:%v", a, b, arr)
}

func constTest() {
	const (
		red = iota
		green
		blue
	)

	fmt.Printf("red:%v,green:%v,blue:%v", red, green, blue)
}

func errorTest() {
	err := errors.New("custom error raise")
	if err != nil {
		fmt.Println(err)
	}
}

func stringTest() {
	s := "hello"
	c := []byte(s)
	c[0] = 'j'
	s2 := string(c)
	fmt.Println(s2)

	s = "c" + s[1:]
	fmt.Printf("slice string %s\n", s)
}

//闭包 匿名函数
func add(x1, x2 int) func(x3, x4 int) (int, int, int) {
	i := 0
	return func(x3, x4 int) (int, int, int) {
		i++
		return i, x1 + x2, x3 + x4
	}
}

func addTest() {
	addFunc := add(1, 2)
	fmt.Println(addFunc(3, 4))
	fmt.Println(addFunc(5, 6))
	fmt.Println(addFunc(7, 8))
}

//类型方法
type Circle struct {
	radius float64
}

func (c Circle) getArea(addtion int) float64 {
	return 3.14*c.radius*c.radius + float64(addtion)
}

func typeFuncTest() {
	var circle Circle
	circle.radius = 10.00
	fmt.Printf("getArea:%f", circle.getArea(1))
}

//数组初始化
var balance = [5]float32{1.0}

var twoDimensionArray = [3][3]int{
	{1, 2, 3},
	{1, 2, 3},
	{1, 2, 3},
}

//ArrayTest instrument
func ArrayTest() {
	fmt.Printf("balance[0]:%f", balance[1])
}

func getAverage(arr []int, size int) float32 {
	var average, sum float32
	for i := 0; i < size; i++ {
		sum += float32(arr[i])
	}
	average = sum / float32(size)
	return average
}

func getAverageTest() {
	arr := []int{111, 3, 533, 3}
	fmt.Printf("getAverage:%f", getAverage(arr, len(arr)))
}

//Books 类型
type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

func booksTest() {
	fmt.Println(Books{"Go language", "foreign", "computer", 2001})
}

//slice
var slice1 = make([]int, 10)

func sliceTest() {
	slice := make([]int, 3, 5)
	fmt.Printf("slice len:%d,cap:%d,content:%v\n", len(slice), cap(slice), slice)
}

func rangeTest() {
	arr := []int{2, 3, 4, 5}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Printf("rangeTest sum:%d\n", sum)

	//map
	mapVal := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range mapVal {
		fmt.Printf("key:%s,val:%s\n", k, v)
	}
}

func mapTest() {
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["Frans"] = "Paris"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["China"] = "Beijing"

	_, ok := countryCapitalMap["America"]
	if ok {
		fmt.Println("America is in map")
	} else {
		fmt.Println("America is not in map")
	}
}

func chanelTest() {
	c := make(chan int)
	go say2("world", c)
	say("hello")
	fmt.Println(<-c)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}
}

func say2(s string, c chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
	c <- 0
	close(c)
}
