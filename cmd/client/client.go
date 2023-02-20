package client

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func StartClient() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("CLIENT_PORT")

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
        log.Fatal("Dial failed:", err.Error())
    }

	fmt.Println("Connected to the server")

	defer conn.Close()

	cmd := "1 LIST abc"
	for {
		conn.Write([]byte(cmd))
		time.Sleep(3 * time.Second)
	}



}