package utils

import "os"

// FileExist 判断文件是否存在
func FileExist(path string) bool {
    fi, err := os.Lstat(path)
    if err == nil {
        return !fi.IsDir()
    }
    return !os.IsNotExist(err)
}

// Byte2File 覆盖
func Byte2File(output []byte, filename string) (err error) {
    err = os.WriteFile(filename, output, 0644)

    return
}
