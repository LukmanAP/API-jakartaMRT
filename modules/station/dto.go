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
