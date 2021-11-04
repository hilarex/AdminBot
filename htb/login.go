package htb

import (
    "adminbot/config"
    
    "io/ioutil"
    "net/http"

    "encoding/json"

//  "regexp"
//  "net/url"
//  "crypto/tls"
    "bytes"
    "time"
//  "strings"
)

func StartLogin(ticker *time.Ticker){
    for {
        select {
            case <- ticker.C:
                Login()
        }
    }
}

func Login() {

/*
	proxyUrl, err := url.Parse("http://127.0.0.1:8080")
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        Proxy: http.ProxyURL(proxyUrl),
    }
*/

	client := &http.Client{
  		Timeout: time.Second * 10,
// 		Transport: tr,
	}
	
	// Request for login
    var jsonData = []byte(`{
        "email" : "`+config.Htb.Email+`",
        "password" : "`+config.Htb.Password+`",
        "remember" : true
    }`)

    req, err := http.NewRequest("POST", "https://www.hackthebox.com/api/v4/login", bytes.NewBuffer(jsonData))
    req.Header.Add("User-Agent", config.USERAGENT)
    req.Header.Set("Content-Type", "application/json; charset=UTF-8")

    resp, err := client.Do(req)
    if err != nil {
        print(err)
        return
    }
    defer resp.Body.Close()

    if resp.Status != "200 OK" {
        print(resp.Status)
        return
    }
    
 	body, _ := ioutil.ReadAll(resp.Body)
    //var htbtoken config.HtbToken

    json.Unmarshal(body, &config.HtbToken)


    //=========================================================

    /*

    // find token 
    r, _ := regexp.Compile("type=\"hidden\" name=\"_token\" value=\"(.+?)\"")
    token := r.FindStringSubmatch(string(body))
    if len(token) == 0{
        return
    }
	crsf_token := token[1]

    // Post request to login
    params := url.Values{}
	params.Set("_token", crsf_token)
	params.Set("email",  config.Htb.Email)
	params.Set("password", config.Htb.Password)
	postData := strings.NewReader(params.Encode())
    req, err = http.NewRequest("POST", "https://www.hackthebox.eu/login", postData)
	req.Header.Add("User-Agent", config.USERAGENT)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") 
	
	resp, err = client.Do(req)
	defer resp.Body.Close()
    if err != nil {
        print(err)
        return
    }
    if resp.StatusCode != 200{
    	fmt.Println("Error connecting to HTB")
        return
    }

    config.Htbcookies = jar
    */

    return
}