package main

import (
    "encoding/json"
    "fmt"
    "reflect"
)

func main () {
    str := `
    {
        "name": "sam",
        "age": 30,
        "married": false,
        "details": {
            "salary": 10000
        },
        "siblings": [
            "john",
            "mary"
        ]
    }
    `
    
    var data map[string]interface {}
    
    err := json.Unmarshal([]byte(str), &data)
    
    if err!=nil {
        panic(err)
    }
    
    fmt.Println(data)
    
    fmt.Println("Name type: ", reflect.TypeOf(data["name"]))
    fmt.Println("typeof typeof: ", reflect.TypeOf(reflect.TypeOf(data["name"])))
    fmt.Println("Name (using some automatic conversion to string) =", data["name"])
    
    var name     string  = data["name"].(string)
    var age      int     = int(data["age"].(float64))
    var married  bool    = data["married"].(bool)
    var salary   int     = int(data["details"].(map[string]interface{})["salary"].(float64))
    var siblings []interface{} = data["siblings"].([]interface{})
    
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Married:", married)
    fmt.Println("Salary:", salary)
    
    for i, s := range siblings {
        fmt.Printf("Sibling[%d] = %s\n", i, s.(string))
    }
}
