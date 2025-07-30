> **Task-api**: This API was designed in the intention of handling list of tasks in and out our database, allowing users to create a task and saving into the database 
> of their choice (for now: either in psql or mysql), get all tasks avaliable in the database, get a task using the ID, update contents of a task (whether the
> title, description or both) and delete task from a database by ID.

**Prerequisites** <br>
Before using the API, make sure to either have either mysql or postgres (psql) in your operating system installed. Also make sure to run the schema below either in mysql
or in psql (recommend using shell or bash). <br>
**Schema**: <br>
*For MySQL* - <br>
```sql
CREATE DATABASE IF NOT EXISTS taskdb;
USE taskdb;

CREATE TABLE IF NOT EXISTS tasks (
id CHAR(36) PRIMARY KEY,
title VARCHAR(150) NOT NULL,
description VARCHAR(300) NOT NULL
);
```
*For psql in bash* - 
```sql
CREATE DATABASE IF NOT EXISTS task_api;
\c task_api

CREATE TABLE IF NOT EXISTS tasks (
id CHAR(36) PRIMARY KEY,
title VARCHAR(150) NOT NULL,
description VARCHAR(300) NOT NULL
);
