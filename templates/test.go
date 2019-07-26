package main

import (
   "net/http"
   "fmt"
   "time"
   "html/template"
)

type Welcome struct {
   Name string
   Time string
}

type Post struct {
  Title string
  Text string
}

func home_page(w http.ResponseWriter, r *http.Request) {
  welcome := Welcome{"Igor", time.Now().Format(time.Stamp)}
  if name := r.FormValue("name"); name != "" {
     welcome.Name = name;
  }

  templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

  if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
     http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func create_post(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
      templates := template.Must(template.ParseFiles("templates/posts/new.html"))

      if err := templates.ExecuteTemplate(w, "new.html", nil); err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
      }
    } else {
      r.ParseForm()
      fmt.Println("title:", r.Form["title"])
      fmt.Println("text:", r.Form["text"])
    }
}

func main() {
   http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
   http.HandleFunc("/" , home_page)
   http.HandleFunc("/posts/new" , create_post)

   fmt.Println("Listening");
   fmt.Println(http.ListenAndServe(":8080", nil));
}
