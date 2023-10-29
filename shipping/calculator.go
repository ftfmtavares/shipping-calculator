package shipping

type checkpoint struct {
	count     int
	excess    int
	packIndex int
}

func CalculateShipping(packSizes []int, order int) (packs []int, excessUnits int, packsCount int) {
	if len(packSizes) == 0 || order <= 0 {
		return make([]int, len(packSizes)), 0, 0
	}

	checkpoints := make([]checkpoint, order)
	packSize := packSizes[0]
	for i := 0; i < order; i++ {
		checkpoints[i].count = (i / packSize) + 1
		checkpoints[i].excess = checkpoints[i].count*packSize - i - 1
		checkpoints[i].packIndex = 0
	}

	for packIndex := 1; packIndex < len(packSizes); packIndex++ {
		packSize = packSizes[packIndex]
		for i := 0; i < order; i++ {
			newCheckPoint := checkpoint{}

			if i < packSize {
				newCheckPoint.count = 1
				newCheckPoint.excess = packSize - i - 1
				newCheckPoint.packIndex = packIndex
			}

			if i >= packSize {
				newCheckPoint.count = checkpoints[i-packSize].count + 1
				newCheckPoint.excess = checkpoints[i-packSize].excess
				newCheckPoint.packIndex = packIndex
			}

			if newCheckPoint.excess < checkpoints[i].excess || (newCheckPoint.excess == checkpoints[i].excess && newCheckPoint.count <= checkpoints[i].count) {
				checkpoints[i] = newCheckPoint
			}
		}
	}

	excessUnits = checkpoints[order-1].excess
	packsCount = checkpoints[order-1].count
	packs = make([]int, len(packSizes))
	for order > 0 {
		packs[checkpoints[order-1].packIndex]++
		order -= packSizes[checkpoints[order-1].packIndex]
	}

	return packs, excessUnits, packsCount
}
