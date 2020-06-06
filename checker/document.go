package checker

import (
	"bytes"
	"encoding/binary"
)

type docType int

const (
	TYPE_DOC docType = iota
	TYPE_DOCX
	TYPE_XLS
	TYPE_XLSX
	TYPE_PPT
	TYPE_PPTX
	TYPE_OOXML
)

// Doc small file will return false
//reference: https://bz.apache.org/ooo/show_bug.cgi?id=111457
func Doc(file string) (bool, error) {
	buf, err := getBytes(file, 0, 514)
	if err != nil {
		return false, err
	}
	return len(buf) > 513 &&
		buf[0] == 0xD0 && buf[1] == 0xCF &&
		buf[2] == 0x11 && buf[3] == 0xE0 &&
		buf[512] == 0xEC && buf[513] == 0xA5, nil
}

func Docx(file string) (bool, error) {
	buf, err := getBytes(file, 0, -1)
	if err != nil {
		return false, err
	}
	typ, ok := msooxml(buf)
	return ok && typ == TYPE_DOCX, nil
}

// Xls small file will return false
func Xls(file string) (bool, error) {
	buf, err := getBytes(file, 0, 514)
	if err != nil {
		return false, err
	}
	return len(buf) > 513 &&
		buf[0] == 0xD0 && buf[1] == 0xCF &&
		buf[2] == 0x11 && buf[3] == 0xE0 &&
		buf[512] == 0x09 && buf[513] == 0x08, nil
}

func Xlsx(file string) (bool, error) {
	buf, err := getBytes(file, 0, -1)
	if err != nil {
		return false, err
	}
	typ, ok := msooxml(buf)
	return ok && typ == TYPE_XLSX, nil
}

// Ppt small file will return false
func Ppt(file string) (bool, error) {
	buf, err := getBytes(file, 0, 514)
	if err != nil {
		return false, err
	}
	return len(buf) > 513 &&
		buf[0] == 0xD0 && buf[1] == 0xCF &&
		buf[2] == 0x11 && buf[3] == 0xE0 &&
		buf[512] == 0xA0 && buf[513] == 0x46, nil
}

func Pptx(file string) (bool, error) {
	buf, err := getBytes(file, 0, -1)
	if err != nil {
		return false, err
	}
	typ, ok := msooxml(buf)
	return ok && typ == TYPE_PPTX, nil
}

func msooxml(buf []byte) (typ docType, found bool) {
	signature := []byte{'P', 'K', 0x03, 0x04}

	// start by checking for ZIP local file header signature
	if ok := compareBytes(buf, signature, 0); !ok {
		return
	}

	// make sure the first file is correct
	if v, ok := checkMSOoml(buf, 0x1E); ok {
		return v, ok
	}

	if !compareBytes(buf, []byte("[Content_Types].xml"), 0x1E) &&
		!compareBytes(buf, []byte("_rels/.rels"), 0x1E) &&
		!compareBytes(buf, []byte("docProps"), 0x1E) {
		return
	}

	// skip to the second local file header
	// since some documents include a 520-byte extra field following the file
	// header, we need to scan for the next header
	startOffset := int(binary.LittleEndian.Uint32(buf[18:22]) + 49)
	idx := search(buf, startOffset, 6000)
	if idx == -1 {
		return
	}

	// now skip to the *third* local file header; again, we need to scan due to a
	// 520-byte extra field following the file header
	startOffset += idx + 4 + 26
	idx = search(buf, startOffset, 6000)
	if idx == -1 {
		return
	}

	// and check the subdirectory name to determine which type of OOXML
	// file we have.  Correct the mimetype with the registered ones:
	// http://technet.microsoft.com/en-us/library/cc179224.aspx
	startOffset += idx + 4 + 26
	if typ, ok := checkMSOoml(buf, startOffset); ok {
		return typ, ok
	}

	// OpenOffice/Libreoffice orders ZIP entry differently, so check the 4th file
	startOffset += 26
	idx = search(buf, startOffset, 6000)
	if idx == -1 {
		return TYPE_OOXML, true
	}

	startOffset += idx + 4 + 26
	if typ, ok := checkMSOoml(buf, startOffset); ok {
		return typ, ok
	} else {
		return TYPE_OOXML, true
	}
}

func compareBytes(slice, subSlice []byte, startOffset int) bool {
	sl := len(subSlice)

	if startOffset+sl > len(slice) {
		return false
	}

	s := slice[startOffset : startOffset+sl]
	for i := range s {
		if subSlice[i] != s[i] {
			return false
		}
	}

	return true
}

func checkMSOoml(buf []byte, offset int) (typ docType, ok bool) {
	ok = true

	switch {
	case compareBytes(buf, []byte("word/"), offset):
		typ = TYPE_DOCX
	case compareBytes(buf, []byte("ppt/"), offset):
		typ = TYPE_PPTX
	case compareBytes(buf, []byte("xl/"), offset):
		typ = TYPE_XLSX
	default:
		ok = false
	}

	return
}

func search(buf []byte, start, rangeNum int) int {
	length := len(buf)
	end := start + rangeNum
	signature := []byte{'P', 'K', 0x03, 0x04}

	if end > length {
		end = length
	}

	if start >= end {
		return -1
	}

	return bytes.Index(buf[start:end], signature)
}