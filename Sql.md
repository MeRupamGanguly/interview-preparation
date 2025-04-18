```bash
rupx@dev:~$ docker volume create rupx-postgres-data
rupx-postgres-data
rupx@dev:~$ docker network create rupx-net
493210af5b0cb70f639c3fd63f088eeec862eb0dcba363c75abe7d3d9ef22ccc
rupx@dev:~$ docker run -d \
   --name postgres \
   --network=rupx-net \
   -v rupx-postgres-data:/var/lib/postgresql/data \
   -e POSTGRES_USER=myuser \
   -e POSTGRES_PASSWORD=mypassword \
   -e POSTGRES_DB=mydatabase \
   -p 5432:5432 \
   postgres:latest

rupx@dev:~$ docker exec -it postgres /bin/bash
root@a447994125b9:/# psql -U myuser -d mydatabase
mydatabase=# \c mydatabase
You are now connected to database "mydatabase" as user "myuser".
mydatabase=# CREATE TABLE Users (
    UserID INT PRIMARY KEY,
    UserName VARCHAR(50),
    ReferrerID INT
);
CREATE TABLE
mydatabase=# CREATE TABLE Products (
    ProductID INT PRIMARY KEY,
    ProductName VARCHAR(50)
);
CREATE TABLE
mydatabase=# CREATE TABLE Orders (
    OrderID INT PRIMARY KEY,
    UserID INT,
    ProductID INT,
    FOREIGN KEY (UserID) REFERENCES Users(UserID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);
CREATE TABLE
mydatabase=# \dt
         List of relations
 Schema |   Name   | Type  | Owner  
--------+----------+-------+--------
 public | orders   | table | myuser
 public | products | table | myuser
 public | users    | table | myuser
(3 rows)

mydatabase=# INSERT INTO Users (UserID, UserName, ReferrerID) VALUES
(1, 'Alice', NULL),
(2, 'Bob', 1),
(3, 'Charlie', 1),
(4, 'Diana', 2);
INSERT 0 4
mydatabase=# INSERT INTO Products (ProductID, ProductName) VALUES
(1001, 'Laptop'),
(1002, 'Phone'),
(1003, 'Tablet'),
(1004, 'Monitor');
INSERT 0 4
mydatabase=# INSERT INTO Orders (OrderID, UserID, ProductID) VALUES
(101, 1, 1001),
(102, 2, 1002),
(103, 2, 1003);
INSERT 0 3
mydatabase=# 
```

# JOINS
1. INNER JOIN : Returns only matching rows from both tables.

2. OUTER JOIN:

    LEFT JOIN : Returns all rows from the left table, and matching rows from the right. If no match, NULLs for the right side.

    RIGHT JOIN : Returns all rows from the right table, and matching rows from the left. NULLs for the left side if no match.

    FULL JOIN : Returns all rows from both tables. If there's no match on either side, fills with NULLs.

3. CROSS JOIN : Returns the Cartesian product – all combinations.

SELF JOIN : Joins a table to itself.

Users Table:

| UserID | UserName | ReferrerID |
|--------|----------|------------|
| 1      | Alice    | NULL       |
| 2      | Bob      | 1          |
| 3      | Charlie  | 1          |
| 4      | Diana    | 2          |

Orders Table:

| OrderID | UserID | ProductID |
|---------|--------|-----------|
| 101     | 1      | 1001      |
| 102     | 2      | 1002      |
| 103     | 2      | 1003      |
| 104     | 5      | 1001      | 

Products Table:

| ProductID | ProductName |
|-----------|-------------|
| 1001      | Laptop      |
| 1002      | Phone       |
| 1003      | Tablet      |
| 1004      | Monitor     |


## Q1: Get users who placed orders
```bash
mydatabase=# SELECT UserName, Orders.OrderID
FROM Orders
INNER JOIN Users ON Orders.UserID = Users.UserID;
```
| UserName | OrderID |
|----------|---------|
| Alice    | 101     |
| Bob      | 102     |
| Bob      | 103     |
(3 rows)

