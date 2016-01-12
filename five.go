package main

import "fmt"

func main() {

	var n1 Notifier

	n := Admin{"MyName", "Myemail"}
	n1 = &Admin{"MyName", "Myemail"}
	n.notify()
	d := 100
	// sendNotofication(n1)
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", n1)
}
func sendNotofication(n Notifier) {
	n.notify()
}

type Notifier interface {
	notify()
}
type Admin struct {
	Name  string
	Email string
}

// func (a Admin) notify() {
// 	fmt.Println(a.Email, a.Name)
// }

func (a *Admin) notify() {
	fmt.Println(a.Email, a.Name)
}

/** interface variable is a pointer type it will accpet any both T and T* receiever type **/
/** interface variable is a value type it will only accpet value type **/
/** Method type T does not consist of Method type T* **/
