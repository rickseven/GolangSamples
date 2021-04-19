# Description
This is a sample REST API with golang, in this example I use the Gorm library as an object relational mapping (ORM) and Gin as the web framework.

There are two endpoints:

`GET /api/v1/products`

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

Example request body:

```
{
    "Id": "7bc76459-9eb6-48a1-a68c-8dfa5013bdca",
    "Name": "BARDI Smart Extension Power Strips",
    "Quantity": 10,
    "Price": 300000
}
```
