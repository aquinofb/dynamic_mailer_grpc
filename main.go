package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	pb "github.com/aquinofb/dynamicmailer/lib"
	"github.com/aquinofb/dynamicmailer/services"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/application.gtpl")
	t.Execute(w, nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/not-found.gtpl")
	t.Execute(w, nil)
}

func dynamicMailerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/dynamic-mailer/index.gtpl")
		t.Execute(w, nil)
	} else {
		status := services.SendEmail(
			pb.Payload{
				Event:          r.FormValue("message[event]"),
				ReferenceUuid: 	r.FormValue("message[reference_uuid]"),
				Email:          r.FormValue("message[email]")},
		)

		if status == 200 {
			http.Redirect(w, r, "/dynamic-mailer", http.StatusFound)
		} else {
			http.Redirect(w, r, "/not-found", http.StatusFound)
		}
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dynamic-mailer", dynamicMailerHandler)
	http.HandleFunc("/not-found", notFoundHandler)
	fmt.Println(".")
	fmt.Println("Server running smoothly on 9090... ( ಠ_ರೃ) ٩(⁎❛ᴗ❛⁎)۶")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
