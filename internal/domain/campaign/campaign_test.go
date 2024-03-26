package campaign_test

import (
	"emailN/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)
	name := "My Campaign"
	content := "My Content"
	emails := []string{"bruno@email.com", "email@email.com"}

	// act - agir
	campaign := campaign.NewCampaign(name, content, emails)

	// assert - assegurar
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
}
