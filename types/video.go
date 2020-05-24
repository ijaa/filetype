package types

var (
	TypeMp4  = NewFileType("mp4", "video/mp4", 0, 12, Mp4)
	TypeAvi  = NewFileType("avi", "video/x-msvideo", 0, 11, Avi)
	TypeWmv  = NewFileType("wmv", "video/x-ms-wmv", 0, 10, Wmv)
	TypeMpeg = NewFileType("mpg", "video/mpeg", 0, 4, Mpeg)
	TypeFlv  = NewFileType("flv", "video/x-flv", 0, 4, Flv)
	Type3gp  = NewFileType("3gp", "video/3gpp", 0, 11, Match3gp)
)

func Mp4(buf []byte) bool {
	return len(buf) > 11 &&
		(buf[4] == 'f' && buf[5] == 't' && buf[6] == 'y' && buf[7] == 'p') &&
		((buf[8] == 'a' && buf[9] == 'v' && buf[10] == 'c' && buf[11] == '1') ||
			(buf[8] == 'd' && buf[9] == 'a' && buf[10] == 's' && buf[11] == 'h') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '2') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '3') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '4') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '5') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == '6') ||
			(buf[8] == 'i' && buf[9] == 's' && buf[10] == 'o' && buf[11] == 'm') ||
			(buf[8] == 'm' && buf[9] == 'm' && buf[10] == 'p' && buf[11] == '4') ||
			(buf[8] == 'm' && buf[9] == 'p' && buf[10] == '4' && buf[11] == '1') ||
			(buf[8] == 'm' && buf[9] == 'p' && buf[10] == '4' && buf[11] == '2') ||
			(buf[8] == 'm' && buf[9] == 'p' && buf[10] == '4' && buf[11] == 'v') ||
			(buf[8] == 'm' && buf[9] == 'p' && buf[10] == '7' && buf[11] == '1') ||
			(buf[8] == 'M' && buf[9] == 'S' && buf[10] == 'N' && buf[11] == 'V') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'A' && buf[11] == 'S') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'C') ||
			(buf[8] == 'N' && buf[9] == 'S' && buf[10] == 'D' && buf[11] == 'C') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'H') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'M') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'P') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'S' && buf[11] == 'S') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'C') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'H') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'M') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'P') ||
			(buf[8] == 'N' && buf[9] == 'D' && buf[10] == 'X' && buf[11] == 'S') ||
			(buf[8] == 'F' && buf[9] == '4' && buf[10] == 'V' && buf[11] == ' ') ||
			(buf[8] == 'F' && buf[9] == '4' && buf[10] == 'P' && buf[11] == ' '))
}

func Avi(buf []byte) bool {
	return len(buf) > 10 &&
		buf[0] == 0x52 && buf[1] == 0x49 &&
		buf[2] == 0x46 && buf[3] == 0x46 &&
		buf[8] == 0x41 && buf[9] == 0x56 &&
		buf[10] == 0x49
}

func Wmv(buf []byte) bool {
	return len(buf) > 9 &&
		buf[0] == 0x30 && buf[1] == 0x26 &&
		buf[2] == 0xB2 && buf[3] == 0x75 &&
		buf[4] == 0x8E && buf[5] == 0x66 &&
		buf[6] == 0xCF && buf[7] == 0x11 &&
		buf[8] == 0xA6 && buf[9] == 0xD9
}

func Mpeg(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x0 && buf[1] == 0x0 &&
		buf[2] == 0x1 && buf[3] >= 0xb0 &&
		buf[3] <= 0xbf
}

func Flv(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x46 && buf[1] == 0x4C &&
		buf[2] == 0x56 && buf[3] == 0x01
}

func Match3gp(buf []byte) bool {
	return len(buf) > 10 &&
		buf[4] == 0x66 && buf[5] == 0x74 && buf[6] == 0x79 &&
		buf[7] == 0x70 && buf[8] == 0x33 && buf[9] == 0x67 &&
		buf[10] == 0x70
}
