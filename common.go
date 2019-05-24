package main

const (
	bulbPowerConsumption = 5
	acPowerConsumption   = 10
)

type HotelInput struct {
	Floor        int
	MainCorridor int
	SubCorridor  int
}

type HotelStructure struct {
	Floor        []int
	MainCorridor []CorridorStructure
	SubCorridor  []CorridorStructure
}

type CorridorStructure struct {
	CorridorCount map[int]CorridorCommon
}

type CorridorCommon struct {
	Light bool
	AC    bool
}

type Handler struct {
	BuildHandler HotelStructure
	InputHandler HotelInput
	PowerHandler FloorPowerRestrict
}

type FloorPowerRestrict struct {
	Floor []CorridorPowerRestrict
}

type CorridorPowerRestrict struct {
	SubCorridor map[int]int
}
