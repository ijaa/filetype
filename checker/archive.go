package checker

func Zip(file string) (bool, error) {
	buf, err := getBytes(file, 0, 4)
	if err != nil {
		return false, err
	}
	return len(buf) > 3 &&
		buf[0] == 0x50 && buf[1] == 0x4B &&
		(buf[2] == 0x3 || buf[2] == 0x5 || buf[2] == 0x7) &&
		(buf[3] == 0x4 || buf[3] == 0x6 || buf[3] == 0x8), nil
}

func Tar(file string) (bool, error) {
	buf, err := getBytes(file, 257, 5)
	if err != nil {
		return false, err
	}
	return len(buf) > 4 &&
		buf[0] == 0x75 && buf[1] == 0x73 &&
		buf[2] == 0x74 && buf[3] == 0x61 &&
		buf[4] == 0x72, nil
}

func Rar(file string) (bool, error) {
	buf, err := getBytes(file, 0, 7)
	if err != nil {
		return false, err
	}
	return len(buf) > 6 &&
		buf[0] == 0x52 && buf[1] == 0x61 && buf[2] == 0x72 &&
		buf[3] == 0x21 && buf[4] == 0x1A && buf[5] == 0x7 &&
		(buf[6] == 0x0 || buf[6] == 0x1), nil
}

func Gz(file string) (bool, error) {
	buf, err := getBytes(file, 0, 3)
	if err != nil {
		return false, err
	}
	return len(buf) > 2 &&
		buf[0] == 0x1F && buf[1] == 0x8B && buf[2] == 0x8, nil
}

func Bz2(file string) (bool, error) {
	buf, err := getBytes(file, 0, 3)
	if err != nil {
		return false, err
	}
	return len(buf) > 2 &&
		buf[0] == 0x42 && buf[1] == 0x5A && buf[2] == 0x68, nil
}

func SevenZ(file string) (bool, error) {
	buf, err := getBytes(file, 0, 6)
	if err != nil {
		return false, err
	}
	return len(buf) > 5 &&
		buf[0] == 0x37 && buf[1] == 0x7A && buf[2] == 0xBC &&
		buf[3] == 0xAF && buf[4] == 0x27 && buf[5] == 0x1C, nil
}

func Pdf(file string) (bool, error) {
	buf, err := getBytes(file, 0, 4)
	if err != nil {
		return false, err
	}
	return len(buf) > 3 &&
		buf[0] == 0x25 && buf[1] == 0x50 &&
		buf[2] == 0x44 && buf[3] == 0x46, nil
}

func Exe(file string) (bool, error) {
	buf, err := getBytes(file, 0, 2)
	if err != nil {
		return false, err
	}
	return len(buf) > 1 &&
		buf[0] == 0x4D && buf[1] == 0x5A, nil
}

func Swf(file string) (bool, error) {
	buf, err := getBytes(file, 0, 3)
	if err != nil {
		return false, err
	}
	return len(buf) > 2 &&
		(buf[0] == 0x43 || buf[0] == 0x46) &&
		buf[1] == 0x57 && buf[2] == 0x53, nil
}

func Rpm(file string) (bool, error) {
	buf, err := getBytes(file, 0, 97)
	if err != nil {
		return false, err
	}
	return len(buf) > 3 &&
		buf[0] == 0xED && buf[1] == 0xAB &&
		buf[2] == 0xEE && buf[3] == 0xDB, nil
}

func Elf(file string) (bool, error) {
	buf, err := getBytes(file, 0, 53)
	if err != nil {
		return false, err
	}
	return len(buf) > 3 &&
		buf[0] == 0x7F && buf[1] == 0x45 &&
		buf[2] == 0x4C && buf[3] == 0x46, nil
}

func Iso(file string) (bool, error) {
	buf, err := getBytes(file, 32769, 5)
	if err != nil {
		return false, err
	}
	return len(buf) > 4 &&
		buf[0] == 0x43 && buf[1] == 0x44 &&
		buf[2] == 0x30 && buf[3] == 0x30 &&
		buf[4] == 0x31, nil
}

// Wasm detects a Web Assembly 1.0 filetype.
func Wasm(file string) (bool, error) {
	// WASM has starts with `\0asm`, followed by the version.
	// http://webassembly.github.io/spec/core/binary/modules.html#binary-magic
	buf, err := getBytes(file, 0, 8)
	if err != nil {
		return false, err
	}
	return len(buf) > 7 &&
		buf[0] == 0x00 && buf[1] == 0x61 &&
		buf[2] == 0x73 && buf[3] == 0x6D &&
		buf[4] == 0x01 && buf[5] == 0x00 &&
		buf[6] == 0x00 && buf[7] == 0x00, nil
}
