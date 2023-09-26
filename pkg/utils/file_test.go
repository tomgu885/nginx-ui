package utils

import (
    "fmt"
    "testing"
)

func TestByte2File(t *testing.T) {
    f := []byte("hello world\n\tbeauty")
    err := Byte2File(f, "test")
    if err != nil {
        fmt.Println("err:", err)
        return
    }

    fmt.Println("success")
}
