/**
 * Tests for whoisgo.go
 *
 * @author Markus Tacker <m@dothiv.org>
 */
package whoisgo

import (
	"io"
	"testing"
)

type MockConnection struct {
	resp      []byte
	recv      []byte
	bytesRead int
}

func (e *MockConnection) Write(b []byte) (bytesWritten int, err error) {
	bytesWritten = len(b)
	e.recv = b
	return
}

func (e *MockConnection) Read(b []byte) (bytesRead int, err error) {
	if e.bytesRead > 0 {
		err = io.EOF
		return
	}
	for i := range e.resp {
		b[i] = e.resp[i]
	}
	bytesRead = len(e.resp)
	e.bytesRead += bytesRead
	return
}

func TestHelp(t *testing.T) {
	expResponse := "WHOISDATA"
	conn := new(MockConnection)
	conn.resp = []byte(expResponse)
	response, err := WhoisQuery(conn, "afilias.info")
	if string(conn.recv) != "afilias.info\n" {
		t.Errorf("Query should be '%s', got '%s'", "afilias.info\n", conn.recv)
	}
	if err != nil {
		t.Errorf("Error should be nil, got %v", err)
	}
	if response != expResponse {
		t.Errorf("Response should be '%v', got '%v'", expResponse, response)
	}
}
