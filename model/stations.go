package model

type Result struct {
	Code        string      `json:"code"`
	Description string   `json:"description"`
	Service     string   `json:"whoAmI"`
	Version     string   `json:"version"`
	Time        string   `json:"time"`
	Data        string   `json:"data"`
}

// {"id":98,"latitude":"40.4257046","longitude":"-3.6893698","name":"Plaza de Colón","light":0,"number":"93","address":"Calle Goya nº 1","activate":1,"no_available":0,"total_bases":24,"dock_bikes":3,"free_bases":19,"reservations_count":0}
type Stations struct {
	Stations []struct {
		ID        int    `json:"id"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
		// Bike station name. PE: "Plaza de Colon"
		Name  string `json:"name"`
		Light int    `json:"light"`
		// Station number
		Number            string `json:"number"`
		Address           string `json:"address"`
		Activate          int    `json:"activate"`
		NoAvailable       int    `json:"no_available"`
		TotalBases        int    `json:"total_bases"`
		DockBikes         int    `json:"dock_bikes"`
		FreeBases         int    `json:"free_bases"`
		ReservationsCount int    `json:"reservations_count"`
	} `json:"stations"`
}
