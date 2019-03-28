package broadcast

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/Kochava/k8s-demo-chat/internal/writerutil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ReadWriteHandler allows an io.ReadWriter to send and receive chats
type ReadWriteHandler struct {
	Writer         io.Writer
	WriterStorage  writerutil.Storage
	InputValidator InputValidator
}

// Handle communicates with a TCP connection
func (handler *ReadWriteHandler) Handle(readWriter io.ReadWriter) {
	var (
		connectionName = handler.getConnectionName()
	)

	handler.WriterStorage.Add(connectionName, readWriter)
	defer handler.WriterStorage.Remove(connectionName)

	for {
		var (
			validInput            bool
			incommingData         []byte
			incommingDataStr, err = bufio.NewReader(readWriter).ReadString('\n')
		)

		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		incommingData = []byte(incommingDataStr)

		if validInput, err = handler.InputValidator.Valid(incommingData); err != nil {
			log.Println("invalid input:", err.Error())
			readWriter.Write([]byte("unable to validate input"))
		}

		if !validInput {
			readWriter.Write([]byte("unable to validate input"))
		}

		handler.Writer.Write(incommingData)
	}
}

func (handler *ReadWriteHandler) getConnectionName() string {
	var (
		nameLength  = 10
		letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	)

	b := make([]rune, nameLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
