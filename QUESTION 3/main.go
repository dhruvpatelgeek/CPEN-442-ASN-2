package main

import (
	"fmt"
	"hash/crc32"
	"log"
	"sync"
	"time"
	"math/rand"
)
var semap sync.Mutex
var MAX_WORKER =8
func main() {
	start := time.Now()
	semap.Lock()
	elapsed := time.Since(start)

	for i:=0;i<MAX_WORKER;i++{
		go crack(i)
	}

	semap.Lock()
	log.Printf("Procress took %s", elapsed)
}


func crack(num int){
	fmt.Println("worker # ",num," online ")
	data := "102E91D22C3795494B378096783BE2A3"
	crc32_val:=crc32.ChecksumIEEE([]byte(data))

	token := make([]byte, 32000000)
	rand.Read(token)
	token_crc:=crc32.ChecksumIEEE([]byte(token))
	var itr=0
	for token_crc!=crc32_val{
		itr+=1;
		if(itr%100==0){
			fmt.Println("finished #",itr)
		}
		rand.Read(token)
		token_crc=crc32.ChecksumIEEE(token)
	}
	fmt.Println(num,": CRC token is",token_crc," and ",crc32_val)
	semap.Unlock()
}