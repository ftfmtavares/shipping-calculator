package database

import (
	"fmt"
	"sort"
)

var defaultPackSizes = []int{
	250,
	500,
	1000,
	2000,
	5000,
}

type Product struct {
	ID        int
	PackSizes []int
}

type Products struct {
	products map[int]Product
}

func NewProducts() Products {
	return Products{
		products: map[int]Product{
			1: {
				ID:        1,
				PackSizes: defaultPackSizes,
			},
		},
	}
}

func (p Products) GetPackSizes(productID int) (empty []int, err error) {
	_, exists := p.products[productID]
	if !exists {
		return empty, fmt.Errorf("product not found")
	}
	return p.products[productID].PackSizes, nil
}

func (p *Products) AddPackSize(productID int, size int) (err error) {
	_, exists := p.products[productID]
	if !exists {
		return fmt.Errorf("product not found")
	}

	for _, packSize := range p.products[productID].PackSizes {
		if packSize == size {
			return fmt.Errorf("product already has that pack size")
		}
	}

	newPackSizes := append(p.products[productID].PackSizes, size)
	sort.Slice(newPackSizes, func(i, j int) bool {
		return newPackSizes[i] < newPackSizes[j]
	})
	p.products[productID] = Product{
		ID:        productID,
		PackSizes: newPackSizes,
	}

	return nil
}

func (p *Products) DeletePackSize(productID int, size int) (err error) {
	_, exists := p.products[productID]
	if !exists {
		return fmt.Errorf("product not found")
	}

	for i, packSize := range p.products[productID].PackSizes {
		if packSize == size {
			newPackSizes := append(p.products[productID].PackSizes[:i], p.products[productID].PackSizes[i+1:]...)
			p.products[productID] = Product{
				ID:        productID,
				PackSizes: newPackSizes,
			}
			return nil
		}
	}

	return fmt.Errorf("product doesn't have that pack size")
}
