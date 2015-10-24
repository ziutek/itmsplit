package main

import (
	"bufio"
	"io"
	"os"
)

func die(msg ...string) {
	for i, m := range msg {
		if i != 0 {
			os.Stderr.WriteString(" ")
		}
		os.Stderr.WriteString(m)
	}
	os.Stderr.WriteString("\n")
	os.Exit(1)
}

func checkErr(err error) {
	if err != nil {
		die("error:", err.Error())
	}
}

func open(name string) *os.File {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	checkErr(err)
	return f
}

func write(f *os.File, b []byte) {
	_, err := f.Write(b)
	checkErr(err)
}

func readByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		if err != io.EOF {
			die(err.Error())
		}
		os.Exit(0)
	}
	return b
}
