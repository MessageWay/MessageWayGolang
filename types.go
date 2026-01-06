package MessageWay

import (
	"net/http"
	"time"
)

type Config struct {
	ApiKey         string
	AcceptLanguage string
	Timeout        time.Duration
}

type App struct {
	config Config
	client *http.Client
}

type Request interface {
	validate() error
}

type Response interface {
	ToString() string
}
