package logging

import (
	"adminbot/framework"
	"adminbot/config"
	"github.com/bwmarrin/discordgo"

	"fmt"
	"encoding/json"
	"io/ioutil"
    "time"
    "sync"
)


var Mu    sync.Mutex
var Log   map[string]int64


func StartLogging(ticker *time.Ticker){
    for {
        select {
            case <- ticker.C:
                Logging()
        }
    }
}

func InitLogging(session *discordgo.Session){
	Mu.Lock()
	defer Mu.Unlock()
    _, err := ioutil.ReadFile("logdata.json")

    if err != nil{
    	guildMembers, err := session.GuildMembers(config.Discord.GuildID, "", 1000)
    	if err != nil {
    		fmt.Println("The bot needs the Server Members Intent authorization")
    	}

        Log = make(map[string]int64)

        for _, member := range guildMembers{
        	Log[member.User.ID] = 0
        }

        initdata, _ := json.Marshal(Log)
        err = ioutil.WriteFile("logdata.json", initdata, 0644)
        if err != nil{
            fmt.Println("[!] error Logging : cannot create logdata.json file")
            return
        }

        return
    }else{
        Log = make(map[string]int64)
    }
}

func Logging(){
    // We update the logdata file with the data from the Log variable, then we clear this variable
    Mu.Lock()
    defer Mu.Unlock()
  
    fileData := make(map[string]int64)
    byteValue, err := ioutil.ReadFile("logdata.json")
    if err != nil{
		fmt.Println("[!] error Logging : cannot read logdata.json file")
    	return
    }
    json.Unmarshal(byteValue, &fileData)

    newdata, _ := json.Marshal( framework.MergeMap(fileData, Log) )

    err = ioutil.WriteFile("logdata.json", newdata, 0644)
    if err != nil{
        fmt.Println("[!] error Logging : cannot write to logdata.json file")
        return
    }

    Log = make(map[string]int64)

}