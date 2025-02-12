#Inventory Management System:

The Inventory Management System (IMS) is a lightweight application developed in Go, 
designed to handle basic inventory operations such as adding, updating, 
and retrieving products. It also interacts with an external API to fetch additional product data.

#Features:
Add New Products: Introduce new items into the inventory with essential details.
Update Product Information: Modify existing product details, including price and quantity.
Retrieve Product Data: Access information about products in the inventory.
Performs certain operations based on the given Endpoints

#Interacting with the System:
Run the code in your local. It will start a local server on port :8080(localhost:8080/)

Endpoints:
product/create - Add product with request body with json data(ID Name Price Quantity Category)
product/:id - Get Product by id
product/all - Get all products
product/:id - Delete a product
product/:name - Get product by name
product/ctg/:category - Get products by category
product/total - Get the total of all products
product/sell/:id - Sells the product and reduces the product quantity
product/restock/:id - Adds the product quantity
product/price/:id - Updates the product price
