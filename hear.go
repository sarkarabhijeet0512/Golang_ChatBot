package main

import (
	"fmt"
	"regexp"
)

type hearStruct struct {
	regex string
	text  string
}

func hear(word string, userid string) string {
	if word == "hello" {
		SendMessage(userid, "Hi there!!")
		// return "In this case,we'll trigger anything that will be handled when 'Hello' is triggered"
	}
	if word == "GET_STARTED" {
		SendMessage(userid, "Hello there! Seems Like this is the first time we're talking! call me GO Bot")
	}
	if word == "How are you?" {
		SendMessage(userid, "Thats Nice of you to ask that to me I am good What about You?!")
	}
	if word == "What are you?" {
		SendMessage(userid, "I am a chat bot named Go bot!")
	}
	if word == "What can you do?" {
		SendMessage(userid, "I can do something")
	} else {
		var re = regexp.MustCompile("(?i)" + word)
		var str = `Hey there! Hi Whats good? Hello! How are you?`

		for _, match := range re.FindAllString(str, -1) {
			// keyValuePair(match)
			key := keyValuePair(match)
			fmt.Println(key)
			SendMessage(userid, key)
			fmt.Println(keyValuePair(match))

		}
	}

	return ""
}
func (h *hearStruct) listen(userid string) {
	if h.text == "" && h.regex == "" {
		panic("Oops! Nothing to listen")
	}
	if h.regex != "" {
		hear(h.regex, userid)
		fmt.Println("REGEX PASSED", h.regex)
	} else if h.text != "" {
		hear(h.text, userid)
		fmt.Println("TEXT PASSED", h.text)
	}
}
func keyValuePair(match string) string {
	response := make(map[string]string)
	response["Hi"] = "Hi there!"
	response["Hello"] = "Hi there!"
	response["What"] = "Ask for something From /What you can do?/What are you?"
	response["How"] = "Ask for something From /How are you?"
	response["Start"] = "Hello there! Seems Like this is the first time we're talking! call me GO Bot"
	return response[match]
}
