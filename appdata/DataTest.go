package appdata

import "fmt"

type Basic struct{
	name string
} 

type Paar struct{
	name string
	vorname string

	} 

type Car struct {
	name string

}

type Listen struct {
	name string
	Items []interface{}
}


// func (c Car) showName (){
// 	fmt.Println(c.name)
// }

func (p Paar) showName() {
	fmt.Println(p.name)
}
// func (b Basic) showName(){
// 	fmt.Println(b.name)
// }


func ListTestData (t interface{}){

	// t.showName()
fmt.Printf("(%v, %T)\n", t, t)
// fmt.Println("Name:", t.name)

}
