package main

import (
	"fmt"
)

// Item yapısı, ürünleri temsil eder
type Item struct {
	Name     string
	Price    float64
	Discount float64
}

// Describable interface'i, Item yapısına özgü bir metot içerir
type Describable interface {
	Description() string
}

// Description metodu, Item yapısı için Describable interface'ini uygular
func (item Item) Description() string {
	discountedPrice := calculatePrice(item)
	if item.Discount > 0 {
		return fmt.Sprintf("%s - %.2f TL (%.0f%% indirimle %.2f TL)", item.Name, item.Price, item.Discount, discountedPrice)
	}
	return fmt.Sprintf("%s - %.2f TL", item.Name, item.Price)
}

// calculatePrice fonksiyonu, bir ürünün fiyatını hesaplar
func calculatePrice(item Item) float64 {
	return item.Price - (item.Price * item.Discount / 100)
}

// totalPrice fonksiyonu, ürün listesinin toplam fiyatını hesaplar
func totalPrice(items []Item) float64 {
	var total float64
	for _, item := range items {
		total += calculatePrice(item)
	}
	return total
}

func main() {
	// Ürün listesini oluştur
	items := []Item{
		{Name: "Elma", Price: 0.75, Discount: 10},
		{Name: "Portakal", Price: 1.00, Discount: 0},
	}

	// Ürünleri ekrana yazdır
	for _, item := range items {
		fmt.Println(item.Description())
	}

	// Toplam fiyatı ekrana yazdır
	fmt.Printf("Toplam Fiyat: %.2f TL\n", totalPrice(items))
}
