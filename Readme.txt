Problem Statement: Inventory Management System (Go Fundamentals Test)

Introduction

You are tasked with building a lightweight inventory management system in Go. This system should allow adding, updating, and retrieving products while interacting with an external API for additional data.

This problem tests your understanding of structs, loops, methods, interfaces, error handling, and interacting with external systems.


---

Requirements

1. Define Product Struct and Basic Methods

Create a Product struct with the following fields:

type Product struct {
    ID          int
    Name        string
    Price       float64
    Quantity    int
    Category    string
}

Implement the following methods:

1. UpdatePrice(newPrice float64)

Updates the product’s price.



2. Sell(quantity int) error

Reduces the quantity if stock is available.

Returns an error if the requested quantity is unavailable.



3. Restock(quantity int)

Increases the quantity of a product.



4. Display() string

Returns a formatted string with product details.





---

2. Create an Inventory System

Define an Inventory struct to manage multiple Product instances.

type Inventory struct {
    Products map[int]Product
}

Implement the following methods:

1. AddProduct(p Product)

Adds a new product to the inventory.

If a product with the same ID exists, return an error.



2. RemoveProduct(id int) error

Removes a product by ID.

Returns an error if the product does not exist.



3. FindProductByName(name string) (*Product, error)

Searches for a product by name.

Returns an error if not found.



4. ListByCategory(category string) []Product

Returns all products in a given category.



5. TotalValue() float64

Computes the total value of all products in inventory.





---

3. Implement Database Storage (Mocked)

To simulate a database, create an interface ProductStorage:

type ProductStorage interface {
    Save(p Product) error
    GetByID(id int) (*Product, error)
    Delete(id int) error
}

Implement two storage options:

1. MemoryStorage

Uses a simple map[int]Product as an in-memory store.



2. MockDatabaseStorage

Simulates database interaction by introducing artificial delays/errors.




Modify Inventory to use a ProductStorage implementation instead of directly managing map[int]Product.


---

4. External API Integration (Simulated API Calls)

Create a struct to interact with an external API:

type ExternalAPI struct {
    BaseURL string
}

Implement:

1. FetchProductDetails(id int) (*Product, error)

Simulates an API call to fetch product details.

Uses time.Sleep() to introduce a random delay (0-3s).

Randomly returns an error for some requests (e.g., 10% failure rate).





---

Bonus Challenges (Advanced Concepts)

1. Concurrency: Fetch multiple products in parallel using goroutines.


2. Unit Testing: Write tests for all methods.


3. JSON Serialization: Save and load products using encoding/json.


4. REST API: Implement CRUD routes using Gin or Chi.




---

Evaluation Criteria

✔ Structs & Methods: Proper definition of structs and methods.
✔ Loops & Error Handling: Correctly iterate and handle edge cases.
✔ Interfaces & Mocking: Use ProductStorage to demonstrate abstraction.
✔ Concurrency & API Calls: Handle delays, errors, and cancellations.
✔ Clean Code: Readability, organization, and proper structuring.


---

Expected Deliverables

Fully implemented Go code with all functions.

Explanation of key design decisions.