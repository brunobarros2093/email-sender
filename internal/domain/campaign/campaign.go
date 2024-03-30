package campaign

import (
	internalerrors "emailN/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"required,email" json:"email"`
}

type Campaign struct {
	ID        string    `validate:"required" json:"id"`
	Name      string    `validate:"min=5,max=24" json:"name"`
	CreatedOn time.Time `validate:"required" json:"created_on"`
	Content   string    `validate:"required,min=5,max=1024" json:"content"`
	Contacts  []Contact `validade:"min=1"json:"contacts"`
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}
	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err
}
