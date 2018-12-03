package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr," PHP Code Optimize\n\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	/*
	arguments := flag.Args()
	fmt.Println(arguments)
	*/
}
