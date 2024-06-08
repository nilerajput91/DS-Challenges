/* in Go is like a guard that ensures a specific action is performed only once, no matter how many times it's called. It's useful for one-time initializations or tasks that should happen exactly once in your program.

Real-life Example: Server Configurations */

package main

import (
	"fmt"
	"sync"
)

func main() {
	config1 := getServerConfig()
	config2 := getServerConfig()

	fmt.Println("Config1:", config1)
	fmt.Println("Config2:", config2)

}

type Configurations struct {
	Port    int
	IsDebug bool
}

var (
	serverConfig Configurations
	once         sync.Once
)

func initalizeServerConfig() {
	fmt.Println("intializing server configuration..")
	serverConfig = Configurations{
		Port:    8080,
		IsDebug: false,
	}
}

func getServerConfig() Configurations {
	once.Do(initalizeServerConfig)
	return serverConfig
}
