package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://orgabor.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := []byte{} //slice
	/*	bs := make([]byte, 99999) //empthy slice with nr of elements

		resp.Body.Read(bs)
		fmt.Println(string(bs)) */
	io.Copy(os.Stdout, resp.Body)

}
