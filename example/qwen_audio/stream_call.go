package main

import (
	"context"
	"fmt"
	"os"

	"github.com/23233/dashscopego"
	"github.com/23233/dashscopego/qwen"
)

func main() {
	model := qwen.QwenAudioTurbo
	token := os.Getenv("DASHSCOPE_API_KEY")

	if token == "" {
		panic("token is empty")
	}

	cli := dashscopego.NewTongyiClient(model, token)

	sysContent := qwen.AudioContentList{
		{
			Text: "You are a helpful assistant.",
		},
	}
	userContent := qwen.AudioContentList{
		{
			Text: "该段对话表达了什么观点? 详细分析该讲话者的语气,展现出什么样的情绪", //nolint:gosmopolitan
		},
		{
			// 使用本地音频文件
			// Audio: "file:///Users/xxx/Desktop/hello_world_female2.wav",
			// 官方文档中的例子
			Audio: "https://dashscope.oss-cn-beijing.aliyuncs.com/audios/2channel_16K.wav",
		},
	}

	input := dashscopego.AudioInput{
		Messages: []dashscopego.AudioMessage{
			{Role: qwen.RoleSystem, Content: &sysContent},
			{Role: qwen.RoleUser, Content: &userContent},
		},
	}

	// callback function:  print stream result
	streamCallbackFn := func(_ context.Context, chunk []byte) error {
		fmt.Print(string(chunk)) //nolint:all
		return nil
	}
	req := &dashscopego.AudioRequest{
		Input:       input,
		StreamingFn: streamCallbackFn,
	}

	ctx := context.TODO()
	resp, err := cli.CreateAudioCompletion(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nnon-stream result: ")                           //nolint:all
	fmt.Println(resp.Output.Choices[0].Message.Content.ToString()) //nolint:all
}
