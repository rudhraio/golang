
package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	// Print system information
	fmt.Printf("\n\033[1;36m╔════════════════════════════════════════════════════════════════╗\033[0m\n")
	fmt.Printf("\033[1;36m║         GO DATA TYPES REFERENCE GUIDE (BEGINNER LEVEL)         ║\033[0m\n")
	fmt.Printf("\033[1;36m╚════════════════════════════════════════════════════════════════╝\033[0m\n")
	fmt.Printf("OS: %s | Architecture: %s\n\n", runtime.GOOS, runtime.GOARCH)

	// --- 1. Signed Integers ---
	// Signed integers can store both positive and negative numbers
	// Use when: counting things that can be negative, storing ages, scores, etc.
	var small_int8 int8 = 127
	var medium_int16 int16 = 32767
	var large_int32 int32 = 2147483647
	var huge_int64 int64 = 9223372036854775807

	fmt.Printf("\n\033[1;33m▶ SIGNED INTEGERS (can be positive or negative)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-30s | %-15s | %-10s\n", "Type", "Range", "Example Value", "Size(bytes)")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "int8", "-128 to 127", small_int8, unsafe.Sizeof(small_int8))
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "int16", "-32,768 to 32,767", medium_int16, unsafe.Sizeof(medium_int16))
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "int32", "-2.1B to 2.1B", large_int32, unsafe.Sizeof(large_int32))
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "int64", "-9.2E18 to 9.2E18", huge_int64, unsafe.Sizeof(huge_int64))
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "int (default)", "Platform dependent", 42, unsafe.Sizeof(42))

	// --- 1.1 Unsigned Integers ---
	// Unsigned integers can only store positive numbers (no negatives)
	// Use when: counting bytes, storing pixel values, indices, etc.
	var small_uint8 uint8 = 255
	var medium_uint16 uint16 = 65535
	var large_uint32 uint32 = 4294967295
	var huge_uint64 uint64 = 18446744073709551615

	fmt.Printf("\n\033[1;33m▶ UNSIGNED INTEGERS (only positive numbers)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-30s | %-15s | %-10s\n", "Type", "Range", "Example Value", "Size(bytes)")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "uint8 (byte)", "0 to 255", small_uint8, unsafe.Sizeof(small_uint8))
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "uint16", "0 to 65,535", medium_uint16, unsafe.Sizeof(medium_uint16))
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "uint32 (rune)", "0 to 4.2B", large_uint32, unsafe.Sizeof(large_uint32))
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "uint64", "0 to 1.8E19", huge_uint64, unsafe.Sizeof(huge_uint64))
	fmt.Printf("%-15s | %-30s | %-15d | %-10d\n", "uint (default)", "Platform dependent", 42, unsafe.Sizeof(uint(42)))

	// --- 2. Floating-point numbers ---
	// Floating-point numbers store decimal values
	// Use when: storing prices, temperatures, scientific calculations, etc.
	var small_float32 float32 = 3.4028235e+38
	var medium_float64 float64 = 1.7976931348623157e+308

	fmt.Printf("\n\033[1;32m▶ FLOATING-POINT NUMBERS (decimal values)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-30s | %-15s | %-10s\n", "Type", "Range (approx)", "Example Value", "Size(bytes)")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-30s | %-15.2f | %-10d\n", "float32", "±1.4e-45 to ±3.4e+38", small_float32, unsafe.Sizeof(small_float32))
	fmt.Printf("%-15s | %-30s | %-15.2e | %-10d\n", "float64", "±5.0e-324 to ±1.7e+308", medium_float64, unsafe.Sizeof(medium_float64))

	// --- 3. Constants ---
	// Constants are immutable values that cannot be changed after declaration
	// Use when: storing fixed values like Pi, configuration settings, etc.
	const Pi float64 = 3.141592653589793
	const E float32 = 2.71828

	fmt.Printf("\n\033[1;35m▶ CONSTANTS (immutable, cannot be changed)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-15s | %-15s | %-10s\n", "Type", "Value", "Precision", "Size(bytes)")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-15.8f | %-15s | %-10d\n", "Pi (float64)", Pi, "14 decimals", unsafe.Sizeof(Pi))
	fmt.Printf("%-15s | %-15.5f | %-15s | %-10d\n", "E (float32)", E, "5 decimals", unsafe.Sizeof(E))

	// --- 4. Booleans ---
	// Boolean values are either true or false
	// Use when: conditional logic, boolean flags, status checks, etc.
	var isTrue bool = true
	isFalse := false

	fmt.Printf("\n\033[1;36m▶ BOOLEANS (true or false)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-15s | %-15s | %-10s\n", "Type", "Value", "Use Case", "Size(bytes)")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-15v | %-15s | %-10d\n", "bool", isTrue, "Control flow", unsafe.Sizeof(isTrue))
	fmt.Printf("%-15s | %-15v | %-15s | %-10d\n", "bool", isFalse, "Flags", unsafe.Sizeof(isFalse))

	// --- 5. Strings ---
	// Strings are sequences of characters (immutable)
	// Use when: storing text, names, messages, file paths, etc.
	var greeting string = "Hello, Go!"
	name := "Gopher"

	fmt.Printf("\n\033[1;31m▶ STRINGS (text data - immutable)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-20s | %-20s | %-10s\n", "Variable", "Value", "Use Case", "Size(bytes)")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-20s | %-20s | %-10d\n", "greeting", greeting, "Messages", unsafe.Sizeof(greeting))
	fmt.Printf("%-15s | %-20s | %-20s | %-10d\n", "name", name, "Names/Identifiers", unsafe.Sizeof(name))

	// --- 6. Zero Values ---
	// Variables declared without initialization get a "Zero Value"
	// Zero value: int=0, float=0.0, bool=false, string=""
	// Use when: need default values without explicit initialization
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string

	fmt.Printf("\n\033[1;33m▶ ZERO VALUES (default values without initialization)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-20s | %-15s | %-10s\n", "Type", "Default Value", "Value", "Size(bytes)")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-20s | %-15d | %-10d\n", "int", "0", zeroInt, unsafe.Sizeof(zeroInt))
	fmt.Printf("%-15s | %-20s | %-15.1f | %-10d\n", "float64", "0.0", zeroFloat, unsafe.Sizeof(zeroFloat))
	fmt.Printf("%-15s | %-20s | %-15v | %-10d\n", "bool", "false", zeroBool, unsafe.Sizeof(zeroBool))
	fmt.Printf("%-15s | %-20s | %-15s | %-10d\n", "string", "\"\" (empty)", "\""+zeroString+"\"", unsafe.Sizeof(zeroString))

	// --- 7. nil and any ---
	// 'any' is an alias for interface{} (available in Go 1.18+)
	// nil represents "no value" for pointers, slices, maps, channels, etc.
	// Use when: need flexible types or optional values
	var ptr *int = nil
	fmt.Printf("\n\033[1;35m▶ NIL AND ANY (special types)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-20s | %-15s | %-20s | %-10s\n", "Type", "Value", "Meaning", "Size(bytes)")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-20s | %-15v | %-20s | %-10d\n", "*int (pointer)", ptr, "No value", unsafe.Sizeof(ptr))

	var anyValue any = "I can be anything!"
	fmt.Printf("%-20s | %-15v | %-20s | %-10d\n", "any (string)", anyValue, "Flexible type", unsafe.Sizeof(anyValue))
	anyValue = 42 // Now I'm an int
	fmt.Printf("%-20s | %-15v | %-20s | %-10d\n", "any (int)", anyValue, "Type changed", unsafe.Sizeof(anyValue))

	// --- 8. Type Aliases ---
	// Common aliases used in Go
	fmt.Printf("\n\033[1;32m▶ TYPE ALIASES (shortcuts for common types)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-30s | %-15s\n", "Alias", "Actual Type", "Common Use")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-30s | %-15s\n", "byte", "uint8", "Raw bytes")
	fmt.Printf("%-15s | %-30s | %-15s\n", "rune", "int32", "Unicode characters")
	var byteVal byte = 255
	var runeVal rune = 'A'
	fmt.Printf("Example: byte(255) = %d, rune('A') = %d\n", byteVal, runeVal)

	// --- 9. Quick Reference: Declaration Styles ---
	fmt.Printf("\n\033[1;36m▶ DECLARATION STYLES (ways to declare variables)\033[0m\n")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-35s | %-30s\n", "Declaration Method", "Example")
	fmt.Printf("\033[90m%s\033[0m\n", "─────────────────────────────────────────────────────────────────────────────")
	fmt.Printf("%-35s | %-30s\n", "var with type", "var x int = 10")
	fmt.Printf("%-35s | %-30s\n", "var with type inference", "var x = 10  // inferred as int")
	fmt.Printf("%-35s | %-30s\n", "short declaration (:=)", "x := 10  // Only inside functions")
	fmt.Printf("%-35s | %-30s\n", "multiple variables", "x, y := 1, 2  // Both same type")
	fmt.Printf("%-35s | %-30s\n", "constant", "const Pi = 3.14  // Immutable")

	fmt.Printf("\n\033[1;36m╔════════════════════════════════════════════════════════════════╗\033[0m\n")
	fmt.Printf("\033[1;36m║                    Reference Guide Complete                    ║\033[0m\n")
	fmt.Printf("\033[1;36m╚════════════════════════════════════════════════════════════════╝\033[0m\n\n")
}