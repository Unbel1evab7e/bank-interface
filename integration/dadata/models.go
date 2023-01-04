package dadata

type SuggestionRequest struct {
	Query string `json:"query"`
}

type SuggestionResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
}

type Suggestion struct {
	Value             string `json:"value"`
	UnrestrictedValue string `json:"unrestricted_value"`
	Data              Data   `json:"data"`
}
type Data struct {
	PostalCode           interface{} `json:"postal_code"`
	Country              string      `json:"country"`
	CountryIsoCode       string      `json:"country_iso_code"`
	FederalDistrict      interface{} `json:"federal_district"`
	RegionFiasID         string      `json:"region_fias_id"`
	RegionKladrID        string      `json:"region_kladr_id"`
	RegionIsoCode        string      `json:"region_iso_code"`
	RegionWithType       string      `json:"region_with_type"`
	RegionType           string      `json:"region_type"`
	RegionTypeFull       string      `json:"region_type_full"`
	Region               string      `json:"region"`
	AreaFiasID           interface{} `json:"area_fias_id"`
	AreaKladrID          interface{} `json:"area_kladr_id"`
	AreaWithType         interface{} `json:"area_with_type"`
	AreaType             interface{} `json:"area_type"`
	AreaTypeFull         interface{} `json:"area_type_full"`
	Area                 interface{} `json:"area"`
	CityFiasID           string      `json:"city_fias_id"`
	CityKladrID          string      `json:"city_kladr_id"`
	CityWithType         string      `json:"city_with_type"`
	CityType             string      `json:"city_type"`
	CityTypeFull         string      `json:"city_type_full"`
	City                 string      `json:"city"`
	CityArea             interface{} `json:"city_area"`
	CityDistrictFiasID   interface{} `json:"city_district_fias_id"`
	CityDistrictKladrID  interface{} `json:"city_district_kladr_id"`
	CityDistrictWithType interface{} `json:"city_district_with_type"`
	CityDistrictType     interface{} `json:"city_district_type"`
	CityDistrictTypeFull interface{} `json:"city_district_type_full"`
	CityDistrict         interface{} `json:"city_district"`
	SettlementFiasID     interface{} `json:"settlement_fias_id"`
	SettlementKladrID    interface{} `json:"settlement_kladr_id"`
	SettlementWithType   interface{} `json:"settlement_with_type"`
	SettlementType       interface{} `json:"settlement_type"`
	SettlementTypeFull   interface{} `json:"settlement_type_full"`
	Settlement           interface{} `json:"settlement"`
	StreetFiasID         string      `json:"street_fias_id"`
	StreetKladrID        string      `json:"street_kladr_id"`
	StreetWithType       string      `json:"street_with_type"`
	StreetType           string      `json:"street_type"`
	StreetTypeFull       string      `json:"street_type_full"`
	Street               string      `json:"street"`
	SteadFiasID          interface{} `json:"stead_fias_id"`
	SteadCadnum          interface{} `json:"stead_cadnum"`
	SteadType            interface{} `json:"stead_type"`
	SteadTypeFull        interface{} `json:"stead_type_full"`
	Stead                interface{} `json:"stead"`
	HouseFiasID          interface{} `json:"house_fias_id"`
	HouseKladrID         interface{} `json:"house_kladr_id"`
	HouseCadnum          interface{} `json:"house_cadnum"`
	HouseType            interface{} `json:"house_type"`
	HouseTypeFull        interface{} `json:"house_type_full"`
	House                interface{} `json:"house"`
	BlockType            interface{} `json:"block_type"`
	BlockTypeFull        interface{} `json:"block_type_full"`
	Block                interface{} `json:"block"`
	Entrance             interface{} `json:"entrance"`
	Floor                interface{} `json:"floor"`
	FlatFiasID           interface{} `json:"flat_fias_id"`
	FlatCadnum           interface{} `json:"flat_cadnum"`
	FlatType             interface{} `json:"flat_type"`
	FlatTypeFull         interface{} `json:"flat_type_full"`
	Flat                 interface{} `json:"flat"`
	FlatArea             interface{} `json:"flat_area"`
	SquareMeterPrice     interface{} `json:"square_meter_price"`
	FlatPrice            interface{} `json:"flat_price"`
	RoomFiasID           interface{} `json:"room_fias_id"`
	RoomCadnum           interface{} `json:"room_cadnum"`
	RoomType             interface{} `json:"room_type"`
	RoomTypeFull         interface{} `json:"room_type_full"`
	Room                 interface{} `json:"room"`
	PostalBox            interface{} `json:"postal_box"`
	FiasID               string      `json:"fias_id"`
	FiasCode             interface{} `json:"fias_code"`
	FiasLevel            string      `json:"fias_level"`
	FiasActualityState   string      `json:"fias_actuality_state"`
	KladrID              string      `json:"kladr_id"`
	GeonameID            string      `json:"geoname_id"`
	CapitalMarker        string      `json:"capital_marker"`
	Okato                string      `json:"okato"`
	Oktmo                string      `json:"oktmo"`
	TaxOffice            string      `json:"tax_office"`
	TaxOfficeLegal       string      `json:"tax_office_legal"`
	Timezone             interface{} `json:"timezone"`
	GeoLat               string      `json:"geo_lat"`
	GeoLon               string      `json:"geo_lon"`
	BeltwayHit           interface{} `json:"beltway_hit"`
	BeltwayDistance      interface{} `json:"beltway_distance"`
	Metro                interface{} `json:"metro"`
	Divisions            interface{} `json:"divisions"`
	QcGeo                string      `json:"qc_geo"`
	QcComplete           interface{} `json:"qc_complete"`
	QcHouse              interface{} `json:"qc_house"`
	HistoryValues        []string    `json:"history_values"`
	UnparsedParts        interface{} `json:"unparsed_parts"`
	Source               interface{} `json:"source"`
	Qc                   interface{} `json:"qc"`
}

type AutoGenerated struct {
	Suggestions []struct {
	} `json:"suggestions"`
}
