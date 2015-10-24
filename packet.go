package main

import (
	"bufio"
)

func readPacket(r *bufio.Reader, buf *[7]byte) int {
	for {
		b := readByte(r)
		for b == 0 {
			b = readByte(r)
		}
		if b == 0x80 {
			continue // Synchronization packet
		}
		if b&0x03 == 0 {
			// Protocol packet
			return readProtoPayload(r, buf, b)
		} else {
			// Source packet
			return readSourcePayload(r, buf, b)
		}
	}
}

func readProtoPayload(r *bufio.Reader, buf *[7]byte, b byte) int {
	var n int
	for {
		buf[n] = b
		if n++; n == len(buf) || b&0x80 == 0 {
			return n
		}
		b = readByte(r)
	}
}

func readSourcePayload(r *bufio.Reader, buf *[7]byte, b byte) int {
	pktlen := 1 + 1<<(b&3-1)
	buf[0] = b
	for n := 1; n < pktlen; n++ {
		buf[n] = readByte(r)
	}
	return pktlen
}
