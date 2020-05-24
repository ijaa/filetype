package types

// archive related filetype
var (
	TypeZip  = NewFileType("zip", "application/zip", 0, 4, Zip)
	TypeTar  = NewFileType("tar", "application/x-tar", 0, 262, Tar)
	TypeRar  = NewFileType("rar", "application/x-rar-compressed", 0, 7, Rar)
	TypeGz   = NewFileType("gz", "application/gzip", 0, 3, Gz)
	TypeBz2  = NewFileType("bz2", "application/x-bzip2", 0, 3, Bz2)
	Type7z   = NewFileType("7z", "application/x-7z-compressed", 0, 6, SevenZ)
	TypePdf  = NewFileType("pdf", "application/pdf", 0, 4, Pdf)
	TypeExe  = NewFileType("exe", "application/x-msdownload", 0, 2, Exe)
	TypeSwf  = NewFileType("swf", "application/x-shockwave-flash", 0, 3, Swf)
	TypeRpm  = NewFileType("rpm", "application/x-rpm", 0, 97, Rpm)
	TypeElf  = NewFileType("elf", "application/x-executable", 0, 53, Elf)
	TypeIso  = NewFileType("iso", "application/x-iso9660-image", 0, 32774, Iso)
	TypeWasm = NewFileType("wasm", "application/wasm", 0, 8, Wasm)
)

func Zip(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x50 && buf[1] == 0x4B &&
		(buf[2] == 0x3 || buf[2] == 0x5 || buf[2] == 0x7) &&
		(buf[3] == 0x4 || buf[3] == 0x6 || buf[3] == 0x8)
}

func Tar(buf []byte) bool {
	return len(buf) > 261 &&
		buf[257] == 0x75 && buf[258] == 0x73 &&
		buf[259] == 0x74 && buf[260] == 0x61 &&
		buf[261] == 0x72
}

func Rar(buf []byte) bool {
	return len(buf) > 6 &&
		buf[0] == 0x52 && buf[1] == 0x61 && buf[2] == 0x72 &&
		buf[3] == 0x21 && buf[4] == 0x1A && buf[5] == 0x7 &&
		(buf[6] == 0x0 || buf[6] == 0x1)
}

func Gz(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x1F && buf[1] == 0x8B && buf[2] == 0x8
}

func Bz2(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x42 && buf[1] == 0x5A && buf[2] == 0x68
}

func SevenZ(buf []byte) bool {
	return len(buf) > 5 &&
		buf[0] == 0x37 && buf[1] == 0x7A && buf[2] == 0xBC &&
		buf[3] == 0xAF && buf[4] == 0x27 && buf[5] == 0x1C
}

func Pdf(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x25 && buf[1] == 0x50 &&
		buf[2] == 0x44 && buf[3] == 0x46
}

func Exe(buf []byte) bool {
	return len(buf) > 1 &&
		buf[0] == 0x4D && buf[1] == 0x5A
}

func Swf(buf []byte) bool {
	return len(buf) > 2 &&
		(buf[0] == 0x43 || buf[0] == 0x46) &&
		buf[1] == 0x57 && buf[2] == 0x53
}

func Rpm(buf []byte) bool {
	return len(buf) > 96 &&
		buf[0] == 0xED && buf[1] == 0xAB &&
		buf[2] == 0xEE && buf[3] == 0xDB
}

func Elf(buf []byte) bool {
	return len(buf) > 52 &&
		buf[0] == 0x7F && buf[1] == 0x45 &&
		buf[2] == 0x4C && buf[3] == 0x46
}

func Iso(buf []byte) bool {
	return len(buf) > 32773 &&
		buf[32769] == 0x43 && buf[32770] == 0x44 &&
		buf[32771] == 0x30 && buf[32772] == 0x30 &&
		buf[32773] == 0x31
}

// Wasm detects a Web Assembly 1.0 filetype.
func Wasm(buf []byte) bool {
	// WASM has starts with `\0asm`, followed by the version.
	// http://webassembly.github.io/spec/core/binary/modules.html#binary-magic
	return len(buf) >= 8 &&
		buf[0] == 0x00 && buf[1] == 0x61 &&
		buf[2] == 0x73 && buf[3] == 0x6D &&
		buf[4] == 0x01 && buf[5] == 0x00 &&
		buf[6] == 0x00 && buf[7] == 0x00
}
