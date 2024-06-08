# Employee Management Backend
Employee management backend

## Enhancements that can be done
1. Redis can be implimented to store the fetched data for a short term so that database hits can be minimized and perfomance can be increased.
2. database indexing can be done to increase datbase perfomance.
3. for creating logs sql triggers can be created 
4. create delete log's who is deleting the accounts 
5. create update logs 
6. flag can be used for showing the returning data where needed will enable the flag so that tha data is retuned.


Note:Tables Get auto Created when project is started
need to add the resource rows.
## Run the project in local 
cd cmd
go run .

## To generate new errors using stringer
cd er
go generate


## Employee Registration 
### POST http://localhost:8765/v1/api/employee
#### Payload
```
{
    "name":"Abhijeet",
    "mobile":"9876543210",
    "position":"Software Engineer",
    "salary":3500000.00
}
```
### Response
```
{
    "success": true,
    "message": "Registration Sucessfully Done",
    "data": {
        "id": 24,
        "name": "Abhijeet",
        "mobile": "9876543210",
        "position": "Software Engineer",
        "salary": 3500000,
        "is_active": true,
        "created_at": "2024-06-09T04:03:36.741183+05:30",
        "updated_at": "2024-06-09T04:03:36.741185+05:30",
    }
}
```
## Error Validations

#### Payload 
```
{
    "name":"Abhijeet",
    "position":"Software Engineer",
    "salary":3500000
}
```
#### Response
```
{
    "code": 4,
    "exception": "employee_backend.InvalidRequestBody",
    "message": "Invalid request body. Mobile is required and cannot be empty"
}
```
#### Payload
```
{
    "name":"Abhijeet",
    "position":"Software Engineer",
    "salary":"3500000"
}
```

#### Response
```
{
    "code": 4,
    "exception": "employee_backend.InvalidRequestBody",
    "message": "Invalid type for field 'salary'. Please check the input value."
}
```

## Get Employee by id
### GET http://localhost:8765/v1/api/employee/1

### Response
```
{
    "success": true,
    "message": "Fetched Sucessfully",
    "data": {
        "id": 24,
        "name": "Abhijeet",
        "mobile": "9876543210",
        "position": "Software Engineer",
        "salary": 3500000,
        "is_active": true,
        "created_at": "2024-06-09T04:03:36.741183+05:30",
        "updated_at": "2024-06-09T04:03:36.741185+05:30",
    }
}
```
#### NO DATA FOUND
```
{
    "success": true,
    "message": "No data found",
    "data": null,
    }
}
```

## Get All Employees with Filter options
### GET http://localhost:8765/v1/api/employee 

#### payload
Note: use limit -1 to display a; data without pagination else default value is set to 20 can be changed by passing value as Limit = 30 
##### FORM-data 
limit:-1
id:1
name:john
position:developer
salary:35000

#### Response
```
{
    "success": true,
    "message": "Fetched Sucessfully",
    "data": [
        {
            "id": 24,
            "name": "Abhijeet",
            "mobile": "9876543210",
            "position": "Software Engineer",
            "salary": 3500000,
            "is_active": true,
            "created_at": "2024-06-09T04:03:36.741183+05:30",
            "updated_at": "2024-06-09T04:03:36.741185+05:30",
            "Mutex": null
        },
        {
            "id": 20,
            "name": "John Doe",
            "mobile": "1234567890",
            "position": "Developer",
            "salary": 50000,
            "is_active": true,
            "created_at": "2024-06-09T03:52:30.344514+05:30",
            "updated_at": "2024-06-09T03:52:30.344515+05:30",
        }
    ],
    "meta": {
        "current_page": 1,
        "total_pages": 1,
        "total_data_count": 2
    }
}
```

#### NO DATA FOUND

```
{
    "success": true,
    "message": "No data found",
    "data": []
}
```

## Update  Employees By ID
### PUT http://localhost:8765/v1/api/employee/24

#### Payload 
```
{
    "salary":6800000.00
}
```
#### Response 
```
{
    "success": true,
    "message": "Updated Sucessfully",
    "data": {
        "id": 24,
        "name": "Abhijeet",
        "mobile": "9876543210",
        "position": "Software Engineer",
        "salary": 6800000,
        "is_active": true,
        "created_at": "2024-06-09T04:03:36.741183+05:30",
        "updated_at": "2024-06-09T04:13:38.045238+05:30",
    }
}

```

## Delete  Employees By ID
### DELETE http://localhost:8765/v1/api/employee/24

#### Response

```
{
    "success": true,
    "message": "Sucessfully Deleted",
    "data": {
        "id": 24,
        "name": "Abhijeet",
        "mobile": "9876543210",
        "position": "Software Engineer",
        "salary": 6800000,
        "is_active": false,
        "created_at": "2024-06-09T04:03:36.741183+05:30",
        "updated_at": "2024-06-09T04:13:38.045238+05:30",
    }
}
```




