/*The "customers" table stores information about the website's customers, including their first and last name, email, and password.*/
CREATE TABLE customers (
    customerID INT AUTO_INCREMENT PRIMARY KEY,
    firstName VARCHAR(50) NOT NULL,
    lastName VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL
);

/*The "categories" table stores information about different product categories, including the name of the category.*/
CREATE TABLE categories (
    categoryID INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

/*The "products" table stores information about the products available for purchase on the website, including their name, price, and the category they belong to.*/
CREATE TABLE products (
    productID INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    categoryID INT NOT NULL,
    FOREIGN KEY (categoryID) REFERENCES categories(categoryID)
);

/*The "orders" table stores information about orders made by customers, including the customer's ID, and the date of the order.*/
CREATE TABLE orders (
    orderID INT AUTO_INCREMENT PRIMARY KEY,
    customerID INT NOT NULL,
    orderDate DATE NOT NULL,
    FOREIGN KEY (customerID) REFERENCES customers(customerID)
);

/*The "orderDetails" table stores information about the products included in each order, including the order ID, product ID, and quantity.*/
CREATE TABLE orderDetails (
    orderID INT NOT NULL,
    productID INT NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (orderID, productID),
    FOREIGN KEY (orderID) REFERENCES orders(orderID),
    FOREIGN KEY (productID) REFERENCES products(productID)
);