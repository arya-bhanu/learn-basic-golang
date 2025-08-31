package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	num := 23
	reflectVal := reflect.ValueOf(num)
	fmt.Printf("Reflectval %+v\n", reflectVal)
	fmt.Printf("Reflectval type %+v\n", reflectVal.Kind())
	if reflectVal.Type().String() == "int" {
		fmt.Printf("Reflectval casted into int %+v\n", reflectVal.Int())
		castedInterface := reflectVal.Interface().(int)
		fmt.Printf("CastedInterface %+v\n", castedInterface)
	}

	student1 := student{
		Nama:  "putu arya",
		Nilai: 100,
	}

	student1.getPropertyInfo()

	student2 := student{
		Nama:  "John Wick",
		Nilai: 100,
	}

	reflectStudent2 := reflect.ValueOf(&student2)
	fmt.Println(reflectStudent2.Type())
	method := reflectStudent2.MethodByName("SetName")
	fmt.Printf("Method student 2: %+v\n", method)
	method.Call([]reflect.Value{
		reflect.ValueOf("Aryo"),
	})
	fmt.Printf("Nama: %+v\n", student2.Nama)

	book1 := Book{
		Title:  "Buku",
		Author: "Putu Gde Arya Bhanuartha",
		Price:  100,
	}
	err := Validate(book1)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func (s *student) getPropertyInfo() {
	reflectVal := reflect.ValueOf(s)

	if reflectVal.Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	reflectType := reflectVal.Type()

	for i := range reflectVal.NumField() {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectVal.Field(i))
		if reflectVal.Field(i).Kind() == reflect.Int {
			fmt.Println("Operasi nilai jika integer (+10): ", reflectVal.Field(i).Interface().(int)+10)
		}
		fmt.Println()
	}

}

type student struct {
	Nama  string
	Nilai int
}

func (s *student) SetName(newName string) {
	s.Nama = newName
}

func Validate(i interface{}) error {
	reflectVal := reflect.ValueOf(i)

	// it must be original (non-pointer) value
	// it must be a struct type object
	if reflectVal.Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	if reflectVal.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, got %s", reflectVal.Kind().String())
	}

	reflectType := reflectVal.Type()

	for i := range reflectVal.NumField() {
		field := reflectType.Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}
		rules := strings.Split(tag, ",")
		for _, r := range rules {
			switch {
			case r == "required":
				if reflectVal.Field(i).IsZero() {
					return fmt.Errorf("field %s is required", field.Name)
				}
			case strings.HasPrefix(r, "min="):
				val := strings.TrimPrefix(r, "min=")
				min, err := strconv.Atoi(val)
				if err != nil {
					return fmt.Errorf("error converting min value into integer: provided (%s)", val)
				}
				if reflectType.Field(i).Type.Kind() != reflect.String {
					return fmt.Errorf("min length should be applied on type string, provided type (%s) on field (%s)", field.Type, field.Name)
				}
				strCompare := reflectVal.Field(i).Interface().(string)
				if len(strCompare) <= min {
					return fmt.Errorf("string length of (%s) shouldn't less than (%d) characters", field.Name, min)
				}
			}

		}
	}

	return nil
}

type Book struct {
	Title  string `validate:"required,min=5"`
	Author string `validate:"required,max=12"`
	Price  int    `validate:"required"`
}
