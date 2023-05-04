SELECT transaction.menu FROM customer JOIN transaction ON customer.id = transaction.customer_id;

SELECT transaction.menu FROM customer JOIN transaction ON customer.id = transaction.customer_id WHERE customer.name = "Ahmad Habibi";

INSERT INTO transaction VALUES ( "bdusdf0eu20", "ddjh0sd02b", "Gaming Chair", 1540700, 70, "BANK BRI", 2, NOW() )