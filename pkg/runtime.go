package pkg

import "runtime"

// GetCodeFilePath returns
func GetCodeFilePath() string {
	_, fileName, _, _ := runtime.Caller(1)
	return fileName
}
