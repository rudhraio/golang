# Go Data Structures: A Comprehensive Reference Guide

Welcome to your one-stop shop for understanding Go's data structures! Think of this as the "IKEA instruction manual" for storing your data – except actually readable and with less confusion about extra screws.

## Overview

Go provides several built-in data structures for organizing your data. Each has its own personality, memory footprint, and special abilities. Let's meet them!

---

## 1. Arrays

**What it is:** A fixed-size container that holds homogeneous data (all the same type). Once created, it's like concrete – immutable in size.

**When to use:** When you know exactly how many items you'll have and want blazing-fast access.

**Time Complexity:**
- Access: O(1)
- Modify: O(1)
- Iterate: O(n)
- Length: O(1)

**Space Complexity:** O(n) where n is the fixed size

**Memory:** Each element takes up space. For `[5]int`, that's 5 × 8 bytes = **40 bytes** (on 64-bit systems). Size is baked into the type itself.

**Operations:**
- **Access**: `array[0]` – Direct index access
- **Modify**: `array[0] = 42` – Update any element
- **Iterate**: Use `for i, v := range array` to loop through
- **Length**: `len(array)` – Always fixed, no surprises

**Example:**
```go
scores := [5]int{10, 20, 30, 40, 50}
scores[2] = 35  // Update third element
fmt.Println(scores[0])  // Access first element
```

**Pro tip:** Arrays are passed by value, so passing them to functions creates a copy. Want to modify? Use pointers!

---

## 2. Slices

**What it is:** Think of slices as "smart arrays" – same concept, but flexible. They can grow, shrink, and generally be more chill about life.

**When to use:** 99% of the time instead of arrays. They're like arrays' cooler cousin who actually shows up to parties.

**Time Complexity:**
- Access: O(1)
- Append: O(1) amortized (O(n) when reallocation needed)
- Slice: O(1)
- Iterate: O(n)
- Length/Capacity: O(1)

**Space Complexity:** O(n) where n is the number of elements

**Memory:** 24 bytes (fixed size on 64-bit systems) regardless of content – they store a pointer, length, and capacity. The actual data lives elsewhere.

**Operations:**
- **Create**: `slice := []int{1, 2, 3}` or `make([]int, 5)`
- **Append**: `slice = append(slice, 4)` – Grows automatically
- **Access**: `slice[0]` – Same as arrays
- **Slice**: `slice[1:3]` – Get a portion (non-destructive)
- **Length**: `len(slice)` – Current elements
- **Capacity**: `cap(slice)` – Room before reallocation

**Example:**
```go
numbers := []int{1, 2, 3}
numbers = append(numbers, 4, 5)  // Now [1, 2, 3, 4, 5]
subset := numbers[1:3]  // [2, 3] - no copy, just a view
```

**Pro tip:** Nil slices (`var s []int`) are different from empty slices (`s := []int{}`). Both have length 0, but nil can be passed through functions without initialization. It's like comparing "not ordered pizza" with "ordered zero pizzas."

---

## 3. Maps

**What it is:** A key-value pair data structure. Like a dictionary or hash table – you have keys, you get values. No judges here.

**When to use:** When you need fast lookups by a key. Need to find a person by their ID? Maps are your friend.

**Time Complexity:**
- Insert: O(1) average, O(n) worst case
- Lookup: O(1) average, O(n) worst case
- Delete: O(1) average, O(n) worst case
- Iterate: O(n)
- Length: O(1)

**Space Complexity:** O(n) where n is the number of key-value pairs

**Memory:** 8 bytes pointer (the actual data is dynamically allocated). Maps are references, so nil maps are common.

**Operations:**
- **Create**: `m := make(map[string]int)` – Mandatory `make()`
- **Insert**: `m["key"] = 42`
- **Lookup**: `value := m["key"]` – Returns zero value if missing
- **Check existence**: `value, exists := m["key"]` – Use comma-ok idiom
- **Delete**: `delete(m, "key")` – Removes the entry
- **Iterate**: `for k, v := range m` – Order is random (by design)
- **Length**: `len(m)` – Number of key-value pairs

**Example:**
```go
prices := make(map[string]int)
prices["apple"] = 5
prices["banana"] = 3

if price, ok := prices["apple"]; ok {
    fmt.Println("Found:", price)  // Found: 5
}

delete(prices, "banana")
```

