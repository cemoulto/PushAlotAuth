package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type GConfig struct {
	Token   string
	Watches []WatchFile
}

type WatchFile struct {
	Path         string
	Banner       string
	TriggerWords []string
}

func GetDefaultConfig() GConfig {
	var tfg GConfig
	tfg.Token = "Fillmein"
	tfg.Watches = make([]WatchFile, 0)
	defaultwatch := WatchFile{
		Path: "/var/log/auth.log",
		TriggerWords: []string{
			"Accepted publickey",
			"Accepted password",
		},
	}
	tfg.Watches = append(tfg.Watches, defaultwatch)
	return tfg
}

func CheckIfResetConfig(args []string) {
	if len(args) == 2 {
		if args[1] == "reset" {
			e := os.Remove("./.pushalotcfg.json")
			if e != nil {
				log.Fatal("Could not remove current config file. Permissions issue?")
			}
			Default := GetDefaultConfig()
			out, e := json.Marshal(Default)
			e = ioutil.WriteFile("./.pushalotcfg.json", out, 600)
			if e != nil {
				log.Fatal("cannot open settings file :(")
			}
			log.Fatal("Built config file. please fill it in.")
		}
	}
}

func GetCFG() GConfig {
	b, e := ioutil.ReadFile("./.pushalotcfg.json")
	tfg := GetDefaultConfig()
	if e != nil {
		out, e := json.Marshal(tfg)
		e = ioutil.WriteFile("./.pushalotcfg.json", out, 600)
		if e != nil {
			log.Fatal("cannot open settings file :(")
		}
		log.Fatal("Built config file. please fill it in.")
	}

	e = json.Unmarshal(b, &tfg)
	if e != nil {
		log.Fatalf("Could not parse config settings. You may have to remove ./.pushalotcfg.json")
	}
	if tfg.Token == "Fillmein" {
		log.Fatal("You need to fill in the config settings in ./.pushalotcfg.json")
	}
	return tfg
}
