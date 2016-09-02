//Starting "hello world" iris server
//See here for how context works: https://kataras.gitbooks.io/iris/content/context.html

package main

import "os"
import "fmt"
import "github.com/kataras/iris"
import "github.com/iris-contrib/middleware/recovery"
import "github.com/torch2424/goSmartHome/jsonHelpers"
import "github.com/torch2424/goSmartHome/routes"

var ApiKeys map[string]interface{}

func main() {

    //Print some spacing
    fmt.Println()

    //Get our keys
    ApiKeys = jsonHelpers.GetKeys()

    //Print some spacing
    fmt.Println()

    //Check our keys
    checkKeys()

    //Print some spacing
    fmt.Println()

    //Initialize our recovery middleware to auto-restart on failure
    iris.Use(recovery.New(iris.Logger))

    //Initialize our api
    api := iris.New()

    api.Get("/", routes.DefaultRoute)
    api.Post("/speak", routes.SpeakPost)
    api.Listen(":4000")
}

//Function to check that we have all of the necessary keys
func checkKeys() {

    //Keys needed: ifttt

    //Save our keys, and err if we are missing any
    //Using type assertion from the map, as our keys will be strings
    iftttKey := ApiKeys["ifttt"].(string)
    if len(iftttKey) < 1 {
        fmt.Println("The ifttt key is blank, exiting...")
        fmt.Println()
        os.Exit(0)
    }

    //Finally Print our keys
    fmt.Printf("Ifttt key: %s\n", iftttKey)
}
