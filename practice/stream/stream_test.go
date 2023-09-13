package stream

import (
	"testing"
	"time"
)

func TestStream(t *testing.T) {
	ready := make(chan bool)
	go func() {
		// Do something
		<-ready
		time.Sleep(1 * time.Second)
		sendFile(1000)
	}()

	server := &FileServer{}
	go server.start()
	time.Sleep(2 * time.Second)
	ready <- true

	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "TestStream",
			filename: "stream.go",
			want:     "Hello",
		},
		{
			name:     "TestStreamLargeFile",
			filename: "stream_large_file.go",
			want:     "Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Do something
		})
	}
}
