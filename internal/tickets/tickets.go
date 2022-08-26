package tickets

import loadData "github.com/sagarciaescobar/mb-integrador-I/internal"

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	tick, _ := loadData.TicketData()
	var length int = len(tick)
	return length, nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	return 2, nil
}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	return 3, nil
}
