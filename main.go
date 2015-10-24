package main

import (
	"bufio"
	//"fmt"
	"os"
)

const usage = `Usage:
  itmsplit [SRC:]FILE ...
where FILE is path to the file where output will be appended,
SRC specifies data source and can be:
  - ITM stimulus port number: p0, p1, ..., p15,
  - DWT source (TODO).
Examples:
1. Read from file and append all to stdout:
	itmsplit /dev/stdout <file
2. From stimulus port 0 to stdout, other to stderr:
	openocd |itmsplit p0:/dev/stdout /dev/stderr
3. Save data from first two stimulus ports, discard other:
	openocd |itmsplit p0:port0.txt p1:port1.txt
`

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString(usage)
		os.Exit(1)
	}
	r := bufio.NewReader(os.Stdin)
	s := newSplitter(os.Args[1:])
	var buf [7]byte
	for {
		n := readPacket(r, &buf)
		h := int(buf[0])
		if h&3 == 0 {
			// Protocol packet
			continue
		} else {
			// Source packet
			if h&4 != 0 {
				// Packet from DWT.
				continue
			}
			// Packet from ITM stimulus port.
			//fmt.Printf("\n%x\n", buf[1:n])
			s.WriteITM(h>>3, buf[1:n])
		}
	}
}
