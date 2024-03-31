package campaign

import (
	"emailN/internal/contract"
	internalerrors "emailN/internal/internalErrors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaignDTO) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}
