package main

import (
	"linweb"
	"linweb/examples/blog/controllers"
	"linweb/interfaces"
	"log"
)

func main() {
	l := linweb.NewLinWeb()
	l.AddMiddlewares(PrintHelloMiddleware)
	l.AddControllers(&controllers.UserController{}, &controllers.BlogController{})
	err := l.Run(":4560")
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func PrintHelloMiddleware(c interfaces.IContext) {
	println("hello linweb!")
	c.Next()
	println("byebye linweb")
}
