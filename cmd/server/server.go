package server

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var clients []int = []int{}

func generateClientId() int {
	return rand.Int()
}

func addClient(id int) {
	clients = append(clients, id)
}

func StartServer() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("SERVER_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("It was not possible to start server: %s", err)
	}
	fmt.Printf("Server running on port %s\n", port)

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Connection failed", err)
	}

	clientId := generateClientId()
	addClient(clientId)

	conn.Write([]byte(strconv.Itoa(clientId)))

	defer conn.Close()
	defer listener.Close()
}
