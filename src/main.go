package main

import "github.com/wxq/metaland-blog/src/app"

//go:generate swag init -g ./main.go -o ../docs --outputTypes go,json
//go:generate swag fmt

// @title						metaland博客接口
// @version						1.0
// @description					metaland博客接口
// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	app.Start()
}
