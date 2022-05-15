package MessageWay

import "net/http"

type Config struct {
	ApiKey string
	AcceptLanguage string
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
