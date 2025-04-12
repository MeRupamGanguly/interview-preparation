In an RDBMS (Relational Database Management System), a relationship refers to the way tables are connected to each other.   It’s based on the concept of foreign keys, which link records from one table to another.

One-to-One: A person can have only one passport, and each passport is assigned to only one person.

One-to-Many: A single customer can place multiple orders, but each order is placed by only one customer.

Many-to-Many: Students can enroll in many courses, and each course can have many students.


Normalization is the process of organizing data to reduce redundancy, Prevent update, insert, and delete anomalies. Updating a single piece of data in multiple places leads to inconsistency. When new data cannot be added to the database without including unwanted or irrelevant information. Deleting a record accidentally removes important data.

Normalization is done by dividing large tables into smaller ones and defining relationships between them, based on certain rules 

First Normal Form (1NF): Ensures that each column contains atomic values, and each row is unique. A table of students with multiple phone numbers in a single column is split into multiple rows where each phone number is stored separately.

2nf If a table has both OrderID and ProductID as the primary key, and ProductName only depends on ProductID, then ProductName should be moved to a separate table. This way, every piece of data depends on the full primary key, not just one part.

Third Normal Form (3NF): Achieved when a table is in 2NF and all attributes are only dependent on the primary key, eliminating transitive dependencies. In a Student table, if the student's department name depends on the department ID, we move department info to a separate table to avoid redundancy.


