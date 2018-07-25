package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	"strconv"
)


func leadingZeroes(data []byte) int {
	zeroes := 0
	for _, v := range data {
		for x:= uint(0); x < 8; x++ {
			if v & (1 << x) != 0 {
				return zeroes
			}
			zeroes++
		}
	}
	return zeroes
}

func ProofOfWork(data []byte, target int) (hash []byte, nonce uint64){

	for i := uint64(0); ; i++{
		combined := bytes.Join([][]byte{data, []byte(strconv.FormatUint(i, 10))}, []byte{})
		sum256 := sha256.Sum256(combined)
		if leadingZeroes(sum256[:]) >= target {
			return sum256[:], i
		}
	}
}

func main() {
	t := time.Now()
	hash, nonce := ProofOfWork([]byte("somedata to hash"), 8)
	since := time.Since(t)
	fmt.Println("hash is sum256", hex.EncodeToString(hash), "and nonce", nonce, "took", since)
}
