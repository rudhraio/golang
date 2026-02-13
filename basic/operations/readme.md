# GO Operations Reference Guide - Your New Best Friend

Welcome to the ultimate cheat sheet for Go operations! Think of this as your personal Operations Wikipedia. We'll cover everything from the basics (that even your grandma could understand) to the fancy stuff that'll make you look smart at tech conferences.

---

## 📚 Table of Contents

1. [Arithmetic Operations](#1-arithmetic-operations)
2. [Relational (Comparison) Operations](#2-relational-operations)
3. [Logical Operations](#3-logical-operations)
4. [Bitwise Operations](#4-bitwise-operations)
5. [String Operations](#5-string-operations)
6. [Array Operations](#6-array-operations)
7. [Slice Operations](#7-slice-operations)
8. [Map Operations](#8-map-operations)
9. [Advanced Math Operations](#9-advanced-math-operations)
10. [Type Conversion](#10-type-conversion)
11. [Practical Examples](#11-practical-examples)

---

## 1. Arithmetic Operations

### What's This?
These are the basic math operations you learned in school (remember third grade math class?). Addition, subtraction, multiplication, division, and modulus (fancy word for "remainder"). We also throw in increment/decrement operators because typing `x = x + 1` is so 2010.

### Operations:
- `+` Addition
- `-` Subtraction
- `*` Multiplication
- `/` Division (integer division in Go drops decimals)
- `%` Modulus (the leftover after division)
- `++` Increment (but you can't use it in expressions like `y = x++`)
- `--` Decrement (same limitations apply)
- `+=, -=, *=, /=, %=` Compound assignments (type less, do more)

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| Basic arithmetic (+, -, *, /) | **O(1)** | **O(1)** |
| Modulus (%) | **O(1)** | **O(1)** |
| Increment/Decrement | **O(1)** | **O(1)** |

**Translation:** All of these are lightning-fast because the CPU handles them natively. They're atomic operations that complete in pretty much zero time. No loops, no allocations. Pure speed!

---

## 2. Relational (Comparison) Operations

### What's This?
These operators let you compare two values and get a `true` or `false` answer. It's like asking your code "Is this bigger than that?" or "Do these two things match?" Simple questions, simple answers.

### Operations:
- `==` Equal to
- `!=` Not equal to
- `>` Greater than
- `<` Less than
- `>=` Greater than or equal to
- `<=` Less than or equal to

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| All comparisons (==, !=, >, <, >=, <=) | **O(1)** | **O(1)** |
| String comparison | **O(n)** where n = string length | **O(1)** |

**Translation:** Number comparisons? Instant! The CPU does them in a nanosecond. String comparisons? They have to check character-by-character, so they take longer based on string length. But it's still blazing fast in practice.

---

## 3. Logical Operations

### What's This?
These are the operators that help your code make decisions based on multiple conditions. Want to check if it's raining AND you have an umbrella? Or if your boss is out OR it's Friday? That's what logical ops are for!

### Operations:
- `!` NOT (negation - flips true to false and vice versa)
- `&&` AND (both conditions must be true)
- `||` OR (at least one condition must be true)

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| NOT (!) | **O(1)** | **O(1)** |
| AND (&&) | **O(1)** | **O(1)** |
| OR (\|\|) | **O(1)** | **O(1)** |

**Bonus: Short-circuit Evaluation**
- `&&` stops as soon as it finds a false (why continue if the first is already false?)
- `\|\|` stops as soon as it finds a true (why keep searching if you already found true?)

**Translation:** These are super efficient! They're essentially just comparisons under the hood. Plus, Go is smart enough to stop evaluating as soon as it knows the answer. Lazy evaluation at its finest!

---

## 4. Bitwise Operations

### What's This?
Welcome to the "low-level coolness" section. These operations work at the binary level (1s and 0s). Remember when computers were just 1s and 0s? Yeah, well, they still are. This is what happens when you talk directly to the CPU in its native language. Useful for flags, permissions, and impressing people at parties (not really, but you can pretend).

### Operations:
- `&` AND (both bits must be 1)
- `|` OR (at least one bit must be 1)
- `^` XOR (bits must be different)
- `~` NOT (flip all bits - goes negative for signed ints)
- `<<` Left Shift (multiply by powers of 2)
- `>>` Right Shift (divide by powers of 2)

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| Bitwise AND, OR, XOR, NOT | **O(1)** | **O(1)** |
| Left Shift, Right Shift | **O(1)** | **O(1)** |

**Translation:** These are literally some of the fastest operations your CPU can do. They're pure hardware instructions. Want to multiply by 16? Use `x << 4` (shift left 4 times). Want to divide by 8? Use `x >> 3`. Modern compilers often optimize multiplication/division to bit shifts anyway!

**Pro Tip:** Bitwise NOT (`~`) on positive integers gives you the negative number minus 1. It's weird but mathematically consistent!

---

## 5. String Operations

### What's This?
Strings are just sequences of characters. In Go, strings are immutable (once created, they can't change). So whenever you "modify" a string, you're actually creating a new one. It's like the butterfly effect but with text!

### Operations:
- Concatenation (`+`) - joining strings
- Length (`len()`) - how many characters
- Indexing (`str[i]`) - get character at position i
- Slicing (`str[i:j]`) - get substring from i to j
- Functions from `strings` package (ToUpper, ToLower, Contains, Index, etc.)

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| String concatenation | **O(n + m)** where n, m = string lengths | **O(n + m)** |
| `len()` | **O(1)** | **O(1)** |
| Indexing | **O(1)** | **O(1)** |
| Slicing | **O(k)** where k = slice length | **O(k)** |
| `strings.Contains()` | **O(n)** where n = string length | **O(1)** |
| `strings.Index()` | **O(n)** | **O(1)** |

**Translation:** 
- Concatenation creates a whole new string, so it takes time proportional to the total length. Don't do this in a loop 1000 times! Use `strings.Builder` instead.
- Length is instant because Go stores it internally.
- Indexing is instant (just pointer arithmetic).
- Slicing creates a new string, proportional to slice size.
- Searching operations scan through the string, so they depend on length.

**Fun Fact:** In Go, strings are UTF-8 by default. That's why indexing with `str[i]` gives you a byte, not a rune (Unicode character). If you want characters, use `range` loops!

---

## 6. Array Operations

### What's This?
Arrays are like your grocery list written in permanent marker. Once you decide the size, you're stuck with it. Fixed size = fixed overhead. They're super efficient and predictable, which makes them perfect for situations where you know exactly how many items you need.

### Operations:
- Create: `[size]Type{values}`
- Length: `len(arr)` - always returns the same value
- Indexing: `arr[i]` - access element at position i
- Modifying: `arr[i] = newValue`
- Iteration: `for i := 0; i < len(arr); i++` or `for index, value := range arr`

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| Array creation | **O(n)** where n = array size | **O(n)** |
| Length | **O(1)** | **O(1)** |
| Indexing (access) | **O(1)** | **O(1)** |
| Modification | **O(1)** | **O(1)** |
| Iteration | **O(n)** | **O(1)** (if not creating new array) |

**Translation:**
- Array creation allocates all memory at once. Safe and predictable.
- Everything is instant once created because arrays are contiguous memory blocks.
- Iteration is just a loop through n elements.

**Gotcha:** Arrays are passed by **copy** in Go. Pass a large array to a function? The whole thing gets copied. Use pointers (`*[size]Type`) to avoid this overhead!

---

## 7. Slice Operations

### What's This?
Slices are arrays' cooler, more flexible cousin. They're dynamic (grow/shrink), and they only point to underlying data instead of copying it. Think of them as a window into an array. You can open the window wider or move it around!

### Operations:
- Create: `[]Type{values}` or `make([]Type, length, capacity)`
- Length: `len(slice)` - current number of elements
- Capacity: `cap(slice)` - how many elements it can hold before needing to reallocate
- Indexing: `slice[i]`
- Slicing: `slice[i:j]` - creates a new slice pointing to the same underlying array
- Append: `slice = append(slice, value)` - adds elements (might cause reallocation)
- Copy: `copy(dest, src)` - copies elements from src to dest

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| Slice creation | **O(n)** | **O(n)** |
| Length | **O(1)** | **O(1)** |
| Capacity | **O(1)** | **O(1)** |
| Indexing | **O(1)** | **O(1)** |
| Slicing | **O(1)** | **O(1)** (no data copied, new header created) |
| Append (amortized) | **O(1)** average | **O(n)** worst case (when reallocation happens) |
| Copy | **O(n)** | **O(n)** (if creating new slice) |

**Translation:**
- Slicing is clever - it doesn't copy data, just creates a new slice header. Instant!
- Append is usually O(1) because of amortization. Many appends share one costly reallocation.
- When append needs more space, Go typically doubles the capacity. This spreads the cost over many operations.

**Important Note:** When you slice, you still reference the original underlying array. Modify element through the slice? The original array changes too! They share the same memory. 🔗

---

## 8. Map Operations

### What's This?
Maps are like dictionary or telephone books. You look up a key and get the value. No ordering, just fast lookups. Perfect for scenarios where you need to find things by name rather than by position.

### Operations:
- Create: `map[KeyType]ValueType{key: value}`
- Access: `map[key]` - get value (returns zero value if key doesn't exist)
- Check existence: `value, exists := map[key]` - exists is boolean
- Add/Update: `map[key] = value`
- Delete: `delete(map, key)`
- Length: `len(map)` - number of key-value pairs
- Iteration: `for key, value := range map`

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| Creation | **O(1)** for empty | **O(n)** for n elements |
| Access | **O(1)** average | **O(1)** |
| Insertion | **O(1)** average | **O(1)** |
| Deletion | **O(1)** average | **O(1)** |
| Lookup (exists check) | **O(1)** average | **O(1)** |
| Iteration | **O(n)** | **O(1)** (not counting iteration itself) |

**Translation:**
- Maps use hash tables under the hood, so lookups are typically constant time.
- If there are hash collisions? Could degrade to O(n) in the worst case, but Go's hash function is really good.
- Iteration visits each element once, so it's linear.

**Wild Behavior:** Maps have random iteration order! Running a loop over a map twice might give different orders. Go does this intentionally to prevent developers from accidentally relying on order. 🎲

**Critical:** Map ordering is randomized in Go (since version 1.0), so if you need consistent ordering, store keys in a separate slice or use a different data structure!

---

## 9. Advanced Math Operations

### What's This?
For when basic arithmetic isn't enough. Power functions, square roots, trigonometry, logarithms. This is where you bust out the heavy math machinery from the `math` package. It's like bringing out the big guns for a serious calculation problem.

### Common Functions:
- `math.Pow(x, y)` - x to the power of y
- `math.Sqrt(x)` - square root
- `math.Abs(x)` - absolute value
- `math.Floor(x)` - round down
- `math.Ceil(x)` - round up
- `math.Round(x)` - round to nearest integer
- `math.Sin(x), math.Cos(x), math.Tan(x)` - trigonometry (uses radians!)
- `math.Log(x)` - natural logarithm
- `math.Log10(x)` - log base 10

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| `math.Pow()` | **O(1)** | **O(1)** |
| `math.Sqrt()` | **O(1)** | **O(1)** |
| `math.Abs()` | **O(1)** | **O(1)** |
| Rounding functions | **O(1)** | **O(1)** |
| Trigonometric functions | **O(1)** | **O(1)** |
| Logarithmic functions | **O(1)** | **O(1)** |

**Translation:** All these are hardware-accelerated floating-point operations. They're fast, but not as fast as integer arithmetic. Floating-point math is a bit slower than integer math, but we're still talking nanoseconds! CPU manufacturers have optimized these to death. 💻

**Pro Tip:** Trigonometric functions expect **radians**, not degrees! Remember: radians = degrees × (π/180). It trips everyone up at first. 📐

---

## 10. Type Conversion

### What's This?
Go is statically typed, which means types are checked at compile time. But sometimes you need to convert from one type to another. This is where explicit type conversion comes in. Go doesn't do implicit conversions (unlike some other languages that are just too helpful for their own good).

### Conversions:
- `int() to float()` - `float64(intValue)`
- `float() to int()` - `int(floatValue)` (loses decimal part)
- `int to string` - `fmt.Sprintf("%d", intValue)`
- `float to string` - `fmt.Sprintf("%.2f", floatValue)`
- Any comparable types - `TargetType(value)`

### Time & Space Complexity:
| Operation | Time | Space |
|-----------|------|-------|
| int ↔ float conversion | **O(1)** | **O(1)** |
| Number to string conversion | **O(log n)** where n = number value | **O(1)** |
| String parsing to number | **O(n)** where n = string length | **O(1)** |

**Translation:**
- Direct type conversions between numeric types are instant bit manipulations.
- Converting to strings requires building a new string, which is logarithmic in number size.
- Parsing strings to numbers requires scanning the string character by character.

**Warning:** Converting float to int truncates (cuts off) the decimal part. It doesn't round! `int(45.9)` gives you `45`, not `46`. If you want rounding, use `math.Round()` first!

---

## 11. Practical Examples

### What's This?
Real-world use cases showing how to combine operations to solve actual problems. Because knowing operations is great, but using them to solve real problems? That's where the magic happens! ✨

### Example 1: Calculate Average
```go
scores := []int{85, 90, 78, 92, 88}
sum := 0
for _, score := range scores {
    sum += score
}
average := float64(sum) / float64(len(scores))
```
**Complexity:** O(n) where n = number of scores (linear time, perfect!)

### Example 2: Count Character Frequencies
```go
word := "programming"
charCount := make(map[rune]int)
for _, char := range word {
    charCount[char]++
}
```
**Complexity:** O(n) where n = string length (we visit each character once)

### Example 3: Filter Even Numbers
```go
allNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var evenNums []int
for _, num := range allNums {
    if num%2 == 0 {
        evenNums = append(evenNums, num)
    }
}
```
**Complexity:** O(n) average time, O(n) space for result

---

## Complexity Cheat Sheet

### Time Complexity Summary:
| Operation Type | Time Complexity | Notes |
|---|---|---|
| Arithmetic | O(1) | All instant |
| Comparisons | O(1) for numbers, O(n) for strings | |
| String operations | O(n) | All involve string length |
| Array indexing | O(1) | Direct memory access |
| Array iteration | O(n) | Visit each element |
| Slice append (amortized) | O(1) | Usually instant |
| Map operations | O(1) avg | Hash table-based |
| Math functions | O(1) | Hardware-accelerated |

### Space Complexity Summary:
| Operation Type | Space | Notes |
|---|---|---|
| Simple operations | O(1) | No extra memory |
| String concatenation | O(n + m) | New string created |
| Slice creation | O(n) | New slice header + data |
| Map creation | O(n) | Hash table overhead |
| Iteration | O(1) | Index/pointer only |

---

## Pro Tips for Maximum Efficiency

1. **Avoid string concatenation in loops** - Use `strings.Builder` instead
2. **Use maps for fast lookups** - O(1) instead of O(n) searching
3. **Slices are better than arrays** - More flexible with similar performance
4. **Bit shifting is faster than multiplication** - `x << 1` vs `x * 2`
5. **Order matters with logical operators** - Put cheap checks first
6. **Short-circuit evaluation saves CPU** - Use it in complex conditions
7. **Pre-allocate slices when you know the size** - `make([]int, 0, expectedSize)`
8. **Range loops create copies** - Use pointers in maps to avoid allocation

---

## Final Thoughts

Now you've got the full operations toolkit! You understand:
- **When** to use each operation
- **Why** they work the way they do
- **How fast** they are (complexity-wise)

Go forth and write efficient code! Your future self (and your CPU) will thank you. And remember: premature optimization is the root of all evil, but understanding your tools helps you make smart decisions from the start! 

Happy coding!

---

**Last Updated:** February 13, 2026  
**Go Version:** 1.21+  
**Tone:** Slightly ridiculous but educational
