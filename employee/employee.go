package employee

type Employee struct {
	ID        int `json:"id"`
	Name      string `json:"name"`
	Age       int `json:"age"`
	Position  string `json:"position"`
	Department string `json:"department"`
	Salary    float64 `json:"salary"`
}