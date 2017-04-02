// Go file for a ticker that will perform jobs depending on the date
// Will tick hourly
package hourlyjobs

import "fmt"
import "time"
import "os/exec"

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
			//Month 1-12
			// Hour 1 - 24 (army time)
      //var month int = int(time.Now().Month())
			var dayOfMonth int = int(time.Now().Day())
      var hour int = int(time.Now().Hour())

      // if it is 10AM on the 1st of the month
			if hour == 10 && dayOfMonth == 1 {
				fmt.Println("It's the 1st of tha month!!!")
				//Play 1st of the month
				exec.Command("aplay", "assets/1stOfThaMonth.mp3")
			}
    }
  }()
}
