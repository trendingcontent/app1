package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
//    "path/filepath"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }

}

func hello(res http.ResponseWriter, req *http.Request) {
    
    b,err := ioutil.ReadFile("go/src/trends/html/index.html")
    if err != nil {
      panic(err)
    }
    fmt.Fprintln(res, string(b))
}
