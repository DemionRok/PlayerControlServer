package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

const potPlayerPath = `C:/Program Files/DAUM/PotPlayer`

var potPlayerProcess *exec.Cmd

func startPotPlayer() {
	potPlayerProcess = exec.Command(potPlayerPath)
	err := potPlayerProcess.Start()
	if err != nil {
		log.Println("Error starting PotPlayer:", err)
	}
}

func handleCommand(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")

	if potPlayerProcess == nil || potPlayerProcess.Process == nil {
		startPotPlayer()
	}

	switch command {
	case "play":
		sendCommandToPotPlayer("/play")
	case "pause":
		sendCommandToPotPlayer("/pause")
	case "stop":
		sendCommandToPotPlayer("/stop")
	case "volumeUp":
		sendCommandToPotPlayer("/vol_up")
	case "volumeDown":
		sendCommandToPotPlayer("/vol_down")
	default:
		http.Error(w, "Invalid command", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendCommandToPotPlayer(command string) {
	if runtime.GOOS == "windows" {
		err := exec.Command("cmd", "/C", "echo "+command+">"+potPlayerPath+".control").Run()
		if err != nil {
			log.Println("Error sending command to PotPlayer:", err)
		}
	} else {
		log.Println("Unsupported operating system")
	}
}

func main() {
	http.HandleFunc("/api/potplayer", handleCommand)

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
