package main

import (
	"fmt"
	"github.com/2dwrld/url-shortener/internal/config"
)

func main() {
	// TODO: init config -> cleanenv
	cfg := config.MustLoad()
	fmt.Printf("%+v\n", cfg)
	// TODO: init logger -> slog
	// TODO: init storage -> SQLite
	// TODO: init router - > "net/http"
	// TODO: run server
}
