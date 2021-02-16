package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	log.SetOutput(fileObj)

	for {
		log.Println("test")
		time.Sleep(time.Second * 3)
	}
}
