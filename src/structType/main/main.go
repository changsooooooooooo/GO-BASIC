package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Name string
	ID string
	Age int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price int
}

type Product struct {
	Name string
	Price int
	ReviewScore float64
}

type Actor struct {
	Name string
	HP int
	Speed float64
}

type Monster struct {
	Actor
	Attack int
	Speed int
}

func main(){
	user := User{"송하나", "hana", 23}
	vip := VIPUser{User{"화랑", "hwarang", 40},
		3,
		250,
	}
	fmt.Printf("유저: %s ID: %s 나이 : %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIPLevel: %d VIP price: %d", vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)


	monster:=Monster{
		Actor{"NPCA", 100, 8.7},
	500,
	200,
	}
	fmt.Println(monster.Speed)
	fmt.Println(monster.Name)
	fmt.Println(monster.Actor.Speed)
	fmt.Println(unsafe.Sizeof(10))
}
