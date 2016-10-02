//Handle sending tweets to apis
package tweeter

import "fmt"
import "time"
import "math/rand"
import "github.com/Jeffail/gabs"

var tweetTicker *time.Ticker
var tweets *gabs.Container
var keys *gabs.Container

//Function to send a tweet to iftttKey
func BackgroundTweet(apiKeys *gabs.Container) {

	//Inform starting of tweet thread
	fmt.Println("Starting Background tweeting!");

	//Save our api keys
	keys = apiKeys

	//Get our available Tweets
	tweets, _ := gabs.ParseJSONFile("jsonFiles/tweets.json")

	//Create our ticker, runs six times a day (4 hours)
	tweetTicker := time.NewTicker(time.Hour * 4)

	//Save package variables to null, so we can use them in multi-threaded sendTweet()
	_ = tweets
	_ = tweetTicker

	go func() {
		for t := range tweetTicker.C {

			//1/10 chance to tweet
	    }
	}()
}
