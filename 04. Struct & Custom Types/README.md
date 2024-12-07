# Essentials

## Session 04

### Struct & Custom Types

A struct (short for "structure") is a composite data type that groups together variables (fields) under one name to create a more complex data type.

Each field in a struct has a **name** and a **type**.

```go
type File struct {
    name    string
    path    string
    size    int
}
```
#### Why we use it?

- Organized Data Grouping: allow you to group related data fields into a single unit, making it easier to manage and understand your data.
- Encapsulation of Behavior: can have methods attached to them, enabling you to define behavior directly related to the data, creating a clear and cohesive structure.
- Customizable Data Types: enable you to define your own types tailored to specific needs, making your code more meaningful and aligned with the problem domain.


#### How to initialize the struct (struct literal)

```go
// With initialize all fields with values
f := File{
    Name:   "test",
    Path:    "./folder1/test.txt",
    Size: 50, // we always need , for the end of all files.
}

// With Field Order (not recommended for readability)
f := File{"test", "./folder1/test.txt", 50}

// Empty struct initialization
f := File{} // If we don't initialize struct, each field gets default value by its type same as:
f := File {
	Name: "",
	Path: "",
	Size: 0,
}

// Partial initialization (other fields get their default values.
f := File{Name: "test"}
```

#### How can we access the field value?
You can access or modify fields using the dot (.) operator. Like:

```go 
fmt.Println(f.Name)
f.Name = 'test2'
f.Path = './folder1/test2.txt'
```

#### Enhance struct functionality by adding methods
Structs can have methods associated with them by defining a function with a receiver type. Like:

```go
// pass the copy of file to Write() function
func (f File) Write(text string) {
  file, err := os.Create(f.Path)
  if err != nil {
      fmt.Println("Error creating file:", err)
      return
  }
  defer file.Close() // we cover it in the previous session
  _, err = file.WriteString(data)
  return
}

// pass the pointer of the file struct to SetName function
func (f *File) SetName(name string) {
f.Name = name
}
```

##### What is receiver?
a receiver is a special parameter that is used to **associate a function** with **a specific type**, typically a struct. This makes the function a **method of that type**, allowing the function to be **called** using **dot notation**.

**Syntax of a receiver:**
  - The name of the receiver (like a variable)
  - The type of the receiver (the type to which the method belongs)

```go
func (r ReceiverType) MethodName() {
    // method body
}
```

**Types of value receivers:**
- _Value Receiver_
  - A copy of the value is passed to the method.
  - The method cannot modify the original value.

  **Note:** Use this when the method does not need to change the receiver's fields.

- _Pointer Receiver_
  - A pointer to the value is passed to the method.
  - The method can modify the original value through the pointer.
  
  **Note:** Use this when the method needs to change the receiver's fields or to avoid copying large structs.

**What types cannot have receivers**
  - Built-in types like int, string, or float64 directly.
  - Types from other packages (you cannot define methods on them directly).

#### Anonymous Structs

```go
file := struct {
    Name string
    Size int
}{
    Name: "test",
    Size: 50,
}
```

##### When we need it?
  - Small scripts or one-off tasks
  - Functions that require a temporary grouping of fields
  - Test cases or mock data structures
  - Rapid prototyping

##### Why we should avoid them?
  - Non-Reusability
  - Reduced readability 
  - Not suitable for large projects
