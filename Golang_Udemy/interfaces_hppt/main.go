package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.index.hu")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := []byte{} //slice
	/*	bs := make([]byte, 99999) //empthy slice with nr of elements

		resp.Body.Read(bs)
		fmt.Println(string(bs)) */

	lg := logWriter{}
	io.Copy(lg, resp.Body)

}

func (lg logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
