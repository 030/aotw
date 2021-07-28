package main

import (
	"path/filepath"

	"github.com/030/aotw/internal/admin"
	sasm "github.com/030/sasm/pkg/slack"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	viper.AddConfigPath(filepath.Join(home, ".aotw"))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %w", err)
	}

	a := viper.GetStringSlice("admins")
	adminOfTheWeek := admin.ThisCurrentWeek(a)

	log.Info(adminOfTheWeek)

	t := sasm.Text{Type: "mrkdwn", Text: adminOfTheWeek}
	b := []sasm.Blocks{{Type: "section", Text: &t}, {Type: "divider"}}
	d := sasm.Data{Blocks: b, Channel: "#jira", Icon: ":robot_face:", Username: "aotw"}

	slackToken := viper.GetString("slack_token")
	if slackToken == "" {
		log.Fatalf("slack_token should not be empty. Check whether these resides in: '%s'", viper.ConfigFileUsed())
	}
	d.PostMessage(slackToken)
}
