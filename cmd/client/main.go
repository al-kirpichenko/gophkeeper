// Package main точка входа
package main

import (
	"fmt"

	"github.com/al-kirpichenko/gophkeeper/internal/client/commands"
)

// Данные о клиенте
var (
	buildVersion string = "0.01"
	buildDate    string = "2023-12-12"
)

func main() {
	fmt.Println("Client version: " + buildVersion)
	fmt.Println("Build: " + buildDate)
	commands.Execute()
}
