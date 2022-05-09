package client

import (
	"flag"
	"github.com/anthodev/gocord/internal/cmds"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

var (
	GuildID   = flag.String("guild", os.Getenv("GUILD_ID"), "Test guild ID. If not passed - bot registers cmds globally")
	BotToken  = flag.String("token", os.Getenv("BOT_TOKEN"), "Bot access token")
	RemoveCmd = flag.Bool("remove", true, "Remove all cmds after shutdowning or not")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	dmPermission                  = true
	defaultMemberPermission int64 = discordgo.PermissionManageServer

	_, _ = dmPermission, defaultMemberPermission

	commands        = cmds.BotCommands()
	commandHandlers = cmds.BotHandlers()
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func RunDiscordApi() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}

	log.Println("Adding the commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Fatalf("Error creating command '%v': %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer func(s *discordgo.Session) {
		err := s.Close()
		if err != nil {
			log.Fatalf("Error closing Discord session: %v", err)
		}
	}(s)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Bot is running. Press Ctrl+C to exit.")
	<-stop

	removeCommands(registeredCommands)
}

func removeCommands(registeredCommands []*discordgo.ApplicationCommand) {
	if *RemoveCmd {
		log.Println("Removing the commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Printf("Error removing command '%v': %v", v.Name, err)
			}
		}
	}
}
