package main

type PhotoResult struct {
	Photo struct {
		Id             string `json: "id"`
		Secret         string `json: "secret"`
		Server         string `json: "server"`
		Farm           int    `json: "farm"`
		DateUploaded   string `json: "dateuploaded"`
		IsFavorite     int    `json: "isfavorite"`
		License        string `json: "license"`
		SafetyLevel    string `json: "safety_level"`
		Rotation       int    `json: "rotation"`
		OriginalSecret string `json: "originalsecret"`
		OriginalFormat string `json: "originalformat"`
		Owner          struct {
			Nsid       string `json: "nssid"`
			Username   string `json: "username"`
			Realname   string `json: "realname"`
			Location   string `json: "location"`
			IconServer string `json: "iconserver"`
			IconFarm   int    `json: "iconfarm"`
			Path_Alias string `json: "path_alias"`
		}
		Title struct {
			Content string `json: "_content"`
		}
		Description struct {
			Content string `json: "_content"`
		}
		Visibility struct {
			IsPublic int `json: "ispublic"`
			IsFriend int `json: "isfriend"`
			IsFamily int `json: "isfamily"`
		}
		Dates struct {
			Posted           string `json: "posted"`
			Take             string `json: "taken"`
			TakenGranularity string `json: "takengranularity"`
			TakenUnknown     string `json: "takenunknown"`
			LastUpdate       string `json: "lastupdate"`
		}
		Views       string `json: "views"`
		Editability struct {
			CanComment int `json: "cancomment"`
			CanAddMeta int `json: "canaddmeta"`
		}
		PublicEditability struct {
			CanComment int `json: "cancomment"`
			CanAddMeta int `json: "canaddmeta"`
		}
		Usage struct {
			CanDownload int `json: "candownload"`
			CanBlog     int `json: "canblog"`
			CanPrint    int `json: "canprint"`
			CanShare    int `json: "canshare"`
		}

		Comments struct {
			Content string `json: "_content"`
		}

		Notes struct {
			Note []struct {
				// TODO: Find example to implement this
			}
		}
		People struct {
			HasPeople int `json: "haspeople"`
		}
		Tags struct {
			Tag []struct {
				Id         string `json: "id"`
				Author     string `json: "author"`
				AuthorName string `json: "authorname"`
				Raw        string `json: "raw"`
				Content    string `json: "_content"`
				MachineTag bool   `json: "machine_tag"`
			}
		}

		Urls struct {
			Url []struct {
				Type    string `json: "type"`
				Content string `json: "_content"`
			}
		}

		Media string `json: "media"`
	}
	Stat string `json: "stat"`
}
