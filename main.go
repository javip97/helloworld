package main

import (
	"log"
	"syscall"
)

func main() {

	fd, err := syscall.Open("something.txt", syscall.O_RDWR, 755)
	if err != nil {
		log.Print(err)
	}
	log.Printf("File Descriptor: %v\n", fd)

	//Read from file using File Descriptor
	data := make([]byte, 12)
	n, err := syscall.Read(fd, data)
	if err != nil {
		log.Print(err)
	}
	log.Printf("Read %d bytes", n)
	log.Printf("Content: %s", data)
}
