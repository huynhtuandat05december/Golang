package builder

type house struct {
	windowType string
	doorType   string
	floor      int
}

func (house *house) GetWindowType() string {
	return house.windowType
}

func (house *house) GetDoorType() string {
	return house.doorType
}

func (house *house) GetFloor() int {
	return house.floor
}
