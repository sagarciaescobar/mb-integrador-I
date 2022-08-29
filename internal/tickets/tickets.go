package tickets

import (
	"time"

	loadData "github.com/sagarciaescobar/mb-integrador-I/internal"
)

func GetTotalTickets(destination string) (int, error) {
	tick, _ := loadData.TicketData()
	var filteredData []loadData.Ticket
	for _, v := range tick {
		if v.Country == destination {
			filteredData = append(filteredData, v)
		}
	}
	var length int = len(filteredData)
	return length, nil
}

// ejemplo 2
func GetBeTimeTickes(option string) (int, error) {
	tick, _ := loadData.TicketData()
	var length int = 0
	for _, v := range tick {
		madrugada, _ := time.ParseDuration("6h")
		manana, _ := time.ParseDuration("12h")
		tarde, _ := time.ParseDuration("19h")
		noche, _ := time.ParseDuration("23h")

		if option == "madrugada" {
			if v.Time.Hours() < madrugada.Hours() {
				length++
			}
		}

		if option == "manana" {
			if v.Time.Hours() < manana.Hours() && v.Time.Hours() > madrugada.Hours() {
				length++
			}
		}
		if option == "tarde" {
			if v.Time.Hours() < tarde.Hours() && v.Time.Hours() > manana.Hours() {
				length++
			}
		}
		if option == "noche" {
			if v.Time.Hours() < noche.Hours() && v.Time.Hours() > tarde.Hours() {
				length++
			}
		}
	}
	return length, nil
}

// ejemplo 3
func AverageDestination(destination string) (float32, error) {
	tick, _ := loadData.TicketData()
	var filteredData []loadData.Ticket
	for _, v := range tick {
		if v.Country == destination {
			filteredData = append(filteredData, v)
		}
	}
	var length float32 = float32(len(filteredData))
	promedio := length / float32(len(tick))
	return promedio, nil
}
