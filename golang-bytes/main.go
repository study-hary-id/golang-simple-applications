package main

import (
	"fmt"

	"golang-bytes/bytesize"
)

func main() {
	fmt.Println(bytesize.ByteSize(2048))
	fmt.Println(bytesize.ByteSize(3292528.64))
	fmt.Println(bytesize.ByteSize(512))
}
