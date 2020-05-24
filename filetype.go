package filetype

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/ijaa/filetype/types"
)

// global errors
var (
	ErrFileContent = errors.New("get file content error")
)

// getBytes get the file n byte from the offset
func getBytes(file string, offset, limit int64) ([]byte, error) {
	fi, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	if limit > 0 {
		buf := make([]byte, limit)
		if count, err := fi.ReadAt(buf, offset); err == nil {
			if int64(count) != limit {
				return nil, ErrFileContent
			}
		}

		return buf, nil
	}
	// limit -1 return the full file content
	return ioutil.ReadFile(file)
}

// Is check the file's type equals the custom FileType
func Is(ft types.FileType, file string) (bool, error) {
	buf, err := getBytes(file, ft.Offset, ft.Limit)
	if err != nil {
		return false, err
	}
	return ft.Checker(buf), nil
}
