package main

import "fmt"
import "github.com/bwmarrin/discordgo"
import "errors"
import "strings"

type PermissionHandler struct {
	Perm *Permission
}

type Permission struct {
	Groups map[string]*PermissionList
	Users  map[string]*PermissionList
}

type PermissionList struct {
	Permissions []string
}

func (p *PermissionHandler) load() (err error) {
	if p.Perm == nil {
		p.Perm = &Permission{
			Groups: make(map[string]*PermissionList),
			Users:  make(map[string]*PermissionList),
		}

		p.Perm.Groups["241843693655425027"] = &PermissionList{
			Permissions: []string{
				"MinatsuGo.command.exampleGroup",
			},
		}

		p.Perm.Users["241841671459962880"] = &PermissionList{
			Permissions: []string{
				"MinatsuGo.command.exampleUser",
			},
		}
	}
	return loadConfiguration("permissions.json", &p.Perm)
}

func (p *PermissionHandler) save() (err error) {
	if p.Perm == nil {
		return errors.New("Trying to save empty config, try load() first.")
	}
	return saveConfiguration("permissions.json", &p.Perm)
}

func (p *PermissionHandler) hasGroup(group string) bool {
	if p.Perm == nil {
		fmt.Println("Error trying to get group, can't find any groups.")
		return false
	}
	if p.Perm.Groups[group] != nil {
		return true
	}
	return false
}

func (p *PermissionHandler) hasUser(id string) bool {
	if p.Perm == nil {
		fmt.Println("Error trying to get user, can't find any users.")
		return false
	}

	if p.Perm.Users[id] != nil {
		return true
	}
	return false
}

func (p *PermissionHandler) checkPermission(permList map[string]*PermissionList, id string, ps string) bool {
	ps = strings.ToLower(ps)
	for _, pm := range permList[id].Permissions {
		pm = strings.ToLower(pm)
		if pm == ps {
			return true
		}
	}
	return false
}

func (p *PermissionHandler) hasPermission(message *discordgo.Message, ps string) bool {
	s := BOT.Session.State

	channel, err := s.Channel(message.ChannelID)
	if err != nil {
		return false
	}

	guild, err := s.Guild(channel.GuildID)
	if err != nil {
		return false
	}

	if message.Author.ID == guild.OwnerID {
		return true
	}

	if p.hasUser(message.Author.ID) {
		if p.checkPermission(p.Perm.Users, message.Author.ID, ps) {
			return true
		}
	}

	member, err := s.Member(guild.ID, message.Author.ID)

	if err != nil {
		return false
	}

	for _, role := range member.Roles {
		if p.hasGroup(role) && p.checkPermission(p.Perm.Groups, role, ps) {
			return true
		}
	}
	return false
}