**Pro tip:** Maps are **not** ordered. If you need ordering, you'll need to sort the keys separately. Also, map access returns zero values for missing keys – no errors thrown. Use the comma-ok pattern to distinguish "missing key" from "key with zero value."

---

## 4. Structs

**What it is:** Custom data types that group related fields together. It's like creating your own blueprint for organizing data.

**When to use:** When you need to represent complex objects with multiple properties. Users, configurations, database records – you name it.

**Time Complexity:**
- Access field: O(1)
- Modify field: O(1)
- Create: O(1)

**Space Complexity:** O(1) - fixed size determined by fields

**Memory:** Sum of all field sizes, plus potential padding for alignment. A struct with three strings and an int? That's `16 + 16 + 16 + 8 = 56 bytes` minimum (with potential padding).

**Operations:**
- **Create**: `person := Person{Name: "Alice", Age: 30}`
- **Access fields**: `person.Name` – Dot notation
- **Modify fields**: `person.Age = 31`
- **Nested access**: `person.Address.City` – Works through nesting
- **Zero value**: `Person{}` – All fields get their zero values

**Example:**
```go
type Person struct {
    Name string
    Age  int
    City string
}

p := Person{Name: "Bob", Age: 25, City: "NYC"}
p.Age = 26  // Modify
fmt.Println(p.Name)  // Access
```

**Pro tip:** Structs are value types by default – copying a struct creates a new copy. Want to modify the original? Use pointers!

---

## 5. Pointers

**What it is:** A memory address. It's like a treasure map – instead of carrying the treasure, you carry directions to it.

**When to use:** When you want to modify data passed to functions, avoid copying large structs, or share data across multiple locations.

**Time Complexity:**
- Dereference: O(1)
- Create pointer: O(1)
- Modify through pointer: O(1)

**Space Complexity:** O(1) - only stores an address

**Memory:** 8 bytes (on 64-bit systems) – just stores an address.

**Operations:**
- **Create**: `ptr := &value` – Get address with `&`
- **Dereference**: `*ptr` – Get the value at that address
- **Modify through pointer**: `*ptr = newValue`
- **Pointer to fields**: `ptr.Field` – Go auto-dereferences for convenience
- **Nil pointers**: `var ptr *int` – Points to nothing

**Example:**
```go
x := 42
ptr := &x      // Get address of x
*ptr = 100     // x is now 100
fmt.Println(x) // 100 – modified through pointer!

var nilPtr *int  // nil pointer, don't dereference!
```

**Pro tip:** Nil pointers will panic if dereferenced. Always check before assuming a pointer is valid (or handle the panic). Also, Go auto-dereferences struct pointers for field access – `ptr.Field` works even though technically it should be `(*ptr).Field`.

---

## 6. Interfaces

**What it is:** A contract that defines which methods a type must implement. It's polymorphism without the headache.

**When to use:** When you want to write flexible code that works with multiple types that share common operations.

**Time Complexity:**
- Type assertion: O(1)
- Method call: O(1)
- Interface assignment: O(1)

**Space Complexity:** O(1) - fixed size header

**Memory:** 16 bytes (pointer to value + pointer to type information).

**Operations:**
- **Define**: `type Reader interface { Read() error }`
- **Implement**: Any type with matching methods automatically implements it (implicit)
- **Type assertion**: `v, ok := iface.(ConcreteType)` – Check actual type
- **Method call**: `iface.MethodName()` – Call interface methods
- **Nil interface**: `var i Reader` – Nil until assigned

**Example:**
```go
type Writer interface {
    Write(string) error
}

type ConsoleWriter struct{}
func (cw ConsoleWriter) Write(s string) error {
    fmt.Println(s)
    return nil
}

var w Writer = ConsoleWriter{}
w.Write("Hello!")  // Works!
```

**Pro tip:** Interfaces are implicitly satisfied in Go. No need to say "implements Writer" – if you have the right methods, you're good. This is beautifully pragmatic and occasionally confusing.

---

## 7. Channels

**What it is:** Goroutine communication pipes. Send data from one goroutine to another safely and synchronously.

**When to use:** For concurrent programming – coordinating work between goroutines without explicit locks.

**Time Complexity:**
- Send: O(1) average
- Receive: O(1) average
- Close: O(1)
- Create: O(1)

**Space Complexity:** O(buffer_size) for buffered channels, O(1) for unbuffered

**Memory:** 8 bytes pointer (internal buffering varies based on buffer size).

