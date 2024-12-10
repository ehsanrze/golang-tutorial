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


Note: We could have receivers for other types as well (not only struct)

```go
type str string

func (text str) log() {
	fmt.Println(text)
}

str.log()
```
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


#### How to have constructor on golang?

In Go, **constructors** are **not a built-in feature** like in object-oriented languages.
Instead, developers typically use functions to initialize and return instances of structs.

##### Use a New function

The idiomatic way to name constructors in Go is to prefix the function name with New. For example:

```go
type File struct {
    name    string
    path    string
    size    int
}

// we could return the copy of the value
func New (name string, path string, size int) File {
	return File{
		Name:   "test", 
		Path:    "./folder1/test.txt", 
		Size: 50,
	}
}

// or pointer -> it is better to avoid copying
func New (name string, path string, size int) *File {
      return &File{
		  Name:   "test", 
		  Path:    "./folder1/test.txt", 
		  Size: 50,
	  }
}


// with input validation -> best practice
func New(name string, path string, size int) (*File, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if path == "" {
		return nil, errors.New("path cannot be empty")
	}
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	
	return &File {
		Name: name, 
		Path: path, 
		Size: size,
	}, nil
}
```

#### What is struct embedding?
struct embedding is a way to include one struct within another.
It allows the embedding struct to inherit fields and methods of the embedded struct, enabling composition over inheritance.

It has two types:
1. within a name
2. without name (anonymous field)

```go

type Address struct {
  City    string
  Country string
}


// former option (within a name)
type Person struct {
  Name    string
  Age     int
  Address
}

// latter option (anonymous field)
type Person struct {
  Name    string
  Age     int
  Addr    Address
}
```

##### Some key points:
**Overriding**

```go
type Person struct {
    Name    string
    Address
    City string // Overrides Address.City
}
```

**Multiple Embeddings**

```go
type Employee struct {
    ID int
    Address
    ContactInfo
}
```

#### What is struct tags? (metadata)

struct tags are special annotations added to struct fields. They give extra details about how the fields should be handled, especially when working with tasks like converting data to JSON, interacting with databases, or validating input.

##### Syntax

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}
```

##### Purpose

**Encoding/Decoding**

Specify how struct fields are marshaled/unmarshaled in formats like JSON, XML, or YAML.

```go
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}
```

**Database Mapping**

Match struct fields with database table columns.

```go
type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
}
```

**Validation**

Define validation rules for struct fields (e.g., using libraries like go-playground/validator).

```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"email"`
}
```