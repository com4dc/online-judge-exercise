package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, e error) {
	for n, e = 0, nil; n < len(b); n++ {
		b[n] = 'A'
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}
