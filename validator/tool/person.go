package tool

type Person struct {
	Name string `check:"string,email"`
	Age  int    `check:"int,1-100"`
}
