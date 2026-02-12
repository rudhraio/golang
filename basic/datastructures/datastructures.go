package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// -- 1. Arrays --
	// Fixed size, homogeneous data structure. Size is part of type definition.
	fmt.Println("\n\033[1;36m=== 1. ARRAYS ===\033[0m")
	var emptyArray [5]int                 // Empty array (zero values)
	filledArray := [5]int{2, 5, 7, 9, 11} // Filled array
	fmt.Println("Empty Array:", emptyArray, "| Len:", len(emptyArray), "| Size:", unsafe.Sizeof(emptyArray), "bytes")
	fmt.Println("Filled Array:", filledArray, "| Len:", len(filledArray), "| Size:", unsafe.Sizeof(filledArray), "bytes")

	// Array operations
	fmt.Println("Operations: Index [0]:", filledArray[0], "| Index [2]:", filledArray[2])
	filledArray[1] = 100
	fmt.Println("After Update [1]=100:", filledArray)

	// -- 2. Slices --
	// Built on top of arrays, but resizable and more flexible. Can be nil.
	fmt.Println("\n\033[1;36m=== 2. SLICES ===\033[0m")
	var nilSlice []string    // Nil slice (no backing array)
	emptySlice := []string{} // Empty slice (backing array exists)
	filledSlice := []string{"this", "is", "a", "slice"}
	filledSlice = append(filledSlice, "example")
	fmt.Println("Nil Slice:", nilSlice, "| Len:", len(nilSlice), "| Cap:", cap(nilSlice), "| Size:", unsafe.Sizeof(nilSlice), "bytes")
	fmt.Println("Empty Slice:", emptySlice, "| Len:", len(emptySlice), "| Cap:", cap(emptySlice), "| Size:", unsafe.Sizeof(emptySlice), "bytes")
	fmt.Println("Filled Slice:", filledSlice, "| Len:", len(filledSlice), "| Cap:", cap(filledSlice), "| Size:", unsafe.Sizeof(filledSlice), "bytes")

	// Slice operations
	fmt.Println("Operations:")
	fmt.Println("  Append:", append(filledSlice, "more"), "| Len:", len(filledSlice))
	fmt.Println("  Index [1]:", filledSlice[1])
	fmt.Println("  Slice [1:3]:", filledSlice[1:3])
	fmt.Println("  Slice [:2]:", filledSlice[:2])
	filledSlice[0] = "modified"
	fmt.Println("  After Update [0]:", filledSlice)

	// -- 3. Maps --
	// Key-value pairs, similar to dictionaries. Can be nil.
	fmt.Println("\n\033[1;36m=== 3. MAPS ===\033[0m")
	var nilMap map[string]int        // Nil map
	emptyMap := make(map[string]int) // Empty map (initialized)
	filledMap := make(map[string]int)
	filledMap["success"] = 200
	filledMap["error"] = 400
	filledMap["failed"] = 500
	fmt.Println("Nil Map:", nilMap, "| Len:", len(nilMap), "| Size:", unsafe.Sizeof(nilMap), "bytes")
	fmt.Println("Empty Map:", emptyMap, "| Len:", len(emptyMap), "| Size:", unsafe.Sizeof(emptyMap), "bytes")
	fmt.Println("Filled Map:", filledMap, "| Len:", len(filledMap), "| Size:", unsafe.Sizeof(filledMap), "bytes")

	// Map operations
	fmt.Println("Operations:")
	fmt.Println("  Index 'success':", filledMap["success"])
	filledMap["notfound"] = 404
	fmt.Println("  After Add:", filledMap)
	delete(filledMap, "failed")
	fmt.Println("  After Delete 'failed':", filledMap)

	// -- 4. Structs --
	// Custom data types that group related data. Cannot be nil (unless pointer).
	fmt.Println("\n\033[1;36m=== 4. STRUCTS ===\033[0m")
	type Address struct {
		Street string
		City   string
		Zip    int
	}

	type Person struct {
		Name    string
		Age     int
		Address Address // Nested struct
	}

	emptyPerson := Person{} // Empty struct (zero values)
	filledPerson := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
			Zip:    12345,
		},
	}
	fmt.Println("Empty Struct:", emptyPerson, "| Size:", unsafe.Sizeof(emptyPerson), "bytes")
	fmt.Println("Filled Struct:", filledPerson, "| Size:", unsafe.Sizeof(filledPerson), "bytes")

	// Struct operations
	fmt.Println("Operations:")
	fmt.Println("  Field Access Name:", filledPerson.Name)
	fmt.Println("  Nested Field Access:", filledPerson.Address.City)
	filledPerson.Age = 31
	fmt.Println("  After Update Age:", filledPerson.Age)
	filledPerson.Address.Zip = 54321
	fmt.Println("  After Update Zip:", filledPerson.Address.Zip)

	// -- 5. Pointers --
	// Reference to a memory address. Can be nil.
	fmt.Println("\n\033[1;36m=== 5. POINTERS ===\033[0m")
	var nilPointer *int // Nil pointer (points to nothing)
	num := 42
	filledPointer := &num // Pointer to integer
	fmt.Println("Nil Pointer:", nilPointer, "| Size:", unsafe.Sizeof(nilPointer), "bytes")
	fmt.Println("Filled Pointer:", filledPointer, "| Value:", *filledPointer, "| Size:", unsafe.Sizeof(filledPointer), "bytes")

	personPtr := &filledPerson // Pointer to struct
	fmt.Println("Struct Pointer:", personPtr, "| Size:", unsafe.Sizeof(personPtr), "bytes")

	// Pointer operations
	fmt.Println("Operations:")
	fmt.Println("  Dereference:", *filledPointer)
	*filledPointer = 100
	fmt.Println("  After Update *ptr=100:", *filledPointer, "| Original num:", num)
	fmt.Println("  Pointer to Field:", personPtr.Name)
	personPtr.Name = "Jane Doe"
	fmt.Println("  After Update via Pointer:", filledPerson.Name)
}