## Q2: Get all users and any orders they made
```bash
mydatabase=# SELECT Users.UserName, Orders.OrderID
FROM Users
LEFT JOIN Orders ON Users.UserID = Orders.UserID;
```
| UserName | OrderID |
|----------|---------|
| Alice    | 101     |
| Bob      | 102     |
| Bob      | 103     |
| Diana    |         |
| Charlie  |         |
(5 rows)


## Q3: Get all orders and the users who made them
```bash
mydatabase=# INSERT INTO Orders (OrderID, UserID, ProductID) VALUES
(105, NULL, 1003);
```
INSERT 0 1
| OrderID | UserID | ProductID |
|---------|--------|-----------|
| 101     | 1      | 1001      |
| 102     | 2      | 1002      |
| 103     | 2      | 1003      |
| 105     |        | 1003      |

```bash
mydatabase=# SELECT Users.UserName, Orders.OrderID
FROM Users
RIGHT JOIN Orders ON Users.UserID = Orders.UserID;
```
| UserName | OrderID |
|----------|---------|
| Alice    | 101     |
| Bob      | 102     |
| Bob      | 103     |
|          | 105     |

## Q4: All users and all orders, matching where possible
```bash
mydatabase=# SELECT Users.UserName, Orders.OrderID
FROM Users
FULL OUTER JOIN Orders ON Users.UserID = Orders.UserID;
```
| UserName | OrderID |
|----------|---------|
| Alice    | 101     |
| Bob      | 102     |
| Bob      | 103     |
|          | 105     |
| Diana    |         |
| Charlie  |         |

## Q5: All possible User-Product combinations
```bash
mydatabase=# SELECT Users.UserName, Products.ProductName
FROM Users
CROSS JOIN Products;
```
| UserName | ProductName |
|----------|-------------|
| Alice    | Laptop      |
| Bob      | Laptop      |
| Charlie  | Laptop      |
| Diana    | Laptop      |
| Alice    | Phone       |
| Bob      | Phone       |
| Charlie  | Phone       |
| Diana    | Phone       |
| Alice    | Tablet      |
| Bob      | Tablet      |
| Charlie  | Tablet      |
| Diana    | Tablet      |
| Alice    | Monitor     |
| Bob      | Monitor     |
| Charlie  | Monitor     |
| Diana    | Monitor     |


## Q6:  who referred whom
```bash
mydatabase=# SELECT u1.UserName AS User , u2.UserName AS Referrer
FROM Users u1
LEFT JOIN Users u2 ON u1.ReferrerID = u2.UserID;
```
| User     | Referrer |
|----------|----------|
| Alice    |          |
| Bob      | Alice    |
| Charlie  | Alice    |
| Diana    | Bob      |


## Q7: Get user's name, order ID, and product name
```bash
mydatabase=# SELECT u.UserName, o.OrderID, p.ProductName
FROM Users u
INNER JOIN Orders o ON u.UserID = o.UserID
INNER JOIN Products p ON o.ProductID = p.ProductID;
```
| UserName | OrderID | ProductName |
|----------|---------|-------------|
| Alice    | 101     | Laptop      |
| Bob      | 102     | Phone       |
| Bob      | 103     | Tablet      |


# ACID Properties in Database Transactions

- Concurrency control in a Database Management System (DBMS) ensures that multiple transactions can execute simultaneously without leading to data inconsistencies. It manages the execution of transactions to maintain the ACID properties—Atomicity, Consistency, Isolation, and Durability.

Atomicity : Eiher every operation of transaction will complete or All failed. If one operation failed then every operation in transaction rolled back.

Consistency : After and before any Transaction, Database should maintain the desired valid state.

Isolation : Transactions can happen simultaneously but each transaction would processed independly. Two people ordering the last item then only one order will sucess and other will failed.

Durability : Once a transaction is complete, even after any kind of crash, in database the transaction will show completed.  

# Relationship
a relationship refers to the way tables are connected to each other. 
One-to-One: A person can have only one passport, and each passport is assigned to only one person.

One-to-Many: A single customer can place multiple orders, but each order is placed by only one customer.

