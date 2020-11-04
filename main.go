package main

import (
	"io/ioutil"
	"encoding/base64"
	"text/template"
	"net/http"
	"fmt"
	"log"
)

func returnAssertion(w http.ResponseWriter, req *http.Request) {
        rawData, _ := ioutil.ReadFile("testdata/assertion.xml")
	b64Data := base64.StdEncoding.EncodeToString([]byte(rawData))
	tpl := template.Must(template.ParseFiles("testdata/samlResponse.tmpl"))
        tpl.Execute(w, b64Data)
}
func returnLoginPage(w http.ResponseWriter, req *http.Request) {
	log.Printf("BB %v", req)
        data, _ := ioutil.ReadFile("testdata/loginpage.html")
        w.Write(data)
}
func returnSamlPage(w http.ResponseWriter, req *http.Request) {
	log.Printf("BB %v", req)
	fmt.Sprintf("Request %s",req)
        data, _ := ioutil.ReadFile("testdata/saml.html")
        w.Write(data)
}

func main() {
	http.HandleFunc("/adfs/ls/idpinitiatedsignon", returnAssertion)
	http.HandleFunc("/adfs/", returnLoginPage)
	http.HandleFunc("/saml/", returnSamlPage)
	http.ListenAndServe(":3000",nil)
}
