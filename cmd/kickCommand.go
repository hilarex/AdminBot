package cmd

import (
	"adminbot/framework"
    "adminbot/config"

	"fmt"
    "encoding/json"
    "io/ioutil"
    "strconv"
    "time"
    "strings"
)

func KickCommand(ctx framework.Context) {
    // TODO [!] In progress 

    if !HasRole(ctx.Discord, ctx.User.ID, "admin"){
        return
    }

    if len(ctx.Args) == 0{
        ctx.Reply("Choose number of days : <days>")
        return
    }
    
    days, _ := strconv.Atoi( ctx.Args[0] )

    //failsafe
    if days < 60{
        ctx.Reply("Numbers of days should be more than 60")
        return   
    }

    dryrun := true
    if len(ctx.Args) > 1{
        if strings.Contains(ctx.Args[1], "-confirm"){
            dryrun = false
        }
    }   

    logdata := make(map[string]int64)
    byteValue, err := ioutil.ReadFile("logdata.json")
    if err != nil{
        fmt.Println("[!] Error Logging, cannot read logdata.json")
        return
    }
    json.Unmarshal(byteValue, &logdata)

    for userid, timestamp := range logdata {

        // check number of days since last message
        if timestamp < (time.Now().Unix() - int64(days * 24*60*60) ){
            member, err := ctx.Discord.GuildMember(config.Discord.GuildID, userid)
            
            if !IsMemberOfTeam(ctx.Discord, userid) && !HasRole(ctx.Discord, userid, "bot"){
                if err == nil{
                    // kicking user
                    ctx.Reply( fmt.Sprintf("Kicking user %s (last message was in %s)", member.User.Username, time.Unix(timestamp, 0)) )
                    if dryrun == false{
                        err = ctx.Discord.GuildMemberDeleteWithReason(config.Discord.GuildID, userid, "inactivity")    
                        if err != nil{
                            fmt.Println("Error kick : ", err)
                            ctx.Reply( fmt.Sprintf("Error kicking user %s ...", member.User.Username) )
                            return
                        }
                    }
                }
            }
        }
    }

    return
}