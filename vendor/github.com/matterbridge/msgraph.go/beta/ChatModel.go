// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Chat undocumented
type Chat struct {
	// Entity is the base model of Chat
	Entity
	// Topic undocumented
	Topic *string `json:"topic,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastUpdatedDateTime undocumented
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`
	// Members undocumented
	Members []ConversationMember `json:"members,omitempty"`
	// Messages undocumented
	Messages []ChatMessage `json:"messages,omitempty"`
	// InstalledApps undocumented
	InstalledApps []TeamsAppInstallation `json:"installedApps,omitempty"`
}