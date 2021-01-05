package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"

	"msteams"
)

const incomingWebhookURL = ""

func main() {
	msclient := msteams.NewWebhook(incomingWebhookURL)

	message := msteams.NewMessage([]msteams.Block{
		msteams.NewContainer(

			msteams.NewTextBlock(
				msteams.TextBlockTextSizeExtraLarge,
				msteams.TextBlockTextWeightBolder,
				fmt.Sprintf("Title"))).
			SetSpacing(msteams.SpacingSmall),

		msteams.NewContainer(
			msteams.NewTextBlock(
				msteams.TextBlockTextSizeSmall,
				msteams.TextBlockTextWeightBolder,
				"TextBlock_1"),
			msteams.NewTextBlock(
				msteams.TextBlockTextSizeSmall,
				msteams.TextBlockTextWeightBolder,
				"TextBlock_1")).
			SetSpacing(msteams.SpacingSmall).
			SeparatorEnable(),
	})

	messageData, _ := jsoniter.Marshal(message)
	fmt.Println("messageData: " + string(messageData))

	res, err := msclient.SendMessage(&message)
	if err != nil {
		return
	}

	fmt.Println("Webhook response")
	fmt.Printf("code: %d\n", res.StatusCode)
	fmt.Printf("body: %s\n" + res.Body)

	return
}
