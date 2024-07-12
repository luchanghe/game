package user

type HeroSeat struct {
	SeatId   int
	SeatName string
}

func NewHeroSeat(seatId int, seatName string) *HeroSeat {
	o := &HeroSeat{
		SeatId:   seatId,
		SeatName: seatName,
	}
	return o
}
