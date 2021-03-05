package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http/cookiejar"
)

// Struct for config.json
type Config struct {
	Prefix      string `json:"Prefix"`
	Htb    		ConfigHtb `json:"HTB"`
	Discord     ConfigDiscord `json:"Discord"`
}
type ConfigHtb struct {
	Email		string `json:"email"`
	Password	string `json:"password"`
	ApiToken	string `json:"api_token"`
}
type ConfigDiscord struct {
	Guild 		string `json:"guild_name"`
	Token 	    string `json:"bot_token"`
	GuildID		string `json:"guild_id"`
	Shoutbox    string `json:"shoutbox_id"`
}

// Struct for Users json file
type User struct {
	DiscordID  int 	`json:"discord_id"`
	UserID     int    `json:"user_id"`
    VIP         bool   `json:"vip"`
    Username string `json:"user_name"`
    Avatar string `json:"avatar"`
    Points string `json:"points"`
    Systems string `json:"systems"`
    Users string `json:"users"`
    Respect string `json:"respect"`
    Country string `json:"country"`
    Team string `json:"team"`
    Level string `json:"level"`
    Rank string `json:"rank"`
    Challs string `json:"challs"`
    Ownership string `json:"ownership"`
    Prolabs map[string]string `json:"prolabs"`
}

// Struct for IppSec json file
type Timestamp struct{
	Minutes 	int 	`json:"minutes"`
	Seconds 	int 	`json:"seconds"`
}
type Video struct{
	Machine 	string 	`json:"machine"`
	VideoId 	string 	`json:"videoId"`
	Timestamp 	Timestamp `json:"timestamp"`
	Line 		string 	`json:"line"`
}

// Struct for Progress json file
type Progress struct{
	Username	string   `json:"user_name"`
	Users 		[]string `json:"user_owns"`
	Roots  		[]string `json:"root_owns"`
	Challs 		[]string `json:"chall_owns"`
}


type Notifs struct{
	Success string 		`json:"success"`
	Html 	[]string 	`json:"html"`
}

// Global variables
var Prefix string
var Htb ConfigHtb
var Discord ConfigDiscord


// Variable for http.Client
const USERAGENT = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.85 Safari/537.36"
var Htbcookies *cookiejar.Jar

func init(){
	
	var conf Config
	values, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(values, &conf)
	
	Prefix = conf.Prefix
	Htb = conf.Htb
	Discord = conf.Discord
}