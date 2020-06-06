package filetype

import "github.com/ijaa/filetype/checker"

// FileType file's mime and extension
type FileType struct {
	MIME      string
	Extension string
	Offset    int64                      // the magic numbers content offset
	Limit     int64                      // the magic numbers content len
	Checker   func(string) (bool, error) // magic number checker func
}

// newFileType create FileType
func newFileType(ext, mime string, checker func(string) (bool, error)) FileType {
	return FileType{
		MIME:      mime,
		Extension: ext,
		Checker:   checker,
	}
}

// archive related filetype
var (
	TypeZip  = newFileType("zip", "application/zip", checker.Zip)
	TypeTar  = newFileType("tar", "application/x-tar", checker.Tar)
	TypeRar  = newFileType("rar", "application/x-rar-compressed", checker.Rar)
	TypeGz   = newFileType("gz", "application/gzip", checker.Gz)
	TypeBz2  = newFileType("bz2", "application/x-bzip2", checker.Bz2)
	Type7z   = newFileType("7z", "application/x-7z-compressed", checker.SevenZ)
	TypePdf  = newFileType("pdf", "application/pdf", checker.Pdf)
	TypeExe  = newFileType("exe", "application/x-msdownload", checker.Exe)
	TypeSwf  = newFileType("swf", "application/x-shockwave-flash", checker.Swf)
	TypeRpm  = newFileType("rpm", "application/x-rpm", checker.Rpm)
	TypeElf  = newFileType("elf", "application/x-executable", checker.Elf)
	TypeIso  = newFileType("iso", "application/x-iso9660-image", checker.Iso)
	TypeWasm = newFileType("wasm", "application/wasm", checker.Wasm)
)

var (
	TypeMp3 = newFileType("mp3", "audio/mpeg", checker.Mp3)
	TypeM4a = newFileType("m4a", "audio/m4a", checker.M4a)
	TypeWav = newFileType("wav", "audio/x-wav", checker.Wav)
)

var (
	TypeDoc  = newFileType("doc", "application/msword", checker.Doc)
	TypeDocx = newFileType("docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document", checker.Docx)
	TypeXls  = newFileType("xls", "application/vnd.ms-excel", checker.Xls)
	TypeXlsx = newFileType("xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", checker.Xlsx)
	TypePpt  = newFileType("ppt", "application/vnd.ms-powerpoint", checker.Ppt)
	TypePptx = newFileType("pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation", checker.Pptx)
)

var (
	TypeJpeg     = newFileType("jpg", "image/jpeg", checker.Jpeg)
	TypeJpeg2000 = newFileType("jp2", "image/jp2", checker.Jpeg2000)
	TypePng      = newFileType("png", "image/png", checker.Png)
	TypeGif      = newFileType("gif", "image/gif", checker.Gif)
	TypeWebp     = newFileType("webp", "image/webp", checker.Webp)
	TypeTiff     = newFileType("tif", "image/tiff", checker.Tiff)
	TypeBmp      = newFileType("bmp", "image/bmp", checker.Bmp)
	TypePsd      = newFileType("psd", "image/vnd.adobe.photoshop", checker.Psd)
	TypeIco      = newFileType("ico", "image/x-icon", checker.Ico)
)

var (
	TypeMp4  = newFileType("mp4", "video/mp4", checker.Mp4)
	TypeAvi  = newFileType("avi", "video/x-msvideo", checker.Avi)
	TypeWmv  = newFileType("wmv", "video/x-ms-wmv", checker.Wmv)
	TypeMpeg = newFileType("mpg", "video/mpeg", checker.Mpeg)
	TypeFlv  = newFileType("flv", "video/x-flv", checker.Flv)
	Type3gp  = newFileType("3gp", "video/3gpp", checker.Match3gp)
)
