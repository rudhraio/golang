package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	separator := strings.Repeat("=", 60)
	fmt.Println(separator)
	fmt.Println("         GO OPERATIONS REFERENCE GUIDE")
	fmt.Println(separator)

	// ============================================================
	// SECTION 1: ARITHMETIC OPERATIONS (Integers & Floats)
	// ============================================================
	fmt.Println("\n\033[1;34m1. ARITHMETIC OPERATIONS (Integers & Floats)\033[0m")
	fmt.Println(strings.Repeat("-", 60))

	a, b := 20, 8
	fmt.Printf("Values: a = %d, b = %d\n\n", a, b)

	fmt.Println("Basic Arithmetic Operations:")
	fmt.Printf("  Addition (a + b):      %d + %d = %d\n", a, b, a+b)
	fmt.Printf("  Subtraction (a - b):   %d - %d = %d\n", a, b, a-b)
	fmt.Printf("  Multiplication (a * b): %d * %d = %d\n", a, b, a*b)
	fmt.Printf("  Division (a / b):      %d / %d = %d\n", a, b, a/b)
	fmt.Printf("  Modulus (a %% b):      %d %% %d = %d (remainder)\n", a, b, a%b)

	fmt.Println("\nIncrement & Decrement Operators:")
	x := 10
	fmt.Printf("  x = %d\n", x)
	x++
	fmt.Printf("  After x++: %d\n", x)
	x--
	fmt.Printf("  After x--: %d\n", x)

	fmt.Println("\nCompound Assignment Operators:")
	y := 20
	fmt.Printf("  Initial: y = %d\n", y)
	y += 5
	fmt.Printf("  y += 5:  %d (same as y = y + 5)\n", y)
	y -= 3
	fmt.Printf("  y -= 3:  %d (same as y = y - 3)\n", y)
	y *= 2
	fmt.Printf("  y *= 2:  %d (same as y = y * 2)\n", y)
	y /= 4
	fmt.Printf("  y /= 4:  %d (same as y = y / 4)\n", y)
	y %= 5
	fmt.Printf("  y %%= 5:  %d (same as y = y %% 5)\n", y)

	fmt.Println("\nFloating Point Operations:")
	f1, f2 := 15.5, 3.2
	fmt.Printf("  Values: f1 = %.1f, f2 = %.1f\n", f1, f2)
	fmt.Printf("  f1 + f2 = %.2f\n", f1+f2)
	fmt.Printf("  f1 - f2 = %.2f\n", f1-f2)
	fmt.Printf("  f1 * f2 = %.2f\n", f1*f2)
	fmt.Printf("  f1 / f2 = %.2f\n", f1/f2)

	// ============================================================
	// SECTION 2: RELATIONAL (COMPARISON) OPERATIONS
	// ============================================================
	fmt.Println("\n\033[1;34m2. RELATIONAL (COMPARISON) OPERATIONS\033[0m")
	fmt.Println("-" + strings.Repeat("-", 59))

	p, q := 15, 10
	fmt.Printf("Values: p = %d, q = %d\n\n", p, q)

	fmt.Println("Comparison Results:")
	fmt.Printf("  p == q (equal):           %v\n", p == q)
	fmt.Printf("  p != q (not equal):       %v\n", p != q)
	fmt.Printf("  p > q (greater than):     %v\n", p > q)
	fmt.Printf("  p < q (less than):        %v\n", p < q)
	fmt.Printf("  p >= q (greater or equal): %v\n", p >= q)
	fmt.Printf("  p <= q (less or equal):   %v\n", p <= q)

	fmt.Println("\nString Comparisons:")
	str1, str2 := "apple", "banana"
	fmt.Printf("  str1 = \"%s\", str2 = \"%s\"\n", str1, str2)
	fmt.Printf("  str1 == str2: %v\n", str1 == str2)
	fmt.Printf("  str1 != str2: %v\n", str1 != str2)

	// ============================================================
	// SECTION 3: LOGICAL OPERATIONS
	// ============================================================
	fmt.Println("\n\033[1;34m3. LOGICAL OPERATIONS\033[0m")
	fmt.Println("-" + strings.Repeat("-", 59))

	isStudent := true
	hasClasses := false
	isWorking := true

	fmt.Printf("Values: isStudent = %v, hasClasses = %v, isWorking = %v\n\n", isStudent, hasClasses, isWorking)

	fmt.Println("Logical NOT (!):")
	fmt.Printf("  !isStudent = %v (negation)\n", !isStudent)
	fmt.Printf("  !hasClasses = %v (negation)\n", !hasClasses)

	fmt.Println("\nLogical AND (&&) - Both must be true:")
	fmt.Printf("  isStudent && hasClasses = %v (true && false)\n", isStudent && hasClasses)
	fmt.Printf("  isStudent && isWorking = %v (true && true)\n", isStudent && isWorking)

	fmt.Println("\nLogical OR (||) - At least one must be true:")
	fmt.Printf("  isStudent || hasClasses = %v (true || false)\n", isStudent || hasClasses)
	fmt.Printf("  hasClasses || isWorking = %v (false || true)\n", hasClasses || isWorking)

	fmt.Println("\nComplex Logical Operations:")
	canGraduate := isStudent && (hasClasses || isWorking)
	fmt.Printf("  canGraduate = isStudent && (hasClasses || isWorking) = %v\n", canGraduate)

	// ============================================================
	// SECTION 4: BITWISE OPERATIONS
	// ============================================================
	fmt.Println("\n\033[1;34m4. BITWISE OPERATIONS (Binary Level)\033[0m")
	fmt.Println("-" + strings.Repeat("-", 59))

	var bitX uint8 = 12 // Binary: 1100
	var bitY uint8 = 10 // Binary: 1010

	fmt.Printf("Values: bitX = %d (binary: 1100), bitY = %d (binary: 1010)\n\n", bitX, bitY)

	fmt.Println("Bitwise Operations:")
	fmt.Printf("  Bitwise AND (&):    %d & %d = %d (binary: 1000)\n", bitX, bitY, bitX&bitY)
	fmt.Printf("  Bitwise OR (|):     %d | %d = %d (binary: 1110)\n", bitX, bitY, bitX|bitY)
	fmt.Printf("  Bitwise XOR (^):    %d ^ %d = %d (binary: 0110)\n", bitX, bitY, bitX^bitY)
	fmt.Printf("  Bitwise NOT (~):    ^%d = %d (flip all bits)\n", bitX, ^bitX)

	fmt.Println("\nBit Shift Operations:")
	fmt.Printf("  Left Shift (<<):    %d << 2 = %d (multiply by 2^2)\n", 5, 5<<2)
	fmt.Printf("  Right Shift (>>):   %d >> 1 = %d (divide by 2^1)\n", 8, 8>>1)

	// ============================================================
	// SECTION 5: STRING OPERATIONS
	// ============================================================
	fmt.Println("\n\033[1;34m5. STRING OPERATIONS\033[0m")
	fmt.Println(strings.Repeat("-", 60))

	fmt.Println("String Concatenation:")
	greeting := "Hello"
	name := "Go"
	fullGreeting := greeting + " " + name
	fmt.Printf("  \"%s\" + \" \" + \"%s\" = \"%s\"\n", greeting, name, fullGreeting)

	fmt.Println("\nString Length:")
	text := "Programming"
	fmt.Printf("  len(\"%s\") = %d\n", text, len(text))

	fmt.Println("\nString Indexing & Slicing:")
	str := "GOLANG"
	fmt.Printf("  String: \"%s\"\n", str)
	fmt.Printf("  str[0] = %c (first character)\n", str[0])
	fmt.Printf("  str[5] = %c (last character)\n", str[5])
	fmt.Printf("  str[0:2] = \"%s\" (first 2 characters)\n", str[0:2])
	fmt.Printf("  str[2:5] = \"%s\" (from index 2 to 5)\n", str[2:5])
	fmt.Printf("  str[2:] = \"%s\" (from index 2 to end)\n", str[2:])

	fmt.Println("\nString Functions (from strings package):")
	text2 := "go programming"
	fmt.Printf("  strings.ToUpper(\"%s\") = \"%s\"\n", text2, strings.ToUpper(text2))
	fmt.Printf("  strings.ToLower(\"GO LANG\") = \"%s\"\n", strings.ToLower("GO LANG"))
	fmt.Printf("  strings.Contains(\"%s\", \"prog\") = %v\n", text2, strings.Contains(text2, "prog"))
	fmt.Printf("  strings.Index(\"%s\", \"prog\") = %d\n", text2, strings.Index(text2, "prog"))

	// ============================================================
	// SECTION 6: ARRAY OPERATIONS
	// ============================================================
	fmt.Println("\n\033[1;34m6. ARRAY OPERATIONS (Fixed Size)\033[0m")
	fmt.Println(strings.Repeat("-", 60))

	var arr [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Array: %v\n\n", arr)

	fmt.Println("Array Operations:")
	fmt.Printf("  Length: len(arr) = %d\n", len(arr))
	fmt.Printf("  First element: arr[0] = %d\n", arr[0])
	fmt.Printf("  Last element: arr[4] = %d\n", arr[4])

	fmt.Println("\nModifying Array Elements:")
	arr[2] = 99
	fmt.Printf("  After arr[2] = 99: %v\n", arr)

	fmt.Println("\nIterating Over Array:")
	fmt.Print("  Elements: ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// ============================================================
	// SECTION 7: SLICE OPERATIONS
	// ============================================================
	fmt.Println("\n\033[1;34m7. SLICE OPERATIONS (Dynamic Size)\033[0m")
	fmt.Println("-" + strings.Repeat("-", 59))

	nums := []int{1, 2, 3}
	fmt.Printf("Initial Slice: %v\n\n", nums)

	fmt.Println("Append Operation:")
	fmt.Printf("  nums = %v\n", nums)
	nums = append(nums, 4)
	fmt.Printf("  After append(nums, 4): %v\n", nums)
	nums = append(nums, 5, 6)
	fmt.Printf("  After append(nums, 5, 6): %v\n", nums)

	fmt.Println("\nSlice Manipulation:")
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("  Slice: %v\n", numbers)
	fmt.Printf("  numbers[1:4] = %v (from index 1 to 4)\n", numbers[1:4])
	fmt.Printf("  numbers[:3] = %v (first 3 elements)\n", numbers[:3])
	fmt.Printf("  numbers[2:] = %v (from index 2 to end)\n", numbers[2:])

	fmt.Println("\nSlice Copy:")
	original := []int{1, 2, 3, 4, 5}
	copied := make([]int, len(original))
	copy(copied, original)
	fmt.Printf("  Original: %v\n", original)
	fmt.Printf("  Copied:   %v\n", copied)
	copied[0] = 999
	fmt.Printf("  After modified copy: Original: %v, Copied: %v\n", original, copied)

	fmt.Println("\nSlice Length & Capacity:")
	slice := make([]int, 3, 5)
	fmt.Printf("  Slice: %v, Length: %d, Capacity: %d\n", slice, len(slice), cap(slice))

	// ============================================================
	// SECTION 8: MAP OPERATIONS
	// ============================================================
	fmt.Println("\n\033[1;34m8. MAP OPERATIONS (Key-Value Pairs)\033[0m")
	fmt.Println("-" + strings.Repeat("-", 59))

	fruits := map[string]int{
		"Apple":  5,
		"Banana": 3,
		"Orange": 7,
	}
	fmt.Printf("Map: %v\n\n", fruits)

	fmt.Println("Accessing Values:")
	fmt.Printf("  fruits[\"Apple\"] = %d\n", fruits["Apple"])
	fmt.Printf("  fruits[\"Banana\"] = %d\n", fruits["Banana"])

	fmt.Println("\nAdding New Key-Value Pairs:")
	fruits["Mango"] = 4
	fmt.Printf("  After fruits[\"Mango\"] = 4: %v\n", fruits)

	fmt.Println("\nUpdating Values:")
	fruits["Apple"] = 10
	fmt.Printf("  After fruits[\"Apple\"] = 10: %v\n", fruits)

	fmt.Println("\nChecking Key Existence:")
	value, exists := fruits["Banana"]
	fmt.Printf("  value, exists := fruits[\"Banana\"]\n")
	fmt.Printf("  Value: %d, Exists: %v\n", value, exists)

	notExist, exists2 := fruits["Grape"]
	fmt.Printf("  value, exists := fruits[\"Grape\"]\n")
	fmt.Printf("  Value: %d, Exists: %v\n", notExist, exists2)

	fmt.Println("\nDeleting Keys:")
	fmt.Printf("  Before delete: %v\n", fruits)
	delete(fruits, "Orange")
	fmt.Printf("  After delete(fruits, \"Orange\"): %v\n", fruits)

	fmt.Println("\nIterating Over Map:")
	fmt.Print("  Items: ")
	for key, value := range fruits {
		fmt.Printf("[%s: %d] ", key, value)
	}
	fmt.Println()

	// ============================================================
	// SECTION 9: ADVANCED MATH OPERATIONS
	// ============================================================
	fmt.Println("\n\033[1;34m9. ADVANCED MATH OPERATIONS\033[0m")
	fmt.Println("-" + strings.Repeat("-", 59))

	fmt.Println("Power & Root Operations:")
	fmt.Printf("  math.Pow(2, 3) = %.0f (2^3)\n", math.Pow(2, 3))
	fmt.Printf("  math.Pow(5, 2) = %.0f (5^2)\n", math.Pow(5, 2))
	fmt.Printf("  math.Sqrt(16) = %.0f (square root of 16)\n", math.Sqrt(16))
	fmt.Printf("  math.Sqrt(25) = %.0f (square root of 25)\n", math.Sqrt(25))

	fmt.Println("\nRounding & Absolute Value:")
	num := -12.7
	fmt.Printf("  math.Abs(%.1f) = %.1f\n", num, math.Abs(num))
	fmt.Printf("  math.Floor(12.7) = %.0f\n", math.Floor(12.7))
	fmt.Printf("  math.Ceil(12.3) = %.0f\n", math.Ceil(12.3))
	fmt.Printf("  math.Round(12.5) = %.0f\n", math.Round(12.5))

	fmt.Println("\nTrigonometric Functions:")
	angle := math.Pi / 4 // 45 degrees in radians
	fmt.Printf("  math.Sin(π/4) = %.2f\n", math.Sin(angle))
	fmt.Printf("  math.Cos(π/4) = %.2f\n", math.Cos(angle))
	fmt.Printf("  math.Tan(π/4) = %.2f\n", math.Tan(angle))

	fmt.Println("\nLogarithmic Functions:")
	fmt.Printf("  math.Log(2.718) = %.2f (natural log of e)\n", math.Log(2.718))
	fmt.Printf("  math.Log10(100) = %.1f (log base 10 of 100)\n", math.Log10(100))

	// ============================================================
	// SECTION 10: TYPE CONVERSION
	// ============================================================
	fmt.Println("\n\033[1;34m10. TYPE CONVERSION\033[0m")
	fmt.Println("-" + strings.Repeat("-", 59))

	fmt.Println("Integer to Float:")
	intVal := 42
	floatVal := float64(intVal)
	fmt.Printf("  %d (int) → %f (float64)\n", intVal, floatVal)

	fmt.Println("\nFloat to Integer:")
	floatVal2 := 45.8
	intVal2 := int(floatVal2)
	fmt.Printf("  %.1f (float64) → %d (int) - Note: decimal part is truncated\n", floatVal2, intVal2)

	fmt.Println("\nInteger to String:")
	fmt.Printf("  Using fmt.Sprintf: %d → \"%s\"\n", 123, fmt.Sprintf("%d", 123))

	fmt.Println("\nFloat to String:")
	fmt.Printf("  Using fmt.Sprintf: %.2f → \"%s\"\n", 45.67, fmt.Sprintf("%.2f", 45.67))

	// ============================================================
	// SECTION 11: PRACTICAL EXAMPLES
	// ============================================================
	fmt.Println("\n\033[1;34m11. PRACTICAL EXAMPLES\033[0m")
	fmt.Println("-" + strings.Repeat("-", 59))

	fmt.Println("Example 1: Calculate Average of Numbers")
	scores := []int{85, 90, 78, 92, 88}
	sum := 0
	for _, score := range scores {
		sum += score
	}
	average := float64(sum) / float64(len(scores))
	fmt.Printf("  Scores: %v\n  Sum: %d, Average: %.2f\n", scores, sum, average)

	fmt.Println("\nExample 2: Count Character Frequencies")
	word := "programming"
	charCount := make(map[rune]int)
	for _, char := range word {
		charCount[char]++
	}
	fmt.Printf("  Word: \"%s\"\n  Character frequencies: %v\n", word, charCount)

	fmt.Println("\nExample 3: Filter Even Numbers")
	allNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNums []int
	for _, num := range allNums {
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		}
	}
	fmt.Printf("  All numbers: %v\n  Even numbers: %v\n", allNums, evenNums)

	fmt.Println("\n" + separator)
	fmt.Println("              END OF REFERENCE GUIDE")
	fmt.Println(separator + "\n")
}
