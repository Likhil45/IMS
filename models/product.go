package models

import "fmt"

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Category string  `json:"category"`
}

func (p *Product) Id() int {
	return p.ID
}

func (p *Product) UpdatePrice(newPrice float64) {
	p.Price = newPrice
}

func (p *Product) Sell(quantity int) error {
	if p.Quantity != 0 {
		p.Quantity--
		return nil
	} else {
		return fmt.Errorf("no stock left")
	}
}

func (p *Product) Restock(quantity int) {
	p.Quantity++
}

func (p *Product) Display() string {
	return fmt.Sprintf("Id: %d\nName: %s\nPrice:%f\nCategory:%s\nQuantity left:%d", p.Id(), p.Name, p.Price, p.Category, p.Quantity)
}
