package list

import (
	"fmt"
	"strings"

	"github.com/sbaildon/bot/commands"
	"github.com/sbaildon/bot/helpers"
	"github.com/bwmarrin/discordgo"
)

func init() {
	commands.AddCommand("list", &Command{})
}

type Command struct {
}

func (c *Command) Execute(session *discordgo.Session, message *discordgo.Message) {
	names := commands.CommandNames()

	response := fmt.Sprintf("Available commands include: %s\n", strings.Join(names, ", "))

	helpers.PrivateMessage(session, message.Author, response)
}


