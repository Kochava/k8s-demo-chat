package broadcast

import (
	"io"
	"log"
	"sync"

	"github.com/Kochava/k8s-demo-chat/internal/writerutil"
)

var (
	_ io.Writer          = &Writer{}
	_ writerutil.Storage = &Writer{}
)

func NewWriter() *Writer {
	return &Writer{
		mutex:   &sync.Mutex{},
		writers: map[string]io.Writer{},
	}
}

type Writer struct {
	mutex   *sync.Mutex
	writers map[string]io.Writer
}

func (broadcaster *Writer) Write(message []byte) (int, error) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	for _, writer := range broadcaster.writers {
		if _, err := writer.Write(message); err != nil {
			log.Println("Error writing:", err.Error())
		}
	}

	return 0, nil
}

// Add will store a writer
func (broadcaster *Writer) Add(name string, writer io.Writer) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	broadcaster.writers[name] = writer
}

// Remove removes a writer from the list
func (broadcaster *Writer) Remove(name string) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	delete(broadcaster.writers, name)
}