**Operations:**
- **Create buffered**: `ch := make(chan int, 5)` – Can hold 5 integers before blocking
- **Create unbuffered**: `ch := make(chan int)` – Synchronous, sender waits for receiver
- **Send**: `ch <- value` – Blocks if full or unbuffered with no receiver
- **Receive**: `value := <-ch` – Blocks if empty
- **Close**: `close(ch)` – Signal that no more data will be sent
- **Check if open**: `value, ok := <-ch` – ok is false if closed
- **Nil channel**: `var ch chan int` – Never use! Sends/receives block forever

**Example:**
```go
ch := make(chan int, 2)
ch <- 10        // Send value
ch <- 20        // Send another
fmt.Println(<-ch)  // Receive: 10
fmt.Println(<-ch)  // Receive: 20
close(ch)       // Done sending
```

**Pro tip:** Don't send on a closed channel (panic!) and don't close a channel that receivers are still waiting on. Only the sender should close channels. Also, receiving from a nil channel blocks forever – not ideal.

---

## 8. Basic Types

**What it is:** Primitive data types provided by Go: integers, strings, booleans, floats.

**When to use:** Always. Every number, text, and boolean you use.

**Time Complexity:**
- Arithmetic operations: O(1)
- String concatenation: O(n) where n is total length
- Comparisons: O(1) for numbers, O(n) for strings
- Type conversion: O(1)

**Space Complexity:** O(1) for numbers/bool, O(n) for strings where n is length

**Memory:**
- `int`: 8 bytes (varies by platform)
- `string`: 16 bytes (pointer + length)
- `bool`: 1 byte
- `float64`: 8 bytes

**Operations:**
- **Arithmetic**: `a + b`, `a - b`, `a * b`, `a / b`
- **String concatenation**: `"Hello" + " World"`
- **Comparisons**: `a == b`, `a < b`, etc.
- **Type conversion**: `int(float32Value)`

**Example:**
```go
name := "Alice"
age := 30
score := 95.5
isActive := true
```

**Pro tip:** Strings in Go are immutable. Concatenating many strings repeatedly is slow – use `strings.Builder` for efficiency.

---

## Quick Comparison Table

| Structure | Size | Mutable | Use Case | Nil-able |
|-----------|------|---------|----------|----------|
| Array | Fixed | Yes (values) | Fixed collections | No |
| Slice | 24B | Yes | Dynamic collections | Yes |
| Map | 8B | Yes | Key-value lookups | Yes |
| Struct | Variable | Yes (fields) | Grouped data | No (unless pointer) |
| Pointer | 8B | Yes | References | Yes |
| Interface | 16B | Yes (methods) | Polymorphism | Yes |
| Channel | 8B | No | Goroutine comm | Yes |

---

## Memory Size Recap (64-bit systems)

- **Pointer**: 8 bytes (just an address)
- **Slice header**: 24 bytes (ptr + len + cap)
- **Map header**: 8 bytes (ptr to internal structure)
- **Interface**: 16 bytes (value ptr + type ptr)
- **Channel**: 8 bytes (ptr to internal structure)
- **String**: 16 bytes (ptr + length)
- **int/float64**: 8 bytes each
- **bool**: 1 byte
- **Array**: Size of elements × count (no header overhead)

---

## Common Gotchas

1. **Arrays are values** – Passing to functions creates a copy. Use slices instead.
2. **Nil vs Empty** – `nil` means no backing data; `{}` means initialized but empty.
3. **Map keys** – Must be comparable (no slices or maps as keys).
4. **String immutability** – Can't modify individual characters; convert to byte slice first.
5. **Pointer dereferencing** – Nil pointers panic! Check first.
6. **Channel send on closed channel** – Panic! Only sender should close.
7. **Interface methods** – Implicitly satisfied; no explicit "implements" needed.

---

## Pro Tips for Success

- Use **slices** instead of arrays (they're more flexible)  
- Use **structs** to organize complex data  
- Use **maps** for fast lookups by key  
- Use **pointers** when you need to modify data passed to functions  
- Use **interfaces** for flexible, decoupled code  
- Use **channels** for safe goroutine communication  
- Check for **nil** before dereferencing pointers  
- Use **comma-ok idiom** for map and channel operations  

---

## Running the Examples

```bash
go run datastructures.go
```

This will show you real examples of each data structure in action with their memory sizes and operations.

---

Happy coding! May your data structures be swift and your nil pointers few!
