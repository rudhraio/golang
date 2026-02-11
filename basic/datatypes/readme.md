# Go Data Types Reference Guide

## Quick Glance: The Cheat Sheet

### Integers (Whole Numbers)
| Category | Types | Vibe | Use When |
|----------|-------|------|----------|
| **Signed** | `int8, int16, int32, int64` | Can be negative (angry) | Storing ages, scores, temperatures |
| **Unsigned** | `uint8, uint16, uint32, uint64` | Always happy (positive only) | Counting things, pixel values, indices |
| **Default** | `int, uint` | Depends on your computer | You're lazy and don't want to think hard |

**Remember:** The bigger the number in the type name, the bigger the number it can hold. `int64` is basically the gym bro of integers.

### Floating-Point Numbers (Decimals)
| Type | Range | Perfect For | Problem |
|------|-------|-------------|---------|
| `float32` | ±1.4e-45 to ±3.4e+38 | Money, temperatures, physics | Might not be precise enough |
| `float64` | ±5.0e-324 to ±1.7e+308 | Scientific stuff, more precision | Uses twice as much memory |

**Warning:** Don't use floats for money in real apps! Use `decimal` libraries instead. Seriously. Trust me on this one.

### Strings (Text Stuff)
```go
var name string = "Gopher"
greeting := "Hello, Go!"
```
- **Immutable:** Can't change once created (like a tattoo, but less permanent)
- **Use for:** Names, messages, file paths, angry programmer comments
- **Size:** Contains memory address to the actual string data (it's complicated)

### Booleans (True/False)
```go
var isAwesome bool = true
var isConfused := false
```
- **Only two options:** `true` or `false` (binary baby!)
- **For:** If statements, loops, checking if your code works (spoiler: it doesn't)
- **Size:** 1 byte (surprisingly spacious for so little information)

### Constants (The Stubborn Ones)
```go
const Pi = 3.14159265
const MAX_ITEMS = 100
```
- **Immutable:** You declare it, that's IT
- **No changing:** Try to change it and Go will slap your wrist
- **When to use:** Pi, CONFIG values, answer to life the universe and everything (42)

### Nil & Any (The Wild Cards)
- **nil:** Means "nothing" (use with pointers, slices, maps, channels)
- **any:** Can be ANYTHING (it's an alias for `interface{}`)

```go
var ptr *int = nil  // No value here, moves on
var mystery any = "Could be anything!"  // The chameleon of types
mystery = 42  // Nope, it's a number now!
```

---

## Zero Values: The Default Laziness

When you declare a variable but don't give it a value, Go assigns it a "zero value". It's Go's way of saying "I'll just put something safe here":

```go
var x int        // Gets 0
var y float64    // Gets 0.0
var z bool       // Gets false
var s string     // Gets "" (empty string)
```

This won't crash your program. Go's got your back.

---

## Type Aliases: Shortcuts for the Lazy

Go gives you some shortcuts:

| Alias | Actually Is | Why It Exists |
|-------|------------|---------------|
| `byte` | `uint8` | Dealing with raw bytes is common, so... shortcut! |
| `rune` | `int32` | Unicode character values (fancy letters from the world) |

```go
var b byte = 'A'      // Represents the letter A as a number
var r rune = '你'     // Can handle fancy Unicode chars!
```

---

## Declaration Styles: Pick Your Flavor

### Inside Functions (Most Common)
```go
// Short declaration - most concise
x := 10

// Long form with type
var x int = 10

// Let Go figure out the type
var x = 10  // Inferred as int
```

### Outside Functions (Package Level)
```go
// Only these two work outside functions
var globalVar = 5
const GlobalConst = 42

// This WON'T work: x := 10  (Short syntax not allowed here)
```

### Multiple Variables (Showing Off)
```go
x, y, z := 1, 2, 3
```

---

## Printing Data Types: Format Specifiers

When you want to print variables, Go uses format specifiers in `fmt` package functions. Here's a quick reference:

### Common Format Specifiers

| Specifier | Use For | Example | Output |
|-----------|---------|---------|--------|
| `%v` | Any value (default format) | `fmt.Printf("%v", 42)` | `42` |
| `%T` | Type of value | `fmt.Printf("%T", 42)` | `int` |
| `%d` | Integer (decimal) | `fmt.Printf("%d", 42)` | `42` |
| `%b` | Integer (binary) | `fmt.Printf("%b", 5)` | `101` |
| `%x` or `%X` | Integer (hex - lowercase/uppercase) | `fmt.Printf("%x", 255)` | `ff` |
| `%o` | Integer (octal) | `fmt.Printf("%o", 8)` | `10` |
| `%f` | Float (decimal notation) | `fmt.Printf("%f", 3.14)` | `3.140000` |
| `%.2f` | Float with precision | `fmt.Printf("%.2f", 3.14159)` | `3.14` |
| `%e` or `%E` | Float (scientific notation) | `fmt.Printf("%e", 3.14)` | `3.140000e+00` |
| `%g` | Float (compact form) | `fmt.Printf("%g", 3.14)` | `3.14` |
| `%s` | String | `fmt.Printf("%s", "Hello")` | `Hello` |
| `%q` | String (quoted) | `fmt.Printf("%q", "Hi")` | `"Hi"` |
| `%c` | Byte/Rune as character | `fmt.Printf("%c", 65)` | `A` |

### Practical Examples

```go
package main
import "fmt"

func main() {
    // Integers
    age := 25
    fmt.Printf("I am %d years old\n", age)        // Output: I am 25 years old
    fmt.Printf("Binary: %b, Hex: %x\n", age, age)  // Output: Binary: 11001, Hex: 19
    
    // Floats
    price := 19.99
    fmt.Printf("Price: %.2f\n", price)             // Output: Price: 19.99
    fmt.Printf("Scientific: %e\n", 1000000.0)      // Output: Scientific: 1.000000e+06
    
    // Strings
    name := "Gopher"
    fmt.Printf("Hello, %s!\n", name)               // Output: Hello, Gopher!
    fmt.Printf("Quoted: %q\n", name)               // Output: Quoted: "Gopher"
    
    // Booleans
    isActive := true
    fmt.Printf("Active: %v\n", isActive)           // Output: Active: true
    fmt.Printf("Type: %T\n", isActive)             // Output: Type: bool
}
```

### Width and Alignment

You can also control spacing:

```go
fmt.Printf("|%5d|\n", 42)    // Right align: |   42|
fmt.Printf("|%-5d|\n", 42)   // Left align:  |42   |
fmt.Printf("|%05d|\n", 42)   // Zero padded:  |00042|
```

---

## The Challenge: Using the Code

Run the `datatypes.go` file to see ALL of this in action:

```bash
go run datatypes.go
```

You'll see:
- All data types with examples
- Size of each type in bytes (yes, Go cares about memory)
- Ranges and limits
- Pretty colored output to make your terminal jealous 

---

## The Golden Rules

1. **Choose the right size:** Don't use `int64` for counting apples (overkill)
2. **Strings are immutable:** Stop trying to change "Hello" into "Hallo"
3. **Constants are sacred:** Don't even THINK about modifying them
4. **nil ≠ empty string:** `nil` is "nothing", `""` is "something empty"
5. **any is NOT magic:** It's just a flexible container that requires type assertion to use
6. **Floats can be weird:** 0.1 + 0.2 might not equal 0.3 (computer math is weird)

---

## When in Doubt, Ask Yourself

When deciding what data type to use, ask yourself these questions: Will my number be negative? Then use a **signed** integer type like `int`. Will it only be positive? Then use an **unsigned** integer type like `uint`. Do you need decimals for calculations? Use **float** and prefer `float64` for better precision. Is it text? Use **string**. Is it a yes/no value? Use **bool**. Is it a fixed value that shouldn't change? Use **const**. Can it be nothing (no value)? Then use **pointers** or **any**. These questions will help you pick the right type almost every time.

---

## Pro Tips from a Gopher

- **`int` is your friend:** Unless you have a good reason, just use `int`
- **`float64` > `float32`:** More precision, only 2x the memory
- **Constants make you confident:** Write them for magic numbers
- **Comments are free:** Explain your types (your future self will thank you)
- **Use `unsafe.Sizeof()`:** To check actual memory usage (spoiler: it's what the code uses!)

---

## Still Confused?

That's normal! Go data types are beginner-friendly but take time to master. Here's what to do:

1. **Run the code:** See examples in action
2. **Edit the code:** Break it on purpose, see what happens
3. **Read error messages:** They're actually helpful (rare in programming!)
4. **Come back here:** Refresh your memory
5. **Experiment:** Try different type combinations

---

## Final Thoughts

Data types might seem boring, but they're the foundation of everything you'll write in Go. Master these, and you're already 20% of the way to being a Go ninja!

Now go forth and type!

---

**P.S.:** If you find a bug in the code, it wasn't me. It was the AI. 
