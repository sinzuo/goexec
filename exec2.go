// server
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, req *http.Request) {

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("jiangyibo"))

}

//var s1 = http.Handle("bobo", handle)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("input file name")
		return
	}

	f1, e := os.Open(os.Args[1])
	if e != nil {
		fmt.Println("input file error")
		return
	}
	defer f1.Close()
	fo, _ := os.OpenFile(os.Args[1]+"temp", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	defer fo.Close()
	f2 := bufio.NewReader(f1)
	for {
		b1, _, e := f2.ReadLine()
		if e != nil {
			break
		}
		fmt.Fprintln(fo, string(b1), "")
	}
	fo.Close()
	//	os.Remove(os.Args[1])
	os.Rename(os.Args[1]+"temp", os.Args[1])

}
