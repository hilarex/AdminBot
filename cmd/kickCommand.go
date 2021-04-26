package cmd

import (
	"adminbot/framework"
    "adminbot/config"

	"fmt"
    "encoding/json"
    "io/ioutil"
    "strconv"
    "time"
)

func KickCommand(ctx framework.Context) {
    // TODO [!] In progress 

    if !HasRole(ctx.Discord, ctx.User.ID, "admin"){
        return
    }

    if len(ctx.Args) == 0{
        ctx.Reply("Choose a command : <days>")
        return
    }
    
    days, _ := strconv.Atoi( ctx.Args[0] )
    fmt.Println(days)
    logdata := make(map[string]int64)

    byteValue, err := ioutil.ReadFile("logdata.json")
    if err != nil{
        fmt.Println("[!] Error Logging, cannot read logdata.json")
        return
    }
    json.Unmarshal(byteValue, &logdata)

    for userid, timestamp := range logdata {
        member, err := ctx.Discord.GuildMember(config.Discord.GuildID, userid)
        if err == nil{
            ctx.Reply( fmt.Sprintf("Last message of %s was in %s", member.User.Username, time.Unix(timestamp, 0)) )    
        }
    }

    return
}