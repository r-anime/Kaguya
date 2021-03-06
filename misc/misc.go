package misc

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"

	"github.com/r-anime/Kaguya/config"
	//"../config"
)

// File for misc. functions, commands and variables

var (
	OptinAbovePosition int
	OptinUnderPosition int
	GlobalMutex	sync.Mutex
)

// HasPermissions sees if a user has elevated permissions. By Kagumi
func HasPermissions(m *discordgo.Member) bool {
	for _, r := range m.Roles {
		for _, goodRole := range config.CommandRoles {
			if r == goodRole {
				return true
			}
		}
	}
	return false
}

// Mentions channel by *discordgo.Channel. By Kagumi
func ChMention(ch *discordgo.Channel) string {
	return fmt.Sprintf("<#%s>", ch.ID)
}

// Mentions channel by channel ID. By Kagumi
func ChMentionID(channelID string) string {
	return fmt.Sprintf("<#%s>", channelID)
}

// Returns a string that shows where the error occured exactly
func ErrorLocation(err error) string {
	_, file, line, _ := runtime.Caller(1)
	errorLocation := fmt.Sprintf("Error is in file [%v] near line %v", file, line)
	return errorLocation
}

// Sends error message to channel command is in. If that throws an error send error message to bot log channel
func CommandErrorHandler(s *discordgo.Session, m *discordgo.Message, err error) {
	_, err = s.ChannelMessageSend(m.ChannelID, err.Error())
	if err != nil {
		_, _ = s.ChannelMessageSend(config.BotLogID, err.Error())
	}
}

// Resolves a userID from a userID, Mention or username#discrim
func GetUserID(s *discordgo.Session, m *discordgo.Message, messageSlice []string) (string, error) {

	var err 	error

	if len(messageSlice) < 2 {
		err = fmt.Errorf("Error: No @user or userID detected.")
		return "", err
	}

	// Pulls the userID from the second parameter
	userID := messageSlice[1]

	// Handles "me" string on whois
	if strings.ToLower(userID) == "me" {
		userID = m.Author.ID
	}

	// Trims fluff if it was a mention. Otherwise check if it's a correct user ID
	if strings.Contains(messageSlice[1], "<@") {
		userID = strings.TrimPrefix(userID, "<@")
		userID = strings.TrimPrefix(userID, "!")
		userID = strings.TrimSuffix(userID, ">")
	}
	_, err = strconv.ParseInt(userID, 10, 64)
	if len(userID) < 17 || err != nil {
		err = fmt.Errorf("Error: Invalid user.")
		return userID, err
	}
	return userID, err
}

