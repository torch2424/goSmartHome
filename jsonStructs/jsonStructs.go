//Define our structs for json here, and functions to decode/encode them
package jsonStructs

import "fmt"
import "io/ioutil"
import "encoding/json"

//Define our response Struct
//Message = name of variable in go
//string is type
//'json:"message"', json defines as json attribute, "message" is key of attribute
type Response struct {
    Message string `json:"message"`
}

func GetKeys() map[string]interface{} {

    //Define our keymap
    var keyMap map[string]interface{}

    //Read in our file
    keysFile, _ := ioutil.ReadFile("../keys.json")

    //Decode our json
    err := json.Unmarshal(keysFile, &keyMap)
    if err != nil {
        panic(err)
    }

    fmt.Println("Keys successfully decoded!")

    return keyMap
}
