package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

type item struct {
    node string
    value int
}

type Node struct {
    head item
    left item
    right item
}

func split(x int) (a, b, c int){
    a = 1
    b = 2
    c = 3
    return
}

func getChildren(a Node) (b, c item) {
    b = a.left
    c = a.right
    return
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}