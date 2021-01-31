package teamsnotifier

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SectionFact struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type NotificationSection struct {
	Title string `json:"activityTitle"`
	Subtitle string `json:"activitySubtitle"`
	Image string `json:"activityImage"`
	Facts []SectionFact `json:"facts"`
	Markdown bool `json:"markdown"`
}

type Notification struct {
	Summary string `json:"summary"`
	Type string `json:"@type"`
	Context string `json:"@context"`
	Color string `json:"themecolor"`
	Sections []NotificationSection `json:"sections"`
}

func NewNotification() Notification{
	return Notification{
		Summary: "Test Notification",
		Type: "MessageCard",
		Context: "http://schema.org/extensions",
		Color: "00EE00",
		Sections: []NotificationSection{
			{
				Title: "CoucouTitle",
				Subtitle: "CoucouSubtitle",
				Image: "https://repository-images.githubusercontent.com/202264544/3ce58c00-19ab-11ea-8a01-81d62334b3ed",
				Facts: []SectionFact{
					{
						Name: "Premier fact",
						Value: "Value du premier fact",
					},
					{
						Name: "Second Fact",
						Value: "Value du second fact",
					},
				},
				Markdown: true,

			},
		},
	}
}

func Notify(e error) {
	var notification = NewNotification()
	notification.Summary = e.Error()
	data, _ := json.Marshal(notification)
	http.Post("https://en97xqx4zsb5.x.pipedream.net","application/json",bytes.NewReader(data))
}
