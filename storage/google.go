package storage

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	googleStorage "cloud.google.com/go/storage"
)

// GoogleStorageAdapter storage adapter for GCP
type GoogleStorageAdapter struct {
	Name        string
	client      *googleStorage.Client
	credentials *googleCredentials
}

type googleCredentials struct {
	ProjectID    string `json:"project_id"`
	PrivateKeyID string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientID     string `json:"client_id"`
}

// GetName provider name
func (s GoogleStorageAdapter) GetName() string {
	return "google"
}

// GetObject get object given a key
func (s GoogleStorageAdapter) GetObject(bucketName, objectName string) (interface{}, error) {
	return nil, errors.New("Not implemented")
}

// GetSignedURL get signed URL for an object
func (s GoogleStorageAdapter) GetSignedURL(bucketName, objectName string) (string, error) {
	return googleStorage.SignedURL(bucketName, objectName, &googleStorage.SignedURLOptions{
		GoogleAccessID: s.credentials.ClientEmail,
		PrivateKey:     []byte(s.credentials.PrivateKey),
		Method:         "GET",
		Expires:        time.Now().Add(24 * time.Hour),
	})
}

// CreateObject insert a new object on GCP bucket
func (s GoogleStorageAdapter) CreateObject(bucketName, objectName string, data []byte) error {
	bucket := s.client.Bucket(bucketName)
	object := bucket.Object(objectName)
	writter := object.NewWriter(context.Background())

	if _, err := writter.Write(data); err != nil {
		return err
	}

	if err := writter.Close(); err != nil {
		return err
	}

	return nil
}

// NewGoogleStorageAdapter initialize GCP storage adapter
func NewGoogleStorageAdapter() *GoogleStorageAdapter {
	credentials := &googleCredentials{}

	env, err := ioutil.ReadFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err == nil {
		if err := json.Unmarshal(env, credentials); err != nil {
			log.Println(err)
			return &GoogleStorageAdapter{}
		}
	}

	client, err := googleStorage.NewClient(context.Background())
	if err != nil {
		log.Println(err)
		return &GoogleStorageAdapter{}
	}

	return &GoogleStorageAdapter{
		client:      client,
		credentials: credentials,
	}
}
