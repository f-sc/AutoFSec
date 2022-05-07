package AutoFSecBackend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
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
		fileInfo, _ := os.Stat(filename)
		readBuffer := make([]byte, fileInfo.Size())
		filePointer.Read(readBuffer)
		return readBuffer
	}
	return []byte{}
}

func CheckLoginState(r *http.Request) bool {
	cookieContents, cookieError := r.Cookie("login")
	if cookieError != nil {
		return false
	}
	return len(cookieContents.Value) > 0
}

func IsBrowserRequest(r http.Request) bool {
	clientTypeCookie, clCookieError := r.Cookie("client-type")
	if clCookieError == nil && clientTypeCookie.Value == "AUTO" {
		return false // handle json here
	}
	return true
}

func LoginHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if CheckLoginState(r) {
			w.Write(ReadWebFile("html/loginOK.html"))
		} else {
			w.Write(ReadWebFile("html/login.html"))
		}
	} else if r.Method == http.MethodPost {
		contents, _ := ioutil.ReadAll(r.Body)
		fmt.Println("Login HTTP got contents: " + string(contents))
		splittedContents := strings.Split((string)(contents), "&")
		if len(splittedContents) > 1 {
			username := strings.Split(splittedContents[0], "=")[1]
			password := strings.Split(splittedContents[1], "=")[1]
			if SystemManager.LogIn(username, CryptoHashString(password)) {
				loginCookie := http.Cookie{
					Name:   "login",
					Value:  username,
					MaxAge: 300,
				}
				http.SetCookie(w, &loginCookie)
				http.Redirect(w, r, "/account", http.StatusSeeOther)
				w.Write(([]byte)("LOGIN OK"))
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				w.Write(([]byte)("LOGIN ERROR"))
			}
		}
	}
}

type CarData struct {
	Username             string
	VehicleManufacturer  string
	SecuritySystemStatus string
	LastCheckStatus      string
}

func ShowAccount(w http.ResponseWriter, r *http.Request) {
	cookieHandle, cookieError := r.Cookie("login")
	if cookieError == nil {
		accountHtmlTemplate, _ := template.ParseFiles("html/account.html")
		wd := CarData{
			Username:             cookieHandle.Value,
			VehicleManufacturer:  SystemManager.GetVehicleInfo(r),
			SecuritySystemStatus: SystemManager.GetSecuritySystemConnectionStatus(r),
			LastCheckStatus:      fmt.Sprint("Date: ", time.Now().Day()) + " " + time.Now().Month().String(),
		}
		if !IsBrowserRequest(*r) {
			jsonConvertedResult, jsonErrorConversion := json.Marshal(wd)
			if jsonErrorConversion == nil {
				w.Write(jsonConvertedResult)
			}
		} else {
			accountHtmlTemplate.Execute(w, &wd)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
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
	http.HandleFunc("/account", ShowAccount)
	ex1()
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./html/css"))))

	http.ListenAndServe(":5555", nil)
}
