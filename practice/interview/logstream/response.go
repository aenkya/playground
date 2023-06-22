package logstream

import (
	"fmt"
	"strings"
)

const (
	ACK   MessageType = "ACK"
	QUERY MessageType = "Q"
	LOG   MessageType = "L"
	MATCH MessageType = "M"
)

type Response struct {
	messageType MessageType
	message     string
	targets     []int
}

type MessageType string

func (r *Response) GetMessageType() MessageType {
	return r.messageType
}

func (r *Response) GetMessage() string {
	return r.message
}

func (r *Response) GetTargets() []int {
	return r.targets
}

func (r *Response) String() string {
	s := string(r.messageType) + ": " + r.message
	if len(r.targets) > 0 {
		s += ", " + strings.Trim(
			strings.Join(
				strings.Fields(fmt.Sprint(r.targets)),
				", ",
			),
			"[]",
		)
	}

	return s
}
