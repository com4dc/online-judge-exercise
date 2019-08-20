package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (n int, e error) {
	n, e = reader.r.Read(b) // readerの中身を読む
	if e == nil {           // errorがなかったら処理してく
		for i, v := range b {
			switch {
			case v >= 'A' && v <= 'Z': // 13個シフト（ROT13だから）
				b[i] = (v-'A'+13)%26 + 'A' // bの中身を書き換えていく
			case v >= 'a' && v <= 'z':
				b[i] = (v-'a'+13)%26 + 'a'
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
