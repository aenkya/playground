package syntax

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type MyReader struct{}

func (r *MyReader) Read(p []byte) (int, error) {
	// Fill the byte slice with 'A' until it is full
	for i := 0; ; i++ {
		p[i%len(p)] = 'A'
		return len(p), nil
	}
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	for k, v := range p[:n] {
		p[k] = rot13(v)
	}

	return n, err
}

func rot13(c byte) byte {
	const rot = 13

	if c >= 'a' && c <= 'z' {
		if c+rot <= 'z' {
			return c + rot
		}

		return c - rot
	} else if c >= 'A' && c <= 'Z' {
		if c+rot <= 'Z' {
			return c + rot
		}

		return c - rot
	}

	return c
}

func Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	if _, err := io.Copy(os.Stdout, &r); err != nil {
		fmt.Print(err)
	}
}
