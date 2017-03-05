package main

import "os"
import "fmt"
import "time"
import "strconv"
import "github.com/xenolf/lego/acme"

func main() {
	fqdn := os.Args[1]
	value := os.Args[2]

	parsedTimeout, err := strconv.Atoi(os.Args[3])
	timeout := time.Duration(parsedTimeout) * time.Second
	interval := 2 * time.Second

	err = acme.WaitFor(timeout, interval, func() (bool, error) {
		return acme.PreCheckDNS(fqdn, value)
	})

	if err != nil {
		fmt.Println("record not found")
		fmt.Println("got error: %v", err)
		os.Exit(1)

	} else {
		fmt.Println("record found")
		os.Exit(0)
	}
}
