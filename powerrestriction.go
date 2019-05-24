package main

func (h *Handler) PowerRestriction(floor, subCorridor int) bool {

	if !h.PowerConsumptionCriteria(floor) {
		if _, ok := h.BuildHandler.SubCorridor[floor].CorridorCount[subCorridor-1]; ok {
			x := h.BuildHandler.SubCorridor[floor].CorridorCount[subCorridor-1]
			x.AC = false
			h.BuildHandler.SubCorridor[floor].CorridorCount[subCorridor-1] = x

			h.PowerHandler.Floor[floor].SubCorridor[subCorridor] = subCorridor - 1
			if !h.PowerConsumptionCriteria(floor) {
				return false
			}
		} else {
			if _, ok := h.BuildHandler.SubCorridor[floor].CorridorCount[subCorridor+1]; ok {
				x := h.BuildHandler.SubCorridor[floor].CorridorCount[subCorridor+1]
				x.AC = false
				h.BuildHandler.SubCorridor[floor].CorridorCount[subCorridor+1] = x

				h.PowerHandler.Floor[floor].SubCorridor[subCorridor] = subCorridor + 1
				if !h.PowerConsumptionCriteria(floor) {
					return false
				}
			}
		}
	} else {
		mappedCorridor := h.PowerHandler.Floor[floor].SubCorridor[subCorridor]
		if mappedCorridor != 0 {
			var x = h.BuildHandler.SubCorridor[floor].CorridorCount[mappedCorridor]
			x.AC = true
			h.BuildHandler.SubCorridor[floor].CorridorCount[mappedCorridor] = x
			h.PowerHandler.Floor[floor].SubCorridor[subCorridor] = 0
			if !h.PowerConsumptionCriteria(floor) {
				return false
			}
		}
	}

	return true
}
