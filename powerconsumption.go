package main

func (h *Handler) PowerConsumptionCriteria(floor int) bool {
	return h.PowerConsumptionPerFloor(floor) <= h.MaxPowerConsumptionAllowedPerFloor()
}

func (h *Handler) MaxPowerConsumptionAllowedPerFloor() int {
	return (h.InputHandler.MainCorridor * 15) + (h.InputHandler.SubCorridor * 10)
}

func (h *Handler) PowerConsumptionPerFloor(floor int) int {

	var powerConsumption int

	for i := 1; i <= h.InputHandler.MainCorridor; i++ {
		if h.BuildHandler.MainCorridor[floor].CorridorCount[i].Light {
			powerConsumption += bulbPowerConsumption
		}

		if h.BuildHandler.MainCorridor[floor].CorridorCount[i].AC {
			powerConsumption += acPowerConsumption
		}
	}

	for i := 1; i <= h.InputHandler.SubCorridor; i++ {
		if h.BuildHandler.SubCorridor[floor].CorridorCount[i].Light {
			powerConsumption += bulbPowerConsumption
		}

		if h.BuildHandler.SubCorridor[floor].CorridorCount[i].AC {
			powerConsumption += acPowerConsumption
		}
	}

	return powerConsumption
}
