package main

import("fmt"
        "time"
)

/*
    This function will print in the console ten times the number 10.
*/
func runFunnyTest(){
    for i:= 0; i<=9; i++{
        go func(){
            fmt.Println(i)
        }()
    }
}

/*
    This function will print in the console the numbers from 0 to 9.
    To fix the funny effect we had to assign i to itself in every
    iteration of the loop.
*/
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
