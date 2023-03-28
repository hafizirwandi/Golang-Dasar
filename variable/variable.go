/*
In Go, there are different types of variables, for example:

- int- stores integers (whole numbers), such as 123 or -123
- float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
- string - stores text, such as "Hello World". String values are surrounded by double quotes
- bool- stores values with two states: true or false

Go variable naming rules:

- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords
*/
package main

import "fmt"

func main(){
	// var b1 bool = true // typed declaration with initial value
	// var b2 = true // untyped declaration with initial value
	// var b3 bool // typed declaration without initial value
	// b4 := true // untyped declaration with initial value
	//membuat variable di go bisa menggunakan kata kunci "var" atau ":="
	//dengan var
	var name ="Hafiz Irwandi"
	fmt.Println("Hai, Nama",name)

	name ="Prestika Dita Sahara"
	fmt.Println("Hai, Nama Saya ",name)

	//dengan :=
	tempat_lahir:="Medan"
	fmt.Println("Saya ",name," Lahir di ",tempat_lahir)
}

