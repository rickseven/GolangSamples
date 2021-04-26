# Description

This is a sample of RESTful API with Golang and MySQL in Docker, in this example I use the Gorm library as an object relational mapping (ORM) and Gin as the web framework. Also, I use phpMyAdmin to manage the database in Docker.

# How To Run

By default, this project created to run in Docker. If you want to run without docker in your local machine directly you need to change database host from `DB_HOST=mysql-service` to `DB_HOST=localhost` in .env file then run this command 
```
go run main.go
```

If you are docker user then run this command
```
docker-compose up --build
```

To try, access this URL [http://localhost:8081/api/v1/products](http://localhost:8081/api/v1/products/).

Then to create the table and insert sample data, run SQL script on `sample_db.sql` file in phpMyAdmin, you can access phpMyAdmin on [http://localhost:9090](localhost:9090).

# API Endpoints

`GET /api/v1/products`

To get all products.

Example response body:

```
[
    {
        "Id": "7b01f186-deae-41a0-90fd-0e67a5958c81",
        "Name": "Laptop Lenovo ThinkPad T470",
        "Quantity": 7,
        "Price": 20000000
    }
]
```

`POST /api/v1/products`

To create a product.

Example request body:

```
{
    "Id": "7bc76459-9eb6-48a1-a68c-8dfa5013bdca",
    "Name": "BARDI Smart Extension Power Strips",
    "Quantity": 10,
    "Price": 300000
}
```
