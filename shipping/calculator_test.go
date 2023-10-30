package shipping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShippingCalculateShipping(t *testing.T) {
	testCases := []struct {
		desc                string
		packSizes           []int
		order               int
		expectedPacks       []int
		expectedExcessItems int
		expectedPacksCount  int
	}{
		{
			desc:                "no defined packsizes",
			packSizes:           []int{},
			order:               1,
			expectedPacks:       []int{},
			expectedExcessItems: 0,
			expectedPacksCount:  0,
		},
		{
			desc:                "no order",
			packSizes:           []int{250, 500, 1000, 2000, 5000},
			order:               0,
			expectedPacks:       []int{0, 0, 0, 0, 0},
			expectedExcessItems: 0,
			expectedPacksCount:  0,
		},
		{
			desc:                "1 ordered item",
			packSizes:           []int{250, 500, 1000, 2000, 5000},
			order:               1,
			expectedPacks:       []int{1, 0, 0, 0, 0},
			expectedExcessItems: 249,
			expectedPacksCount:  1,
		},
		{
			desc:                "250 ordered items",
			packSizes:           []int{250, 500, 1000, 2000, 5000},
			order:               250,
			expectedPacks:       []int{1, 0, 0, 0, 0},
			expectedExcessItems: 0,
			expectedPacksCount:  1,
		},
		{
			desc:                "251 ordered items",
			packSizes:           []int{250, 500, 1000, 2000, 5000},
			order:               251,
			expectedPacks:       []int{0, 1, 0, 0, 0},
			expectedExcessItems: 249,
			expectedPacksCount:  1,
		},
		{
			desc:                "501 ordered items",
			packSizes:           []int{250, 500, 1000, 2000, 5000},
			order:               501,
			expectedPacks:       []int{1, 1, 0, 0, 0},
			expectedExcessItems: 249,
			expectedPacksCount:  2,
		},
		{
			desc:                "12001 ordered items",
			packSizes:           []int{250, 500, 1000, 2000, 5000},
			order:               12001,
			expectedPacks:       []int{1, 0, 0, 1, 2},
			expectedExcessItems: 249,
			expectedPacksCount:  4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			packs, excessUnits, packsCount := CalculateShipping(tC.packSizes, tC.order)
			assert.Equal(t, tC.expectedPacks, packs, "packs selection should match")
			assert.Equal(t, tC.expectedExcessItems, excessUnits, "excess units should match")
			assert.Equal(t, tC.expectedPacksCount, packsCount, "packs count should match")
		})
	}
}
