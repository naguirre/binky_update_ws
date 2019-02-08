package main


import (
  "html/template"
  "log"
  "net/http"
  "os"
  "path/filepath"
  "fmt"
  "io/ioutil"
)

func main() {
  fs := http.FileServer(http.Dir("fw"))
  http.Handle("/fw/", http.StripPrefix("/fw/", fs))
  http.HandleFunc("/latest", serveLatest)

  log.Println("Listening...")
  http.ListenAndServe(":3005", nil)
}

func serveLatest(w http.ResponseWriter, r *http.Request) {
     data, err := ioutil.ReadFile("fw/latest.json")
     if err != nil {
	fmt.Fprint(w, err)
     } else {
	fmt.Fprintf(w, string(data))
     }
}
