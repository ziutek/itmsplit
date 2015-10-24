package main

import (
	"os"
	"strconv"
	"strings"
)

type splitter struct {
	itm  map[int]*os.File
	dflt *os.File
}

func newSplitter(args []string) *splitter {
	s := &splitter{itm: make(map[int]*os.File)}
	for _, arg := range args {
		i := strings.IndexByte(arg, ':')
		if i == -1 {
			if s.dflt != nil {
				die(arg, "redefines default output")
			}
			s.dflt = open(arg)
			continue
		}
		src, out := arg[:i], arg[i+1:]
		if len(src) < 2 || src[0] != 'p' {
			die("unknown source:", src)
		}
		port := src[1:]
		p, err := strconv.Atoi(port)
		if err != nil || uint(p) > 32 {
			die("bad ITM stimulus port number:", port)
		}
		if s.itm[p] != nil {
			die(out, "redefines output for", src)
		}
		s.itm[p] = open(out)
	}
	return s
}

func (s *splitter) WriteITM(p int, buf []byte) {
	if f := s.itm[p]; f != nil {
		f.Write(buf)
		return
	}
	if s.dflt != nil {
		s.dflt.Write(buf)
	}
}
