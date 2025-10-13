package repository

import (
	"Lab1/intermal/app/ds"
	"strings"
)

func (r *Repository) GetDraftRequestChronicleResearchInfo() (ds.RequestChronicleResearch, []ds.ChronicleResearch, error) {
	creatorID := uint(1)

	var requestChronicleResearch ds.RequestChronicleResearch
	err := r.db.Where("creator_id = ? AND status = ?", creatorID, ds.RequestStatusDraft).First(&requestChronicleResearch).Error
	if err != nil {
		return ds.RequestChronicleResearch{}, nil, err
	}

	var chronicleResearch []ds.ChronicleResearch
	err = r.db.Where("id_request_research = ?", requestChronicleResearch.ID).Find(&chronicleResearch).Error
	if err != nil {
		return ds.RequestChronicleResearch{}, nil, err
	}

	return requestChronicleResearch, chronicleResearch, nil
}

func (r *Repository) CreateRequestChronicleResearch(resourceID uint) (ds.RequestChronicleResearch, error) {
	requestChronicleResearch := ds.RequestChronicleResearch{
		Name:        "",
		SearchEvent: "",
		Status:      ds.RequestStatusDraft,
		CreatorID:   1,
	}
	err := r.db.Create(&requestChronicleResearch).Error
	if err != nil {
		return ds.RequestChronicleResearch{}, err
	}

	err = r.AddChronicleResearchToRequest(requestChronicleResearch.ID, resourceID)
	if err != nil {
		return ds.RequestChronicleResearch{}, err
	}

	return requestChronicleResearch, nil
}

func (r *Repository) checkIsMatch(quote string, searchEvent string) bool {
	eventLower := strings.ToLower(strings.TrimSpace(searchEvent))
	quoteLower := strings.ToLower(strings.TrimSpace(quote))

	if quoteLower == "" || eventLower == "" {
		return false
	}

	return strings.Contains(quoteLower, eventLower)
}

func (r *Repository) AddChronicleResearchToRequest(requestID uint, resourceID uint) error {
	var request ds.RequestChronicleResearch
	err := r.db.Where("id = ?", requestID).First(&request).Error
	if err != nil {
		return err
	}

	quote := ""
	isMatched := r.checkIsMatch(quote, request.SearchEvent)

	chronicleResearch := ds.ChronicleResearch{
		IDRequestResearch: requestID,
		IDResource:        resourceID,
		Quote:             quote,
		IsMatched:         isMatched,
	}

	err = r.db.Create(&chronicleResearch).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteRequestChronicleResearch(requestID int) error {
	result := r.db.Exec("UPDATE request_chronicle_researches SET status = ? WHERE id = ?", ds.RequestStatusDeleted, requestID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
