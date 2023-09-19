package db

type Assignee struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Card struct {
	Id               string     `json:"id"`
	Title            string     `json:"Name"`
	Status           string     `json:"Status"`
	Progress         float64    `json:"progress"`
	AsWho            string     `json:"En tant que"`
	IWant            string     `json:"Je veux"`
	DefinitionOfDone string     `json:"Definition Of Done"`
	Description      string     `json:"Description"`
	Jh               float64    `json:"JH"`
	Assignees        []Assignee `json:"Assignees"`
	OrderedJH        string     `json:"Ordered JH"`

	Category []string `json:"Category"`
	Sprint   []string `json:"Sprint"`
	Sector   []string `json:"Secteur"`
}
