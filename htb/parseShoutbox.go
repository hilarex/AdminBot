package htb

import (
    "adminbot/config"
    //"adminbot/framework"

    "github.com/bwmarrin/discordgo"
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "regexp"
    "strconv"
    //"strings"
    "time"
)

var (
    lastNotifProcess string
    categoryID string
    newHtbChannelID string
)

func StartParseShoutbox(ticker *time.Ticker, session *discordgo.Session){
    
    ParseShoutbox(session)
    for {
        select {
            case <- ticker.C:
                ParseShoutbox(session)
        }
    }
}

func ParseShoutbox(session *discordgo.Session) {
/*
function to parse shoutbox api
*/

    regexList := map[string]string{
            "box_pwn": `(?:.*)profile\/(\d+)\">(?:.*)<\/a> owned (.*) on <a(?:.*)profile\/(?:\d+)\">(.*)<\/a> <a(?:.*)`,
            "chall_pwn": `(?:.*)profile\/(\d+)\">(?:.*)<\/a> solved challenge <(?:.*)>(.*)<(?:.*)><(?:.*)> from <(?:.*)>(.*)<(?:.*)><(?:.*)`,
            //OLD "prolab_pwn": `(?:.*)profile\/(\d+)\">(?:.*)<\/a> got flag <(?:.*)>(.*)<\/span> from <(?:.*)>(.+?)<\/span> <`,
            "prolab_pwn": `(?:.*)profile\/(\d+)\">(?:.*)just got flag (.+?) from (.+?)!`,
            //OLD "new_box_incoming": `(?:.*)Get ready to spill some (?:.* blood .*! <.*>)(.*)<(?:.* available in <.*>)(.*)<(?:.*)><(?:.*)`,
            //OLD "new_box_out": `(?:.*)>(.*)<(?:.*) is mass-powering on! (?:.*)`,
            "vip_upgrade": `(?:.*)profile\/(\d+)\">(?:.*)<\/a> became a <(?:.*)><(?:.*)><(?:.*)> V.I.P <(?:.*)`,
    }



    client := &http.Client{
    		Timeout: time.Second * 10,
    }

    // Request shouts API
    req, _ := http.NewRequest("POST", "https://www.hackthebox.eu/api/shouts/get/initial/html/30?api_token="+ config.Htb.ApiToken, nil)
    req.Header.Add("User-Agent", config.USERAGENT)
    resp, err := client.Do(req)
    if err != nil{
        fmt.Println("[!] parseShoutbox, cannot do request")
        return
    }
    if resp.StatusCode != 200{
        fmt.Println("[!] parseShoutbox, error no 200")
        return
    }
    defer resp.Body.Close()

    // List of all notifs 
    var notifs config.Notifs
    body, _ := ioutil.ReadAll(resp.Body)
    _ = json.Unmarshal(body, &notifs)

    // Check if there is new notifs
    if lastNotifProcess == notifs.Html[len(notifs.Html)-1]{
        return
    }

    // Only parse new notifs
    notifs.Html = notifs.Html[getNewNotifPos(lastNotifProcess, notifs.Html):]
    lastNotifProcess = notifs.Html[len(notifs.Html)-1]

    // Get users list
    var users []config.User
    byteValue, err := ioutil.ReadFile("users.json")
    if err != nil{
        fmt.Println("[!] parseShoutbox, no users.json file")
        return
    }

    // Create map of userID : username 
    json.Unmarshal(byteValue, &users)
    var usersId = make(map[int]int)
    for _, user := range users{
        usersId[user.UserID] = user.DiscordID
    }


    // Go over the regex list on each line of new notifs
    var r *regexp.Regexp
    for _, notif := range notifs.Html{

        for typeOfNotif, reg := range regexList{
            r = regexp.MustCompile(reg)
            match := r.FindStringSubmatch(notif)

            if len(match) > 0{
                switch typeOfNotif{
                    case "box_pwn":
                        id, _ := strconv.Atoi(match[1]) 
                        if usersId[id] != 0{
                            // Get the discord ID that correspond to the HTB ID
                            member, _ := session.GuildMember(config.Discord.GuildID, strconv.Itoa(usersId[id]))
                            msg, _ := session.ChannelMessageSend(config.Discord.Shoutbox, fmt.Sprintf("👏 %v owned %v of %v !", member.Mention(), match[2], match[3]))
                            session.MessageReactionAdd(msg.ChannelID, msg.ID, "👏")
                        }
                    case "chall_pwn":
                        id, _ := strconv.Atoi(match[1]) 
                        if usersId[id] != 0{
                            member, _ := session.GuildMember(config.Discord.GuildID, strconv.Itoa(usersId[id]))
                            msg, _ := session.ChannelMessageSend(config.Discord.Shoutbox, fmt.Sprintf("👏 %v solved challenge %v from %v !", member.Mention(), match[2], match[3]))
                            session.MessageReactionAdd(msg.ChannelID, msg.ID, "👏")
                        }
                    case "prolab_pwn":
                        id, _ := strconv.Atoi(match[1]) 
                        if usersId[id] != 0{
                            member, _ := session.GuildMember(config.Discord.GuildID, strconv.Itoa(usersId[id]))
                            msg, _ := session.ChannelMessageSend(config.Discord.Shoutbox, fmt.Sprintf("🚩 %v got flag %v from %v !", member.Mention(), match[2], match[3]))
                            session.MessageReactionAdd(msg.ChannelID, msg.ID, "👏")
                        }
                    case "vip_upgrade":
                        id, _ := strconv.Atoi(match[1]) 
                        if usersId[id] != 0{
                            member, _ := session.GuildMember(config.Discord.GuildID, strconv.Itoa(usersId[id]))
                            msg, _ := session.ChannelMessageSend(config.Discord.Shoutbox, fmt.Sprintf("🍾 %v became VIP ! Take out the champagne 🥂", member.Mention()))
                            session.MessageReactionAdd(msg.ChannelID, msg.ID, "🍾")
                        }
                    /*case "new_box_incoming":
                        timeRemaining := strings.Split(match[2], ":")
                        // match hours and minutes to prevent spam
                        if framework.IsInSlice(timeRemaining[0], []string{"19", "05", "00"}){
                            if framework.IsInSlice(timeRemaining[1], []string{"59"}){
                                manageHtbChannel(session, strings.ToLower(match[1]))
                                session.ChannelMessageSend(newHtbChannelID, fmt.Sprintf("⏱ box %v is coming in %v ! ⏱", match[1], match[2]))
                            }
                        }
                    case "new_box_out":
                        session.ChannelMessageSend(newHtbChannelID, fmt.Sprintf("🚨 new box %v is live ! 🚨\nWill you get first blood ?", match[1]))
                    */
                    default:
                        session.ChannelMessageSend(config.Discord.Shoutbox, fmt.Sprintf(typeOfNotif,"=",match[1:]))
                }
            }
        }
    }
}


