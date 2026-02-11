# GoLang Practice

Welcome to my Go playground! 

This repo is where my Go code comes to play, trip, fall, and (sometimes) run successfully. Expect lots of experiments, a few bugs, and maybe a giggle or two. Enjoy the chaos!

---

## Beautiful Terminal Output: ANSI Escape Codes

### What Are ANSI Escape Codes?

ANSI (American National Standards Institute) escape codes are special sequences that tell your terminal to display text in colors, bold, italic, and other styles. In Go, they're strings that start with `\033[` and end with `m`.

### How They Work in Go

An ANSI escape code has this structure:
```
\033[<CODE>m
```

- `\033` is the escape character (tells terminal something special is coming)
- `[` starts the code sequence
- `<CODE>` is the action (color, style, etc.)
- `m` ends the sequence

### Common ANSI Codes for Text

**Text Styles:**
| Code | Effect | Example |
|------|--------|---------|
| `\033[1m` | Bold | Makes text thick |
| `\033[3m` | Italic | Slants text |
| `\033[4m` | Underline | Underlines text |
| `\033[0m` | Reset | Turns off all styling |

**Foreground Colors (Text Color):**
| Code | Color |
|------|-------|
| `\033[30m` | Black |
| `\033[31m` | Red |
| `\033[32m` | Green |
| `\033[33m` | Yellow |
| `\033[34m` | Blue |
| `\033[35m` | Magenta |
| `\033[36m` | Cyan |
| `\033[37m` | White |
| `\033[90m` | Gray (bright black) |

**Background Colors:**
| Code | Color |
|------|-------|
| `\033[40m` | Black background |
| `\033[41m` | Red background |
| `\033[42m` | Green background |
| `\033[43m` | Yellow background |
| `\033[44m` | Blue background |

### Practical Examples

**Simple Colored Text:**
```go
package main
import "fmt"

func main() {
    // Basic colors
    fmt.Printf("\033[31mRed text\033[0m\n")           // Red text
    fmt.Printf("\033[32mGreen text\033[0m\n")         // Green text
    fmt.Printf("\033[34mBlue text\033[0m\n")          // Blue text
}
```

**Bold and Colors Combined:**
```go
fmt.Printf("\033[1;31mBold Red\033[0m\n")             // Bold red text
fmt.Printf("\033[1;32mBold Green\033[0m\n")           // Bold green text
fmt.Printf("\033[1;36mBold Cyan\033[0m\n")            // Bold cyan text
```

**Multiple Codes Together:**
```go
fmt.Printf("\033[1;35;44mBold Magenta on Blue\033[0m\n")
// 1 = bold, 35 = magenta text, 44 = blue background
```

**Practical Use Case (Like in datatypes.go):**
```go
fmt.Printf("\033[1;33m▶ INTEGERS\033[0m\n")           // Bold yellow header
fmt.Printf("\033[90m%s\033[0m\n", "─────────────────") // Gray divider line
fmt.Printf("Normal text here\n")
```

### The Key to Beautiful Output

1. **Use colors for structure:** Headers, separators, categories
2. **Don't overdo it:** Too many colors = ugly and hard to read
3. **Always reset:** End colored sections with `\033[0m` so later text isn't affected
4. **Use semantic colors:** Green for success, Red for errors, Blue for headers, etc.

### How ANSI Codes Differ Across Languages

**Go:**
```go
fmt.Printf("\033[31mRed\033[0m\n")
```

**Python:**
```python
print("\033[31mRed\033[0m")
# Or use: print(f"\033[31mRed\033[0m")
```

**JavaScript/Node.js:**
```javascript
console.log("\033[31mRed\033[0m");
// Or use a library like 'chalk'
```

**Java:**
```java
System.out.println("\033[31mRed\033[0m");
```

**C/C++:**
```c
printf("\033[31mRed\033[0m\n");
```

The interesting part: ANSI codes work the SAME way in almost all languages! The difference is only in how you output them (printf, print, console.log, etc.). This makes them very portable.

### Pro Tips for Terminal Beauty

1. **Test on your terminal:** Some old terminals don't support colors. Test before shipping.

2. **Use reset carefully:**
   ```go
   fmt.Printf("\033[31mRed\033[0m Normal again\n")
   ```
   Always reset after colors to avoid affecting following text.

3. **Don't rely only on color:** Colorblind users might not see red/green. Use text or symbols too:
   ```go
   fmt.Printf("\033[31m✗ Error\033[0m\n")      // Red with symbol
   fmt.Printf("\033[32m✓ Success\033[0m\n")    // Green with symbol
   ```

4. **Use libraries for complex coloring:**
   - Python: `colorama`, `termcolor`
   - JavaScript: `chalk`, `colors`
   - Go: `fatih/color` package
   
   These libraries handle edge cases and make code more readable.

5. **See the datatypes.go for real examples:** Our reference code uses ANSI codes to create beautiful tables with colors and borders.

### Copy-Paste Ready: Common Patterns

**Header with box:**
```go
fmt.Printf("\033[1;36m╔════════════════════════════════════════╗\033[0m\n")
fmt.Printf("\033[1;36m║           MY BEAUTIFUL HEADER          ║\033[0m\n")
fmt.Printf("\033[1;36m╚════════════════════════════════════════╝\033[0m\n")
```

**Section title:**
```go
fmt.Printf("\033[1;33m▶ SECTION TITLE\033[0m\n")
```

**Divider line:**
```go
fmt.Printf("\033[90m─────────────────────────────────────────\033[0m\n")
```

**Error message:**
```go
fmt.Printf("\033[1;31m✗ Something went wrong\033[0m\n")
```

**Success message:**
```go
fmt.Printf("\033[1;32m✓ Everything is great\033[0m\n")
```