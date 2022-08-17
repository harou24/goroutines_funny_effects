package main

import("fmt"
        "time"
)

func runFunnyTest(){
    for i:= 0; i<=9; i++{
        go func(){
            fmt.Println(i)
        }()
    }
}

func runExpectedTest(){
    for i:= 0; i<=9; i++{
        i := i
        go func(){
            fmt.Println(i)
        }()
    }
}

func main(){
    fmt.Println("Running funny test...")
    runFunnyTest()
    time.Sleep(time.Second)
    fmt.Println("Running expected test...")
    runExpectedTest()
    time.Sleep(time.Second)
}
