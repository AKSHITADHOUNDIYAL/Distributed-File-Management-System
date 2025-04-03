package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const masterURL = "http://localhost:8080/register" // Master node address

type Worker struct {
	ID   string `json:"id"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func registerWithMaster() {
	worker := Worker{
		ID:   "worker-1",
		IP:   "127.0.0.1",
		Port: "8081",
	}

	data, err := json.Marshal(worker)
	if err != nil {
		fmt.Println("Error encoding worker data:", err)
		return
	}

	resp, err := http.Post(masterURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Failed to register with master:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Registered with master successfully")
}

func main() {
	fmt.Println("Worker Node Starting...")

	// Register with master
	registerWithMaster()

	// Keep running (Placeholder for file handling logic)
	select {}
}
