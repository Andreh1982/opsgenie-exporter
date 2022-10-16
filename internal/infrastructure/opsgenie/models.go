package opsgenie

import "time"

type Configuration struct {
	ApiUrl   string `json:"apiurl"`
	GenieKey string `json:"geniekey"`
}

type IncidentList struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		Actions         []interface{} `json:"actions"`
		CreatedAt       time.Time     `json:"createdAt"`
		Description     string        `json:"description"`
		ExtraProperties struct {
		} `json:"extraProperties"`
		ID               string        `json:"id"`
		ImpactStartDate  time.Time     `json:"impactStartDate"`
		ImpactedServices []interface{} `json:"impactedServices"`
		Links            struct {
			API string `json:"api"`
			Web string `json:"web"`
		} `json:"links"`
		Message    string `json:"message"`
		OwnerTeam  string `json:"ownerTeam"`
		Priority   string `json:"priority"`
		Responders []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"responders"`
		Status    string        `json:"status"`
		Tags      []interface{} `json:"tags"`
		TinyID    string        `json:"tinyId"`
		UpdatedAt time.Time     `json:"updatedAt"`
	} `json:"data"`
	Paging struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"paging"`
	Took      float64 `json:"took"`
	RequestID string  `json:"requestId"`
}

type IncidentTimeline struct {
	Data struct {
		Entries []struct {
			ID        string    `json:"id"`
			Group     string    `json:"group"`
			Type      string    `json:"type"`
			EventTime time.Time `json:"eventTime"`
			Hidden    bool      `json:"hidden"`
			Actor     struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"actor"`
			Title struct {
				Type    string `json:"type"`
				Content string `json:"content"`
			} `json:"title"`
			Description struct {
				Type    string `json:"type"`
				Content string `json:"content"`
			} `json:"description,omitempty"`
		} `json:"entries"`
		NextOffset string `json:"nextOffset"`
	} `json:"data"`
	Took      float64 `json:"took"`
	RequestID string  `json:"requestId"`
}

type Incident struct {
	Data struct {
		ID               string        `json:"id"`
		Description      string        `json:"description"`
		ImpactedServices []string      `json:"impactedServices"`
		TinyID           string        `json:"tinyId"`
		Message          string        `json:"message"`
		Status           string        `json:"status"`
		Tags             []interface{} `json:"tags"`
		CreatedAt        time.Time     `json:"createdAt"`
		UpdatedAt        time.Time     `json:"updatedAt"`
		Priority         string        `json:"priority"`
		OwnerTeam        string        `json:"ownerTeam"`
		Responders       []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"responders"`
		ExtraProperties struct {
		} `json:"extraProperties"`
		Links struct {
			Web string `json:"web"`
			API string `json:"api"`
		} `json:"links"`
		ImpactStartDate time.Time     `json:"impactStartDate"`
		ImpactEndDate   time.Time     `json:"impactEndDate"`
		Actions         []interface{} `json:"actions"`
	} `json:"data"`
	Took      float64 `json:"took"`
	RequestID string  `json:"requestId"`
}

type TeamsList struct {
	Data []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"data"`
	Took      float64 `json:"took"`
	RequestID string  `json:"requestId"`
}

type TeamInfo struct {
	Data struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Members     []struct {
			User struct {
				ID       string `json:"id"`
				Username string `json:"username"`
			} `json:"user"`
			Role string `json:"role"`
		} `json:"members"`
	} `json:"data"`
	Took      float64 `json:"took"`
	RequestID string  `json:"requestId"`
}
