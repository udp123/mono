package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "welcome!\n")
}

func Hello(w http.ResponseWriter, r * http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func ExParams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hey there %s! what is: %s\n", ps.ByName("named"), ps.ByName("catchall"))
}

func CatchInput(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "content in: %s", ps.ByName("catchall"))
}


func main () {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
    router.GET("/dev/:named/*catchall", ExParams)
    router.GET("/catch/*catchall", CatchInput)
    
    log.Fatal(http.ListenAndServe(":6677", router))

    //http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    //})
}
