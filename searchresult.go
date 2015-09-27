package main 

type SearchResult struct {
	Photos struct {
		Page    int    `json: "page"`
		Pages   int    `json: "pages"`
		PerPage int    `json: "perpage"`
		Total   string `json: "total"`
		Photo   []struct {
			Id       string `json: "id"`
			Owner    string `json: "owner"`
			Secret   string `json: "secret"`
			Server   string `json: "server"`
			Farm     int    `json: "farm"`
			Title    string `json: "title"`
			IsPublic int    `json: "ispublic"`
			IsFriend int    `json: "isfriend"`
			IsFamily int    `json: "isfamily`
		} `json: "photo"`
	} `json: "photos"`
	Stat string `json: "stat"`
}
