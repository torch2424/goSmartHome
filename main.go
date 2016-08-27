//Starting "hello world" iris server

package main

import "github.com/kataras/iris"

func main() {
    api := iris.New()
    api.Get("/hi", hi)
    api.Listen(":4000")
}

func hi(ctx *iris.Context){
   ctx.Write("Hi %s", "iris")
}
