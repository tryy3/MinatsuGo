package main

import (
	"fmt"
	"regexp"
	"strings"
)

func getGods(settings map[string][]string) []*Gods {
	outGod := []*Gods{}

	for _, God := range GODS {
		for _, setting := range args {
			if (settings["Name"] != nil && stringInSlice(God.Name, settings["Name"])) ||
				((settings[""]))
			if God.Attack == setting ||
				God.Class == setting ||
				God.Difficulty == setting ||
				God.Name == setting ||
				God.Pantheon == setting ||
				God.Power == setting {
				outGod = append(outGod, God)
				break
			}
		}
	}
	return outGod
}

func stringInSlice(a string, l []string) bool {
	for _, b := range l {
		if b == a {
			return true
		}
	}
	return false
}
func addToMap(m map[string][]string, k string, v string) map[string][]string {
	fmt.Println(m)
	fmt.Println(k)
	fmt.Println(v)

	if m[k] == nil {
		m[k] = []string{v}
		return m
	}

	if stringInSlice(v, m[k]) {
		return m
	}

	m[k] = append(m[k], v)
	return m
}

func parseTypes(args []string) map[string][]string {
	var types = make(map[string][]string)
	args = regexp.MustCompile(", | ,|,| ").Split(strings.ToLower(strings.Join(args, " ")), -1)
	for _, t := range args {
		switch t {
		case "chinese":
			fallthrough
		case "egyptian":
			fallthrough
		case "greek":
			fallthrough
		case "hindu":
			fallthrough
		case "japanese":
			fallthrough
		case "mayan":
			fallthrough
		case "norse":
			fallthrough
		case "roman":
			addToMap(types, "Pantheon", t)

		case "melee":
			fallthrough
		case "ranged":
			addToMap(types, "Attack", t)

		case "magical":
			fallthrough
		case "physical":
			addToMap(types, "Power", t)

		case "assassin":
			fallthrough
		case "guardian":
			fallthrough
		case "hunter":
			fallthrough
		case "mage":
			fallthrough
		case "warrior":
			addToMap(types, "Class", t)

		case "easy":
			fallthrough
		case "average":
			fallthrough
		case "hard":
			addToMap(types, "Difficulty", t)

		default:
			addToMap(types, "Name", t)
		}
	}
	return types
}
