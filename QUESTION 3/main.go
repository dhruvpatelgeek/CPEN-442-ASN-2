package main

import (
	"fmt"
	"hash/crc32"
	"log"
	"math/rand"
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
	//data := "102E91D22C3795494B378096783BE2A3"
	//crc32_val:=crc32.ChecksumIEEE([]byte(data))
	var cache=make(map[string]string)

	var token=make([]byte,80000)

	var itr=0
	var a,b []byte
	for {
		rand.Read(token)
		token_crc:=crc32.ChecksumIEEE([]byte(token))
		itr+=1
		if itr%1000==0{
			fmt.Println("#",itr)
		}
		if val, ok := cache[strconv.Itoa(int(token_crc))];ok {
			a=token
			b=[]byte(val)
			break;
		} else {
			cache[strconv.Itoa(int(token_crc))]= string(token)
		}
	}
	fmt.Println(num,": CRC token is")
	fmt.Println(a)
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println(b)


	semap.Unlock()
}