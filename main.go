package main

import (
	"database/sql/driver"
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/paulsmith/gogeos/geos"
)

type Geo struct {
	geos.Geometry
}

func GeoFromWKT(wkt string) *Geo {
	geo, err := geos.FromWKT(wkt)
	if err != nil {
		fmt.Println(err)
	}
	return &Geo{Geometry: *geo}
}

func (g *Geo) Value() (driver.Value, error) {
	str, err := g.ToWKT()
	if err != nil {
		return nil, err
	}

	return str, nil
}

func (g *Geo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid value")
	}

	str := string(bytes)

	geom, err := geos.FromHex(str)
	if err != nil {
		return errors.Wrap(err, "Value could not be converted to geo")
	}

	geometry := Geo{Geometry: *geom}
	*g = geometry

	return nil
}

type City struct {
	Name       string
	State      string
	Population int
	Center     *Geo `gorm:"type:geography(point)"`
}

var cities = []City{
	{
		Name:       "Boston",
		State:      "MA",
		Population: 675_647,
		Center:     GeoFromWKT("POINT (42.350000 -71.050000)"),
	},
	{
		Name:       "Bellingham",
		State:      "WA",
		Population: 80_885,
		Center:     GeoFromWKT("POINT (48.751944 -122.478611)"),
	},
	{
		Name:       "Chicago",
		State:      "IL",
		Population: 2_746_388,
		Center:     GeoFromWKT("POINT (41.881944 -87.627778)"),
	},
	{
		Name:       "Houston",
		State:      "TX",
		Population: 2_304_580,
		Center:     GeoFromWKT("POINT (29.780000 -95.390000)"),
	},
	{
		Name:       "Los Angeles",
		State:      "CA",
		Population: 3_898_747,
		Center:     GeoFromWKT("POINT (34.010000 -118.410000)"),
	},
	{
		Name:       "New York",
		State:      "NY",
		Population: 8_804_190,
		Center:     GeoFromWKT("POINT (40.712778 -74.006111)"),
	},
	{
		Name:       "Portland",
		State:      "OR",
		Population: 652_503,
		Center:     GeoFromWKT("POINT (45.520000 -122.681944)"),
	},
	{
		Name:       "San Francisco",
		State:      "CA",
		Population: 873_965,
		Center:     GeoFromWKT("POINT (37.777500 -122.416389)"),
	},
	{
		Name:       "Seattle",
		State:      "WA",
		Population: 737_015,
		Center:     GeoFromWKT("POINT (47.609722 -122.333056)"),
	},
	{
		Name:       "Walla Walla",
		State:      "WA",
		Population: 31_731,
		Center:     GeoFromWKT("POINT (46.065000 -118.330278)"),
	},
}

func insertData(db *gorm.DB) {
	for _, city := range cities {
		db.Create(&city)
	}
}

func queryCities(db *gorm.DB) {
	var closeCities []City
	db.Raw(
		"select * from cities where ST_DWithin(center, ST_MakePoint(?,?)::geography, ?)",
		47.6,
		-122.33,
		500*1000,
	).Scan(&closeCities)

	for _, city := range closeCities {
		fmt.Println(city)
	}
}

func connect(host, database, username, password string) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("sslmode=disable host=%s dbname=%s  user=%s password=%s", host, database, username, password)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Println("Failed to connect to PostgreSQL database: %s", dbURI)
		return nil, err
	}

	return db.Debug(), nil
}

func main() {
	db, err := connect("localhost", "gorm-postgis", "gorm-postgis", "gorm-postgis")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	db.AutoMigrate(&City{})

	insertData(db)
	queryCities(db)
}
