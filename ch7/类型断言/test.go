package main

import (
	"bytes"
	//"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	//c := w.(*bytes.Buffer)
	//fmt.Println(f, c)
	f.Write([]byte("123"))


	w = os.Stdout
	rw := w.(io.ReadWriter)
	fmt.Println(rw)
	rw.Write([]byte("123"))

	var w1 io.Writer = os.Stdout
	f, ok := w.(*os.File) //成功: ok, f == os.Stdout
	b, ok := w.(*bytes.Buffer) //失败: !ok, b == nil
	if f, ok := w.(*os.File); ok {
		// ...使用 f...
	}

	if w, ok := w.(*os.File); ok {
		// ...use w...
	}
}
