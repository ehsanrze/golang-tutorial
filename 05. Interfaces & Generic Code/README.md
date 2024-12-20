
## **1. Polymorphism**

### **Concept Overview**
Polymorphism is a core concept in programming that allows objects of different types to be treated as objects of a common super type. The term originates from Greek, meaning "many forms."

### **Types of Polymorphism**
1. **Subtype Polymorphism (Inclusion Polymorphism):**
   - Allows a subtype to be treated as its supertype.
   - **Example:** A `Dog` and a `Cat` can both be treated as `Animal`.

2. **Parametric Polymorphism:**
   - Enables functions or data types to be written generically, so they can handle values uniformly without depending on their type.
   - **Example:** Although Go doesn't have traditional generics (prior to Go 1.18), similar behavior can be achieved using the empty interface (interface{}).

3. **Ad-hoc Polymorphism:**
   - Allows functions to operate differently based on their input types, typically through function overloading or operator overloading.
   - **Note:** Go does not support traditional function overloading, but similar flexibility can be achieved using interfaces.

### **Polymorphism in Go**
- **Implementation:**  
  In Go, polymorphism is primarily achieved through **interfaces**, enabling different types to be treated uniformly based on shared behaviors rather than their concrete implementations.
- **Benefit:**  
  Facilitates writing flexible and reusable code by allowing functions to operate on any type that satisfies a particular interface.

---

## **2. Interfaces in Go**

### **What is an Interface?**
- **Definition:**  
  An interface in Go is a type that defines a set of method signatures. It specifies **what** methods a type must implement, without dictating **how** those methods are executed.
  
- **Implicit Implementation:**  
  Unlike some other languages, Go does not require explicit declarations that a type implements an interface. If a type defines all the methods declared in an interface, it implicitly satisfies that interface.

### **Why Use Interfaces?**

1. **Abstraction:** Hides the implementation details and exposes only the necessary behaviors.

2. **Flexibility and Extensibility:** Allows different types to be used interchangeably as long as they satisfy the same interface.

3. **Decoupling Components:** Reduces dependencies between different parts of a system by relying on interfaces rather than concrete types.

4. **Facilitating Testing:** Enables the creation of mock implementations that satisfy interfaces for testing purposes.

5. **Promoting Reusability:** Encourages writing generic functions and libraries that operate on any type satisfying a particular interface.

### **Key Characteristics of Go Interfaces**
- **Method Signatures Only:**  
  Interfaces contain only method signatures without any implementation.
  
- **Dynamic Dispatch:**  
  When a method is called on an interface, Go determines at runtime which concrete type’s method to execute based on the underlying type that satisfies the interface.

- **Composition Over Inheritance:**  
  Go favors interface composition (combining multiple interfaces) rather than traditional inheritance, promoting more flexible and modular designs.

---

### Example of Interface in Go

Imagine you want to model different entities that can produce sounds, such as dogs, cats, and birds. You can achieve this using an interface.

#### **Interface Definition**

```go
// Speaker is an interface that requires a Speak method
type Speaker interface {
    Speak() string
}
```

#### **Implementing the Interface with Different Types**

1. **Dog Type**

```go
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}
```

2. **Cat Type**

```go
type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}
```

3. **Bird Type**

```go
type Bird struct {
    Name string
}

func (b Bird) Speak() string {
    return "Tweet!"
}
```

#### **Using the Interface to Achieve Polymorphism**

With the `Speaker` interface defined, you can create a collection of different `Speaker` types and interact with them uniformly.

```go
package main

import (
	"fmt"
)

// Speaker is an interface that requires a Speak method
type Speaker interface {
	Speak() string
}

// Dog struct
type Dog struct {
	Name string
}

// Speak method for Dog
func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

// Cat struct
type Cat struct {
	Name string
}

// Speak method for Cat
func (c Cat) Speak() string {
	return c.Name + " says: Meow!"
}

// Bird struct
type Bird struct {
	Name string
}

// Speak method for Bird
func (b Bird) Speak() string {
	return b.Name + " says: Tweet!"
}

// PrintVoice is a function that takes a Speaker and prints its voice
func PrintVoice(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	// Define separate Speaker objects
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	bird := Bird{Name: "Tweety"}

	// Use the PrintVoice function to print each Speaker's voice
	PrintVoice(dog)   // Output: Buddy says: Woof!
	PrintVoice(cat)   // Output: Whiskers says: Meow!
	PrintVoice(bird)  // Output: Tweety says: Tweet!
}
```

#### **Expected Output**

```
Woof!
Meow!
Tweet!
```
---

### Interface Embedding in Go

**Interface Embedding** in Go allows you to compose new interfaces by including one or more existing interfaces within them. This promotes reusability and modularity by building complex interfaces from simpler, focused ones.


#### **Example**

Building upon our previous example where we have a `Speaker` interface, let's introduce a new behavior: movement. We'll create a `Mover` interface and then embed both `Speaker` and `Mover` into a new `Animal` interface. This allows us to define more comprehensive behaviors for our entities.

##### **Interface Definitions**

```go
// Speaker is an interface that requires a Speak method
type Speaker interface {
    Speak() string
}

// Mover is an interface that requires a Move method
type Mover interface {
    Move() string
}

// Animal is an interface that embeds Speaker and Mover
type Animal interface {
    Speaker
    Mover
}
```

##### **Complete Code Example Including Interface Embedding**

Here's the complete Go program incorporating interface embedding:

