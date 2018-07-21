package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "yourtion",
		"course":  "golang",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName) // get Zero value
	} else {
		fmt.Println("ket does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
