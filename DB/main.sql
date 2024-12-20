-- Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    deposit_amount NUMERIC DEFAULT 0
);

-- Books Table
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    stock_availability INTEGER NOT NULL,
    rental_costs NUMERIC NOT NULL,
    category VARCHAR(255) NOT NULL
);

-- Rental History Table
CREATE TABLE rental_history (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    book_id INTEGER NOT NULL REFERENCES books(id),
    rented_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    returned_on TIMESTAMP
);
