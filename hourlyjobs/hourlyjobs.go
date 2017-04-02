// Go file for a ticker that will perform jobs depending on the date
// Will tick hourly
package hourlyjobs

import "fmt"
import "time"

var timeTicker *time.Ticker

func HourlyJobs() {
  //Inform starting of tweet thread
	fmt.Println("Starting Hourly Jobs!");

  //Create our ticker, runs every hour
	timeTicker := time.NewTicker(time.Minute)

  // Start a background thread
  go func() {
    for _ = range timeTicker.C {

      // Get our current time
      var month int = int(time.Now().Month())
      var hour int = int(time.Now().Hour())
      fmt.Println(month)
      fmt.Println(hour)

      // if it is 10AM on the 1st of the month

    }
  }()
}
