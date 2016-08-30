//Starting "hello world" iris server
//See here for how context works: https://kataras.gitbooks.io/iris/content/context.html

package main

import "fmt"
import "log"
import "os/exec"
import "github.com/kataras/iris"
import "github.com/iris-contrib/middleware/recovery"

var iftttKey = ""

//Define our response Struct
//Message = name of variable in go
//string is type
//'json:"message"', json defines as json attribute, "message" is key of attribute
type Response struct {
    Message string `json:"message"`
}

func main() {

    //Initialize our recovery middleware to auto-restart on failure
    iris.Use(recovery.New(iris.Logger))

    //Get our keys


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

//The /speak Post. Reads the statement field from json
func speakPost(ctx *iris.Context) {

    //Get our Json values
    testField := ctx.FormValueString("statement")

    //Log the event
	speakLog := fmt.Sprintf("/speak post | Speaking the following statement: %s\n", testField)
    fmt.Printf(speakLog)

    //Run the espeak command, and catch any errors
    //exec.Command(comman, commandArguments)
    cmd := exec.Command("espeak", testField);
    err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

    //Send Okay and respond
    response := Response{fmt.Sprintf("Success! Speaking the following statement: %s", testField)}
    ctx.JSON(iris.StatusOK, response)
}
