package main

import (
    "money-send-api/infrastructure"
)

// @title Send Money API
// @version 1.0
// @description This api is send money to user.
// @host localhost:1323
// @BasePath /
//@securityDefinitions.apikey Bearer
//@in header
//@name Authorization
func main() {
    infrastructure.Init()
}
