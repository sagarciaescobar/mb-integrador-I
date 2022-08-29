package loadData

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ticket struct {
	Id      int64
	Name    string
	Email   string
	Country string
	Time    time.Duration
	Flight  string
}

func readData() ([]byte, error) {
	data, err := os.ReadFile("tickets.csv")
	if err != nil {
		fmt.Printf("%v", err)
		log.Fatal("no se pudo leer el archivo")
		return nil, err
	}
	return data, nil
}

func TicketData() ([]Ticket, error) {
	data, _ := readData()
	var formatedData []Ticket
	dataSlice := strings.Split(string(data), "\n")
	for i, v := range dataSlice {
		if i < len(dataSlice)-1 {
			var dato []string = strings.Split(v, ",")
			id, _ := strconv.ParseInt(dato[0], 10, 64)
			var timeData []string = strings.Split(dato[4], ":")
			hour, _ := time.ParseDuration(timeData[0] + "h" + timeData[1] + "m")
			ticket := Ticket{Id: id, Name: dato[1], Email: dato[2], Country: dato[3], Time: hour, Flight: dato[5]}
			formatedData = append(formatedData, ticket)
		}
	}

	return formatedData, nil
}
