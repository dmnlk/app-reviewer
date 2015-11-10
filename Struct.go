package main

type ITunesResponse struct {
	Feed struct {
		Author struct {
			Name struct {
				Label string `json:"label"`
			} `json:"name"`
			Uri struct {
				Label string `json:"label"`
			} `json:"uri"`
		} `json:"author"`
		Entry []struct {
			Category struct {
				Attributes struct {
					Im_Id  string `json:"im:id"`
					Label  string `json:"label"`
					Scheme string `json:"scheme"`
					Term   string `json:"term"`
				} `json:"attributes"`
			} `json:"category"`
			ID struct {
				Attributes struct {
					Im_BundleId string `json:"im:bundleId"`
					Im_Id       string `json:"im:id"`
				} `json:"attributes"`
				Label string `json:"label"`
			} `json:"id"`
			Im_Artist struct {
				Attributes struct {
					Href string `json:"href"`
				} `json:"attributes"`
				Label string `json:"label"`
			} `json:"im:artist"`
			Im_ContentType struct {
				Attributes struct {
					Label string `json:"label"`
					Term  string `json:"term"`
				} `json:"attributes"`
			} `json:"im:contentType"`
			Im_Image []struct {
				Attributes struct {
					Height string `json:"height"`
				} `json:"attributes"`
				Label string `json:"label"`
			} `json:"im:image"`
			Im_Name struct {
				Label string `json:"label"`
			} `json:"im:name"`
			Im_Price struct {
				Attributes struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"attributes"`
				Label string `json:"label"`
			} `json:"im:price"`
			Im_ReleaseDate struct {
				Attributes struct {
					Label string `json:"label"`
				} `json:"attributes"`
				Label string `json:"label"`
			} `json:"im:releaseDate"`
			Link struct {
				Attributes struct {
					Href string `json:"href"`
					Rel  string `json:"rel"`
					Type string `json:"type"`
				} `json:"attributes"`
			} `json:"link"`
			Rights struct {
				Label string `json:"label"`
			} `json:"rights"`
			Title struct {
				Label string `json:"label"`
			} `json:"title"`
		} `json:"entry"`
		Icon struct {
			Label string `json:"label"`
		} `json:"icon"`
		ID struct {
			Label string `json:"label"`
		} `json:"id"`
		Link []struct {
			Attributes struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
				Type string `json:"type"`
			} `json:"attributes"`
		} `json:"link"`
		Rights struct {
			Label string `json:"label"`
		} `json:"rights"`
		Title struct {
			Label string `json:"label"`
		} `json:"title"`
		Updated struct {
			Label string `json:"label"`
		} `json:"updated"`
	} `json:"feed"`
}
