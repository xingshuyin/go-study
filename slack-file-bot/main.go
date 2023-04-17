package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

// https://app.slack.com/client/T04KL7TUPQS/C04KL4HGYFM
// https://api.slack.com/apps/A04KKC882P8/oauth?success=1
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4666265975842-4674492672836-lX0yi0ilLMUUHpzbMSjMKUPB")
	os.Setenv("CHANNEL_ID", "C04KL4HGYFM")
	os.Setenv("SLACK_APP_TOKEN", "")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"file.txt"}
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s/n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s/n", file.Name, file.URL)
	}
}
