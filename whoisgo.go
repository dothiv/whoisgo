/**
 * Query WHOIS servers
 *
 * @author Markus Tacker <m@dothiv.org>
 */
package whoisgo

import (
	"errors"
	"io"
	"net"
)

func WhoisQuery(conn io.ReadWriter, domain string) (response string, err error) {
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
		if bytesRead == 0 {
			err = errors.New("Zero-byte length response received.")
			return
		}
		response += string(status[:bytesRead])
	}
	return
}

func WhoisConnect(server string) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", server+":43")
	return
}
