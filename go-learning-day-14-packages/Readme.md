# 🐹 Go Learning — Day 14

# Go Modules

> **Goal:** Understand Go Modules from absolute zero — what they are, why they exist, how they work, and how to manage dependencies in a real Go project.

---

# 📚 Table of Contents

1. [What Problem Are Go Modules Solving?](#1-what-problem-are-go-modules-solving)
2. [What Is a Go Module?](#2-what-is-a-go-module)
3. [Module vs Package](#3-module-vs-package)
4. [The `go.mod` File](#4-the-gomod-file)
5. [`go mod init`](#5-go-mod-init)
6. [Module Path](#6-module-path)
7. [How Packages Are Imported](#7-how-packages-are-imported)
8. [Why Do We Need `go.mod`?](#8-why-do-we-need-gomod)
9. [Adding External Dependencies](#9-adding-external-dependencies)
10. [`go get`](#10-go-get)
11. [`go get` With a Version](#11-go-get-with-a-version)
12. [`go.sum`](#12-gosum)
13. [`go mod tidy`](#13-go-mod-tidy)
14. [`go list -m all`](#14-go-list-m-all)
15. [`go mod download`](#15-go-mod-download)
16. [`go mod verify`](#16-go-mod-verify)
17. [Direct and Indirect Dependencies](#17-direct-and-indirect-dependencies)
18. [Complete Module Workflow](#18-complete-module-workflow)
19. [Practical Project](#19-practical-project)
20. [Important Commands Cheat Sheet](#20-important-commands-cheat-sheet)
21. [Final Mental Model](#21-final-mental-model)

---

# 1. What Problem Are Go Modules Solving?

Before understanding Modules, imagine we create a Go project:

```text
my-project/
│
├── main.go
├── calculator/
│   └── calculator.go
└── user/
    └── user.go
```

Our project contains our own packages.

But eventually we may want to use external libraries:

```text
Gin
JWT
MongoDB Driver
PostgreSQL Driver
OpenTelemetry
```

Now Go needs to know:

* What is this project?
* Which dependencies does it use?
* Which versions of those dependencies?
* Where can those dependencies be found?
* How can another developer reproduce the same project?

This is where **Go Modules** come in.

---

# 2. What Is a Go Module?

A **Go Module** is a collection of Go packages that are developed and versioned together as one project.

Think about it like a house:

```text
                🏠 MODULE
                   │
        ┌──────────┼──────────┐
        │          │          │
        ↓          ↓          ↓
   Calculator    User       Product
    Package     Package      Package
```

The:

```text
🏠 Module
```

is the entire project.

The:

```text
🚪 Packages
```

are the individual parts inside the project.

A module also keeps track of the project's external dependencies.

---

# 3. Module vs Package

This is one of the most important concepts.

## 📦 Package

A package is a collection of related Go source files.

Example:

```text
calculator/
│
├── add.go
├── subtract.go
└── multiply.go
```

Inside those files:

```go
package calculator
```

This tells Go:

> These files belong to the `calculator` package.

---

## 🏠 Module

A module is the larger project that can contain multiple packages.

Example:

```text
my-project/
│
├── go.mod
├── main.go
│
├── calculator/
│   └── calculator.go
│
├── user/
│   └── user.go
│
└── product/
    └── product.go
```

Here:

```text
Module
  │
  ├── main package
  ├── calculator package
  ├── user package
  └── product package
```

### Simple rule

> **Package = group of Go code**

> **Module = entire Go project + its dependency information**

---

# 4. The `go.mod` File

When we create a Go Module, Go creates:

```text
go.mod
```

Example:

```go
module github.com/tahashafiq175/my-project

go 1.26
```

The `go.mod` file tells Go important information about the module.

At the basic level:

```text
go.mod
  │
  ├── Module identity
  │
  ├── Go version
  │
  └── Dependency requirements
```

Later, dependencies may appear:

```go
module github.com/tahashafiq175/my-project

go 1.26

require github.com/fatih/color v1.18.0
```

Now Go knows that this project requires the `fatih/color` module.

---

# 5. `go mod init`

To turn a normal directory into a Go Module, use:

```bash
go mod init MODULE_PATH
```

Example:

```bash
go mod init github.com/tahashafiq175/my-project
```

Go creates:

```text
my-project/
│
├── go.mod
└── ...
```

And `go.mod` contains:

```go
module github.com/tahashafiq175/my-project

go 1.26
```

## What does `go mod init` actually do?

It tells Go:

> **"This directory is the root of a Go Module, and this is the module's identity."**

It does NOT:

* install every dependency
* create every package
* write your application code
* automatically download every library

It mainly **initializes the module**.

---

# 6. Module Path

Consider:

```go
module github.com/tahashafiq175/my-project
```

This is called the **Module Path**.

It identifies the module.

For a GitHub project, it commonly matches:

```text
github.com/username/repository
```

For example:

```text
github.com/tahashafiq175/go-learning-day-14-packages
```

can identify:

```text
GitHub
   ↓
tahashafiq175
   ↓
go-learning-day-14-packages
```

The module path becomes important when packages inside the module are imported or when another project uses your module.

---

# 7. How Packages Are Imported

Suppose our project looks like:

```text
my-project/
│
├── go.mod
├── main.go
│
└── calculator/
    └── calculator.go
```

`go.mod`:

```go
module github.com/tahashafiq175/my-project
```

`calculator.go`:

```go
package calculator

func Add(a, b int) int {
    return a + b
}
```

Now `main.go` can import it:

```go
package main

import (
    "fmt"

    "github.com/tahashafiq175/my-project/calculator"
)

func main() {
    result := calculator.Add(10, 20)

    fmt.Println(result)
}
```

Output:

```text
30
```

---

# 8. Why Do We Need `go.mod`?

This is where an important question appears:

> "If I can access my package using the path, why do I need `go.mod`?"

Excellent question.

The import path:

```text
github.com/tahashafiq175/my-project/calculator
```

is based on the module path:

```text
github.com/tahashafiq175/my-project
```

plus the package directory:

```text
calculator
```

So:

```text
MODULE PATH
     +
PACKAGE DIRECTORY
     ↓
IMPORT PATH
```

Therefore:

```text
github.com/tahashafiq175/my-project
                    +
               /calculator
                    ↓
github.com/tahashafiq175/my-project/calculator
```

The `go.mod` file tells Go:

> **"This is where my module begins."**

---

# 9. Adding External Dependencies

Our own packages aren't the only packages we can use.

Suppose we want to use:

```text
github.com/fatih/color
```

This is an external module.

We can add it using:

```bash
go get github.com/fatih/color
```

Go resolves the module and updates the module information.

Our project may now contain:

```text
my-project/
│
├── go.mod
├── go.sum
└── main.go
```

And `go.mod` may contain:

```go
require github.com/fatih/color v1.18.0
```

The exact version can differ depending on what Go resolves.

---

# 10. `go get`

The basic command is:

```bash
go get PACKAGE
```

Example:

```bash
go get github.com/fatih/color
```

Think of it as:

> **"Add this dependency to my Go Module."**

The general syntax is:

```bash
go get MODULE_PATH
```

---

# 11. `go get` With a Version

We can also request a specific version:

```bash
go get github.com/fatih/color@v1.18.0
```

The structure is:

```text
go get MODULE_PATH@VERSION
```

Example:

```bash
go get github.com/example/library@v1.2.3
```

This means:

> "Add this module and use version `v1.2.3`."

---

# 12. `go.sum`

When dependencies are used, Go may create:

```text
go.sum
```

A very simple way to understand the two files is:

```text
go.mod
   ↓
"What dependencies and versions does my project need?"

go.sum
   ↓
"What cryptographic checksums should those downloaded modules have?"
```

---

## 🔐 Simple Analogy

Imagine ordering a product.

```text
go.mod
   ↓
"I ordered Product X version 1.0"
```

```text
go.sum
   ↓
"This is the expected fingerprint of Product X"
```

Go uses checksums to help verify the integrity of module contents.

You normally **do not manually edit `go.sum`**.

Go manages it.

---

# 13. `go mod tidy`

One of the most useful commands is:

```bash
go mod tidy
```

Think:

> 🧹 **"Go, clean up my dependencies."**

Go examines your source code and makes the module dependency information consistent with what the project actually uses.

It can:

* Add missing dependencies
* Remove unnecessary dependencies
* Update `go.mod`
* Update `go.sum`

Conceptually:

```text
              go mod tidy
                   │
                   ↓
            Analyze source code
                   │
          ┌────────┴────────┐
          ↓                 ↓
    Missing deps       Unused deps
          ↓                 ↓
        Add               Remove
          └────────┬────────┘
                   ↓
            Update go.mod
                   ↓
            Update go.sum
```

### Easy memory trick

```text
go mod tidy = dependency cleanup 🧹
```

---

# 14. `go list -m all`

To list the modules involved in the current module:

```bash
go list -m all
```

Notice the order:

```bash
go list -m all
```

NOT:

```bash
go list all -m
```

Here:

```text
-m = option
all = module pattern
```

The command means:

> **"List all modules involved in this project."**

You may see:

```text
github.com/tahashafiq175/my-project
github.com/fatih/color v1.18.0
...
```

---

# 15. `go mod download`

This command downloads the modules required by your project:

```bash
go mod download
```

A common use case is CI/CD.

For example:

```text
GitHub Actions
      │
      ↓
Clone repository
      │
      ↓
go mod download
      │
      ↓
Dependencies available
      │
      ↓
go test
      │
      ↓
go build
```

---

# 16. `go mod verify`

You can verify downloaded modules:

```bash
go mod verify
```

It checks the downloaded module contents against the expected information.

Think:

```text
Downloaded dependency
        ↓
Check its integrity
        ↓
Pass / Fail
```

---

# 17. Direct and Indirect Dependencies

You may see something like:

```go
require (
    github.com/fatih/color v1.18.0
    some-library v1.2.0 // indirect
)
```

## Direct Dependency

Your code directly uses it.

```go
import "github.com/fatih/color"
```

Therefore:

```text
Your Application
       ↓
      color
```

---

## Indirect Dependency

A dependency of one of your dependencies.

Example:

```text
Your Application
       ↓
      Gin
       ↓
   Library X
```

Your application directly depends on Gin.

Gin depends on Library X.

Therefore Library X is an **indirect dependency** of your project.

---

# 18. Complete Module Workflow

A typical Go project can follow this workflow:

```text
             Create Project
                   │
                   ↓
             go mod init
                   │
                   ↓
              go.mod created
                   │
                   ↓
              Write Code
                   │
                   ↓
        Need External Dependency?
                   │
                  YES
                   ↓
                go get
                   │
                   ↓
             go.mod updated
                   │
                   ↓
             go.sum updated
                   │
                   ↓
             go mod tidy
                   │
                   ↓
         Dependency information
                cleaned
                   │
                   ↓
             go test / run
                   │
                   ↓
               go build
```

---

# 19. Practical Project

Let's create a simple module.

## Step 1 — Create the project

```bash
mkdir go-module-demo
cd go-module-demo
```

---

## Step 2 — Initialize the module

```bash
go mod init github.com/tahashafiq175/go-module-demo
```

Now:

```text
go-module-demo/
└── go.mod
```

---

## Step 3 — Create a package

```text
go-module-demo/
│
├── go.mod
├── main.go
│
└── calculator/
    └── calculator.go
```

`calculator.go`:

```go
package calculator

func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}
```

---

## Step 4 — Use the package

`main.go`:

```go
package main

import (
    "fmt"

    "github.com/tahashafiq175/go-module-demo/calculator"
)

func main() {

    sum := calculator.Add(10, 20)
    product := calculator.Multiply(10, 20)

    fmt.Println("Sum:", sum)
    fmt.Println("Product:", product)
}
```

Run:

```bash
go run .
```

Output:

```text
Sum: 30
Product: 200
```

---

# 20. Important Commands Cheat Sheet

| Command                  | Purpose                                       |
| ------------------------ | --------------------------------------------- |
| `go mod init MODULE`     | Create a new module                           |
| `go get PACKAGE`         | Add/change a dependency                       |
| `go get PACKAGE@VERSION` | Add/change a dependency to a specific version |
| `go mod tidy`            | Clean and synchronize dependencies            |
| `go list -m all`         | List all modules                              |
| `go mod download`        | Download dependencies                         |
| `go mod verify`          | Verify downloaded modules                     |

---

# 21. Final Mental Model

This is the most important diagram in this lesson:

```text
                         GO PROJECT
                             │
                             ↓
                         GO MODULE
                             │
                         go.mod
                             │
              ┌──────────────┴──────────────┐
              ↓                             ↓
        Your Packages                External Modules
              │                             │
       ┌──────┼──────┐              ┌───────┼───────┐
       ↓      ↓      ↓              ↓       ↓       ↓
   calculator user  product        Gin      JWT     DB
       │
       ↓
 Import Path
       │
       ↓
MODULE PATH + PACKAGE DIRECTORY
```

For example:

```text
Module Path:

github.com/tahashafiq175/my-project

        +

Package Directory:

calculator

        ↓

Import Path:

github.com/tahashafiq175/my-project/calculator
```

---

# 🧠 The Three Most Important Files/Concepts

## `go.mod`

Think:

> **"What is my project, and what modules does it require?"**

---

## `go.sum`

Think:

> **"What are the expected cryptographic fingerprints of my dependencies?"**

---

## Packages

Think:

> **"How is my Go code organized into reusable groups?"**

---

# ⭐ One-Minute Revision

### What is a Module?

A module is the **project-level unit of dependency management** in Go.

### What is `go.mod`?

The file that defines the module's identity, Go version, and dependency requirements.

### What does `go mod init` do?

Creates `go.mod` and initializes the current directory as a Go module.

### What does `go get` do?

Adds or changes a dependency.

### What does `go get package@version` do?

Adds/changes a dependency to a specified version.

### What is `go.sum`?

A file containing checksums used to verify module contents.

### What does `go mod tidy` do?

Cleans and synchronizes the module's dependency information.

### What does `go list -m all` do?

Lists the modules involved in the project.

---

# 🎯 Final Takeaway

Don't think of Go Modules as:

> "Just a `go.mod` file."

Think of them as:

> **The system Go uses to identify a project, organize its dependency requirements, manage dependency versions, and make builds reproducible.**

The relationship is:

```text
                    MODULE
                       │
                       ↓
                    go.mod
                       │
            ┌──────────┴──────────┐
            ↓                     ↓
        Packages            Dependencies
            │                     │
            ↓                     ↓
       Import Paths          Versions
                                  │
                                  ↓
                              go.sum
                                  │
                                  ↓
                          Integrity Checks
```

That is the core idea behind **Go Modules**.
