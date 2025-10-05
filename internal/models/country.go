package models

type Country struct {
	Name       Name                `json:"name"`
	Capital    []string            `json:"capital"`
	Region     string              `json:"region"`
	Subregion  string              `json:"subregion"`
	Population int64               `json:"population"`
	Demonyms   Demonyms            `json:"demonyms"`
	Currencies map[string]Currency `json:"currencies"`
	Languages  map[string]string   `json:"languages"`
	Flags      Flags               `json:"flags"`
	Maps       Maps                `json:"maps"`
	CCA2       string              `json:"cca2"` //ISO 3166-1 alfa-2, necessary for the news api
}

type Name struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}
type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Flags struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
}

type Demonyms struct {
	Eng map[string]string `json:"eng"`
}

type Maps struct {
	GoogleMaps     string `json:"googleMaps"`
	OpenStreetMaps string `json:"openStreetMaps"`
}
