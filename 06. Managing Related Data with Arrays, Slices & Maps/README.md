# A. Arrays

## 1. Introducing Arrays

An **array** is a fixed-size collection of elements of the same type. The size of an array is part of its type.

```go
var arr [5]int // An array of five integers, initialized to zero
```

### Declaring and Initializing Arrays

```go
// Declaration with zero values
var numbers [3]int

// Declaration with initial values
primes := [5]int{2, 3, 5, 7, 11}

// Inferring size
letters := [...]string{"a", "b", "c"}

// Specifying index
mixed := [5]int{0: 10, 2: 30, 4: 50}
```

### Accessing and Modifying Elements

```go
fmt.Println(primes[0]) // Outputs: 2
primes[1] = 4
fmt.Println(primes)    // Outputs: [2 4 5 7 11]
```
---

## 2. Slicing Arrays

A **slice** is a dynamically-sized, flexible **view** into the elements of an array. Unlike arrays, slices are reference types.

### Creating Slices from Arrays

```go
arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4] // Elements 20, 30, 40
fmt.Println(slice) // Outputs: [20 30 40]
```

### Slicing Functionality

- **Reference Type**: Slices reference an underlying array.
- **No Copy**: Slicing does not copy the array; it creates a view. So they are memory efficient.
- **Memory Efficient:** Multiple slices can share the same underlying array without extra memory overhead.

### Example (Modifying Slice Reflects in Array)

```go
original := [5]int{1, 2, 3, 4, 5}
s1 := original[1:3] // [2, 3]
s2 := original[2:5] // [3, 4, 5]

s1[1] = 20
fmt.Println(original) // [1 2 20 4 5]
fmt.Println(s2)      // [20, 4, 5]
```

---

## 3. `len()` and `cap()` Functions on Arrays and Slices

### Arrays

- `len(array)` returns the number of elements.
- `cap(array)` also returns the number of elements since the size is fixed.

```go
arr := [3]int{1, 2, 3}
fmt.Println(len(arr)) // 3
fmt.Println(cap(arr)) // 3
```

### Slices

- `len(slice)` returns the number of elements in the slice.
- `cap(slice)` returns the number of elements from the start of the slice to the end of the underlying array.

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:3]
fmt.Println(len(slice)) // 2
fmt.Println(cap(slice)) // 4
```

---

## 4. Creating Dynamic Arrays with Slices

Slices are dynamic and can grow or shrink. They are built on top of arrays but provide flexibility. When we declare an array without specified length, actually we create a slice.

### Creating a Slice

```go
var s []int // nil slice

s = []int{1, 2, 3} // Non-nil slice
```

### Resizing with `append`

```go
s = append(s, 4, 5)
fmt.Println(s) // [1 2 3 4 5]
```

### Removing an Item from a Slice

To remove an element at index `i`, create a new slice excluding that element.

#### Example

```go
s := []int{10, 20, 30, 40, 50}
i := 2 // Remove 30

s = append(s[:i], s[i+1:]...)
fmt.Println(s) // [10 20 40 50]
```

### Appending a Slice to an Existing Slice

Use `append` with the `...` operator to append one slice to another.

#### Example

```go
s1 := []int{1, 2, 3}
s2 := []int{4, 5, 6}

s1 = append(s1, s2...) // Append all elements of s2 to s1
fmt.Println(s1) // [1 2 3 4 5 6]
```
---
# B. Maps

## 1. Introduction to Maps

### What Are Maps?

A **map** in Go is a built-in data type that associates keys with values. Unlike arrays or slices, which use integer indices, maps allow you to use any comparable type as keys, providing a flexible way to store and access data.

### Why Use Maps Instead of Arrays?

- **Flexible Keys**: Maps allow non-integer keys (e.g., strings, structs).
- **Efficient Lookups**: Average constant time complexity for key lookups.
- **Dynamic Size**: Maps can grow or shrink as needed, unlike fixed-size arrays.

### Example

**Use Case:** Managing company website URLs.

#### Using an Array

Imagine you have a list of company website URLs. Using an array requires you to access them via integer indices, which can be less intuitive and harder to manage, especially as the list grows or changes.

```go
package main

import "fmt"

func main() {
    // Using an array to store company URLs
    var companyURLs [3]string
    companyURLs[0] = "https://www.google.com"
    companyURLs[1] = "https://www.microsoft.com"
    companyURLs[2] = "https://www.apple.com"

    // Accessing a URL by index
    fmt.Println("Company at index 1:", companyURLs[1]) // Output: Company at index 1: https://www.microsoft.com
}
```

**Limitations:**
- **Unclear Association:** It's not immediately clear which URL corresponds to which company without maintaining parallel arrays or additional context.
- **Rigid Structure:** Adding or removing companies requires managing array indices, which can lead to errors.

#### Using a Map

Using a map allows you to associate each company name directly with its URL, making the data more intuitive and easier to manage.

```go
package main

import "fmt"

