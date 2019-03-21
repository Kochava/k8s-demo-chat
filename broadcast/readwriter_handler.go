package broadcast

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"time"

	"git.dev.kochava.com/notter/distchat/writerutil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ReadWriteHandler allows an io.ReadWriter to send and receive chats
type ReadWriteHandler struct {
	Writer        io.Writer
	WriterStorage writerutil.Storage
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
			return
		} else if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		handler.Writer.Write([]byte(incommingData))
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
