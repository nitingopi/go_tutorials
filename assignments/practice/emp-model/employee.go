package main

import (
	"fmt"
	"reflect"
	"time"
)

type TeamMember interface {
	PrintName()
	PrintDetails()
}

type Employee struct {
	FirstName, LastName string
	Dob					time.Time 
	JobTitle, Location  string
}

func (e Employee) PrintName() {
	fmt.Printf("\n%s %s\n", e.FirstName, e.LastName)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Date of birth: %s, Job: %s, Location: %s", e.Dob, e.JobTitle, e.Location)
}

type Developer struct {
	Employee //type embedding for composition
	Skills []string
}

// override print details
func (d Developer) PrintDetails() {
	d.Employee.PrintDetails()
	fmt.Println("Technical skills:")
	for _, v := range d.Skills {
		fmt.Println(v)
	}
}

type Manager struct {
	Employee //type embedding for composition
	Projects  []string
	Locations []string
}

func ( m Manager) PrintDetails() {
	m.Employee.PrintDetails()
	fmt.Println("Projects :")
	for _, v := range m.Projects{
		fmt.Println(v)
	}
	fmt.Println("Managing teams for locations :")
	for _, v := range m.Locations{
		fmt.Println(v)
	}
}

func (m Manager) ToString(){
	fmt.Printf("%+v\n", m)
}

type Team struct {
	Name, Description string 
	TeamMembers       []TeamMember
}

func (t Team) PrintTeamDetails(){
	fmt.Printf("Team: %s - %s\n", t.Name, t.Description)
	fmt.Println("Details of team members: ")
	for _, v := range t.TeamMembers {
		//Get the concrete type name
		t := reflect.TypeOf(v)
		fmt.Println("\nType : ",t)
		v.PrintName()
		v.PrintDetails()
		if t.String() == "main.Manager" {
			m := v.(Manager) // Type assertion
			m.ToString()
		}
	}
}

func main(){
	steve := Developer{
		Employee: Employee{
			FirstName: "Steve",
			LastName:  "John",
			Dob:       time.Date(1990, time.February, 17, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Software Engineer",
			Location:  "San Fancisco",
		},
		Skills: []string {"Go", "Docker", "Kubernetes"},
	}
	irene := Developer{
		Employee: Employee{
			FirstName: "Irene",
			LastName:  "Rose",
			Dob:       time.Date(1991, time.January, 13, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Software Engineer",
			Location:  "Santa Clara",
		},
		Skills: []string{"Go", "MongoDB"},
	}
	alex := Manager{
		Employee: Employee{
			FirstName: "Alex",
			LastName:  "Williams",
			Dob:       time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Program Manger",
			Location:  "Santa Clara",
		},
		Projects:  []string{"CRM", "e-Commerce"},
		Locations: []string{"San Fancisco", "Santa Clara"},
	}

	// Create team
	team := Team{
		"Go",
		"Golang Engineering Team",
		[]TeamMember{steve, irene, alex},
	}
	// Get details of Team
	team.PrintTeamDetails()

}
