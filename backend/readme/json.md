* Method: get (all records)
http://localhost:4000/get
data:
{
    "idheader": 1,
    "companyname": "abc1",
    "address": "st angeles",
    "numberinvoice": 12,
    "datetime": "2023-02-12 15:00:00",
    "createdat": "2023-02-12 15:00:00"
}

* Method: get (one)
http://localhost:4000/getOne/2

* Method: post
http://localhost:4000/createOne
data:
{
    "companyname": "test",
    "address": "stqw ole",
    "numberinvoice": 34,
    "datetime": "2023-03-12 15:00:00",
    "createdat": "2023-04-12 15:00:00"
}

* Method: put (one)
http://localhost:4000/updateOne/1
data:
{
    "idheader": 1,
    "companyname": "abc1",
    "address": "st street",
    "numberinvoice": 412,
    "datetime": "2022-02-12 15:00:00",
    "createdat": "2022-02-02 15:00:00"
}

