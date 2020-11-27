package main

import (
	"flag"
	"log"
	"syscall"
)

func main() {

	path := flag.String("path", "testfile.txt", "set the path variable to exampleOutput.txt")
	flag.Parse()

	fd, err := syscall.Open(*path, syscall.O_CREAT|syscall.O_RDWR|syscall.O_TRUNC, 755)

	if err != nil {
		log.Println("creating file error", err)
		return
	}

	log.Println(fd)
	log.Println(*path)

	// fd, err = syscall.Open(*path, syscall.O_RDWR, 755)
	// if err != nil {
	// 	log.Println("there has been an error")
	// 	log.Print(err)
	// 	return
	// }

	// log.Println("no error keep executing")
	// log.Printf("File Descriptor: %v\n", fd)

	bytesWritten, err := syscall.Write(fd, []byte("hello world from fd"))
	if err != nil {
		log.Println("error with writting fd")
		log.Println(err)
		return
	}
	log.Println(bytesWritten)

	// fd, err = syscall.Open(*path, syscall.O_RDWR, 755)
	// log.Println("bytes written", bytesWritten)
	// //Read from file using File Descriptor
	// data := make([]byte, 50)
	// n, err := syscall.Read(fd, data)
	// if err != nil {
	// 	log.Println("error syscall read")
	// 	log.Print(err)
	// 	return
	// }
	// log.Printf("Read %d bytes", n)
	// log.Printf("Content: %s", data)
}
