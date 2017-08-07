package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// func handler(rw http.ResponseWriter, req *http.Request) {

// 	fmt.Fprintf(rw, "Hi there, I love %s!", req.URL.Path[1:])

// }

var ClientID = "c164a88d1422407fa9d17dd31178ce18"
var ClientSecret = "7043a5921b634bd08addea668367ff4e"
var RedirectURI = "http://libshaft.szenyo.com:9292/redirect"

func redirect(res http.ResponseWriter, req *http.Request) {

	code := req.FormValue("code")

	if len(code) != 0 {

		formResponse, err := http.PostForm("https://api.instagram.com/oauth/access_token", url.Values{"client_id": {ClientID}, "client_secret": {ClientSecret}, "grant_type": {"authorization_code"}, "redirect_uri": {RedirectURI}, "code": {code}})
		if err != nil {
			fmt.Println(err)
			http.NotFound(res, req)
			return
		}
		defer formResponse.Body.Close()

		if formResponse.StatusCode == 200 {

			body, _ := ioutil.ReadAll(formResponse.Body)

			res.Write(body)
			return
		}
		fmt.Println(formResponse.StatusCode)
		http.NotFound(res, req)
	}

}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, os.Getenv("GOPATH")+"/bin/resources/index.html")
}

func main() {

	fmt.Println("starting service")
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)

}
