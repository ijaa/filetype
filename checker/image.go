package checker

func Jpeg(file string) (bool, error) {
	buf, err := getBytes(file, 0, 3)
	if err != nil {
		return false, err
	}
	return len(buf) > 2 && buf[0] == 0xFF &&
		buf[1] == 0xD8 &&
		buf[2] == 0xFF, nil
}

func Jpeg2000(file string) (bool, error) {
	buf, err := getBytes(file, 0, 13)
	if err != nil {
		return false, err
	}
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
		buf[12] == 0x0, nil
}

func Png(file string) (bool, error) {
	buf, err := getBytes(file, 0, 4)
	if err != nil {
		return false, err
	}
	return len(buf) > 3 &&
		buf[0] == 0x89 && buf[1] == 0x50 &&
		buf[2] == 0x4E && buf[3] == 0x47, nil
}

func Gif(file string) (bool, error) {
	buf, err := getBytes(file, 0, 3)
	if err != nil {
		return false, err
	}
	return len(buf) > 2 &&
		buf[0] == 0x47 && buf[1] == 0x49 && buf[2] == 0x46, nil
}

func Webp(file string) (bool, error) {
	buf, err := getBytes(file, 0, 12)
	if err != nil {
		return false, err
	}
	return len(buf) > 11 &&
		buf[8] == 0x57 && buf[9] == 0x45 &&
		buf[10] == 0x42 && buf[11] == 0x50, nil
}

func Tiff(file string) (bool, error) {
	buf, err := getBytes(file, 0, 11)
	if err != nil {
		return false, err
	}
	return len(buf) > 10 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) || // Little Endian
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A)) && // Big Endian
		!CR2(buf), nil // To avoid conflicts differentiate Tiff from CR2
}

func Bmp(file string) (bool, error) {
	buf, err := getBytes(file, 0, 2)
	if err != nil {
		return false, err
	}
	return len(buf) > 1 &&
		buf[0] == 0x42 &&
		buf[1] == 0x4D, nil
}

func Psd(file string) (bool, error) {
	buf, err := getBytes(file, 0, 4)
	if err != nil {
		return false, err
	}
	return len(buf) > 3 &&
		buf[0] == 0x38 && buf[1] == 0x42 &&
		buf[2] == 0x50 && buf[3] == 0x53, nil
}

func Ico(file string) (bool, error) {
	buf, err := getBytes(file, 0, 4)
	if err != nil {
		return false, err
	}
	return len(buf) > 3 &&
		buf[0] == 0x00 && buf[1] == 0x00 &&
		buf[2] == 0x01 && buf[3] == 0x00, nil
}

func CR2(buf []byte) bool {
	return len(buf) > 10 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) || // Little Endian
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A)) && // Big Endian
		buf[8] == 0x43 && buf[9] == 0x52 && // CR2 magic word
		buf[10] == 0x02 // CR2 major version
}
