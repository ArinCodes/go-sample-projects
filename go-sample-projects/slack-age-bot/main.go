package main

import(
	"context"
	"fmt"
"log"
"os"
"github.com/shomali11/slacker/v2"
"strconv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main(){
	os.Setenv("SLACK_BOT_TOKEN","xoxb-5866054697168-5844808229700-sjh14BKnztlktkchtHrE1S2o")
	os.Setenv("SLACk_APP_TOKEN","xapp-1-A05QPCULCQM-5866058451296-4b9bd890833d59ef9a6ed7fd0070d4251d8730f86f117d7b3e1b297e4068007c")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

    go printCommandEvents(bot.printCommandEvents())
    
   bot.Command("my yob is <year>" , &slacker.CommandDefinition{
	Description : "yob calculator",
	
	Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
       year := request.Param("year")
	   yob ,err := strconv.Atoi(year)
	   if err!=nil{
		println("error") 
	   }
      age := 2023-yob
	  r := fmt.Sprintf("age is %d" ,age)
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