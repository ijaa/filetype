package types

// FileType file's mime and extension
type FileType struct {
	MIME      string
	Extension string
	Offset    int64             // the magic numbers content offset
	Limit     int64             // the magic numbers content len
	Checker   func([]byte) bool // magic number checker func
}

// NewFileType create FileType
func NewFileType(ext, mime string, offset, limit int64, checker func([]byte) bool) FileType {
	return FileType{
		MIME:      mime,
		Extension: ext,
		Offset:    offset,
		Limit:     limit,
		Checker:   checker,
	}
}
