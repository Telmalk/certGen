package main

import (
	"fmt"
	"os"
	"training.go/gencert/cert"
	"training.go/gencert/pdf"
)

func main()  {
	c, err := cert.New("GoLang", "Clement Haller", "2020-06-16")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var saver cert.Saver
	saver, err =  pdf.New("output")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = saver.Save(*c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
