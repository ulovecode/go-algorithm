package main

type ticketInfo struct {
	stationName string
	checkInTime int
}

type peopeleAge struct {
	avgTime int
	number  int
}
type UndergroundSystem struct {
	statisticalSys map[string]map[string]peopeleAge
	ticketSys      map[int]ticketInfo
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		statisticalSys: make(map[string]map[string]peopeleAge),
		ticketSys:      make(map[int]ticketInfo, 0),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.ticketSys[id] = ticketInfo{
		stationName: stationName,
		checkInTime: t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	ticketInfo := this.ticketSys[id]
	delete(this.ticketSys, id)
	if _, ok := this.statisticalSys[ticketInfo.stationName]; !ok {
		this.statisticalSys[ticketInfo.stationName] = make(map[string]peopeleAge, 0)
	}
	avgTime := t - ticketInfo.checkInTime
	if prePeopleAvgTime, ok := this.statisticalSys[ticketInfo.stationName][stationName]; ok {
		this.statisticalSys[ticketInfo.stationName][stationName] = peopeleAge{
			avgTime: prePeopleAvgTime.avgTime + avgTime,
			number:  prePeopleAvgTime.number + 1,
		}

	} else {
		this.statisticalSys[ticketInfo.stationName][stationName] = peopeleAge{
			avgTime: avgTime,
			number:  1,
		}
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	avgPeople := this.statisticalSys[startStation][endStation]
	return float64(avgPeople.avgTime) / float64(avgPeople.number)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
