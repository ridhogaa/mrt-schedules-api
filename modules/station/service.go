package station

import (
	"encoding/json"
	"mrt-schedules-api/common/client"
	"net/http"
	"time"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	GetScheduleByStation(id string) (response StationResponse, err error)
}

type ServiceImpl struct {
	client *http.Client
}

func NewService() Service {
	return &ServiceImpl{client: &http.Client{
		Timeout: 10 * time.Second,
	}}
}

func (s *ServiceImpl) GetAllStation() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	res, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	var stations []Station
	err = json.Unmarshal(res, &stations)

	for _, station := range stations {
		response = append(response, StationResponse{
			Id:   station.Id,
			Name: station.Name,
		})
	}

	return
}

func (s *ServiceImpl) GetScheduleByStation(id string) (response StationResponse, err error) {

	return
}
