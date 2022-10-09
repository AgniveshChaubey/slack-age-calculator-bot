package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4196648929089-4186528606788-Pf3dBbAtaBMp3jhsojAb13VO")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A044ZDXRBM5-4184036007234-319de33a8fa894a03befb279fca925828706abfecabaee136651e2a3a7b3cab6")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	//writing go routines
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil{
				fmt.Println("error")
			}
			age := 2022-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}


func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		//these steps are optional as per need
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
		
	}
}