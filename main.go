//Starting "hello world" iris server
//See here for how context works: https://kataras.gitbooks.io/iris/content/context.html

package main

import "fmt"
import "log"
import "os/exec"
import "io/ioutil"
import "github.com/kataras/iris"
import "github.com/iris-contrib/middleware/recovery"
import "github.com/torch2424/goSmartHome/jsonStructs"

var iftttKey = ""

func main() {

    //Initialize our recovery middleware to auto-restart on failure
    iris.Use(recovery.New(iris.Logger))

    //Get our keys


    //Initialize our api
    api := iris.New()

    api.Get("/", defaultRoute)
    api.Post("/speak", speakPost)
    api.Listen(":4000")
}

//DefaultRoute
func defaultRoute(ctx *iris.Context) {

    //Read our markdown from our views
    resMarkdown, err := ioutil.ReadFile("views/defaultRoute.md")

    if err != nil {
        ctx.Write("Could not read from views...")
        return;
    }

    //Render our markdown
   ctx.Markdown(iris.StatusOK, string(resMarkdown))
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
    response := jsonStructs.Response{fmt.Sprintf("Success! Speaking the following statement: %s", testField)}
    ctx.JSON(iris.StatusOK, response)
}
