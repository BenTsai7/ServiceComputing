package main

import (
    "io/ioutil"
	"os"
	"fmt"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
		fmt.Fprintf(os.Stdout,"pipe error")
    } else {
	   fmt.Fprintf(os.Stdout,"Printer: Receive data: %s",data)
	}
}