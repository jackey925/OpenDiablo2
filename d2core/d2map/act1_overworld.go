package d2map

import (
	"math/rand"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2wilderness"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func (m *MapEngine) GenerateAct1Overworld(cacheTiles bool) {
	rand.Seed(m.seed)
	region, entities := loadRegion(m.seed, 0, 0, d2enum.RegionAct1Town, 1, -1, cacheTiles)
	m.regions = append(m.regions, region)
	m.entities.Add(entities...)
	if strings.Contains(region.regionPath, "E1") {
		region, entities := loadRegion(m.seed, region.tileRect.Width, 0, d2enum.RegionAct1Town, 2, -1, cacheTiles)
		m.AppendRegion(region)
		m.entities.Add(entities...)
	} else if strings.Contains(region.regionPath, "S1") {
		yOffset := region.tileRect.Height
		waterXOffset := region.tileRect.Width - 16
		region, entities := loadRegion(m.seed, 0, yOffset, d2enum.RegionAct1Town, 3, -1, cacheTiles)
		m.AppendRegion(region)
		m.entities.Add(entities...)
		yOffset += region.tileRect.Height

		for i := 0; i < 8; i++ {
			// West Border
			region, entities = loadRegion(m.seed, 0, yOffset, d2enum.RegionAct1Wilderness, d2wilderness.TreeBorderWest, 0, cacheTiles)
			m.AppendRegion(region)
			m.entities.Add(entities...)

			// East Border
			region, entities = loadRegion(m.seed, waterXOffset, yOffset, d2enum.RegionAct1Wilderness, d2wilderness.WaterBorderEast, 0, cacheTiles)
			m.AppendRegion(region)
			m.entities.Add(entities...)

			yOffset += region.tileRect.Height
		}

	}

}
