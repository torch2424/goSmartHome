//Simply banner components to be printed
package banner

import "fmt"
import "io/ioutil"

//Functions to print our banner components
func PrintBanner() {

    //Read in our files
    bannerImage, _ := ioutil.ReadFile("banner/bannerImage.txt")

    bannerText, _ := ioutil.ReadFile("banner/bannerText.txt")

    //Print the ascii
    fmt.Println()
    fmt.Println(string(bannerImage))
    fmt.Println()
    fmt.Println(string(bannerText))
}