Many-to-Many: Students can enroll in many courses, and each course can have many students.


# DELETE vs TRUNCATE vs DROP 
The DELETE statement is used to remove one or more rows from a table based on a condition using the WHERE clause. It is a Data Manipulation Language (DML) command. Each row deletion is logged in the transaction log, making it slower for large datasets. Rollback is possible if used within a transaction. The table structure and its schema remain unchanged after deletion. `DELETE FROM Employees WHERE Department = 'HR';`

The TRUNCATE command is used to remove all rows from a table without logging individual row deletions. It is a Data Definition Language (DDL) command. Truncation is faster than DELETE, especially on large tables. Rollback is not possible in most databases once the truncate operation is committed. The table structure, constraints, and schema remain intact. `TRUNCATE TABLE Employees;`

The DROP statement is used to completely remove a table, view, or database. It is also a DDL command. It not only deletes all the rows in the table but also removes the table structure, constraints, indexes, and permissions. It cannot be rolled back in most systems. `DROP TABLE Employees;`


# Normalization
Normalization is the process of organizing data to reduce redundancy, Prevent update, insert, and delete anomalies. Updating a single piece of data in multiple places leads to inconsistency. When new data cannot be added to the database without including unwanted or irrelevant information. Deleting a record accidentally removes important data.

1NF – Eliminate repeating groups : Ensures that each column contains atomic values, and each row is unique. A table of students with multiple phone numbers in a single column is split into multiple rows where each phone number is stored separately.

2NF – Remove partial dependencies : If a table has both OrderID and ProductID as the primary key, and ProductName only depends on ProductID, then ProductName should be moved to a separate table. This way, every piece of data depends on the full primary key, not just one part.

3NF – Remove transitive dependencies : In a Student table, if the student's department name depends on the department ID, we move department info to a separate table to avoid redundancy.

BCNF – Stronger version of 3NF

# Primary Key 
A Primary Key is a column (or a set of columns) that uniquely identifies each record in a table. It cannot be NULL, and it must be unique. 

# Foreign Key 
A Foreign Key is a field (or collection of fields) in one table that refers to the Primary Key in another table.
```bash
CREATE TABLE Orders (
  OrderID INT PRIMARY KEY,
  CustomerID INT,
  FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID)
);
```

# Triggers
A Trigger is a set of SQL statements that automatically executes in response to a specific event on a table or view.

We want to automatically update a LastModified timestamp field any time a row is updated
```bash
CREATE TRIGGER trg_UpdateTimestamp # creates a new trigger named trg_UpdateTimestamp.
AFTER UPDATE ON Employees
FOR EACH ROW # he trigger will run once per row that is updated, not just once per statement. If your UPDATE statement affects 5 rows, this trigger will run 5 times — once for each row.
BEGIN # trigger body
   UPDATE Employees 
   SET LastModified = NOW() # This updates the same row that was just updated, setting the LastModified column to the current timestamp using NOW()
   WHERE EmployeeID = OLD.EmployeeID; # OLD.EmployeeID refers to the value of EmployeeID before the update.
END;
```

# CAP :
The CAP Theorem states that a distributed database system can provide only two out of the following three guarantees:
- Consistency: Every read receives the most recent write.
- Availability: Every request (read/write) receives a response, even if some nodes are unavailable.
- Partition Tolerance: The system continues to operate despite network partitions.

# Indexing
Indexing is a technique used in databases to improve the speed of data retrieval operations on a table at the cost of additional storage and write performance. Indexes are created on columns that are frequently used in WHERE, JOIN, ORDER BY, or GROUP BY clauses.

- Without an index, the DBMS performs a full table scan. 
- With an index, it can quickly look up the location of matching rows, improving query performance significantly.

### Indexing Strategies :-

      - B-tree (default) Good for =, <, >, BETWEEN, ORDER BY.

      - GIN (Generalized Inverted Index) Stores multiple index entries per row. Used for: full-text search, arrays, JSONB, tsvector.

      - GiST (Generalized Search Tree) Good for nearest-neighbor, ranges.

      - Hash Optimized for = only.

