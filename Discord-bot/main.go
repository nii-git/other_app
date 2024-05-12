package main

import (
	"discord/config"
	"discord/infra"

	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/exp/slog"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// config取得
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println("ERROR: config file read error")
		return
	}

	// mysql接続
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))

	db, err := infra.NewDB(config, logger)

	discord, err := infra.NewDiscordClient(config, db)

	// メッセージ受信時のハンドラを登録します。
	discord.Session.AddHandler(messageCreate)
	discord.Session.AddHandler(discord.VoiceStateUpdate)

	// Discordに接続します。
	err = discord.Session.Open()
	if err != nil {
		fmt.Println("Error opening Discord connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// CTRL+Cを受信するまで待機します。
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	fmt.Println("Received termination signal. Closing Discord connection...")
	err = discord.Session.Close()
	if err != nil {
		fmt.Println("Error closing Discord connection:", err)
	}
	fmt.Println("Bot disconnected. Exiting...")
}

// func voiceStateUpdate(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
// 	// 新しいVoiceStateを取得します。
// 	newState := vsu.VoiceState

// 	// ユーザーがボイスチャンネルに参加した場合の処理
// 	if newState.ChannelID != "" && vsu.BeforeUpdate == nil {
// 		fmt.Printf("User %s joined voice channel %s\n", newState.UserID, newState.ChannelID)
// 		_, err := s.ChannelMessageSend(TEXTCHANELID, "ボイチャに入った！")
// 		if err != nil {
// 			fmt.Printf("Error: %s", err)
// 		}
// 		return
// 	}

// 	// ユーザーがボイスチャンネルから退出した場合の処理
// 	if newState.ChannelID == "" {
// 		fmt.Printf("User %s left voice channel\n", newState.UserID)
// 		_, err := s.ChannelMessageSend(TEXTCHANELID, "ボイチャから抜けた！")
// 		if err != nil {
// 			fmt.Printf("Error: %s", err)
// 		}
// 		return
// 	}

// }

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Get a message!!")

	fmt.Println(m.ChannelID)

	// メッセージがボットのものであるかを確認します。
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "" {
		// メッセージがコマンドであるかを確認します。
		_, err := s.ChannelMessageSend(m.ChannelID, "呼んだ？")
		if err != nil {
			fmt.Printf("Error: %s", err)
		}

		_, err = s.ChannelMessageSend(m.ChannelID, m.Message.Content+"って言ってた！")
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}

}
