package station

type Station struct {
	Id   string `json:"nid"`
	Name string `json:"title"`
}

type StationResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Scedules struct {
	StationId         string `json:"nid"`
	StationName       string `json:"title"`
	SceduleBundaranHI string `json:"jadwal_hi_biasa"`
	SceduleLebakBulus string `json:"jadwal_lb_biasa"`
}

type SceduleResponse struct {
	StationName string `json:"station"`
	Time        string `json:"time"`
}

// Swagger-only typed wrappers (to document responses)
// These types mirror the common response format but with concrete Data fields
// so they can be referenced in Swagger docs without changing runtime code.
type GetStationsAPIResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    []StationResponse  `json:"data"`
}

type GetScedulesAPIResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    []SceduleResponse  `json:"data"`
}

type ErrorAPIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
