package main

func (h *Handler) BuildHotelStructure() {

	h.BuildHandler = HotelStructure{}

	h.BuildHandler.BuildHotelFloor(h.InputHandler.Floor)
	h.BuildHandler.BuildHotelMainCorridor(h.InputHandler.Floor, h.InputHandler.MainCorridor)
	h.BuildHandler.BuildHotelSubCorridor(h.InputHandler.Floor, h.InputHandler.SubCorridor)
	h.BuildHandler.PrintHotelStructure(h.InputHandler.Floor, h.InputHandler.MainCorridor, h.InputHandler.SubCorridor)

}

func (h *Handler) PowerHotelStructure() {

	h.PowerHandler = FloorPowerRestrict{}
	h.PowerHandler.PowerBuilder(h.InputHandler.Floor, h.InputHandler.SubCorridor)
}
