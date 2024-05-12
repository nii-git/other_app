package infra

import (
	"discord/config"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	db      *DB
	config  *config.Config
	Session *discordgo.Session
}

func NewDiscordClient(config *config.Config, db *DB) (*Discord, error) {
	ses, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return nil, err
	}
	d := Discord{
		db:      db,
		config:  config,
		Session: ses,
	}
	return &d, nil
}

func (d Discord) VoiceStateUpdate(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
	// 新しいVoiceStateを取得します。
	newState := vsu.VoiceState

	mention := fmt.Sprintf("<@%s>", newState.UserID)

	// ユーザーがボイスチャンネルに参加した場合の処理
	if newState.ChannelID != "" && vsu.BeforeUpdate == nil {
		fmt.Printf("User %s joined voice channel %s\n", newState.UserID, newState.ChannelID)
		_, err := s.ChannelMessageSend(d.config.TextChannelID, mention+"がボイチャに入った！")
		if err != nil {
			fmt.Printf("Error: %s", err)
		}

		err = d.db.InsertTransaction(newState.UserID)
		if err != nil {
			_, err := s.ChannelMessageSend(d.config.TextChannelID, fmt.Sprintf("DBエラーだよ！%s", err.Error()))
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
		}

		return
	}

	// ユーザーがボイスチャンネルから退出した場合の処理
	if newState.ChannelID == "" {
		fmt.Printf("User %s left voice channel\n", newState.UserID)
		_, err := s.ChannelMessageSend(d.config.TextChannelID, mention+"がボイチャから抜けた！")
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		transaction, err := d.db.GetTransaction(newState.UserID)
		if err != nil {
			fmt.Printf("Error: %s", err)
			_, _ = s.ChannelMessageSend(d.config.TextChannelID, fmt.Sprintf("DBエラーだよ！%s", err.Error()))
			return
		}
		err = d.db.DeleteTransaction(newState.UserID)
		if err != nil {
			fmt.Printf("Error: %s", err)
			_, _ = s.ChannelMessageSend(d.config.TextChannelID, fmt.Sprintf("DBエラーだよ！%s", err.Error()))
			return
		}

		studyTime, err := d.db.GetStudyTime(newState.UserID)
		if err != nil {
			fmt.Printf("Error: %s", err)
			_, _ = s.ChannelMessageSend(d.config.TextChannelID, fmt.Sprintf("DBエラーだよ！%s", err.Error()))
			return
		}

		var allStudyTimeMinutes int
		currentStudyTimeMinutes := int(time.Since(transaction.CreatedAt).Minutes())
		if studyTime == nil {
			allStudyTimeMinutes = currentStudyTimeMinutes
		} else {
			allStudyTimeMinutes += studyTime.StudyTimeMinutes + currentStudyTimeMinutes
		}

		err = d.db.UpsertStudyTime(newState.UserID, allStudyTimeMinutes)
		if err != nil {
			fmt.Printf("Error: %s", err)
			_, _ = s.ChannelMessageSend(d.config.TextChannelID, fmt.Sprintf("DBエラーだよ！%s", err.Error()))
			return
		}

		_, _ = s.ChannelMessageSend(d.config.TextChannelID, fmt.Sprintf(
			"勉強時間は%d分で、総勉強時間は%.1f時間になったよ！\nおつかれ様でした！",
			currentStudyTimeMinutes, float64(allStudyTimeMinutes)/60))

		return
	}

}
