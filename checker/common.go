package checker

import (
	"io"
	"io/ioutil"
	"os"
)

// getBytes get the len bytes from the offset, if get bytes less than length, return error
// len<=0: get all bytes
func getBytes(file string, offset uint64, length int64) ([]byte, error) {
	fi, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	if length > 0 {
		buf := make([]byte, length)
		_, err := fi.ReadAt(buf, int64(offset))
		if err != nil && err != io.EOF {
			return nil, err
		}
		return buf, nil
	}

	return ioutil.ReadFile(file)
}
