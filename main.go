package main

import (
    "flag"
    "github.com/bwmarrin/discordgo"
    "strings"
    
)

func init() {
    flag.StringVar(&token, "t", "", "Bot token")
    flag.Parse()
}

var token string

func main() {
    // do nothing
    return
}
