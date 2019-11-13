package controller

import (
	"html/template"
	"log"
	"net/http"
)

func SendLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	user := Account{
		UserName: r.FormValue("UserName"),
		Password: r.FormValue("Password"),
	}
	
	if r.Method==http.MethodPost{
		tpl,err:=template.ParseFiles("./../views/Login.html")
		if err!=nil{
			log.Fatal(err)
		}
		_=tpl.Execute(w,user)
	}
}

func SendClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	m := Message{
		Message: r.FormValue("UserName"),
	}
	tpl,err:=template.ParseFiles("./../views/client.html")
	if err!=nil{
		log.Fatal(err)
	}
	_=tpl.Execute(w,m)
}