- Single-column index – Fast, simple.
```bash
CREATE INDEX idx_email ON users(email);
SELECT * FROM users WHERE email = 'x@example.com';
```
- Multicolumn index – Useful when queries filter on multiple fields.
PostgreSQL uses the index efficiently only if the WHERE clause includes the leftmost column(s).
```bash
CREATE INDEX idx_name_dob ON employees(last_name, date_of_birth);
#-- Uses index (filters on both)
SELECT * FROM employees WHERE last_name = 'Smith' AND date_of_birth = '1990-01-01';
#-- Uses index (filters on first column only)
SELECT * FROM employees WHERE last_name = 'Smith';
#-- ❌ Will NOT use index (skips first column)
SELECT * FROM employees WHERE date_of_birth = '1990-01-01';
```
- Partial index – Index only a subset of rows.
```bash
CREATE INDEX idx_active_email ON users(email) WHERE is_active = true;
#-- Query that uses the index
SELECT * FROM users WHERE email = 'john@example.com' AND is_active = true;
#-- ❌ Query that won't use the index (missing the WHERE clause match)
SELECT * FROM users WHERE email = 'john@example.com';
```

# POINTS 
- Transaction : A transaction is a sequence of operations performed as a single logical unit of work. It must satisfy ACID properties.

- Deadlock : A deadlock occurs when two or more transactions are waiting for each other’s resources, creating a cycle where no transaction can proceed.

- Cascading : Cascading actions define what happens when a referenced row in the parent table is updated or deleted.

      ON DELETE CASCADE – Deletes related rows in child table
      ON UPDATE CASCADE – Updates foreign keys in child table
      SET NULL / SET DEFAULT – Sets foreign key to NULL or default

    ```bash
    FOREIGN KEY (DeptID) REFERENCES Departments(DeptID)
    ON DELETE CASCADE
    ON UPDATE CASCADE;
    ```
- A surrogate key is a system-generated unique identifier, often an auto-incremented number. It's used instead of natural keys to uniquely identify a record.
- A composite key is a primary key that consists of two or more columns. It is used when a single column cannot uniquely identify a record.
- A schema is the structure of a database, defined by a collection of tables, views, indexes, procedures, and other database objects. It acts as a blueprint for how the data is organized.
- Referential Integrity ensures that foreign key values in a table match primary key values in the related table, maintaining valid references between tables.
- The WHERE clause is used to filter individual rows before any grouping or aggregation happens. It operates on raw data in the table.
- The HAVING clause is used to filter aggregated results after the GROUP BY has grouped the data. It works on grouped records, often with aggregate functions like COUNT, SUM, AVG, etc.
    ```bash
    SELECT department, COUNT(*) AS emp_count
    FROM employees
    WHERE status = 'active'         -- Filters individual rows
    GROUP BY department
    HAVING COUNT(*) > 10;   
    ```


- Clustered index : 
A clustered index means the actual table data is stored in the same order as the index.
    ```bash
    CREATE INDEX idx_emp_id ON employees(employee_id);
    CLUSTER employees USING idx_emp_id;
    ```
  Future inserts won't follow this order.It's not maintained automatically. You’d need to manually recluster again later.

  All indexes in PostgreSQL are non-clustered by default. When you: `CREATE INDEX idx_email ON users(email);` You're building a separate structure that helps look up users by email quickly, but it doesn’t change how the rows are stored in the table.

- CTID : Every index in PostgreSQL uses CTID to point to the actual row.
It allows fast lookups, but changes if the row is updated/moved.
`"Alice" → (0,1)` 

      SELECT * FROM employees WHERE name = 'Alice';

      Looks up "Alice" in the B-tree index.

      Finds the CTID: (0,1)

      Uses that CTID to jump directly to the row in the table heap.

- View : A view is a virtual table based on the result of a SQL query. It does not store data physically. Views can be used to present data in a specific way without altering the underlying tables.
  ```bash
  CREATE VIEW ActiveEmployees AS
  SELECT Name, Department FROM Employees WHERE Status = 'Active';
  ```




