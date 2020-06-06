package filetype

// Is check the file's type equals the custom FileType
// return error when get the file content io operation failed
func Is(ft FileType, file string) (bool, error) {
	return ft.Checker(file)
}

// IsIn check the file's type in one of the types
func IsIn(fts []FileType, file string) (bool, error) {
	for _, ft := range fts {
		is, err := Is(ft, file)
		if err != nil {
			continue
		}
		if is {
			return true, nil
		}
	}
	return false, nil
}
