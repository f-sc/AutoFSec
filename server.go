package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

const indexPage2 = `
<h1>Login done!</h1>
`

func ReadWebFile(filename string) []byte {
	filePointer, error := os.OpenFile(filename, os.O_RDWR, 0666)
	if error == nil {
		fileInfo, _ := os.Stat("html/index.html")
		readBuffer := make([]byte, fileInfo.Size())
		filePointer.Read(readBuffer)
		return readBuffer
	}
	return []byte{}
}

func LoginHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write(ReadWebFile("html/index.html"))
	} else if r.Method == http.MethodPost {
		contents, _ := ioutil.ReadAll(r.Body)
		fmt.Println("Login HTTP got contents: " + string(contents))
		splittedContents := strings.Split((string)(contents), "&")
		if len(splittedContents) > 1 {
			username := strings.Split(splittedContents[0], "=")[1]
			password := strings.Split(splittedContents[1], "=")[1]
			fmt.Println("*** TRYING to login: " + username + " : " + password)
			var newUser User
			if newUser.LogIn(username, CryptoHashString(password)) {
				w.Write(ReadWebFile("html/loginOK.html"))
			} else {
				w.Write(ReadWebFile("html/index.html"))
			}
		}
	}
}

func LogRet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(indexPage2))
}

func ex1() {
	d := http.Dir(".")
	f, err := d.Open("./html/css/styles.css")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func InitializeServer(port string) {
	http.HandleFunc("/login", LoginHttp)
	http.HandleFunc("/ret", LogRet)
	//http.HandleFunc("/css/styles.css", FileStylesCSS)
	fmt.Println(os.Getwd())
	ex1()
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./html/css"))))

	http.ListenAndServe(":5555", nil)
}
