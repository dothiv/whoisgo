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
	"time"
)

func whoisQuery(whoisServer string, domain string) (response string, err error) {
	// Connect
	conn, connErr := net.DialTimeout("tcp", whoisServer+":43", time.Duration(3)*time.Second)
	if connErr != nil {
		err = connErr
		return
	}
	defer conn.Close()

	// Send query
	_, writeErr := conn.Write([]byte(domain + "\n"))
	if writeErr != nil {
		err = writeErr
		return
	}

	// Read response
	var bytesRead int
	var status [1024]byte
	var readErr error
	for {
		bytesRead, readErr = conn.Read(status[0:])
		if readErr != nil {
			if readErr == io.EOF {
				break
			}
			err = readErr
			return
		}
		response += string(status[:bytesRead])
	}
	return
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
	response, queryError := whoisQuery(os.Args[1], os.Args[2])
	if queryError != nil {
		fmt.Fprintln(os.Stderr, queryError.Error())
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, response)
}
