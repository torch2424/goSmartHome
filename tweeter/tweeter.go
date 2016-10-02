//Handle sending tweets to apis
package tweeter

import "fmt"
import "github.com/Jeffail/gabs"

//Function to send a tweet to iftttKey
func BackgroundTweet() {

	tweets, _ := gabs.ParseJSONFile("jsonFiles/tweets.json")
	fmt.Println(tweets)
	fmt.Println("hello!");
}
