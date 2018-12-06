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
	flag.Usage = usage
	InputFile := flag.String("input","","path of a file as input")
	OutputFile := flag.String("output","","path of a file as output")
	MinifyFile := flag.Bool("minify",true,"")
	//OptimizeFile := flag.Bool("optimize",true,"")
	flag.Parse()
	/////////////////////////////////////////////////////////////////
	if *InputFile == "" {
		fmt.Println("Arguments is missing.\n");
		usage();
		os.Exit(1);
	}
	/////////////////////////////////////////////////////////////////
	//fmt.Println(*InputFile)
	//fmt.Println(*OutputFile)
	//fmt.Println(*MinifyFile)
	//fmt.Println(*OptimizeFile)
}
