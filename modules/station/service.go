package station

import (
	"encoding/json"
	"errors"
	"mrt-schedules-api/common/client"
	"net/http"
	"strings"
	"time"
)

const url = "https://jakartamrt.co.id/id/val/stasiuns"

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	GetScheduleByStation(id string) (response []ScheduleResponse, err error)
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

func (s *ServiceImpl) GetScheduleByStation(id string) (response []ScheduleResponse, err error) {

	res, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var schedules []Schedule
	err = json.Unmarshal(res, &schedules)

	var scheduleSelected Schedule

	for _, schedule := range schedules {
		if schedule.StationId == id {
			scheduleSelected = schedule
		}
	}

	if scheduleSelected.StationId == "" {
		err = errors.New("stations not found")
		return
	}
	scheduleSplitHI := strings.Split(scheduleSelected.ScheduleHI, ",")
	scheduleSplitLB := strings.Split(scheduleSelected.ScheduleLB, ",")
	for _, item := range scheduleSplitHI {
		response = append(response, ScheduleResponse{
			StationName: scheduleSelected.StationName,
			Time:        item,
		})
	}
	for _, item := range scheduleSplitLB {
		response = append(response, ScheduleResponse{
			StationName: scheduleSelected.StationName,
			Time:        item,
		})
	}
	return
}
