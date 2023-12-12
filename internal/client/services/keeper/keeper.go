package keeper

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"

	"github.com/al-kirpichenko/gophkeeper/internal/client/models"
	"github.com/al-kirpichenko/gophkeeper/internal/client/storage"
)

// KeeperService is an interface that provides methods for authentication and registration.
type KeeperService interface {
	Create(title string, secret *models.Secret) error
	Get(title string) (string, error)

	// GetClient returns the service's client.
	GetClient() *resty.Client
}

// keeperService is a concrete implementation of KeeperService.
type keeperService struct {
	client *resty.Client
}

// newConfiguredClient returns a client configured for https (if required).
func newConfiguredClient(baseURL string) *resty.Client {
	client := resty.New().SetBaseURL(baseURL)
	if strings.Contains(baseURL, "https") {
		client = client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return client
}

// NewKeeperService creates a new instance of authService with the given baseURL and returns it as an AuthService.
func NewKeeperService(baseURL string) KeeperService {
	client := newConfiguredClient(baseURL)
	return &keeperService{client: client}
}

func encodeToBase64(v *models.Secret) ([]byte, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return nil, err
	}
	encoder.Close()
	return buf.Bytes(), nil
}

func (k *keeperService) Create(title string, secret *models.Secret) error {

	r := &models.CreateResponse{}

	tokenStorage := storage.New()

	token := tokenStorage.GetToken()

	content, err := encodeToBase64(secret)

	if err != nil {
		return err
	}

	resp, err := k.client.R().
		SetResult(r).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(fmt.Sprintf(`{"title":"%s","content":"%s"}`, title, content)).
		Post("/api/secret/create")
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusCreated {
		return errors.New(r.Message)
	}
	return nil

}

// Get authenticates a user with the given username and password and returns an authentication token if successful.
func (k *keeperService) Get(title string) (string, error) {

	r := &models.SecretResponse{}

	tokenStorage := storage.New()

	token := tokenStorage.GetToken()

	resp, err := k.client.R().
		SetResult(r).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(fmt.Sprintf(`{"title":"%s"}`, title)).
		Post("/api/secret/read")
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != http.StatusOK {
		return "", err
	}

	myString := strings.ReplaceAll(string(r.Content), "{", "")
	myString = strings.ReplaceAll(myString, "}", "")

	if err != nil {
		return "", err
	}
	return myString, nil
}

// GetClient returns the service's client.
func (k *keeperService) GetClient() *resty.Client {
	return k.client
}
