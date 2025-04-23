package bufferedwriter

import (
	"bytes"
	"testing"
)

type mockFile struct {
	writtenData bytes.Buffer
}

func (mf *mockFile) write(data []byte) {
	mf.writtenData.Write(data)
}

func TestBufferedWriter_WriteAndFlush(t *testing.T) {
	mock := &mockFile{}
	bufferSize := 5
	bw := &bufferedWriter{
		file:   mock,
		buffer: make([]byte, bufferSize),
		index:  0,
	}

	input := []byte("hello world")
	bw.Write(input)
	bw.Flush()

	// Check if the buffer is flushed correctly
	expectedFirstFlush := "hello"
	if mock.writtenData.String()[:len(expectedFirstFlush)] != expectedFirstFlush {
		t.Errorf("expected first flush to write '%s', got '%s'", expectedFirstFlush, mock.writtenData.String()[:len(expectedFirstFlush)])
	}

	// Check if remaining data is in the buffer
	expectedRemaining := " world"
	bw.Flush()
	if mock.writtenData.String()[len(expectedFirstFlush):] != expectedRemaining {
		t.Errorf("expected second flush to write '%s', got '%s'", expectedRemaining, mock.writtenData.String()[len(expectedFirstFlush):])
	}
}

func TestBufferedWriter_Flush(t *testing.T) {
	mock := &mockFile{}
	bufferSize := 5
	bw := &bufferedWriter{
		file:   mock,
		buffer: make([]byte, bufferSize),
		index:  0,
	}

	input := []byte("test")
	bw.Write(input)
	bw.Flush()

	// Check if the buffer is flushed correctly
	expected := "test"
	if mock.writtenData.String() != expected {
		t.Errorf("expected flush to write '%s', got '%s'", expected, mock.writtenData.String())
	}

	// Ensure index is reset
	if bw.index != 0 {
		t.Errorf("expected index to be reset to 0, got %d", bw.index)
	}
}
