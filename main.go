package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"git.sr.ht/gotoxy/illchat/src"
	"github.com/sirupsen/logrus"
)

const figlet = `

W E L C O M E  T O
  
      ,a8a,                    ,gggg,                               
     ,8" "8, ,dPYb, ,dPYb,   ,88"""Y8b,,dPYb,                  I8   
     d8   8b IP' Yb IP' Yb  d8"      Y8IP' Yb                  I8   
     88   88 I8  8I I8  8I d8'   8b  d8I8  8I               88888888
     88   88 I8  8' I8  8',8I    "Y88P'I8  8'                  I8   
     Y8   8P I8 dP  I8 dP I8'          I8 dPgg,     ,gggg,gg   I8   
      8, ,8' I8dP   I8dP  d8           I8dP" "8I   dP"  "Y8I   I8   
8888  "8,8"  I8P    I8P   Y8,          I8P    I8  i8'    ,8I  ,I8,  
 8b,  ,d8b, ,d8b,_ ,d8b,_  Yba,,_____,,d8     I8,,d8,   ,d8b,,d88b, 
  "Y88P" "Y88P'"Y888P'"Y88   "Y888888888P      Y8P"Y8888P" Y88P""Y8 
               
                                 Illfonic     P 2 P  C H A T  P O C
`

func init() {
	// Log as Text with color
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	// Log to stdout
	logrus.SetOutput(os.Stdout)
}

func main() {
	// Define input flags
	username := flag.String("user", "", "username to use in the chatroom.")
	chatroom := flag.String("room", "", "chatroom to join.")
	loglevel := flag.String("log", "", "level of logs to print.")
	discovery := flag.String("discover", "", "method to use for discovery.")
	// Parse input flags
	flag.Parse()

	// Set the log level
	switch *loglevel {
	case "panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "info", "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Display the welcome figlet
	fmt.Println(figlet)
	fmt.Println("IllChat is initializing...")
	fmt.Println("This may take upto 30 seconds.")
	fmt.Println()

	// Create a new P2PHost
	p2phost := src.NewP2P()
	logrus.Infoln("Completed P2P Initialization")

	// Connect to peers with the chosen discovery method
	switch *discovery {
	case "announce":
		p2phost.AnnounceConnect()
	case "advertise":
		p2phost.AdvertiseConnect()
	default:
		p2phost.AdvertiseConnect()
	}
	logrus.Infoln("Connected to Service Peers")

	// Join the chat room
	chatapp, _ := src.JoinChatRoom(p2phost, *username, *chatroom)
	logrus.Infof("Joined the '%s' chatroom as '%s'", chatapp.RoomName, chatapp.UserName)

	// Wait for network setup to complete
	time.Sleep(time.Second * 5)

	// Create the Chat UI
	ui := src.NewUI(chatapp)
	// Start the UI system
	ui.Run()
}
