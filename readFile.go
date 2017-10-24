package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	foo()
	return
}

// 按行读取
func readLine() {
	f, _ := os.Open("/etc/passwd")
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		s, _, e := r.ReadLine()
		fmt.Println(string(s))

		if e != nil {
			break
		}
	}
}

// 按字节读取到缓冲区
func readByte() {
	f, _ := os.Open("/etc/passwd")
	defer f.Close()

	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	buf := make([]byte, 1024)
	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}

		w.Write(buf[:n])
	}
}
