package structutils

import (
	"fmt"
	"reflect"
	structs "github.com/fatih/structs"
)

func ListType (object interface{}){
	fmt.Println("\nObject Type: ", reflect.TypeOf(object))
}

func ListMethods(object interface{} ) {
	objectType := reflect.TypeOf(object)
	for i := 0; i < objectType.NumMethod(); i++ {
		method := objectType.Method(i)
		fmt.Println(method.Name)
	}
}

func ListFields (object interface{}) {
	//List fields of a Struct. // https://stackoverflow.com/questions/24337145/get-name-of-struct-field-using-reflection
	fields := structs.Names(&object)
	fmt.Println(fields)
}