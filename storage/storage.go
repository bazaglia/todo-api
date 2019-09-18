package storage

import (
	"os"
)

// Adapter interface for storage in cloud providers
type Adapter interface {
	GetName() string
	GetObject(bucket, name string) (interface{}, error)
	GetSignedURL(bucket, name string) (string, error)
	CreateObject(bucket, name string, data []byte) error
}

// NewAdapter initialize storage adapter for a cloud provider
func NewAdapter() Adapter {
	switch name := os.Getenv("STORAGE_PROVIDER"); name {
	case "aws":
		// return NewAWSAdapter{client: client}
	case "google":
		return NewGoogleStorageAdapter()
	}

	return NewGoogleStorageAdapter()
}
