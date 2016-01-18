package main

import (
    "fmt"
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

func main(){
    
    a := item{"1", 2}
    b := item{"2", 3}
    c := item{"3", 4}
      
    d := Node{a, b, c}
   
    
    left, right := getChildren(d)
    
    fmt.Println(left)
    fmt.Println(right)
}