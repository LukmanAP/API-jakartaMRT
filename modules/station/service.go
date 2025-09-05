package station

import (
	"encoding/json"
	"jadwal-mrt/common/client"
	"net/http"
	"time"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
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
