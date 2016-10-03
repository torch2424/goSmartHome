//Handle sending tweets to apis
package tweeter

import "fmt"
import "time"
import "math/rand"
import "github.com/Jeffail/gabs"
import "github.com/parnurzeal/gorequest"

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
	tweets, err := gabs.ParseJSONFile("jsonFiles/tweets.json")
	if err != nil {
		panic(err)
	}

	//Create our ticker, runs six times a day (4 hours)
	tweetTicker := time.NewTicker(time.Hour * 4)

	//Function in seperate thread to tweet
	go func() {
		for _ = range tweetTicker.C {

			//~ 1/10 chance to tweet
			chance := rand.Intn(10)
			if chance != 3 {
				//Return if not 3. Golang uses continue to skip to the next iteration of for loop
				continue
			}

			//Send a request to ifttt to tweet

			//Some Spacing
			fmt.Println()

			//First construct our url
			postUrl := fmt.Sprintf("https://maker.ifttt.com/trigger/%s/with/key/%s", "send_tweet", keys.Path("ifttt").Data().(string))

			//Get our random Tweet
			tweetMap, err := tweets.ChildrenMap()
			if err != nil {
				fmt.Println("Error parsing tweet json")
				continue
			}
			mapKeys := make([]string, 0, len(tweetMap))
		    for k := range tweetMap {
		        mapKeys = append(mapKeys, k)
		    }
			tweetArray, err := tweets.Search(mapKeys[rand.Intn(len(mapKeys))]).Children()
			if err != nil {
				fmt.Println("Error parsing tweet array")
				continue
			}
			//Get the final tweet string
			reqTweet := tweetArray[rand.Intn(len(tweetArray))]

			//Log we are making the tweet
			fmt.Println("Sending tweet to IFTTT: ", reqTweet)

			//Create our response with gabs
			tweetJson := gabs.New()
			tweetJson.Set(reqTweet.Data().(string), "value1")

			//Make our request!, use the tweetJson.String() to get it to work
			// Response, Body, Err
			_, resBody, _ := gorequest.New().Post(postUrl).Set("Content-Type", "application/json").Send(tweetJson.String()).End()
			fmt.Println("IFTTT Response: ", resBody)

			//Some Spacing
			fmt.Println()
	    }
	}()
}
