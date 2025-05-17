# Standard Library Package

The stdlib package provides built-in functions and utilities for the toy language.

## Purpose

This package offers:
- Essential built-in functions
- Type manipulation utilities
- I/O operations
- String and collection helpers

## Built-in Functions

### Core Functions

#### `print(...args)`
Prints arguments to stdout with spaces between them.
```
print("Hello", "World")  // Output: Hello World
print(42, true)          // Output: 42 true
```

#### `println(...args)`
Same as print but adds a newline.
```
println("Hello")         // Output: Hello\n
```

#### `len(collection)`
Returns the length of a string, array, or map.
```
len("hello")             // 5
len([1, 2, 3])          // 3
len({a: 1, b: 2})       // 2
```

#### `type(value)`
Returns the type of a value as a string.
```
type(42)                 // "INTEGER"
type("hello")           // "STRING"
type([1, 2])            // "ARRAY"
```

### String Functions

#### `str(value)`
Converts value to string representation.
```
str(42)                  // "42"
str(true)                // "true"
```

#### `concat(...args)`
Concatenates arguments as strings.
```
concat("Hello", " ", "World")  // "Hello World"
```

#### `substr(str, start, length)`
Extracts substring.
```
substr("hello", 1, 3)    // "ell"
```

### Array Functions

#### `push(array, value)`
Adds element to end of array.
```
let arr = [1, 2]
push(arr, 3)             // arr = [1, 2, 3]
```

#### `pop(array)`
Removes and returns last element.
```
let arr = [1, 2, 3]
pop(arr)                 // returns 3, arr = [1, 2]
```

#### `slice(array, start, end)`
Returns array slice.
```
slice([1, 2, 3, 4], 1, 3) // [2, 3]
```

### Math Functions

#### `abs(number)`
Absolute value.
```
abs(-42)                 // 42
abs(42)                  // 42
```

#### `min(...args)`
Returns minimum value.
```
min(5, 2, 8, 1)         // 1
```

#### `max(...args)`
Returns maximum value.
```
max(5, 2, 8, 1)         // 8
```

## Implementation Structure

### Builtin Type
```go
type Builtin struct {
    Fn func(args ...interface{}) (interface{}, error)
}
```

### Registration
```go
var Builtins = map[string]*Builtin{
    "print": {
        Fn: printFn,
    },
    "len": {
        Fn: lenFn,
    },
    // ...
}
```

### Error Handling

Built-ins should validate:
- Argument count
- Argument types
- Operation validity

```go
func lenFn(args ...interface{}) (interface{}, error) {
    if len(args) != 1 {
        return nil, fmt.Errorf("wrong number of arguments")
    }
    // Type checking and implementation
}
```

## Adding New Built-ins

1. Define the function signature
2. Implement validation logic
3. Implement core functionality
4. Add to Builtins map
5. Write tests

### Example: Adding `reverse`

```go
func reverseFn(args ...interface{}) (interface{}, error) {
    if len(args) != 1 {
        return nil, fmt.Errorf("reverse expects 1 argument")
    }
    
    switch arg := args[0].(type) {
    case string:
        // Reverse string
        runes := []rune(arg)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        return string(runes), nil
    case []interface{}:
        // Reverse array
        result := make([]interface{}, len(arg))
        for i, j := 0, len(arg)-1; j >= 0; i, j = i+1, j-1 {
            result[i] = arg[j]
        }
        return result, nil
    default:
        return nil, fmt.Errorf("reverse not supported for %T", arg)
    }
}
```

## Testing Strategy

1. **Argument validation tests**
2. **Type checking tests**
3. **Functionality tests**
4. **Edge case tests**
5. **Performance benchmarks**

## Future Enhancements

1. **I/O Functions**
   - File operations
   - Network requests
   - User input

2. **Advanced Collections**
   - Map/filter/reduce
   - Sorting
   - Set operations

3. **Concurrency**
   - Goroutine equivalents
   - Channels
   - Mutexes

4. **Error Handling**
   - Try/catch equivalents
   - Error types
