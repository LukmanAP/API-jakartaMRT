package station

import (
	"encoding/json"
	"errors"
	"jadwal-mrt/common/client"
	"net/http"
	"strings"
	"time"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	CheckScedulesByStation(id string) (response []SceduleResponse, err error)
}

type service struct {
	Client *http.Client
}

func NewService() Service {
	return &service{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetAllStation implements Service.
func (s *service) GetAllStation() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(*s.Client, url)
	if err != nil {
		return
	}

	var station []Station
	err = json.Unmarshal(byteResponse, &station)

	for _, item := range station {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}
	return
}

func (s *service) CheckScedulesByStation(id string) (response []SceduleResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns/"

	byteResponse, err := client.DoRequest(*s.Client, url)
	if err != nil {
		return
	}

	var scedules []Scedules
	err = json.Unmarshal(byteResponse, &scedules)
	if err != nil {
		return
	}

	var sceduleSelected Scedules
	for _, item := range scedules {
		if item.StationId == id {
			sceduleSelected = item
			break
		}
	}

	if sceduleSelected.StationId == "" {
		err = errors.New("station not found")
		return
	}

	response, err = ConvertDataToResponse(sceduleSelected)
	if err != nil {
		return
	}
	return

}

func ConvertDataToResponse(scedule Scedules) (response []SceduleResponse, err error) {
	var (
		LebakBulusTripName = "Stasiun Lebak Bulus Grab"
		BundaranHITripName = "Stasiun Bundaran HI Bank DKI"
	)
	sceduleLebakBulus := scedule.SceduleLebakBulus
	sceduleBundaranHI := scedule.SceduleBundaranHI

	sceduleLebakBulusParsed, err := ConvertSceduleToTimeFormat(sceduleLebakBulus)
	if err != nil {
		return
	}

	sceduleBundaranHIParsed, err := ConvertSceduleToTimeFormat(sceduleBundaranHI)
	if err != nil {
		return
	}

	for _, item := range sceduleLebakBulusParsed {
		if item.Format("15:05") > time.Now().Format("15:05") {
			response = append(response, SceduleResponse{
				StationName: LebakBulusTripName,
				Time:        item.Format("15:05"),
			})
		}
	}
	for _, item := range sceduleBundaranHIParsed {
		if item.Format("15:05") > time.Now().Format("15:05") {
			response = append(response, SceduleResponse{
				StationName: BundaranHITripName,
				Time:        item.Format("15:05"),
			})
		}
	}
	return
}

func ConvertSceduleToTimeFormat(scedule string) (response []time.Time, err error) {
	var (
		parsedTime time.Time
		scedules   = strings.Split(scedule, ",")
	)

	for _, item := range scedules {
		trimnedTime := strings.TrimSpace(item)
		if trimnedTime == "" {
			continue
		}

		parsedTime, err = time.Parse("15:05", trimnedTime)
		if err != nil {
			err = errors.New("invalid time format " + trimnedTime)
			return
		}
		response = append(response, parsedTime)
	}
	return

}
