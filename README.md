# VKTEST: REST API for VK internship
---

<img src='https://fanibani.ru/wp-content/uploads/2021/07/milie004.jpg' width='350' height='300'>

## Description
This is a REST API application that manages a database and simulates a movie library.

## Features
- RESTful API endpoints for basic CRUD operations.
- Docker Compose configuration for easy deployment.
- Environment variable setup for configuring database connection.

## Prerequisites
- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Docker Compose: [Installation Guide](https://docs.docker.com/compose/install/)

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/dkshi/vktest.git
    ```

2. Navigate to the project directory:
    ```bash
    cd vktest
    ```

3. Create a `.env` file in the project root with the following content:
    ```plaintext
    DB_NAME=postgres
    DB_HOST=db
    DB_PORT=5432
    DB_USERNAME=postgres
    DB_PASSWORD=qwerty
    DB_SSLMODE=disable
    ```

## Usage
1. Start the application using Docker Compose:
    ```bash
    docker-compose up
    ```

2. Access Swagger at `http://localhost:8080/swagger/index.html#`.

## Configuration
- To change database connection settings, modify the values in the `.env` file.
- Additional environment variables can be added to customize other aspects of the application.

