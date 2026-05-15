# Sift library
## Welcome to my library, and thanks if you using it, there is how to install it and use it

## Installation: 
### Run this command to install: 
```go
go get github.com/nik251021/sift
```
## There is how to use it
### Examples:
* Work with slice: 
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Task: Skip first 2 numbers, and get next 5, 
// And leave only even numbers
res := sift.From(numbers).
    Skip(2).
    Take(5).
    Where(func(n int) bool { return n % 2 == 0 }).
    ToSlice()

// Output: [4, 6]
```
* Work with custom structures:
```go
type User struct {
    Name string
    Age  int
}

users := []User{
    {"Alice", 25},
    {"Bob", 15},
    {"Charlie", 30},
}

// Task: find first user that have at least 18
adult, found := sift.From(users).Find(func(u User) bool {
    return u.Age >= 18
})

if found {
    fmt.Println("Found adult:", adult.Name)
}
```
* Fast validation
```go
scores := []int{85, 92, 78, 100, 64}

// Do all students have more then 60 score?
isEveryonePassed := sift.From(scores).All(func(s int) bool {
    return s > 60
})

// Is there 1 student with perfect score?
hasPerfectScore := sift.From(scores).Any(func(s int) bool {
    return s == 100
})
```

# Update on version v1.0.1
## More features:
* Distinct Easily remove all duplicate values from your collection. It works with any comparable types (numbers, strings, etc.).
```go
numbers := []int{1, 2, 2, 3, 1, 4, 4, 5}
res := sift.From(numbers).Distinct().ToSlice()

// Output: [1, 2, 3, 4, 5]
```
* Select: Transformation (Mapping)
    The Select method allows you to transform elements from one type to another (e.g., from a struct to a string or a number).

    ⚠️ Important Note on Chaining:
    Due to Go's generic limitations, methods cannot introduce new type parameters. This means Select breaks the fluent chain (you can't call it with a dot after .Where()). It must be called as a standalone function.
```go
type User struct {
    ID   int
    Name string
}

users := []User{
    {ID: 1, Name: "Ivan"},
    {ID: 2, Name: "Oleg"},
}

query := sift.From(users)
namesQuery := sift.Select(query, func(u User) string {
    return u.Name
})

names := namesQuery.ToSlice()
// Output: ["Ivan", "Oleg"]
```

# Contributing: "Feel free to open issues or submit pull requests!"