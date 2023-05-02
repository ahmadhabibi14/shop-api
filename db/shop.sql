CREATE TABLE Customer (
   id VARCHAR(50) PRIMARY KEY NOT NULL,
   name VARCHAR(100) NOT NULL
);

CREATE TABLE Transaction (
   id VARCHAR(50) PRIMARY KEY NOT NULL,
   customer_id VARCHAR(50) NOT NULL,
   menu VARCHAR(100) NOT NULL,
   price INT NOT NULL,
   qty INT NOT NULL,
   payment VARCHAR(100) NOT NULL,
   total INT NOT NULL,
   created_at DATETIME,
   FOREIGN KEY (customer_id) REFERENCES Customer(id)
);


CREATE TABLE User (
   id VARCHAR(50) PRIMARY KEY NOT NULL,
   username VARCHAR(100) NOT NULL
);

CREATE TABLE Post (
   id VARCHAR(50) PRIMARY KEY NOT NULL,
   user_id VARCHAR(50) NOT NULL,
   text_post TEXT NOT NULL,
   created_at DATETIME,
   FOREIGN KEY (user_id) REFERENCES User(id)
);