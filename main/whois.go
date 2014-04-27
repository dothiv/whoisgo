/**
 * Query WHOIS servers
 *
 * Usage: whois-go whois-server domainname
 *
 * @author Markus Tacker <m@dothiv.org>
 */
package main

import (
	"fmt"
	"github.com/dothiv/whoisgo"
	"os"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s whois-server domainname\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Example: %s whois.afilias.info afilias.info\n", os.Args[0])
}

func main() {
	if len(os.Args) != 3 {
		Usage()
		os.Exit(0)
	}
	// Connect
	fmt.Fprintf(os.Stdout, "Connecting to %s\n", os.Args[1])
	conn, connErr := whoisgo.WhoisConnect(os.Args[1])
	if connErr != nil {
		fmt.Fprintln(os.Stderr, connErr.Error())
		os.Exit(1)
	}
	defer conn.Close()
	// Query
	fmt.Fprintf(os.Stdout, "Querying %s on %s\n", os.Args[1], os.Args[2])
	response, queryError := whoisgo.WhoisQuery(conn, os.Args[2])
	if queryError != nil {
		fmt.Fprintln(os.Stderr, queryError.Error())
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, response)
}