// Print fluff message on bot ping
func OnBotPing(s *discordgo.Session, m *discordgo.MessageCreate) {
	if (m.Content == fmt.Sprintf("<@%v>", s.State.User.ID) || m.Content == fmt.Sprintf("<@!%v>", s.State.User.ID)) && m.Author.ID == "128312718779219968" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Ah, creator. Please have a nice day.")
		if err != nil {
			_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
			if err != nil {
				return
			}
			return
		}
		return
	}
	if (m.Content == fmt.Sprintf("<@%v>", s.State.User.ID) || m.Content == fmt.Sprintf("<@!%v>", s.State.User.ID)) && m.Author.ID == "66207186417627136" {
		randomNum := rand.Intn(5)
		if randomNum == 0 {
			_, err := s.ChannelMessageSend(m.ChannelID, "Scum of the earth.")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
		if randomNum == 1 {
			_, err := s.ChannelMessageSend(m.ChannelID, "Don't touch  me! I might get something.")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
		if randomNum == 2 {
			_, err := s.ChannelMessageSend(m.ChannelID, "How disgusting of a person can you be.")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
		if randomNum == 3 {
			_, err := s.ChannelMessageSend(m.ChannelID, "Could you please leave the room? All the good air is running away from you.")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
		if randomNum == 4 {
			_, err := s.ChannelMessageSend(m.ChannelID, "I'm in a good mood today. I will allow you to be my chair.")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
		return
	}
	if m.Content == fmt.Sprintf("<@%v>", s.State.User.ID) || m.Content == fmt.Sprintf("<@!%v>", s.State.User.ID) {
		randomNum := rand.Intn(5)
		if randomNum == 0 {
			_, err := s.ChannelMessageSend(m.ChannelID, "You dare address me? How... cute...")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
		if randomNum == 1 {
			_, err := s.ChannelMessageSend(m.ChannelID, "Ugh, you're such a pig. Err, I mean, gokigenyou.")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
		if randomNum == 2 {
			_, err := s.ChannelMessageSend(m.ChannelID, "The life of a plebeian must be so hard. Here, have 500 yen.")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
		if randomNum == 3 {
			_, err := s.ChannelMessageSend(m.ChannelID, "Dogs bark and are loyal. Cats laze around but are regal. You buzz around, accomplishing nothing. Just like an insect.")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
		if randomNum == 4 {
			_, err := s.ChannelMessageSend(m.ChannelID, "Is pinging me really the most fun thing you can do right now? You have my condolences.")
			if err != nil {
				_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
				if err != nil {
					return
				}
				return
			}
			return
		}
	}
}

// Discord Playing status
func StatusReady(s *discordgo.Session, e *discordgo.Ready) {
	err := s.UpdateStatus(0, "Love is War")
	if err != nil {
		_, err = s.ChannelMessageSend(config.BotLogID, err.Error()+"\n"+ErrorLocation(err))
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Parses a string for a channel and returns its ID and name
func ChannelParser(s *discordgo.Session, channel string) (string, string) {
	var (
		channelID 	string
		channelName string
		flag		bool
	)

	// If it's a channel ping remove <# and > from it to get the channel ID
	if strings.Contains(channel, "#") {
		channelID = strings.TrimPrefix(channel,"<#")
		channelID = strings.TrimSuffix(channelID,">")
	}

	// Check if it's an ID by length and save the ID if so
	_, err := strconv.Atoi(channel)
	if len(channel) >= 17 && err == nil {
		channelID = channel
	}

	// Find the channelID if it doesn't exists via channel name, else find the channel name
	channels, err := s.GuildChannels(config.ServerID)
	if err != nil {
		_, err = s.ChannelMessageSend(config.BotLogID, err.Error() + "\n" + ErrorLocation(err))
		if err != nil {
			return channelID, channelName
		}
		return channelID, channelName
	}
	for _, cha := range channels {
		if channelID == "" {
			if strings.ToLower(cha.Name) == strings.ToLower(channel) {
				channelID = cha.ID
				channelName = cha.Name
				flag = true
				break
			}
		}
		if cha.ID == channelID {
			channelName = cha.Name
			flag = true
			break
		}
	}

	if !flag {
		return "", ""
	}

	return channelID, channelName
}

// Parses a string for a category and returns its ID and name
func CategoryParser(s *discordgo.Session, category string) (string, string) {
	var (
		categoryID 		string
		categoryName 	string
		flag			bool
	)

	// Check if it's an ID by length and save the ID if so
	_, err := strconv.Atoi(category)
	if len(category) >= 17 && err == nil {
		categoryID = category
	}

	// Find the categoryID if it doesn't exists via category name, else find the category name
	channels, err := s.GuildChannels(config.ServerID)
	if err != nil {
		_, err = s.ChannelMessageSend(config.BotLogID, err.Error() + "\n" + ErrorLocation(err))
		if err != nil {
			return categoryID, categoryID
		}
		return categoryID, categoryName
	}
	for _, cha := range channels {
		if categoryID == "" {
			if cha.Type != discordgo.ChannelTypeGuildCategory {
				continue
			}
			if strings.ToLower(cha.Name) == strings.ToLower(category) {
				categoryID = cha.ID
				categoryName = cha.Name
				flag = true
				break
			}
		}
		if cha.ID == categoryID {
			categoryName = cha.Name
			flag = true
			break
		}
	}

	if !flag {
		return "", ""
	}

	return categoryID, categoryName
}