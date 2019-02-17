package main

import ("fmt";"net";"bufio";"os";"encoding/json";"io/ioutil")

var (
	ConfigFile	= "config.json"
	Address		= ""
	Port		= ""
	Protocol	= ""
	c 			Config
)



// struct for loading configuration information
type Config struct {

    Address             string
    Port	            string
    Protocol			string
}

func init() {
 // open our config and parse
   confBytes, e := ioutil.ReadFile(ConfigFile)
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(1)
    }

    if e := json.Unmarshal(confBytes, &c); e != nil {
        fmt.Println(e.Error())
        os.Exit(2)
    }

}

func main() {

// connect to the server
b, err := net.Dial(c.Protocol, c.Address + ":" + c.Port)
if err != nil {
	fmt.Println(err)
	return
	}

//send the message
reader := bufio.NewReader(os.Stdin)
fmt.Println("Text to send:")
msg, _:= reader.ReadString('\n')

//send to socket
fmt.Fprintf(b, msg + "\n")
// close the connection

// listen for reply
    message, _ := bufio.NewReader(b).ReadString('\n')
    fmt.Print("Message from server: "+message)
}

