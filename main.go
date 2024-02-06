package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
)

func clear_terminal() {
	// function to clear terminal
	cmd := exec.Command("clear")
	
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	// parse the filename from command line arguments
	var filename string
	
	flag.StringVar(&filename, "f", "", "Specify your output file name")

	flag.Parse()

	if filename == "" {
		args := flag.Args()
		if len(args) > 0 {
			filename = args[0]
		} else {
			fmt.Println("Please provide a file that you want to edit")
			return
		}
	}
	
	// read content from specified filename 	
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("The following error has occurred when attempting to open %s:", filename)
		fmt.Println(err)
		return
	}

	// clearing terminal and printing file content
	clear_terminal()
	fmt.Println(string(content))
	return
}
