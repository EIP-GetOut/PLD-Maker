package db

type Assignee struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CardFields struct {
	Title            string     `json:"Name"`
	Status           string     `json:"Status"`
	Progress         float64    `json:"progress"`
	AsWho            string     `json:"En tant que"`
	IWant            string     `json:"Je veux"`
	DefinitionOfDone string     `json:"Definition Of Done"`
	Description      string     `json:"Description"`
	Jh               float64    `json:"JH"`
	Assignee         []Assignee `json:"Assignees"`
	OrderedJH        string     `json:"Ordered JH"`

	Category []string `json:"Category"`
	Sprint   []string `json:"Sprint"`
	Secteur  []string `json:"Secteur"`
}

type Card struct {
	Id          string     `json:"id"`
	CreatedTime string     `json:"createdTime"`
	Fields      CardFields `json:"fields"`
}

type Cards struct {
	Cards  []Card `json:"records"`
	Offset string `json:"offset"`
}
