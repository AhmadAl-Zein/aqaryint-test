package main

const bufferSize = 8

func main() {
	buffer := make([]byte, bufferSize)
	readChannel := make(chan byte)
	writeChannel := make(chan byte)
}
