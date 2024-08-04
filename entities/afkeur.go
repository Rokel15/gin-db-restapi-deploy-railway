package entities

type Afkeur struct {
	ID         int    `json:"id"`
	JumlahAyam int    `json:"jumlahAyam"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	Mortalitas int    `json:"Mortalitas"`
}
