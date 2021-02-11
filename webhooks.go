package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"lastName"`
	ProilePic string `json:"proile_pic"`
}
type TextReplyRecipientstruct struct {
	ID string `json:"id"`
}
type TextReplystruct struct {
	Text string `json:"text"`
}
type CallSendAPIResponse struct {
	RecID     string `json:"recipient_id"`
	MessageID string `json:"message_id"`
}
type Vertex struct {
	X string                   `json:"message_type"`
	I TextReplyRecipientstruct `json:"recipient"`
	Y TextReplystruct          `json:"message"`
}

// Body is the struct used for sending the data as formaed
type Body struct {
	Object string
	Entry  []struct {
		ID        string
		Time      int64
		Messaging []struct {
			Timestamp int64
			Sender    Senderstruct
			Recipient Recipientstruct
			Message   *Messagestruct
			Postback  *Postbackstruct `json:"postback"`
		}
	}
}
type Messagestruct struct {
	Text        string
	Attachments []*Attachmentstruct
}
type MessagingStruct struct {
}
type Attachmentstruct struct {
	Type    string
	Payload Payloadstruct
}
type Recipientstruct struct {
	ID string
}
type Senderstruct struct {
	ID string
}
type Payloadstruct struct {
	URL      string
	Reusable bool `json:"reusable"`
}
type Postbackstruct struct {
	Payload string `json:"payload"`
}

var tk Config
var v = getToken()
var _ = json.Unmarshal([]byte(v), &tk)

func webhookGetHandler(response http.ResponseWriter, request *http.Request) {
	token := tk.VerifyToken
	tokenTrue := request.URL.Query().Get("hub.verify_token")
	hubChallange := request.URL.Query().Get("hub.challenge")

	if tokenTrue == token {
		response.Header().Set("content-type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte(hubChallange))
	} else {
		fmt.Fprint(response, "token does not match")
	}
}
func webhookPostHandler(response http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("Error parsing the body")
	}
	var body Body
	json.Unmarshal([]byte(data), &body)
	if body.Object == "page" {
		fmt.Println("the Whole object recieved:", string(data))
		for _, entires := range body.Entry {
			for _, messaging := range entires.Messaging {
				if messaging.Message != nil {
					attachments := messaging.Message.Attachments
					text := messaging.Message.Text
					// regex := messaging.Message.Regex
					if attachments != nil {
						SendMessage(messaging.Sender.ID, "Oops!Cant do that yet")
						fmt.Println("Attachment.Cannot process")
					} else if messaging.Message.Text != "" {
						fmt.Println("THE USER ID IS:", messaging.Sender.ID)
						vr := hearStruct{text: text}
						vr.listen(messaging.Sender.ID)
					}
				} else if messaging.Postback != nil {
					fmt.Println("Yay!we have a postback event!!")
					if messaging.Postback.Payload == "GET_STARTED" {
						SendMessage(messaging.Sender.ID, "Hello! It seems thisis the first time we're taling! Call me Go bot")
					}
				}
			}
			response.WriteHeader(http.StatusOK)
		}

	} else {
		fmt.Println("Seems like the postback of the get started button", data)
	}
}
func SendMessage(UserID string, text string) {
	userProfile := getUserProfile(UserID)
	var prof Profile
	json.Unmarshal([]byte(userProfile), &prof)
	i := TextReplyRecipientstruct{UserID}
	t := TextReplystruct{text}
	send, err := json.Marshal(Vertex{"RESPONSE", i, t})
	if err != nil {
		log.Println(">>ERROR SENDING MESSAGE")
	}
	callSendAPI(send)
}
func callSendAPI(data []byte) {
	accessToken := tk.AccessToken
	fmt.Println(accessToken)
	response, err := http.Post("https://graph.facebook.com/v2.6/me/messages?access_token="+accessToken, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("Error here", err)
		panic("callSend API Err:" + err.Error())
	}
	res, _ := ioutil.ReadAll(response.Body)
	var rs CallSendAPIResponse
	json.Unmarshal([]byte(res), &rs)
	if rs.MessageID == "" {
		fmt.Println("Error wih CallsendAPI here", string(res))
	}
}
func getUserProfile(userID string) string {
	accessToken := tk.AccessToken
	profileField := []string{"first_name", "last_name", "name", "Profile_pic"}
	fmt.Println(profileField)
	separatedUserFields := strings.Join(profileField, ",")
	response, err := http.Get("https://graph.facebook.com/v3.1/" + userID + "?fields=" + separatedUserFields + "&acsess_token=" + accessToken)
	if err != nil {
		fmt.Println(">>ERROR ACCESSING THE USER PROFILE:", err)
	}
	res, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(res))
	return string(res)
}
