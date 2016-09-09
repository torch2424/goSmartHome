//Starting "hello world" iris server
//See here for how context works: https://kataras.gitbooks.io/iris/content/context.html

package main

import "os"
import "fmt"
import "github.com/kataras/iris"
import "github.com/iris-contrib/middleware/recovery"
import "gopkg.in/alecthomas/kingpin.v2"
import "github.com/torch2424/goSmartHome/jsonHelpers"
import "github.com/torch2424/goSmartHome/routes"
import "github.com/torch2424/goSmartHome/banner"

//Our json map for api keys
var ApiKeys map[string]interface{}

//Command Line Parser (Kingpin) Setup
var (
	app = kingpin.New("Karen Smart Home", "A talking smart home server written in go")
	userIp = kingpin.Flag("server", "IP address to ping, including port").Short('s').Default("0.0.0.0:80").TCP()
)

func main() {

    //Start by printing our server banner
    banner.PrintBanner()

    //Print some spacing
    fmt.Println()

    //Parse our input
	kingpin.Parse()

    //Get our keys
    fmt.Println("Parsing API keys from keys.json, beep, bop, boop, beep...")
    fmt.Println()
    ApiKeys = jsonHelpers.GetKeys()

    //Print some spacing
    fmt.Println()

    //Check our keys
    checkKeys()

    //Print some spacing
    fmt.Println()

    //Initialize our recovery middleware to auto-restart on failure
    iris.Use(recovery.Handler)

    //Initialize our api and routes
    api := iris.New()
    api.Get("/", routes.DefaultRoute)
    api.Post("/speak", routes.SpeakPost)

    //Serve the app
    serverIp := *userIp
    fmt.Println(serverIp.String())
    api.Listen(serverIp.String())
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
