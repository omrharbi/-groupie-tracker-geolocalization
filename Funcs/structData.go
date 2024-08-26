package Groupie_tracker

type JsonData struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Artist struct {
	Id             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	DatesLocations map[string][]string `json:"datesLocations"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Date           []string            `json:"dates"`
	Location       []string
}

type GeoResponse struct {
	Items []struct {
		Position struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"position"`
	} `json:"items"`
}

type Location struct {
	Location []string `json:"locations"`
}
type Date struct {
	Date []string `json:"dates"`
}
type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Errors struct {
	Message     string
	Code        int
	Description string
}
type AllMessageErrors struct {
	NotFound                    string `json:"notfound"`
	BadRequest                  string `json:"badrequest"`
	InternalError               string `json:"internalerror"`
	MethodNotAllowed            string `json:"methodnotallowed"`
	DescriptionNotFound         string `json:"description_notfound"`
	DescriptionBadRequest       string `json:"description_badrequest"`
	DescriptionInternalError    string `json:"description_internalerror"`
	DescriptionMethodNotAllowed string `json:"description_methodnotallowed"`
}
