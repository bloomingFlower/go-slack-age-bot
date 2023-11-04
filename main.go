package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

/*
명령어 이벤트를 출력하는 함수입니다.
명령어 이벤트는 사용자가 명령어를 입력할 때마다 발생합니다.
이 함수는 명령어 이벤트를 받아서 출력합니다.
*/
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

// main 함수는 슬랙 봇을 생성하고, 명령어를 정의하고, 이벤트를 출력하며, 슬랙 클라이언트를 시작합니다.
func main() {
	/*
		1. 슬랙 클라이언트를 생성합니다.
		2. 명령어를 정의합니다.
		3. 명령어 이벤트를 출력합니다.
		4. 슬랙 클라이언트를 시작합니다.
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	slackAppToken := os.Getenv("SLACK_APP_TOKEN")
	//이 함수는 슬랙 클라이언트를 생성하는데 사용됩니다.
	bot := slacker.NewClient(slackBotToken, slackAppToken)
	// 여기서는 명령어를 정의합니다.
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2020 - yob + 1
			r := fmt.Sprintf("your age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)

	}
}
