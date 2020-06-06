package checker

func Mp3(file string) (bool, error) {
	buf, err := getBytes(file, 0, 3)
	if err != nil {
		return false, err
	}
	return len(buf) > 2 &&
		((buf[0] == 0x49 && buf[1] == 0x44 && buf[2] == 0x33) ||
			(buf[0] == 0xFF && buf[1] == 0xfb)), nil
}

func M4a(file string) (bool, error) {
	buf, err := getBytes(file, 0, 11)
	if err != nil {
		return false, err
	}
	return len(buf) > 10 &&
		((buf[4] == 0x66 && buf[5] == 0x74 && buf[6] == 0x79 &&
			buf[7] == 0x70 && buf[8] == 0x4D && buf[9] == 0x34 && buf[10] == 0x41) ||
			(buf[0] == 0x4D && buf[1] == 0x34 && buf[2] == 0x41 && buf[3] == 0x20)), nil
}

func Wav(file string) (bool, error) {
	buf, err := getBytes(file, 0, 12)
	if err != nil {
		return false, err
	}
	return len(buf) > 11 &&
		buf[0] == 0x52 && buf[1] == 0x49 &&
		buf[2] == 0x46 && buf[3] == 0x46 &&
		buf[8] == 0x57 && buf[9] == 0x41 &&
		buf[10] == 0x56 && buf[11] == 0x45, nil
}
