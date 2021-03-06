package main

import (
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
)

type Page struct {
  Title string
  Body []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

// func loadPage(title string) *Page {
//   filename := title + ".txt"
//   body, _ := ioutil.ReadFile(filename)
//   return &Page{Title: title, Body:body}
// }

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, error := ioutil.ReadFile(filename)
  if error != nil {
    return nil, error
  }
  return &Page{Title: title, Body:body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// func main() {
//   p1 := Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
//   p1.save()
//   p2, _ := loadPage("TestPage")
//   fmt.Println(string(p2.Body))
// }

func main() {
  p1 := Page{Title: "test-page-1", Body: []byte("This is a sample Page.")}
  p1.save()
  http.HandleFunc("/view/", viewHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
