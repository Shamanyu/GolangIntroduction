package main

import (
  // "fmt"
  "log"
  "io/ioutil"
  "net/http"
  "html/template"
)

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, error := ioutil.ReadFile(filename)
  if error != nil {
    return nil, error
  }
  return &Page{Title: title, Body:body}, nil
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
//   title := r.URL.Path[len("/view/"):]
//   p, _ := loadPage(title)
//   fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  // t, _ := template.ParseFiles("view.html")
  // t.Execute(w, p)
  renderTemplate(w, "view", p)

}

// func editHandler(w http.ResponseWriter, r *http.Request) {
//   title := r.URL.Path[len("/edit"):]
//   p, error := loadPage(title)
//   if error != nil {
//     p = &Page{Title: title}
//   }
//   fmt.Fprintf(w, "<h1>Editing %s</h1>"+
//     "<form action=\"/save/%s\" method=\"POST\">"+
//     "<textarea name=\"body\">%s</textarea><br>"+
//     "<input type=\"submit\" value=\"Save\">"+
//     "</form>",
//     p.Title, p.Title, p.Body)
// }

func editHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit"):]
  p, error := loadPage(title)
  if error != nil {
    p = &Page{Title: title}
  }
  // t, _ := template.ParseFiles("edit.html")
  // t.Execute(w, p)
  renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  t, _ := template.ParseFiles(tmpl + ".html")
  t.Execute(w, p)
}

func main() {
  p1 := Page{Title: "hosted-wiki-2", Body: []byte("This is a sample page created by hosted-wiki-2 program")}
  p1.save()
  // http.HandleFunc("/save/", saveHandler)
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
