package main

import (
	"fmt"
	"os"

	"github.com/lisn-rocks/lisn/configs"
)

func init() {
	configs.Init()
}

func main() {
	writeDbEnv()
}

func writeDbEnv() {
	file, err := os.Create("db.env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file,
		`
POSTGRES_USER="%s"
POSTGRES_PASSWORD="%s"
POSTGRES_DB="%s"
		`,
		configs.Database.User,
		configs.Database.Password,
		configs.Database.Name,
	)
}
