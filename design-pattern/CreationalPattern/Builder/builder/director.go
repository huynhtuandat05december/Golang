package builder

type director struct {
	builder iBuilder
}

func GetNewDirector(builder iBuilder) *director {
	return &director{
		builder: builder,
	}

}

func (d *director) SetBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) BuildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
