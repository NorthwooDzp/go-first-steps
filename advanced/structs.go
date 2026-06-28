package advanced

import "fmt"

type Skill struct {
	name  string
	level int
}

type Person struct {
	Name      string   // Exported field, accessible outside current package
	Age       uint     // Exported field, accessible outside current package
	sports    []string // private field, accessible only inside advanced pkg
	skills    []Skill  // embedded structs slice
	mainSkill Skill
}

func getName(p Person) string {
	return p.Name
}

func (p Person) printName() {
	p.Name = p.Name + " changed" // this change affects only this function's copy of the value
	fmt.Println("Name inside function without referrence =>", p.Name)
}

func (p *Person) changeAge(newAge uint) {
	p.Age = newAge
	fmt.Println("New age in correct method =>", p.Age)
}

func Structs() {
	var p1 Person = Person{Name: "Volodymyr", Age: 39, sports: []string{"Tennis", "Swimming"}} // order can be any, some fields could be not specified and will take default values
	fmt.Println("Person 1 =>", p1)

	p2 := Person{"Someone", 32, []string{"Football", "Basketball"}, []Skill{{"Gameing", 10}}, Skill{"Science", 2}} // require correct order for props
	fmt.Println("Person 2 =>", p2)
	p2.Name = "Someanother"
	fmt.Println("Person 2 after reassignment name =>", p2)

	fmt.Println("Sending params to func =>", getName(p1), "|", getName(p2))

	/* Methods that do not allow to change the original struct (value receivers) */
	p1.printName()
	p2.printName()
	fmt.Println("Trying to change name without referrences =>", p1, "|", p2)

	/* Pointer receivers struct methods, change original struct */
	p1.changeAge(37)
	p2.changeAge(34)
	fmt.Println("Changed structs => ", p1, "|", p2)
}
