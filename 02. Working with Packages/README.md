#### Package

1. Explanation of Splitting Code Across Files in the Same Package in Go
In Go, you can split your code across multiple files within the same package. This allows you to organize and manage
large codebases by breaking down the code into smaller, more manageable parts, while still keeping everything within
the same package. All the files within a package can share the same namespace and can access each other’s functions,
types, and variables.

- Key Points:
Files Must Be in the Same Directory: All files that belong to the same package must be in the same directory. This
means
that you can organize your code into multiple files, but they must all be part of the same directory and package.

- Consistent Package Name: All files in the same package must declare the same package name at the top of each file.
This
ensures that the files are treated as part of the same package.

- Benefits of Splitting Code Across Multiple Files:
Better Code Organization: Splitting large programs into smaller files helps keep related code together. This makes it
easier to navigate the project and manage different parts of the code.
Improved Readability: Dividing code into logical units (e.g., separating functions for different entities like user
management, database interaction, etc.) improves the readability and maintainability of your codebase.
Easier to Extend and Test: When the code is divided into separate files, it’s easier to add new functionality and test
individual components.

#### Example: Package:

- File 1: circle.go

```go
package shapes

import "math"

func AreaOfCircle(radius float64) float64 {
return math.Pi * radius * radius
}
```

- File 2: rectangle.go

```go
package shapes

func AreaOfRectangle(length, width float64) float64 {
return length * width
}
```

- File 3: main.go

```go
package main

import (
"fmt"
"your_project/shapes"
)

func main() {

circleArea := shapes.AreaOfCircle(5.0)
rectangleArea := shapes.AreaOfRectangle(4.0, 6.0)

fmt.Printf("Area of Circle: %.2f\n", circleArea)
fmt.Printf("Area of Rectangle: %.2f\n", rectangleArea)
}
```

2. Why Would You Use More Than One Package in Go?
In Go, using multiple packages allows you to organize your code in a more modular and maintainable way. As your
project grows, breaking it into separate packages can help you manage different responsibilities and features without
cluttering your code. Each package can encapsulate a specific set of functionalities, making your codebase more
readable, scalable, and easier to maintain.

#### Reasons to Use More Than One Package:

- Separation of Concerns: As projects grow, it’s essential to separate different aspects of the code into logical units.
For example, you can have one package for handling database operations, another for managing HTTP requests, and
another for utility functions. This helps keep each part of your application focused on a single responsibility.

- Better Dependency Management: By using packages, you can manage dependencies more effectively. For instance, one
package might deal with external APIs, while another handles database interactions. This separation ensures that
changes to one package don’t inadvertently affect others.

- Scalability: In large applications, breaking your code into multiple packages makes it easier to scale your
application. You can add new functionality or modify existing parts of the application independently without affecting
the whole system.

- Testability and Maintainability: With separate packages, you can write unit tests for each package independently,
ensuring that individual components work as expected. It also becomes easier to modify or extend features without
touching unrelated parts of the code.

- Reduce Complexity: When you organize your code into well-defined packages, you can reduce the overall complexity of
your codebase. Each package focuses on a specific set of functionalities, making the code easier to understand and
maintain.

3. Preparing Code for Multiple Packages in Go
When using multiple packages in Go, it's essential to structure your code in a way that allows for proper interaction
between different parts of the code. Preparing code for multiple packages involves organizing your project files
correctly, managing dependencies effectively, and ensuring that each package has a clear responsibility. This helps
keep
your codebase modular, scalable, and easy to maintain.

```
your_project/
│
├── main.go
├── db/
│   └── db.go
├── web/
│   └── web.go
└── utils/
└── utils.go
```

4. Importing Packages in Go

- Basic Structure of the import Statement

```
import "path/to/package"
```

- Multi Structure of the import packages

```
import (
"fmt"
"math"
"myapp/db"
)

```

4. Exporting & Importing Identifiers (Variables, Functions & More) in Go

```
package mypackage

var ExportedVar = "I am exported!"  // This variable is accessible outside the package.
var unexportedVar = "I am not exported."  // This variable is private to the package.

```

```
package mypackage

// Exported function
func ExportedFunction() {
fmt.Println("This is an exported function!")
}

// Unexported function
func unexportedFunction() {
fmt.Println("This function is private to the package.")
}

```

5. Using Third-Party Packages

```
go get github.com/sirupsen/logrus
```

```go
package main

import (
"github.com/sirupsen/logrus"
)

func main() {

var log = logrus.New()

log.SetLevel(logrus.InfoLevel)

log.Info("This is an info log message.")

log.Warn("This is a warning log message.")

log.Error("This is an error log message.")
}
```
