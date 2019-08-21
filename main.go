package main

import (
	"fmt"
	"reflect"
)

func main() {
	var dog dog

	//Pass in reference to dog
	err := Get("Rover", &dog)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(dog.String())
	return
}

//Get takes an name and an interface{} to populate if found and
//returns an error if nothing is found
func Get(name string, entity interface{}) error {
	m := map[string]dog{
		"Rover": dog{Name: "Rover"},
	}

	doggy, ok := m[name]
	if !ok {
		return fmt.Errorf("Dog called %s has not been found", name)
	}

	//takes the dog struct and converts it to a reflect.Value{} struct
	val := reflect.ValueOf(doggy)

	//the passed in interface is converted to a reflect.Value struct
	//which allows the updating of the passed in dog (entity) with the value
	//of the doggy returned from the map
	reflect.ValueOf(entity).Elem().Set(val)

	return nil
}

type dog struct {
	Name string
}

func (d *dog) String() string {
	return fmt.Sprintf("My name is %s", d.Name)
}
