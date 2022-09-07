package structs

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Organization struct {
	Id		           int       `json:"id"`
    Title 	           string    `json:"title"`
    Password           string    `json:"password"`
    Email              string    `json:"email"`
    EmailNotifications bool      `json:"emailNotifications"`
    Key                string    `json:"key"`
    UserLimit          int       `json:"userLimit"`
    Statistics         bool      `json:"statistics"`
    ProtectionCheating bool      `json:"protectionCheating"`
    Date               string `json:"date"`
    Time               string `json:"time"`
    ThemeId            int       `json:"themeId"`
}

type Theme struct {
	Id    int
    Title string
}

type Payment struct {
	Number             int        `json:"number"`
    Money              int        `json:"money"`
    Date               timestamp.Timestamp  `json:"date"`
    Time               timestamp.Timestamp  `json:"time"`
    Users              int        `json:"users"`
    Statistics         bool       `json:"statistics"`
    ProtectionCheating bool       `json:"protectionCheating"`
    OrganizationId     int        `json:"organizationId"`
}

type DeadLine struct {
	Id	           int       `json:"Id"`
    Date           timestamp.Timestamp `json:"Date"`
    Time           timestamp.Timestamp `json:"Time"`
    OrganizationId int       `json:"OrganizationId"`
}

type Session struct {
	Id		             int                  `json:"id"`
    UserId               int                  `json:"userId"` 
    IpAddress            string               `json:"ipAddress"`
    BackendStartDateTime timestamp.Timestamp  `json:"backendStartDateTime"`
    StateChangeDateTime  timestamp.Timestamp  `json:"stateChangeDateTime"`
    StateId              int                  `json:"stateId"`
}

type State struct {
	Id		int    `json:"id"`
    Title   string `json:"title"`
}