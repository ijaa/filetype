package types

var (
	TypeJpeg     = NewFileType("jpg", "image/jpeg", 0, 3, Jpeg)
	TypeJpeg2000 = NewFileType("jp2", "image/jp2", 0, 13, Jpeg2000)
	TypePng      = NewFileType("png", "image/png", 0, 4, Png)
	TypeGif      = NewFileType("gif", "image/gif", 0, 3, Gif)
	TypeWebp     = NewFileType("webp", "image/webp", 8, 4, Webp)
	TypeTiff     = NewFileType("tif", "image/tiff", 0, 11, Tiff)
	TypeBmp      = NewFileType("bmp", "image/bmp", 0, 2, Bmp)
	TypePsd      = NewFileType("psd", "image/vnd.adobe.photoshop", 0, 4, Psd)
	TypeIco      = NewFileType("ico", "image/x-icon", 0, 4, Ico)
)

func Jpeg(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0xFF &&
		buf[1] == 0xD8 &&
		buf[2] == 0xFF
}

func Jpeg2000(buf []byte) bool {
	return len(buf) > 12 &&
		buf[0] == 0x0 &&
		buf[1] == 0x0 &&
		buf[2] == 0x0 &&
		buf[3] == 0xC &&
		buf[4] == 0x6A &&
		buf[5] == 0x50 &&
		buf[6] == 0x20 &&
		buf[7] == 0x20 &&
		buf[8] == 0xD &&
		buf[9] == 0xA &&
		buf[10] == 0x87 &&
		buf[11] == 0xA &&
		buf[12] == 0x0
}

func Png(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x89 && buf[1] == 0x50 &&
		buf[2] == 0x4E && buf[3] == 0x47
}

func Gif(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x47 && buf[1] == 0x49 && buf[2] == 0x46
}

func Webp(buf []byte) bool {
	return len(buf) > 11 &&
		buf[8] == 0x57 && buf[9] == 0x45 &&
		buf[10] == 0x42 && buf[11] == 0x50
}

func Tiff(buf []byte) bool {
	return len(buf) > 10 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) || // Little Endian
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A)) && // Big Endian
		!CR2(buf) // To avoid conflicts differentiate Tiff from CR2
}

func Bmp(buf []byte) bool {
	return len(buf) > 1 &&
		buf[0] == 0x42 &&
		buf[1] == 0x4D
}

func Psd(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x38 && buf[1] == 0x42 &&
		buf[2] == 0x50 && buf[3] == 0x53
}

func Ico(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x00 && buf[1] == 0x00 &&
		buf[2] == 0x01 && buf[3] == 0x00
}

func CR2(buf []byte) bool {
	return len(buf) > 10 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) || // Little Endian
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A)) && // Big Endian
		buf[8] == 0x43 && buf[9] == 0x52 && // CR2 magic word
		buf[10] == 0x02 // CR2 major version
}