```go
package main

import (
    "fmt"
)

// Speaker is an interface that requires a Speak method
type Speaker interface {
    Speak() string
}

// Mover is an interface that requires a Move method
type Mover interface {
    Move() string
}

// Animal is an interface that embeds Speaker and Mover
type Animal interface {
    Speaker
    Mover
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return d.Name + " says: Woof!"
}

func (d Dog) Move() string {
    return d.Name + " runs swiftly."
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return c.Name + " says: Meow!"
}

func (c Cat) Move() string {
    return c.Name + " walks gracefully."
}

type Bird struct {
    Name string
}

func (b Bird) Speak() string {
    return b.Name + " says: Tweet!"
}

func (b Bird) Move() string {
    return b.Name + " flies elegantly."
}

// PrintVoice is a function that takes a Speaker and prints its voice
func PrintVoice(s Speaker) {
    fmt.Println(s.Speak())
}

// PrintMovement is a function that takes a Mover and prints its movement
func PrintMovement(m Mover) {
    fmt.Println(m.Move())
}

// PrintAnimalInfo is a function that takes an Animal and prints its voice and movement
func PrintAnimalInfo(a Animal) {
    fmt.Println(a.Speak())
    fmt.Println(a.Move())
    fmt.Println() // For better readability
}

func main() {
    // Define separate Animal objects
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    bird := Bird{Name: "Tweety"}

    // Use the PrintVoice function to print each Speaker's voice
    PrintVoice(dog)   // Output: Buddy says: Woof!
    PrintVoice(cat)   // Output: Whiskers says: Meow!
    PrintVoice(bird)  // Output: Tweety says: Tweet!

    fmt.Println() // Separator

    // Use the PrintMovement function to print each Mover's movement
    PrintMovement(dog)   // Output: Buddy runs swiftly.
    PrintMovement(cat)   // Output: Whiskers walks gracefully.
    PrintMovement(bird)  // Output: Tweety flies elegantly.

    fmt.Println() // Separator

    // Use the PrintAnimalInfo function to print both voice and movement
    PrintAnimalInfo(dog)
    // Output:
    // Buddy says: Woof!
    // Buddy runs swiftly.

    PrintAnimalInfo(cat)
    // Output:
    // Whiskers says: Meow!
    // Whiskers walks gracefully.

    PrintAnimalInfo(bird)
    // Output:
    // Tweety says: Tweet!
    // Tweety flies elegantly.
}
```

#### **Output**

```
Buddy says: Woof!
Whiskers says: Meow!
Tweety says: Tweet!

Buddy runs swiftly.
Whiskers walks gracefully.
Tweety flies elegantly.

Buddy says: Woof!
Buddy runs swiftly.

Whiskers says: Meow!
Whiskers walks gracefully.

Tweety says: Tweet!
Tweety flies elegantly.
```
---
## 3. Introduction to `interface{}` (Any)

In Go, `interface{}` is the empty interface, capable of holding values of any type. It’s akin to `Any` in other languages.

```go
var anyValue interface{} // var anyValue any
anyValue = 42
anyValue = "Hello, Go!"
anyValue = 3.14
```

### Pros
- **Flexibility:** Can store values of any type.
- **Generic Data Structures:** Useful for collections holding multiple types. 
- **Extendability:** Useful for preventing duplication for generic behaviors like print().

### Cons
- **Type Safety:** Loses type information, leading to potential runtime errors.
- **Performance Overhead:** Requires type assertions or reflection, which are slower.
- **Less Readable Code:** Makes the code harder to understand and maintain.

**Why Not Use Everywhere?**
Using `interface{}` everywhere undermines Go’s strong type system, leading to fragile and error-prone code.

---

### Type Switches

Type switches allow you to determine the dynamic type of an `interface{}` value.

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case float64:
        fmt.Printf("Float64: %f\n", v)
    default:
        fmt.Printf("Unknown type\n")
    }
}

describe(10)        // Integer: 10
describe("GoLang")  // String: GoLang
describe(3.14)      // Float64: 3.140000
```

---

### Extracting Types

To use the underlying value, perform a type assertion.

```go
var anyValue interface{} = "Hello"

str, ok := anyValue.(string)
if ok {
    fmt.Println("String value:", str)
} else {
    fmt.Println("Not a string")
}
```

---

### Limitations of `interface{}` 

- **Type Safety:** Operations on `interface{}` require type assertions, risking panics.
- **Code Clarity:** Less clear what types are expected.
- **Performance:** Type assertions and reflection can be costly.

#### Solution?
Generics provide type safety and performance benefits by allowing functions and data structures to operate on any specified type without sacrificing type information.

---

## 4. Introduction to Generics

Generics, introduced in Go 1.18, allow writing flexible and type-safe code.

### Example: Adding Two Variables

Using `interface{}` And Without Generics:

```go
func Add(a, b interface{}) interface{} {
    switch a := a.(type) {
    case int:
        if b, ok := b.(int); ok {
            return a + b
        }
    case float64:
        if b, ok := b.(float64); ok {
            return a + b
        }
    case string:
        if b, ok := b.(string); ok {
            return a + b
        }
    }
    return nil
}

fmt.Println(Add(2, 3))           // 5
fmt.Println(Add(2.5, 3.1))       // 5.6
fmt.Println(Add("Hello, ", "Go")) // Hello, Go
```

Using Generics:

```go
package main

import "fmt"

// Define a type constraint
type Addable interface {
    int | float64 | string
}

// Generic Add function
func Add[T Addable](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Add(2, 3))                   // 5
    fmt.Println(Add(2.5, 3.1))               // 5.6
    fmt.Println(Add("Hello, ", "Go"))        // Hello, Go
}
```
