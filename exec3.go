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
	b1 := bufio.NewReader(os.Stdin)
	c1, _, e := b1.ReadLine()
	if e != nil {
		fmt.Println("error")
	}
	fmt.Println(string(c1))
	fmt.Println("Hello World!")
}
