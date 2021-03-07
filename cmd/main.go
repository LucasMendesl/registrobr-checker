package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lucasmendesl/registrobr-checker/fqdn"
)

func exit(err string) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s <hostname>\n", os.Args[0])
	}

	if len(os.Args) < 2 {
		exit("hostname is required argument")
	}

	var hostname = os.Args[1]

	fqdn, err := fqdn.Search(hostname)

	if err != nil {
		exit(err.Error())
	}

	fmt.Println(fqdn.String())
}
