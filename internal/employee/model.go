package employee

type EmployeeName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type EmployeeSalaryEstimate struct {
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Salary2016   float64 `json:"salary2016"`
	TotalReviews int     `json:"totalReviews"`
}
