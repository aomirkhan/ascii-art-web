package main

import (
	"fmt"
	"git/rzhampeis/ascii-art-web/pkg"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/background.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "https://wallpaperaccess.com/download/1mb-626330")
	})

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/ascii-art", formHandler)
	log.Println("Server start on http://127.0.0.1:8000")
	log.Println("OK(200)")
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmp, err := template.ParseFiles("templates/form.html")
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	} else {
		tmp.Execute(w, nil)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}
	name := r.FormValue("convert")
	for _, i := range name {
		if (i < 32 || i > 126) && i != 10 && i != 13 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			return
		}
	}
	banner := r.FormValue("fonts")
	if banner == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
	text := pkg.Converter(banner, name)

	tmp, err := template.ParseFiles("templates/form.html")
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	} else {
		tmp.Execute(w, text)
	}
}
