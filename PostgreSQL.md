```bash
docker volume create rupx-postgres-data

docker network create rupx-net

docker run -d \
   --name postgres \
   --network=rupx-net \
   -v rupx-postgres-data:/var/lib/postgresql/data \
   -e POSTGRES_USER=myuser \
   -e POSTGRES_PASSWORD=mypassword \
   -e POSTGRES_DB=mydatabase \
   -p 5432:5432 \
   postgres:latest

docker exec -it postgres /bin/bash

psql -U myuser -d mydatabase

\c mydatabase

CREATE TABLE customers (
    customer_id INT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    country VARCHAR(50),
    signup_date DATE
);

INSERT INTO customers (customer_id, name, email, country, signup_date) VALUES
(1, 'Alice Smith', 'alice@gmail.com', 'USA', '2022-01-10'),
(2, 'Bob Johnson', 'bob@yahoo.com', 'Canada', '2023-03-15'),
(3, 'Charlie Lee', 'charlie@outlook.com', 'UK', '2021-11-05'),
(4, 'Diana Patel', 'diana@protonmail.com', 'USA', '2022-07-20'),
(5, 'Ethan Zhang', 'ethan@163.com', 'China', '2023-01-01');

CREATE TABLE orders (
    order_id INT PRIMARY KEY,
    customer_id INT,
    order_date DATE,
    status VARCHAR(20),
    total_amount DECIMAL(10,2),
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

INSERT INTO orders (order_id, customer_id, order_date, status, total_amount) VALUES
(101, 1, '2023-01-10', 'Completed', 150.00),
(102, 2, '2023-01-15', 'Pending', 300.00),
(103, 1, '2023-02-20', 'Completed', 200.00),
(104, 3, '2023-03-05', 'Cancelled', 0.00),
(105, 4, '2023-03-10', 'Completed', 500.00),
(106, 5, '2023-04-01', 'Shipped', 250.00),
(107, 3, '2023-04-10', 'Completed', 125.00);

CREATE TABLE order_items (
    item_id INT PRIMARY KEY,
    order_id INT,
    product_name VARCHAR(100),
    quantity INT,
    price DECIMAL(10,2),
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
);

INSERT INTO order_items (item_id, order_id, product_name, quantity, price) VALUES
(1, 101, 'Laptop', 1, 150.00),
(2, 102, 'Monitor', 2, 100.00),
(3, 102, 'Mouse', 1, 100.00),
(4, 103, 'Keyboard', 2, 100.00),
(5, 105, 'Smartphone', 1, 500.00),
(6, 106, 'Tablet', 1, 250.00),
(7, 107, 'Webcam', 1, 125.00);

```













Create Table

CREATE TABLE Department (
    department_id INT PRIMARY KEY,
    department_name VARCHAR(50)
);


Add Data

INSERT INTO Department (department_id, department_name) VALUES
(1, 'Engineering'),
(2, 'Marketing'),
(3, 'HR'),
(4, 'Finance');


Get Data

SELECT * FROM Department;



CREATE TABLE Employee (
    employee_id INT PRIMARY KEY,
    name VARCHAR(100),
    department_id INT,
    join_date DATE,
    FOREIGN KEY (department_id) REFERENCES Department(department_id)
);


INSERT INTO Employee (employee_id, name, department_id, join_date) VALUES
(101, 'Alice', 1, '2020-01-10'),
(102, 'Bob', 1, '2021-06-15'),
(103, 'Charlie', 2, '2019-03-22'),
(104, 'David', 3, '2022-07-01'),
(105, 'Eve', 4, '2018-11-30'),
(106, 'Frank', 1, '2023-02-14');


CREATE TABLE Salary (
    salary_id INT PRIMARY KEY,
    employee_id INT,
    base_salary DECIMAL(10, 2),
    bonus DECIMAL(10, 2),
    effective_from DATE,
    FOREIGN KEY (employee_id) REFERENCES Employee(employee_id)
);