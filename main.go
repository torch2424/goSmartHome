//Starting "hello world" iris server
//See here for how context works: https://kataras.gitbooks.io/iris/content/context.html

package main

import "github.com/kataras/iris"

func main() {
    //Initialize our api
    api := iris.New()

    api.Get("/", defaultRoute)
    api.Get("/hi", hiGet)
    api.Post("/speak", speakPost)
    api.Listen(":4000")
}

//DefaultRoute
func defaultRoute(ctx *iris.Context) {
   ctx.Write("Welcome to the torch2424 goSmartHome!")
}

//The /hi route GET
func hiGet(ctx *iris.Context){
   ctx.Write("Hi %s", "iris")
}

//The /speak Post
func speakPost(ctx *iris.Context) {
    testField := ctx.FormValueString("test")
	// myDb.InsertUser(...)
	println(testField)
	println("Post from /speak")
}
