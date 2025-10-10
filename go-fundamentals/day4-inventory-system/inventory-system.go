package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

type Inventory struct {
	Products map[string]Product // key = Product ID
}

func (r *Inventory) AddProduct(id string, product Product) error {
	if _, exists := r.Products[id]; exists {
		return fmt.Errorf("⚠️ Ya existe un producto con ID %s: %s", id, product.Name)
	}
	r.Products[id] = product
	fmt.Printf("✅ Producto agregado: %s (ID: %s, Precio: $%.2f, Stock: %d)\n",
		product.Name, id, product.Price, product.Stock)
	return nil
}

func (r *Inventory) UpdateStock(id string, quantity int) error {
	product, exists := r.Products[id]
	if !exists {
		return fmt.Errorf("producto con ID %s no encontrado", id)
	}
	product.Stock = quantity
	r.Products[id] = product
	return nil
}

func (r *Inventory) GetTotalValue() float64 {
	var total float64
	for _, product := range r.Products {
		total += product.Price * float64(product.Stock)
	}
	return total
}

func (r *Inventory) RemoveProduct(id string) error {
	_, exists := r.Products[id]
	if !exists {
		return fmt.Errorf("producto con ID %s no encontrado", id)
	}
	delete(r.Products, id)
	return nil
}

func (r *Inventory) ListLowStock(threshold int) ([]string, error) {
	var ids []string
	for id, product := range r.Products {
		if product.Stock < threshold {
			ids = append(ids, id)
			fmt.Printf("%s, con stock: %d\n", product.Name, product.Stock)
		}
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("no hay productos con stock bajo")
	}
	return ids, nil
}

func NewInventory() *Inventory {
	return &Inventory{Products: make(map[string]Product)}
}

func main() {
	inv := NewInventory()
	// Slice de productos con datos épicos
	products := []Product{
		{Name: "Teclado del Trueno", Price: 12500.99, Stock: 10},
		{Name: "Mouse de las Sombras", Price: 8999.50, Stock: 5},
		{Name: "Monitor del Vacío", Price: 45200.00, Stock: 3},
		{Name: "Auriculares del Caos", Price: 15999.99, Stock: 8},
		{Name: "Silla del Dominio", Price: 74999.00, Stock: 2},
	}

	for i, p := range products {
		id := fmt.Sprintf("P%03d", i+1)
		err := inv.AddProduct(id, p)
		if err != nil {
			fmt.Println("Error al agregar producto:", err)
		}
	}

	fmt.Printf("\n💰 Valor total del inventario: $%.2f\n", inv.GetTotalValue())

	lowStock, err := inv.ListLowStock(5)
	if err != nil {
		fmt.Println("✅ Todo con buen stock.")
	} else {
		for _, item := range lowStock {
			fmt.Println("⚠️ Bajo stock:", item)
		}
	}

	if err != nil {
		fmt.Println("✅ Todo con buen stock.")
	} else {
		for _, id := range lowStock {
			inv.UpdateStock(id, 10)
			fmt.Println("🔧 Reposición aplicada a:", id)
		}
	}

	fmt.Println(inv.ListLowStock(5))

}
