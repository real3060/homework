package api

import (
	"demo/password/config"
	"fmt"
)

func Api() {
	config := config.NewConfig()
	fmt.Println(config)
}
