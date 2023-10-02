package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

const potPlayerPath = `C:/Program Files/DAUM/PotPlayer`

func handleCommand(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")

	switch command {
	case "play":
		if runtime.GOOS == "windows" {
			exec.Command(potPlayerPath, "PotPlayerMini64.exe /play").Start()
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
		}
	case "pause":
		if runtime.GOOS == "windows" {
			exec.Command(potPlayerPath, "PotPlayerMini64.exe /pause").Start()
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
		}
	case "stop":
		if runtime.GOOS == "windows" {
			exec.Command(potPlayerPath, "PotPlayerMini64.exe /stop").Start()
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
		}
	case "volumeUp":
		if runtime.GOOS == "windows" {
			exec.Command(potPlayerPath, "PotPlayerMini64.exe /vol_up").Start()
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
		}
	case "volumeDown":
		if runtime.GOOS == "windows" {
			exec.Command(potPlayerPath, "PotPlayerMini64.exe /vol_down").Start()
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Invalid command", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api/potplayer", handleCommand)

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
