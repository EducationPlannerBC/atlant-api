package main

import (
	//"bytes"
	//"io"
	//"net/http"

	"fmt"
	"net"
	"os"
	"time"
	//"strconv"
	//"strings"
)

const (
	serverIP   = "192.168.100.110"
	serverPort = "1345"
	connType   = "tcp"
)

const defaultPort = 1345 // usally 1344 but EPBC uses 1345
const eicar = "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"

const version = "1.0"
const useragent = "EPBC-Client/1.1"

// const ICAPTERMINATOR = "\r\n\r\n"
// const HTTPTERMINATOR = "0\r\n\r\n"

var icapService = "respmod"

func errorExit(err string) {
	os.Exit(1)
}

func getServiceOptions() string {
	//requestHeader := "OPTIONS icap://" + serverIP + "/" + icapService + " ICAP/" + version + "\r\n" + "Host: " + serverIP + "\r\n" + "User-Agent: " + useragent + "\r\n" + "Encapsulated: null-body=0\r\n" + "\r\n"
	requestHeader := "OPTIONS icap://" + serverIP + ":" + serverPort + "/ " +
		"ICAP/" + version + "\r\n" + "Host: " + serverIP + ":" + serverPort + "\r\n" +
		"Encapsulated: null-body=0\r\n\r\n"
	return requestHeader
}

func main() {
	// if len(os.Args) != 4 {
	// 	errorExit("usage: icap --scan file serverhost port")
	// }
	conn, err := net.DialTimeout("tcp", serverIP+":"+serverPort, 1*time.Second)
	checkError(err)
	_, err = conn.Write([]byte(getServiceOptions()))
	checkError(err)
	bufferSize := 4096
	serverResponseBuffer := make([]byte, bufferSize)
	fmt.Println("Response to Request for Server Options:")
	var fullResponse []byte
	numBytesRead := bufferSize
	for numBytesRead == bufferSize {
		numBytesRead, err = conn.Read(serverResponseBuffer)
		checkError(err)
		//fmt.Printf("%s\n", serverResponseBuffer[0:numBytesRead])
		fullResponse = append(fullResponse, serverResponseBuffer[0:numBytesRead]...)
	}
	fmt.Printf("%s\n", fullResponse)
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
