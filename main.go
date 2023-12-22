package main

import (
	"github.com/micmonay/keybd_event"
	"log"
	"net/http"
	"runtime"
)

func sendKey(keyCode ...int) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, key := range keyCode {
		kb.SetKeys(key)
		err := kb.Press()
		if err != nil {
			log.Println("Press error:", err)
			return
		}
	}

	for _, keyCode := range keyCode {
		kb.SetKeys(keyCode)
		err := kb.Release()
		if err != nil {
			log.Println("Release error:", err)
			return
		}
	}
}

func handleCommand(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")

	switch command {
	case "startPause":
		if runtime.GOOS == "windows" {
			sendKey(keybd_event.VK_LMENU, keybd_event.VK_HOME)
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
			return
		}
	case "moveForward":
		if runtime.GOOS == "windows" {
			sendKey(keybd_event.VK_LMENU, keybd_event.VK_RIGHT)
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
		}
	case "moveBackward":
		if runtime.GOOS == "windows" {
			sendKey(keybd_event.VK_LMENU, keybd_event.VK_LEFT)
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
		}
	case "volumeUp":
		if runtime.GOOS == "windows" {
			sendKey(keybd_event.VK_LMENU, keybd_event.VK_UP)
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
			return
		}
	case "volumeDown":
		if runtime.GOOS == "windows" {
			sendKey(keybd_event.VK_LMENU, keybd_event.VK_DOWN)
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
			return
		}
	case "nextFile":
		if runtime.GOOS == "windows" {
			sendKey(keybd_event.VK_LMENU, keybd_event.VK_PAGEDOWN)
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
			return
		}
	case "prevFile":
		if runtime.GOOS == "windows" {
			sendKey(keybd_event.VK_LMENU, keybd_event.VK_PAGEUP)
		} else {
			log.Println("Unsupported OS")
			http.Error(w, "Unsupported OS", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Invalid command", http.StatusBadRequest)
		return
	}

	log.Println("Command executed:", command)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api/wmp", handleCommand)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe("192.168.1.103:8080", nil))
}