func main() {
    // Using a map to store company names and their URLs
    companyURLs := map[string]string{
        "Google":    "https://www.google.com",
        "Microsoft": "https://www.microsoft.com",
        "Apple":     "https://www.apple.com",
    }

    // Accessing a URL by company name
    fmt.Println("Microsoft URL:", companyURLs["Microsoft"]) // Output: Microsoft URL: https://www.microsoft.com

    // Adding a new company
    companyURLs["Amazon"] = "https://www.amazon.com"
    fmt.Println("Amazon URL added:", companyURLs["Amazon"]) // Output: Amazon URL added: https://www.amazon.com
}
```

**Advantages:**
- **Clear Associations:** Each company name is directly linked to its URL.
- **Ease of Access:** Retrieve URLs using meaningful keys (company names) instead of numerical indices.
- **Dynamic Management:** Easily add, update, or remove companies without worrying about array indices.

---

## 2. Mutating Maps

### Creating a Map

```go
// Using make
ages := make(map[string]int)

// Using a map literal
scores := map[string]int{
    "Alice": 90,
    "Bob":   85,
}
```

### Adding and Updating Entries

```go
ages["Charlie"] = 35       // Add
ages["Alice"] = 26         // Update
fmt.Println(ages)          // Output: map[Alice:26 Charlie:35]
```

### Deleting Entries

```go
delete(ages, "Bob")        // Removes "Bob" from the map
fmt.Println(ages)          // Output: map[Alice:26 Charlie:35]
```

### Accessing Values

```go
age, exists := ages["Alice"]
if exists {
    fmt.Println("Alice is", age, "years old")
} else {
    fmt.Println("Alice not found")
}
// Output: Alice is 26 years old
```

---

## 3. Maps vs. Structs

Both **maps** and **structs** are used to group data, but they serve different purposes.

### Maps

- **Dynamic Keys**: Keys can be of any comparable type.
- **Flexible Structure**: Suitable for scenarios where the data structure isn't fixed.
- **Use Case**: Collections with custom or dynamic labels.

**Example:**

```go
userData := map[string]interface{}{
    "Name":    "Alice",
    "Age":     25,
    "Country": "USA",
}
fmt.Println(userData["Country"]) // Output: USA
```

### Structs

- **Fixed Fields**: Fields are predefined and typed.
- **Predefined Structure**: Ensures a consistent data shape.
- **Use Case**: When the data structure is well-defined and unlikely to change.

**Example:**

```go
type User struct {
    Name    string
    Age     int
    Country string
}

user := User{
    Name:    "Alice",
    Age:     25,
    Country: "USA",
}
fmt.Println(user.Country) // Output: USA
```

---
# C. Array & Map Features

## 1. The `make()` Function in Go

The `make()` function in Go is used to initialize slices, maps, and channels. It allocates and initializes the internal data structures, returning the initialized type.

### Using `make()` with Maps

```go
package main

import "fmt"

func main() {
    // Initialize a map with make
    myMap := make(map[string]int)

    // Add key-value pairs
    myMap["apple"] = 5
    myMap["banana"] = 3

    fmt.Println(myMap) // Output: map[apple:5 banana:3]
}
```

**With Capacity:**

```go
myMap := make(map[string]int, 10) // Preallocate space for 10 key-value pairs but it extendable
```

### Using `make()` with Slices

```go
package main

import "fmt"

func main() {
    // Initialize a slice with make
    mySlice := make([]int, 5) // Length 5, default value 0

    fmt.Println(mySlice) // Output: [0 0 0 0 0]
}
```

**With Capacity:**

```go
mySlice := make([]int, 5, 10) // Length 5, Capacity 10
```

**Note:** While `make()` is not used for arrays (which have fixed sizes), slices are more commonly used and dynamic.

### Benefits of Preallocation

- **Performance Improvement:** Fewer memory reallocations and copies.
- **Predictable Memory Usage:** Helps manage memory more efficiently.
- **Avoids Runtime Overhead:** Reduces the need for dynamic resizing.

---
## 2. `for` Loops Over Arrays and Maps

### Iterating Over Arrays and Slices

Use a traditional `for` loop or the `range` keyword.

#### Using Traditional `for` Loop:

```go
for i := 0; i < len(myArray); i++ {
    fmt.Printf("Index: %d, Value: %s\n", i, myArray[i])
}
```

#### Using `range`:

```go
package main

import "fmt"

func main() {
    myArray := [3]string{"Go", "Python", "Java"}

    for index, value := range myArray {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }
}
```

#### Output:
```
Index: 0, Value: Go
Index: 1, Value: Python
Index: 2, Value: Java
```

---

### Iterating Over Maps

When iterating over maps, the order is not guaranteed.

```go
package main

import "fmt"

func main() {
    myMap := map[string]int{
        "apple":  5,
        "banana": 3,
        "cherry": 7,
    }

    for key, value := range myMap {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}
```

#### Output:
```
Key: apple, Value: 5
Key: banana, Value: 3
Key: cherry, Value: 7
```

#### Iterating Only Keys or Values:

```go
// Only keys
for key := range myMap {
    fmt.Println("Key:", key)
}

// Only values
for _, value := range myMap {
    fmt.Println("Value:", value)
}
```