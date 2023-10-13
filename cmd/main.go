package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/CyCoreSystems/ari/v6/client/native"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("env error!?, %v", err)
	}

	var (
		application = os.Getenv("ASTERISK_ARI_APP")
		serverURL   = os.Getenv("ASTERISK_URL")
		username    = os.Getenv("ASTERISK_ARI_USERNAME")
		password    = os.Getenv("ASTERISK_ARI_PASSWORD")
		callerID    = os.Getenv("CALLER_ID")
	)

	logger := log.New()
	logger.AddHook(&writer.Hook{
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
		},
	})

	logger.AddHook(&writer.Hook{
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.WarnLevel,
			log.InfoLevel,
			log.DebugLevel,
			log.TraceLevel,
		},
	})

	logger.SetOutput(io.Discard)
	logger.SetFormatter(&log.JSONFormatter{})

	//get client
	client := NewAriClient(
		application,
		callerID,
		username,
		password,
		serverURL,
		logger,
	)

	pingErr := client.PingAsteriskWebsocket(time.Second * 5)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		logger.Info("Waiting if an error happens")
		err := <-pingErr
		logger.Errorf("Ping error: %v", err)
		wg.Done()
	}()
	wg.Wait()
	logger.Info("Exit")
}

func NewAriClient(
	application,
	callerID,
	username,
	password,
	serverURL string,
	logger *log.Logger,
) *native.Client {
	url := strings.TrimPrefix(serverURL, "http://")
	options := native.Options{
		Application:  application,
		URL:          fmt.Sprintf("http://%s/ari", url),
		WebsocketURL: fmt.Sprintf("ws://%s/ari/events", url),
		Username:     username,
		Password:     password,
	}

	logger.Info("Connecting to Asterisk.")

	client := native.New(&options)

	err := client.Connect()
	if err != nil {
		logger.Fatalf("Failed to connect to Asterisk. err: %v", err)
	} else {
		logger.Info("Connected to Asterisk!")
	}

	return client
}
