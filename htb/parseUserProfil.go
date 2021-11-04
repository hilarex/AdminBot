package htb

import (
    "adminbot/config"

    "net/http"
    "io/ioutil"
    "time"
    "encoding/json"
//    "regexp"
    "strconv"
    "fmt"
//    "strings"
    "sync"
)

func ParseUserProfil(wg *sync.WaitGroup, user *config.User, progress *config.Progress){
/*
Get information about an user by scrapping his HTB profil
*/
    if wg != nil{
        defer wg.Done()    
    }
    

    client := &http.Client{
        Timeout: time.Second * 10,
    }

    // Get request for userId
    req, err := http.NewRequest("GET", "https://www.hackthebox.com/api/v4/user/profile/basic/"+strconv.Itoa(user.UserID), nil)
    req.Header.Add("User-Agent", config.USERAGENT)
    req.Header.Add("Authorization", "Bearer "+config.HtbToken.Message.AccessToken)
    resp, err := client.Do(req)
    if err != nil {
        print(err)
        return
    }
    defer resp.Body.Close()
    
    // If valid
    if resp.StatusCode != 200{
        fmt.Println("[!] error getting info on: "+strconv.Itoa(user.UserID))
        return
    }

    // Read response
    body, _ := ioutil.ReadAll(resp.Body)

    var infos map[string](map[string](interface{}))
    json.Unmarshal([]byte(body), &infos)

    if  fmt.Sprintf("%s", infos["profile"]["isVip"]) == "true" {
        user.VIP = true
    } else {
        user.VIP = false
    }
    user.Username = fmt.Sprintf("%s", infos["profile"]["name"])
    user.Ownership = fmt.Sprintf("%s", infos["profile"]["rank_ownership"])
    user.Avatar = fmt.Sprintf("%s", infos["profile"]["avatar"])
    user.Points = fmt.Sprintf("%s", infos["profile"]["points"])
    user.Systems = fmt.Sprintf("%s", infos["profile"]["system_owns"])
    user.Users = fmt.Sprintf("%s", infos["profile"]["user_owns"])
    user.Respect = fmt.Sprintf("%s", infos["profile"]["respects"])
    user.Country = fmt.Sprintf("%s", infos["profile"]["country_name"])
    user.Level = fmt.Sprintf("%s", infos["profile"]["rank"])
    user.Rank = fmt.Sprintf("%s", infos["profile"]["ranking"])

    if(infos["profile"]["team"] != nil){
        team, _ := infos["profile"]["team"].(map[string]interface{})
        user.Team = fmt.Sprintf("%s", team["name"])
    }else{
        user.Team = ""
    }
    

  //  user.Team = team["name"]
    // TODO PROGRESS
/*
    if progress != nil{
        r = regexp.MustCompile(username+` owned (root|user) (?:.*?)(?:\d)">(.*?)<\/a>`)
        matches := r.FindAllStringSubmatch(html, -1)
        progress.Username = username
        progress.Users = nil
        progress.Roots = nil
        progress.Challs = nil
        for _, match := range matches{
            switch string(match[1]){
                case "root": 
                    progress.Roots = append(progress.Roots, strings.ToLower(match[2]))
                case "user":
                    progress.Users = append(progress.Users, strings.ToLower(match[2]))
            }
        }
    }
*/
    return 
}