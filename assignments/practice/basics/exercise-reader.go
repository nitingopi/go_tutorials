package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (myReader MyReader) Read(s []byte) (int, error) {
	s = s[:cap(s)]
	for i := range s {
		s[i] = 'A'
	}
	return cap(s), nil
	// return 0, nil
}

func main() {
	reader.Validate(MyReader{})
}
