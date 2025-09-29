package repository

import (
	"Lab1/intermal/app/ds"
	"fmt"
)

func (r *Repository) GetChronicleResources() ([]ds.ChronicleResource, error) {
	var resources []ds.ChronicleResource
	err := r.db.Find(&resources).Error
	if err != nil {
		return nil, err
	}
	if len(resources) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}
	return resources, nil
}

func (r *Repository) GetChronicleResource(id int) (ds.ChronicleResource, error) {
	var resource ds.ChronicleResource
	err := r.db.Where("id = ?", id).First(&resource).Error
	if err != nil {
		return ds.ChronicleResource{}, err
	}
	return resource, nil
}

func (r *Repository) GetChronicleResourcesByTitle(title string) ([]ds.ChronicleResource, error) {
	var resources []ds.ChronicleResource
	err := r.db.Where("title ILIKE ?", "%"+title+"%").Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

// &&&&
func (r *Repository) GetChronicleResearchForRequest(requestID int) (ds.RequestChronicleResearch, []ds.ChronicleResearch, error) {
	var request ds.RequestChronicleResearch
	err := r.db.Where("id = ?", requestID).First(&request).Error
	if err != nil {
		return ds.RequestChronicleResearch{}, nil, fmt.Errorf("такой заявки нет")
	}

	var research []ds.ChronicleResearch
	err = r.db.Where("id_request_research = ?", requestID).Find(&research).Error
	if err != nil {
		return request, nil, err
	}

	return request, research, nil
}
