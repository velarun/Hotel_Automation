package main

import "fmt"

func (hs *HotelStructure) BuildHotelFloor(floorCount int) {
	for i := 1; i <= floorCount; i++ {
		hs.Floor = append(hs.Floor, i)
	}
}

func (hs *HotelStructure) BuildHotelMainCorridor(floorCount int, mainCorridorCount int) {

	for j := 1; j <= floorCount; j++ {
		cc := &CorridorStructure{CorridorCount: map[int]CorridorCommon{}}
		for i := 1; i <= mainCorridorCount; i++ {
			cc.CorridorCount[i] = CorridorCommon{Light: true, AC: true}
		}
		hs.MainCorridor = append(hs.MainCorridor, *cc)
	}
}

func (hs *HotelStructure) BuildHotelSubCorridor(floorCount int, subCorridorCount int) {
	for j := 1; j <= floorCount; j++ {
		cc := &CorridorStructure{CorridorCount: map[int]CorridorCommon{}}
		for i := 1; i <= subCorridorCount; i++ {
			mc := &CorridorCommon{Light: false, AC: true}
			cc.CorridorCount[i] = *mc
		}
		hs.SubCorridor = append(hs.SubCorridor, *cc)
	}
}

func boolCheck(boolVal bool) string {
	if boolVal {
		return "ON"
	}

	return "OFF"
}

func (hs *HotelStructure) PrintHotelStructure(floorCount int, mainCorridorCount int, subCorridorCount int) {
	fmt.Println()
	fmt.Println("Output from controller:")
	fmt.Println()
	for i := 0; i < floorCount; i++ {
		fmt.Printf("\t\tFloor %d\t\t\n", hs.Floor[i])
		for j := 1; j <= mainCorridorCount; j++ {
			fmt.Printf("Main corridor %d Light %d : %s AC : %s\n", j, j, boolCheck(hs.MainCorridor[i].CorridorCount[j].Light), boolCheck(hs.MainCorridor[i].CorridorCount[j].AC))
		}
		for k := 1; k <= subCorridorCount; k++ {
			fmt.Printf("Sub corridor %d Light %d : %s AC : %s\n", k, k, boolCheck(hs.SubCorridor[i].CorridorCount[k].Light), boolCheck(hs.SubCorridor[i].CorridorCount[k].AC))
		}
	}
}
