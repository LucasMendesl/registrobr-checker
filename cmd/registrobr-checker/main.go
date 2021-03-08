package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/lucasmendesl/registrobr-checker/fqdn"
)

func exit(err string) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func isValidHost(host string) bool {
	var rx = regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
	return rx.MatchString(host)
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	var hostname string

	flag.StringVar(&hostname, "hostname", "", "example.com")
	flag.Parse()

	if hostname == "" {
		exit("hostname is required parameter, please use -help command to verify hostname parameter usage")
	}

	if !isValidHost(hostname) {
		exit("invalid hostname, please enter with a valid hostname")
	}

	fqdn, err := fqdn.Search(hostname)

	if err != nil {
		exit(err.Error())
	}

	fmt.Println(fqdn.String())
}
