package main

import("fmt"
        "time"
)

func main(){
    go func(){
        fmt.Println("Helllo")
    }()
    for i:= 0; i<=9; i++{
        go func(){
            fmt.Println(i)
        }()
    }
    time.Sleep(time.Second)
}
