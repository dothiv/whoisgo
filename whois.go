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
	"io"
	"net"
	"os"
)

func whoisQuery(whoisServer string, domain string) string {
	conn, connErr := net.Dial("tcp", whoisServer+":43")
	if connErr != nil {
		fmt.Fprintln(os.Stderr, connErr.Error())
		os.Exit(1)
	}
	_, writeErr := conn.Write([]byte(domain + "\n"))
	if writeErr != nil {
		fmt.Fprintln(os.Stderr, writeErr.Error())
		os.Exit(2)
	}
	var bytesRead int
	var status [1024]byte
	var readErr error
	var response string
	for {
		bytesRead, readErr = conn.Read(status[0:])
		if readErr != nil {
			if readErr == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, readErr.Error())
			os.Exit(3)
		}
		response += string(status[:bytesRead])
	}
	return response
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s whois-server domainname\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Example: %s whois.afilias.info afilias.info\n", os.Args[0])
}

func main() {
	if len(os.Args) != 3 {
		Usage()
		os.Exit(0)
	}
	fmt.Fprintf(os.Stdout, "Querying %s on %s\n", os.Args[1], os.Args[2])
	fmt.Fprintf(os.Stdout, whoisQuery(os.Args[1], os.Args[2]))
}
