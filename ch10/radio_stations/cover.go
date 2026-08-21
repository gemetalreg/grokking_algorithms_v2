package radiostations

func Cover() {

	for len(states_needed) > 0 {
		best_station := ""
		states_covered := make(set)
		for station, states_for_station := range stations {
			covered := *(states_needed.Intersect(&states_for_station))
			if len(covered) > len(states_covered) {
				best_station = station
				states_covered = covered
			}
		}
		final_stations[best_station] = struct{}{}
		for state := range states_covered {
			delete(states_needed, state)
		}
	}
}
