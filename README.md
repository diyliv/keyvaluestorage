# Key:Value storage

## Add

| Property      | Value         |
| ------------- | ------------- |
| API           | /api/add      |
| Method        | 'POST'        |
| Description   | This endpoint is responsible for adding new values |

Input:
```
    key: user's key
    value: user's value 
    exp_time: how long will the data be stored
```
Sample request:
```json
{
    "key":"hello",
    "value":"world"
}
```

Sample response: 
```json
{
    "key": "hello",
    "value": "world",
    "added_status": true,
    "added_time": "2022-09-03 18:27:55.8310409 +0300 MSK",
    "deleted_time": "",
    "expiration_info": {
        "exp_key": null,
        "key_value": null,
        "exp_time": 0,
        "deleted_time": ""
    }
}
```

## Get 
| Property      | Value         |
| ------------- | ------------- |
| API           | /api/get      |
| Method        | 'POST'        |
| Description   | This endpoint is responsible for getting stored values |

Input:
```
    key: user's key
```

Sample request:
```json
{
    "key":"hello"
}
```

Sample response:

```json
    {
    "key": "hello",
    "value": "world",
    "added_status": true,
    "added_time": "2022-09-03 18:27:55.8310409 +0300 MSK",
    "deleted_time": "",
    "expiration_info": {
        "exp_key": null,
        "key_value": null,
        "exp_time": 0,
        "deleted_time": ""
    }
}
```

## Delete

| Property      | Value         |
| ------------- | ------------- |
| API           | /api/delete   |
| Method        | 'POST'        |
| Description   | This endpoint is responsible for deleting stored values |


Input:
```
    key: user's key
```

Sample request:

```json
{
    "key":"hello"
}
```

Sample response:

```json
{
    "key": "hello",
    "value": "world",
    "added_status": false,
    "added_time": "2022-09-03 18:27:55.8310409 +0300 MSK",
    "deleted_time": "2022-09-03 18:34:48.6261362 +0300 MSK",
    "expiration_info": {
        "exp_key": null,
        "key_value": null,
        "exp_time": 0,
        "deleted_time": ""
    }
}
```