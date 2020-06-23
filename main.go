package main

import (
	"flag"
	"fmt"
	"os"
	"training.go/gencert/cert"
	"training.go/gencert/html"
	"training.go/gencert/pdf"
)


func main()  {
	file := flag.String("file", "student.csv", "CSV file input")
	if len(*file) <= 0 {
		fmt.Printf("Invalid file. got=%v", *file)
		os.Exit(1)
	}
	outputType := flag.String("type", "pdf", "Output type of the certificate")
	flag.Parse()
	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err =  html.New("output")
	case "pdf":
		saver, err =  pdf.New("output")
	default:
		fmt.Printf("Unknow output type. got'%v'\n", *outputType)
	}
	if err != nil {
		fmt.Printf("Could not create generator: %v", err)
		os.Exit(1)
	}
	// c, err := cert.New("GoLang", "Boby", "2020-06-16")
	certs, err := cert.ParseCSV(*file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Println(err)
		}
	}
}
