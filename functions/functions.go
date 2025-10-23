package functions

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var startTime = time.Now()
var serverStartTime time.Time
var serverRunning bool

func BotUptime() string {
	uptime := time.Since(startTime)
	hours := int(uptime.Hours())
	minutes := int(uptime.Minutes()) % 60
	seconds := int(uptime.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func ServerStatus() string {
	godotenv.Load()
	log.Print("Getting bot token from .env file")
	server := os.Getenv("SERVERADD")

	log.Print("Fetching server information")
	url := fmt.Sprintf("http://%s", server)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("could not create request:", err)
		return "error"
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("could not fetch server:", err)
		return "error"
	}
	defer resp.Body.Close()

	return fmt.Sprintf("%d", resp.StatusCode)
}

func StartServer() string {
	serverStartTime = time.Now()
	serverRunning = true
	return "Starting Server..."
}

func StopServer() string {
	serverRunning = false
	return "Stopping Server..."
}

func ServerUptime() string {
	if serverRunning {
		uptime := time.Since(serverStartTime)
		hours := int(uptime.Hours())
		minutes := int(uptime.Minutes()) % 60
		seconds := int(uptime.Seconds()) % 60
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	} else {
		return "server is down..."
	}
}

func ColorStatus() int {
	if ServerStatus() == "200" {
		return 0x57F287
	} else {
		return 0xFF0000
	}
}
