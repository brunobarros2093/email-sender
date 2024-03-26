package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name, content string, emails []string) *Campaign {
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}
	return &Campaign{
		ID:        "1",
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}
}