func getNewNotifPos(last string, notifs []string) int{
    for i, notif := range notifs{
        if notif == last{
            return i + 1
        }
    }
    return 0
}

/*
func manageHtbChannel(session *discordgo.Session, boxName string){
    var channelsInCategory []*discordgo.Channel
    boxName = strings.ToLower(boxName)
    
    channels, _ := session.GuildChannels(config.Discord.GuildID)
    
    // Parse channels to get the category ID if it is not set
    if categoryID == ""{
        for _, c := range channels{
            // category type is 4
            if c.Type == 4 {
                if strings.ToLower(c.Name) == "htb" {
                    categoryID = c.ID
                    break
                }
            }
        }
    }

    // Get list of channels in this category
    for _, c := range channels{
        if c.ParentID == categoryID && c.Name == boxName {
            newHtbChannelID = c.ID
            return
        }
        if c.ParentID == categoryID{
            channelsInCategory = append(channelsInCategory, c)
        }
    }


    new := discordgo.GuildChannelCreateData{
        Name     : boxName,
        Type     : 0,
        ParentID : categoryID,
    }

    // Delete old channel
    session.ChannelDelete(channelsInCategory[len(channelsInCategory)-2].ID)
    // Create new channel
    channel, _ := session.GuildChannelCreateComplex(config.Discord.GuildID, new)
    newHtbChannelID = channel.ID

    return
}
*/