package main

import (
    "fmt"
    "os/exec"
)

func split(x int) (a, b int){
    a := 1
    b := 2
    return
}

func execute() {
    out, err := exec.Command("pwd").Output()
    
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Command Successfully Executed")
    fmt.Println(out)
}


func main(){
    
    fmt.Println("Simple Shell")
    fmt.Println("---------------------")
    
    execute()    
    fmt.Println(split(5))
}