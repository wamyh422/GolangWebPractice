package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

type Page struct {
  Title string
  Body []byte
}

// method save and its receiver is p, a pointer to Page
func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

// read content in file and set title as Title
func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{ Title: title, Body: body }, nil
}

// Basic io test
// func main() {
//   p1 := &Page{ Title: "TestPage", Body: []byte("This is a simple Page.") }
//   p1.save()
//   p2, _ := loadPage("TestPage")
//   fmt.Println(string(p2.Body))
// }

// web handler responseWriter and Request
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hi there, I love %s!!</h1>", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  if(p == nil) {
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "No data", "Empty content.")
  } else {
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
  }
}



// Basic web test
func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/view/", viewHandler)
  http.HandlerFunc("/edit/", editHandler)
  http.HandlerFunc("/save/", saveHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
