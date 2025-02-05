package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	SoftwareEngineer = "Software Engineer"
	BackEnd          = "Backend"
	FrontEnd         = "Frontend"
	UiUx             = "UI/UX"
	ProductManager   = "ProductManager"
)

type Employees struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	DoB       time.Time `json:"dob"`
	Position  string    `json:"position"`
	Salary    int       `json:"salary"`
	ManagerID int       `json:"manager_id"`
}

type Point struct {
	X, Y int
}

func main() {
	var dilbert Employees

	dilbert.ID = 0
	dilbert.Name = "Dilbert"
	dilbert.Address = "Surabaya Indonesia"
	dilbert.DoB = time.Date(1997,
		time.November,
		11,
		0,
		0,
		0,
		0,
		time.UTC)
	dilbert.Position = SoftwareEngineer
	dilbert.Salary = 15_000_000
	dilbert.ManagerID = 0

	EmployeeById := func(id int) *Employees {
		if id == dilbert.ID {
			return &dilbert
		}
		return nil
	}

	EmployeeById(0).Salary += 5_000_000

	result, err := json.MarshalIndent(dilbert, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

	point := Point{
		X: 1,
		Y: 5,
	}

	if p, err := json.MarshalIndent(point, "", "   "); err == nil {
		fmt.Printf("Point %s\n", p)
	}
}
