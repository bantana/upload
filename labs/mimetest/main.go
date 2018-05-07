package main

import (
	"fmt"
	"log"
	"mime"
)

func main() {
	fmt.Println("vim-go")
	src := "您好!"

	fmt.Printf("src : %#v\n", src)
	strutf8 := mime.QEncoding.Encode("utf-8", src)
	fmt.Printf("strutf8: %#v\n", strutf8)
	dec := new(mime.WordDecoder)
	str, err := dec.Decode(strutf8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("str : %#v\n", str)
	if src == str {
		fmt.Printf("str: %#v is equal src: %#v\n", str, src)
	}
}
