//go:build !solution

package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	if len(guests) == 0 {
		return []Load{}
	}

	mapp_arrival := make(map[int]int)
	mapp_departure := make(map[int]int)

	max_departure_day := 0
	for _, guest := range guests {
		if guest.CheckOutDate > max_departure_day {
			max_departure_day = guest.CheckOutDate
		}

		mappDates(guest.CheckInDate, mapp_arrival)
		mappDates(guest.CheckOutDate, mapp_departure)
	}

	day := 0
	count := 0
	loads := make([]Load, 0)

	for day != max_departure_day+1 {

		arrived, _ := mapp_arrival[day]
		depatured, _ := mapp_departure[day]

		if arrived == depatured {
			day++
			continue
		}

		count = count + (arrived - depatured)
		loads = append(loads, Load{
			StartDate:  day,
			GuestCount: count,
		})
		day++
	}

	return loads
}

func mappDates(date int, mapDate map[int]int) {
	mapDate[date]++
}
