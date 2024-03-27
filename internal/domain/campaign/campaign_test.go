package campaign_test

import (
	"emailN/internal/domain/campaign"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "My Campaign"
	content = "My Content"
	emails  = []string{"bruno@email.com", "email@email.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)

	// act - agir
	campaign := campaign.NewCampaign(name, content, emails)

	// assert - assegurar
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
}

func Test_NewCampaignIDIsNotNil(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)

	// act - agir
	campaign := campaign.NewCampaign(name, content, emails)
	assert.NotNil(campaign.ID)
}
func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)

	now := time.Now().Add(-time.Second)

	// act - agir
	campaign := campaign.NewCampaign(name, content, emails)

	assert.Greater(campaign.CreatedOn, now)
}
