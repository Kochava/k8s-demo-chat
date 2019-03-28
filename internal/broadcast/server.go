package broadcast

import "context"

// Server describes a common ingestion system
type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}
