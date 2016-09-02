//Go file for all of our route functions
package routes

import "fmt"
import "log"
import "os/exec"
import "io/ioutil"
import "github.com/kataras/iris"
import "github.com/torch2424/goSmartHome/jsonHelpers"

//DefaultRoute
func DefaultRoute(ctx *iris.Context) {

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
func SpeakPost(ctx *iris.Context) {

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
    response := jsonHelpers.Response{fmt.Sprintf("Success! Speaking the following statement: %s", testField)}
    ctx.JSON(iris.StatusOK, response)
}
