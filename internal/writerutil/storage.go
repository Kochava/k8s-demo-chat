package writerutil

import "io"

type Storage interface {
	Add(string, io.Writer)
	Remove(string)
}
