# AdminBot

AdminBot is a discord bot to manage your discord server with some functionnality for HackTheBox users. It can notify you and your friends when you flag a challenge or own a new box.


In the shoutbox, display :

    The pwn user / root of the boxes and challenges.
    The pwn flags of the prolabs and fortress
    If a member gets a VIP pass

In the htb category, display :

    How long before the new box arrives
    When the box arrived
    
Other :
    Unlike Hack The Box's Discord server bot, this bot automatically updates your rank.
    It automatically manage channel when a new box is out.


## Install

You need a config.json in the config directory like this
```
{
  "Prefix" : ">",
  "HTB" : {
    "email" : "<htb email>",
    "password" : "<htb password>",
    "api_token" : "<htb api token>"
  },

  "discord" : {
    "guild_name" : "<name of your guild>",
    "bot_token" : "<token of your bot>",
    "guild_id" : "<id of your guild>",
    "shoutbox_id" : "<id of your shoutbox channel>"
  }
}
```

Then create users.json, ippsec.json, progress.json, challs.json and boxes.json files

## Channels gestion

You need to create a "htb" category on your discord server with two channels. The bot will then automatically manage theses channels like this :

    - when a new htb box is comming, it will create a new channel for this box in the "htb" category
    - the bot will delete the penultimate (so there is always two boxes channels in this category)
    - the bot will send a countdown and a notification when the new box is out
