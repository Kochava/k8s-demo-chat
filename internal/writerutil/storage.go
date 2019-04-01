package writerutil

import "io"

// Storage defines a storage mechanism for io.Writers
type Storage interface {
	Add(string, io.Writer)
	Remove(string)
}
