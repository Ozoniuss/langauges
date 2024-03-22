package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

type Laptop struct {
	Id      int
	Brand   string
	Price   float64
	Details string
}

func (l Laptop) toCSVString() string {
	return fmt.Sprintf("%d,%s,%.2f,%s", l.Id, l.Brand, l.Price, l.Details)
}

const (
	MAX_PRICE = 50_000
	MIN_PRICE = 1_000

	NUM_LAPTOPS = 100_000
)

// stackoverflow copy paste
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.IntN(len(letterBytes))]
	}
	return string(b)
}

var brands = []string{"lenovo", "dell", "hp", "system76", "tuxedo", "acer", "microsoft", "macbook", "framework", "asus", "samsung", "huawei"}

func generate() []Laptop {
	var laptops = make([]Laptop, 0, 100_000)
	for i := range 999_999 {
		brand_idx := rand.IntN(len(brands))
		l := Laptop{
			Id:      i + 1,
			Brand:   brands[brand_idx],
			Details: RandStringBytes(64),
			Price:   rand.Float64()*8000 + 1000,
		}
		laptops = append(laptops, l)
	}
	return laptops
}

func main() {
	f, err := os.Create("laptops.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("id,brand,price,details\n")

	laptops := generate()
	for idx, l := range laptops {
		if idx != len(laptops)-1 {
			f.WriteString(l.toCSVString() + "," + "\n")
		} else {
			f.WriteString(l.toCSVString() + "\n")
		}
	}
}
