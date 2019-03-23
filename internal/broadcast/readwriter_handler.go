package broadcast

import (
	"bufio"
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
	InputValidator InputValidator
	Writer         io.Writer
	WriterStorage  writerutil.Storage
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
			incommingData, err = bufio.NewReader(readWriter).ReadString('\n')
		)

		if err == io.EOF {
			log.Println("recieved EOF")
			return
		} else if err != nil {
			log.Println("failed to read input:", err.Error())
			return
		}

		handler.handleInput([]byte(incommingData), readWriter)
	}
}

func (handler *ReadWriteHandler) handleInput(input []byte, writer io.Writer) {
	var (
		err        error
		validInput bool
	)

	if validInput, err = handler.InputValidator.Validate(input); err != nil {
		writer.Write([]byte("unable to validate input"))
		return
	}

	if !validInput {
		writer.Write([]byte("invalid input"))
		return
	}

	handler.Writer.Write(input)
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
