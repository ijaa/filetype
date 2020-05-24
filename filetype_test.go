package filetype

import (
	"testing"

	"github.com/ijaa/filetype/types"
)

func TestIs(t *testing.T) {
	ok, err := Is(types.TypeDocx, "./testfiles/sample.docx")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypeDocx.Extension, err)
	}
	ok, err = Is(types.TypeGif, "./testfiles/sample.gif")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypeGif.Extension, err)
	}
	ok, err = Is(types.TypeJpeg, "./testfiles/sample.jpg")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypeJpeg.Extension, err)
	}
	ok, err = Is(types.TypeMp4, "./testfiles/sample.mp4")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypeMp4.Extension, err)
	}
	ok, err = Is(types.TypePng, "./testfiles/sample.png")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypePng.Extension, err)
	}
	ok, err = Is(types.TypePptx, "./testfiles/sample.pptx")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypePptx.Extension, err)
	}
	ok, err = Is(types.TypeTar, "./testfiles/sample.tar")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypeTar.Extension, err)
	}
	ok, err = Is(types.TypeTiff, "./testfiles/sample.tif")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypeTiff.Extension, err)
	}
	ok, err = Is(types.TypeXlsx, "./testfiles/sample.xlsx")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypeXlsx.Extension, err)
	}
	ok, err = Is(types.TypeZip, "./testfiles/sample.zip")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypeZip.Extension, err)
	}
	ok, err = Is(types.TypeMp3, "./testfiles/sample.mp3")
	if !ok {
		t.Fatalf("Invalid match: %s, err:%s", types.TypeMp3.Extension, err)
	}

}
