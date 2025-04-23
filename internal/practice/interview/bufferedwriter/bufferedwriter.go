package bufferedwriter

type BufferedWriter interface {
	Write()
	Flush()
}

type File interface {
	write([]byte)
}

type bufferedWriter struct {
	file   File
	buffer []byte
	index  int
}

func (bw *bufferedWriter) Write(input []byte) {
	for _, b := range input {
		if bw.index >= len(bw.buffer)-1 {
			bw.Flush()
		}

		bw.buffer[bw.index] = b
		bw.index++
	}
}

func (bw *bufferedWriter) Flush() {
	bw.file.write(bw.buffer[:bw.index])
	bw.index = 0
}
