package filetype

import (
	"testing"
)

func TestIs(t *testing.T) {
	ok, err := Is(TypeDocx, "./testfiles/sample.docx")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeDocx.Extension, err)
	}
	ok, err = Is(TypeGif, "./testfiles/sample.gif")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeGif.Extension, err)
	}
	ok, err = Is(TypeJpeg, "./testfiles/sample.jpg")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeJpeg.Extension, err)
	}
	ok, err = Is(TypeMp4, "./testfiles/sample.mp4")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeMp4.Extension, err)
	}
	ok, err = Is(TypePng, "./testfiles/sample.png")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypePng.Extension, err)
	}
	ok, err = Is(TypePptx, "./testfiles/sample.pptx")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypePptx.Extension, err)
	}
	ok, err = Is(TypeTar, "./testfiles/sample.tar")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeTar.Extension, err)
	}
	ok, err = Is(TypeTiff, "./testfiles/sample.tif")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeTiff.Extension, err)
	}
	ok, err = Is(TypeXlsx, "./testfiles/sample.xlsx")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeXlsx.Extension, err)
	}
	ok, err = Is(TypeZip, "./testfiles/sample.zip")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeZip.Extension, err)
	}
	ok, err = Is(TypeMp3, "./testfiles/sample.mp3")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeMp3.Extension, err)
	}

	ok, err = IsIn([]FileType{TypeAvi, TypeMp3, TypeMp4}, "./testfiles/sample.mp3")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", TypeMp3.Extension, err)
	}

}
