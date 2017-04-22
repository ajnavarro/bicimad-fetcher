package bicimad_fetcher

import (
	"github.com/inconshreveable/log15"
	"time"
)

type Fetcher struct {
	Client *Client
	DAO    *StationsDAO
	TTS    time.Duration
}

func (f *Fetcher) Fetch() error {
	if err := f.DAO.CreateTables(); err != nil {
		return err
	}

	log15.Info("tables created", "tts", f.TTS)

	firstTime := true
	for {
		if !firstTime {
			<-time.After(f.TTS)
		}
		firstTime = false

		stations, err := f.Client.Stations()
		if err != nil {
			log15.Error("error fetching stations from API", "error", err)
			continue
		}

		if err := f.DAO.SaveStations(stations); err != nil {
			log15.Error("error saving stations to database", "error", err)
			continue
		}

		log15.Info("data saved correctly", "time", time.Now())
	}
}
