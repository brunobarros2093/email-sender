package campaign

import (
	"emailN/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaignDTO) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	s.Repository.Save(campaign)
	return campaign.ID, nil
}
