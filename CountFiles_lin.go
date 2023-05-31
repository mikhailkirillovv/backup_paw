// CountFiles
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func backup() {
	cmd := exec.Command("/bin/sh", "/home/opc/scripts/test.sh")
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	logFile, err := os.OpenFile("my.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
}

func main() {
	a, _ := os.ReadDir("/home/opc/t1")
	fmt.Println(len(a))
	if len(a) < 4 {
		log.Println("without_redis")
		time.Sleep(7 * time.Second)
		backup()
	else {
		log.Println("full_backup")
	}
}
