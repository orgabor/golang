/*
Can return multiple values from a function. Usually used to communicate errors

Can ignore a return value by assigning to blank identifier, _.

We can get it to compile by adding blank identifiers.
But this isn't always a good idea. In this case, checking the size of "existent.txt" works, because the file exists. But our users get a panic, a runtime error, when we try to get the size of "nonexistent.txt", because the file doesn't exist.
We can fix this by checking to see if the error value is anything other than nil, and handling the error if it is.


*/



package main

import (
    "fmt"
   "log"
    "math"
    "os"
)

func main() {
   squareRoot, err := squareRoot(29)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(squareRoot)
    
    fileInfo, error := os.Stat("existent.txt")
    if error !=nil {
        fmt.Println(error)
    }else{
         fmt.Println(fileInfo.Size())
    }

   
    fileInfo, error = os.Stat("nonexistent.txt")
        if error !=nil {
        fmt.Println(error)
    }else{
         fmt.Println(fileInfo.Size())
    }
}

func squareRoot(x float64) (float64, error) {
    if x < 0 {
        return 0, fmt.Errorf("can't take square root of negative number")
    }
    return math.Sqrt(x), nil
}
