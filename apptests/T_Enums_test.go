package apptests

import (
	"fmt"
	"testing"
)

// import data "trustkbb.de/tools/dbtools/appdata"

// func (s Season) show() string{
// 	fmt.Println("Jahreszeit")
// }

type Season int

const (
	Spring Season = iota + 1
	Sommer
	Autumn
	Winter
)

type Capacity int

const (
	KB Capacity = 1024
	MB          = 1024 * 1024
	GB          = 1024 * 1024 * 1024
	TB          = 1024 * 1024 * 1024 * 1024
)

func (c Capacity) show() {
	fmt.Printf("Capacity (in Byte) %-15d ", c)
	switch {
	case c < KB:
		fmt.Printf("enspricht Byte\n")
	case c < MB && c >= KB:
		fmt.Printf("enspricht KB\n")
	case c < GB && c >= MB:
		fmt.Printf("enspricht MB\n")
	case c < TB && c >= GB:
		fmt.Printf("enspricht GB\n")
	case c >= TB:
		fmt.Printf("enspricht TB oder größer\n")

	}

}

func (s Season) String() string {
	return [...]string{"Frühling", "Sommer", "Herbst", "Winter"}[s-1]
}

func (s Season) Show() {
	switch s {
	case 1:
		fmt.Println("Frühling")
	case 2:
		fmt.Println("Sommer")
	case 3:
		fmt.Println("Herbst")
	case 4:
		fmt.Println("Winter")
	}
}

func TestEnumsS(t *testing.T) {
	var s Season
	s = Spring
	fmt.Println("Jahreszeit:", s.String())

	s = Autumn
	fmt.Println("Jahreszeit:", s.String())
	s.Show()

	fmt.Println("Season", Spring, Sommer, Autumn, Winter)
	// if r != zeichen[0] {
	// 	t.Errorf("Expected Error, result must be %s\n", zeichen[0])
	// }

}

func TestEnumsC(t *testing.T) {
	
	values := []Capacity{25400, 20000000, 3450, 3456700000000,56700000000}
	for _, v := range values {
		v.show()
	}

}
