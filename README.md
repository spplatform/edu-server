# Education Server

Fast and simple server

## Endpoints

- `api/login`

Request:
```json
{
    "login": "my-login",
    "password": "my-password"
}
```

Response:
```json
{
    "user": {
        "id": 1,
        "name": "my-name", //in development
        "roadmap-ids": [1,2,3]
    },
    "new": true,
    "first-poll": {
        "id": 5,
        "description": "some poll",
        "questions": [
            {
                "id": 11,
                "description": "who are you?",
                "answers": [
                    {
                        "id": 111,
                        "description": "programmer",
                    },
                    {
                        "id": 112,
                        "description": "designer",
                    }                   
                ]
            },
            {
                "id": 12,
                "description": "what is your age?",
                "answers": [
                    {
                        "id": 121,
                        "description": "under 18",
                    },
                    {
                        "id": 122,
                        "description": "over 18",
                    }                   
                ]
            }
        ]
    },
    "second-poll": {...} // same structure as a first-poll
}
```

- `api/user/{id}`

Request:

id: numerical user id

Response:
```json
{
    "id": 1,
    "name": "my-name", //in development
    "roadmap-ids": [1,2,3]
}
```

- `api/user/{id}/process-poll`

Request:

id: numerical user id

```json
{
    "answers-first": {
        "1":[1],
        "2":[3]
    }, 
    "answers-second": {
        "4":[2],
        "5":[1]
    }
}
```

Response:
```json
{
    "roadmap-id": 22
}
```

