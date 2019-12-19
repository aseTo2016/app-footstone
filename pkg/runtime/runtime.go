package runtime

import "runtime"

// GetCodeFilePath returns
func GetCodeFilePath() string {
	_, fileName, _, _ := runtime.Caller(1)
	return fileName
}

// Stack returns stack
func Stack() string {
	data := make([]byte, 0, 1024)
	runtime.Stack(data, true)
	return string(data)
}
