package server

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/AdrianoSantana/tcp-server/cmd/server/dto"
	"github.com/joho/godotenv"
)

var clients []int = []int{}
const ACTION_LIST string = "LIST"
const ACTION_RELAY string = "RELAY"

func generateClientId() int {
	return rand.Int()
}

func addClient(id int) {
	clients = append(clients, id)
}

func handleIncomingRequest(conn net.Conn) dto.Request {
    // store incoming data
    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }

	var r dto.Request
	rawRequest := strings.Split(string(buffer), " ")
	if len(rawRequest) >= 2 {
		id, _ := strconv.Atoi(rawRequest[0])
		r.Id = id

		r.Action = rawRequest[1]
		r.Body = rawRequest[2]
	}
	return r
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

	sendMessage(conn, strconv.Itoa(clientId))
	defer conn.Close()
	defer listener.Close()

	for {
		r := handleIncomingRequest(conn)
		if r.Action == ACTION_LIST {
			response := createListString(clients)
			conn.Write([]byte(response))
		}
	}
	
}

func createListString(clients []int) string {
	return strings.Join(strings.Fields(fmt.Sprint(clients)), "")
}

func sendMessage(conn net.Conn, message string) {
	conn.Write([]byte(message + "\n"))
}
