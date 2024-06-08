package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var host string
var port int
var timeOutStr string

func copyTo(target io.Writer, source io.Reader) {

	buff := make([]byte, 1024)
	for {

		n, err := source.Read(buff[:])
		if errors.Is(err, io.EOF) {
			fmt.Println("Connection is closed")
			os.Exit(0)
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = target.Write(buff[:n])
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			return
		}

		//buff = buff[:0]
	}
}

// _, err := io.Copy(target, source)
// if errors.Is(err, io.EOF) {
// 	return
// }
// if err != nil {
// 	log.Fatal(err)
// }
//}

func main() {
	flag.StringVar(&host, "h", "127.0.0.1", "host for connection")
	flag.IntVar(&port, "p", -1, "port for connection")
	flag.StringVar(&timeOutStr, "t", "10s", "timeout connection")
	flag.Parse()

	fmt.Println(port)
	if port < 0 || port > 65535 {
		log.Fatal("Port is not found")
	}

	timeOut, err := time.ParseDuration(timeOutStr)
	if err != nil {
		log.Fatal(err)
	}
	// if strings.EqualFold(host, "localhost") {
	// 	host = "127.0.0.1"
	// }
	// hostTest := strings.Split(host, ".")
	// if len(hostTest) != 4 {
	// 	log.Fatal("Host is not found")
	// }
	// for _, el := range hostTest {
	// 	res, err := strconv.Atoi(el)
	// 	if err != nil || res < 0 || res > 255 {
	// 		log.Fatal("Host is not found")
	// 	}
	// }

	conn, err := net.DialTimeout("tcp", host+":"+fmt.Sprint(port), timeOut)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	defer fmt.Println("test")
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		//for{
		<-sigChan
		os.Exit(0)
		//}
	}()

	go copyTo(os.Stdout, conn)

	copyTo(conn, os.Stdin)
}
