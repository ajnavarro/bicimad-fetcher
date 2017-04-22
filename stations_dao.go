package bicimad_fetcher

import (
	"database/sql"
	"time"

	"github.com/ajnavarro/bicimad-fetcher/model"

	_ "github.com/mattn/go-sqlite3"
)

type StationsDAO struct {
	DB *sql.DB
}

func NewStationsDAO(dbfile string) (*StationsDAO, error) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	return &StationsDAO{
		DB: db,
	}, nil
}

func (s *StationsDAO) CreateTables() error {
	stationsData := "create table if not exists stations(" +
		"id integer not null, " +
		"createdAt integer," +
		"latitude string, " +
		"longitude string, " +
		"name string, " +
		"number integer," +
		"address string," +
		"light integer," +
		"activate integer," +
		"noAvailable integer," +
		"totalBases integer," +
		"dockBikes integer," +
		"freeBases integer," +
		"reservationsCount integer" +
		");"
	_, err := s.DB.Exec(stationsData)
	return err
}

func (s *StationsDAO) SaveStations(m *model.Stations) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into stations(" +
		"id,createdAt,latitude,longitude,name,number,address,light,activate,noAvailable,totalBases,dockBikes," +
		"freeBases,reservationsCount) " +
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	fetchedAt := time.Now().Unix()
	for _, s := range m.Stations {
		_, err := stmt.Exec(
			s.ID,
			fetchedAt,
			s.Latitude,
			s.Longitude,
			s.Name,
			s.Number,
			s.Address,
			s.Light,
			s.Activate,
			s.NoAvailable,
			s.TotalBases,
			s.DockBikes,
			s.FreeBases,
			s.ReservationsCount,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
