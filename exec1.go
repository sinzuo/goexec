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
	c1 := exec.Command("echo", "jiang")
	_, err := c1.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	var out bytes.Buffer
	c1.Stdout = &out
	c1.Run()
	fmt.Println(out.String())
	//	http.HandleFunc("/jiang", handle)
	//http.ListenAndServe(":9999", nil)
	fmt.Println("Hello World!")
}
