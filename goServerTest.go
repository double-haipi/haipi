package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func goSvr() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/loginSession",loginSession)
	http.HandleFunc("/count",count)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

func sayHellonName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("path:", r.URL.Path)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello haipi")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(curTime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	t, err := template.ParseFiles("D:/study/excises/Go/src/haipi/resources/login.gtpl")
	if err != nil {
		fmt.Println("parse err:", err)
	} else {
		t.Execute(w, token)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "POST" {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("token is valid")
		} else {
			fmt.Println("token is invalid")
		}

		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("escaped username:", template.HTMLEscapeString(r.Form.Get("username")))
		template.HTMLEscape(w, []byte("login success"))

	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request method:", r.Method)
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("uploadedFile.txt", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

}
