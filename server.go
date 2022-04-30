package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type Server struct {
	m_serverSocket     net.Listener
	m_waitGroup        sync.WaitGroup
	m_connectedClients map[string]net.Conn
	m_isServerWorking  bool
}

func (serverEntity *Server) StartServer(ip string, port string) bool {
	var serverInitializationError error
	serverEntity.m_serverSocket, serverInitializationError = net.Listen("tcp", ip+":"+port)

	return serverInitializationError != nil
}

func InitializeConnection() {
	var mainServerEntity *Server = new(Server)
	mainServerEntity.StartServer("localhost", "5555")

	defer mainServerEntity.StopServer()

	mainServerEntity.m_connectedClients = make(map[string]net.Conn, 0)

	for {
		newClient, clientAcceptError := mainServerEntity.m_serverSocket.Accept()
		if clientAcceptError != nil {
			fmt.Println("Error while establishing connection with new client!")
			return
		}

		if newClient != nil {
			mainServerEntity.m_connectedClients[newClient.LocalAddr().String()] = newClient

			mainServerEntity.m_waitGroup.Add(1)

			go mainServerEntity.ProcessClientConnection(newClient.LocalAddr().String())
		}

	}
}

func (serverEntity *Server) StopServer() {
	if serverEntity.m_isServerWorking {
		fmt.Println("Stopping server!")
		serverEntity.m_serverSocket.Close()
		serverEntity.m_isServerWorking = false
	}
}

func (mainServer *Server) ProcessClientConnection(userIp string) {
	defer mainServer.m_waitGroup.Done()

	buffer := make([]byte, 1024)

	var errorCounter = 0
	for {
		_, err := mainServer.m_connectedClients[userIp].Read(buffer)
		if err != nil {
			fmt.Println("Error while processing user with ip: ", userIp)
			errorCounter++
			if errorCounter > 5 {
				return
			}
		}

		mainServer.m_connectedClients[userIp].Write([]byte("Got it"))
		fmt.Println("New message from ", userIp, " : ", string(buffer))

		if strings.Contains(string(buffer), "stop") {
			mainServer.StopServer()
			return
		}
	}
}
