package main

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)
var semap sync.Mutex
var MAX_WORKER =1
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

	var token=make([]byte,80000)
	rand.Read(token)
	token_crc:=crc32.ChecksumIEEE([]byte(token))
	var itr=0
	for token_crc!=crc32_val{
		itr+=1;
		if(itr%100000==0){
			fmt.Println("finished #",itr)
		}
		rand.Read(token)
		token_crc=crc32.ChecksumIEEE([]byte(token))
	}
	fmt.Println(num,": CRC token is",token_crc," and ",crc32_val)
	fmt.Println(num,": CRC val is",token[:1000])
	var file_name = "RAND_CRYPT#"
	file_name+=strconv.Itoa(num)
	file_name+=".log"
	f, err := os.Create(file_name)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = f.WriteString(string(token))
	defer f.Close()
	semap.Unlock()
}