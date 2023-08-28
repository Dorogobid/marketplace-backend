package main

import "github.com/Dorogobid/marketplace-backend/cmd"

//	@title						Marketplace application API
//	@version					0.0.1
//	@host						api.sellspot.com.ua
//	@description				Swagger documentation for Marketplace application API
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						X-API-KEY
//	@description				Authorization API key
//	@BasePath					/

func main() {
	cmd.Execute()
}
