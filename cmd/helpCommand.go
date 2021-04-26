package cmd

import (
	"adminbot/framework"
	"adminbot/config"
)

func HelpCommand(ctx framework.Context) {

	ctx.Reply( "```LaPiraterie's Bot"+`

Command     Options             Description                              
----------  ---------------     --------------------------------------------
ping                            Want to ping pong ?
echo        <sentence>          A simple echo command
help                            Shows this message

ippsec      <search_term>       Search through Ippsec videos
role        <role>              Add or remove Discord role

get_user    <htb_user_name>     Stalk your competitors

Restricted to team's members :
----------- -----------------   --------------------------------------------
verify      <api_token>         Verify your HTB account 
me                              Get your HTB info
leaderboard                     Get the leaderboard of the guild
prolab      <name>              Get progress of prolabs
progress 	<box_name>			Get progress of boxes
----------- -----------------   --------------------------------------------
Type `+config.Prefix+`help command for more info on a command.
You can also type `+config.Prefix+"help category for more info on a category."+"```" )
	
	return
}