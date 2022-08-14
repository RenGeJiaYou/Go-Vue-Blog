package main

import (
	"go-vue-blog/model"
	"go-vue-blog/routers"
)

func main() {
	model.InitDB()
	routers.InitRouter()
}
