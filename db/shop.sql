CREATE TABLE users (
   id VARCHAR(50) PRIMARY KEY NOT NULL,
   username VARCHAR(100) NOT NULL,
   password VARCHAR(100) NOT NULL
);

CREATE TABLE transaction (
   id VARCHAR(50) PRIMARY KEY NOT NULL,
   customer_id VARCHAR(50) NOT NULL,
   menu VARCHAR(100) NOT NULL,
   price INT NOT NULL,
   qty INT NOT NULL,
   payment VARCHAR(100) NOT NULL,
   total INT NOT NULL,
   created_at DATETIME,
   FOREIGN KEY (customer_id) REFERENCES users(id)
);