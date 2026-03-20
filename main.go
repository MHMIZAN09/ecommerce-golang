package main

import "ecommerce/cmd"

func main() {
	// cnf := config.GetConfig()
	// println("Version:", cnf.Version)
	// println("Service Name:", cnf.ServiceName)
	// println("HTTP Port:", cnf.HttpPort)

	cmd.Server()
}
