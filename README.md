[![Deploy to GCE](https://github.com/Ulas-Scan/UlaScan_BE/actions/workflows/deploy.yml/badge.svg?branch=main)](https://github.com/Ulas-Scan/UlaScan_BE/actions/workflows/deploy.yml)

# Ulascan: Bangkit 2024 Batch 6 Capstone Project Backend

Welcome to the backend repository for the Bangkit 2024 Batch 6 capstone project! This backend serves the mobile app for our project.

## Features

- Authentication: User authentication and authorization.
- ...
- ...

## Installation

To set up the backend for the capstone project, follow these steps:

1. **Clone the Repository**: 
   ```sh
   git clone https://github.com/javakanaya/ulascan.git
   ```
2. **Set Up PostgreSQL Database**:
    Connect to your PostgreSQL database:
    ```sh
    psql -U <your_user> 
    ```
    Create the database
    ```SQL
    CREATE DATABASE ulasacan";    
    ```
    Run the following SQL command to enable the uuid-ossp extension:
    ```SQL
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    ```

3. **Set Environment Variables**:
    Create a ```.env``` file in the root directory of the project based on the ```.env.example``` file Define the following environment variables in the ```.env``` file:
    ```env
    DB_HOST=<your_database_host>
    DB_USER=<your_database_user>
    DB_PASSWORD=<your_database_password>
    DB_NAME=<your_database_name>
    DB_PORT=<your_database_port>
    ```

4. **Run the Application**:
    ```sh
    go run main.go
    ```
