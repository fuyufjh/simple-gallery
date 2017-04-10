package obs

type Photo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Category struct {
	Name   string `json:"name"`
	Photos []*Photo `json:"photos"`
}
