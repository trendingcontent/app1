package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    /*
    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
    */
    s := &http.Server{
        Addr:           ":"+os.Getenv("PORT"),
        Handler:        myHandler,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())

}

func myHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hello, world")
}


