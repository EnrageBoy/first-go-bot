package telegram

import (
	"net/http"
)

type Client struct {
	host     string
	basePath string
	Client   http.Client
}

func New(host string, token string) {

}
