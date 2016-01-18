package main

import (
    "fmt"
    "bufio"
    "os"
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
    
<<<<<<< HEAD
    a := item{"1", 2}
    b := item{"2", 3}
    c := item{"3", 4}
      
    d := Node{a, b, c}
   
    
    left, right := getChildren(d)
    
    fmt.Println(left)
    fmt.Println(right)
=======
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Simple Shell")
    fmt.Println("---------------------")
    
    for {
        fmt.Print("-> ")
        text, _ := reader.ReadString('\n')        
        fmt.Println(text)
    }
    
>>>>>>> 388c74c83e9ee93e7a0befe6775010fe907cfa59
}