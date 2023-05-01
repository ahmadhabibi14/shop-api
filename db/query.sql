SELECT transaction.menu FROM customer JOIN transaction ON customer.id = transaction.customer_id;

SELECT transaction.menu FROM customer JOIN transaction ON customer.id = transaction.customer_id WHERE customer.name = "Ahmad Habibi";