package main

func (p *FloorPowerRestrict) PowerBuilder(floor, subCorridor int) {
	for j := 1; j <= floor; j++ {
		cp := make(map[int]int)
		cc := &CorridorPowerRestrict{}
		for i := 1; i <= subCorridor; i++ {
			cp[i] = 0
		}
		cc.SubCorridor = cp
		p.Floor = append(p.Floor, *cc)
	}
}
