package main

import (
	"fmt"

	"./Backend"
)

func main() {
	if AutoFSecBackend.SystemManager.InitSystem() {
		fmt.Println("\nSystem init OK")
		AutoFSecBackend.SystemManager.CreateNewUser("fes", "123")
		AutoFSecBackend.InitializeServer("5555")
	}
}
