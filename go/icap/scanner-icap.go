package main

import (
	//"bytes"
	//"io"
	//"net/http"

	"fmt"
	"net"
	"os"
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
	requestHeader := "OPTIONS icap://" + serverIP + "/" + icapService + " ICAP/" + version + "\r\n" + "Host: " + serverIP + "\r\n" + "User-Agent: " + useragent + "\r\n" + "Encapsulated: null-body=0\r\n" + "\r\n"
	return requestHeader
}

func main() {
	// if len(os.Args) != 4 {
	// 	errorExit("usage: icap --scan file serverhost port")
	// }
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverIP+":"+serverPort)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte(getServiceOptions()))
	checkError(err)
	serverResponseBuffer := make([]byte, 4096)
	numBytesRead, err := conn.Read(serverResponseBuffer)
	checkError(err)
	fmt.Println("Response to Request for Server Options:")
	fmt.Printf("%s\n", serverResponseBuffer[0:numBytesRead])
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
