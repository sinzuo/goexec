// server
package main

import (
	"bytes"
	"fmt"
	"net/http"

	"os/exec"
)

func handle(w http.ResponseWriter, req *http.Request) {

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("jiangyibo"))

}

//var s1 = http.Handle("bobo", handle)

func main() {
	c1 := exec.Command("/bin/sh", "-f", "jiang.txt")
	b1 := &bytes.Buffer{}
	c1.Stdout = b1
	c1.Run()

	fmt.Println("Hello World!" + string(b1.Bytes()))

	c1 = exec.Command("/bin/sh", "-c", "echo jiang")
	b1 = &bytes.Buffer{}
	c1.Stdout = b1
	c1.Run()

	fmt.Println("Hello World!" + string(b1.Bytes()))
}
