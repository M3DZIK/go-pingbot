package config

import "time"

// backend
var PingBot_Ticker = 2 * time.Minute

// update
var GH_Repo = "MedzikUser/go-pingbot"
var Latest_Version_Check = 2 * time.Minute

// website
var Port = ":8080"
