package main

import "fmt"
import "github.com/awnumar/memguard"

func main() {
	key := memguard.NewEnclaveRandom(32)
	buf, _ := key.Open()

	// []byte{0xd8, 0x69, 0xb4, 0x65, 0x42, 0x99, 0xc0, 0x65, 0xce, 0x24, 0x9c, 0x2f, 0x27, 0xcf, 0xb8, 0x28, 0xdc, 0xc, 0xa9, 0xdd, 0x23, 0x30, 0x35, 0xfe, 0x91, 0x90, 0xc2, 0x4, 0x8f, 0x8e, 0x12, 0x7c}
	fmt.Printf("%#v\n", buf.Bytes())

	select {}
}
