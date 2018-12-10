/*My solution*/

package main

import (

"fmt"

"io/ioutil"

"os"

)



func main() {



data, err := ioutil.ReadFile(os.Args[1])

if err != nil {

fmt.Println("Can't read file:", os.Args[1], err)



}



fmt.Println(string(data))

/*Teachers solution*/
/*
 import (
 	"fmt"
 	"os"
 	"io"
 )

 func main() {
 	f, err := os.Open(os.Args[1])

 	if err !=nil {
 		fmt.Println("Error", err)
 		os.Exit(1)
 	}

 	io.Copy(os.Stdout, f)
 }
 */