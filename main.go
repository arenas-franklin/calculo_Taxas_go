package main

import "github.com/arenas-franklin/go-intensivo-01/internal/entity"

type Car struct {
	Model string
	Color string
}

//metodo
func (c Car)Start(){
	println(c.Model + " has been Start")
}

func (c *Car) ChangeColor(color string){
	c.Color = color
	println("new color: " + c.Color)
}


//função
// func soma(x, y int)int{
// 	return x + y
// }
// 
func main(){

	order,err := entity.NewOrder("1", -10, 1)
	if err != nil{
		println(err.Error())
	}else{
		println(order.ID)
	}
}