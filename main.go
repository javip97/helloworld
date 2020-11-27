package main

import (
	"flag"
	"log"
	"syscall"
)

func main() {

	path := flag.String("path", "/home/guest/go/helloworld-main/testfile.txt", "usage")
	flag.Parse()
	fd, err := syscall.Open(*path, syscall.O_RDWR, 755)
	log.Println(*path)
	if err != nil {
		log.Println("there has been an error")
		log.Print(err)
		return
	}

	log.Println("no error keep executing")
	log.Printf("File Descriptor: %v\n", fd)

	bytesWritten, err := syscall.Write(fd, []byte("hello world from fd"))
	if err != nil {
		log.Println("error with writting fd")
		log.Println(err)
		return
	}

	fd, err = syscall.Open("something.txt", syscall.O_RDWR, 755)
	log.Println("bytes written", bytesWritten)
	//Read from file using File Descriptor
	data := make([]byte, 50)
	n, err := syscall.Read(fd, data)
	if err != nil {
		log.Println("error syscall read")
		log.Print(err)
		return
	}
	log.Printf("Read %d bytes", n)
	log.Printf("Content: %s", data)
}
