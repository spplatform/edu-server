# Education Server

Fast and simple server

## Constants

### Statuses

- 0 not started
- 1 in progress
- 2 finished
- 3 failed

## Endpoints

- [POST: `api/login`](#api/login)
- [GET: `api/user/{id}`](#api/user/{id})
- [POST: `api/user/{id}/process-poll`](#api/user/{id}/process-poll)
- [GET: `api/roadmap/{id}`](#api/roadmap/{id})
- [GET: `api/badge/{id}`](#api/badge/{id})
- [POST: `api/badge`](#api/badge)
- [GET: `api/certificate/{id}`](#api/certificate/{id})
- [POST: `api/certificate`](#api/certificate)

Request:
```json
{
    "login":    "my-login",
    "password": "my-password"
}
```

Response:
```json
{
    "user": {
        "id": 1,
        "name": "my-name", //in development
        "roadmap-ids": [1,2,3],
        "badge-ids": [4,5,6],
        "certificate-ids": [7,8,9]
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

### api/user/{id}

Request:

id: numerical user id

Response:
```json
{
    "id": 1,
    "name": "my-name", //in development
    "roadmap-ids": [1,2,3],
	"badge-ids": [4,5,6],
	"certificate-ids": [7,8,9]
}
```

### api/user/{id}/process-poll

Request:

id: numerical user id

```json
{
    "answers-first": {
        "1":[1],
        "2":[2]
    }, 
    "answers-second": {
        "1":[2,3]
    }
}
```

Response:
```json
{
    "roadmap-id": 22
}
```

### api/roadmap/{id}

Request:

id: numerical roadmap id

Response:
```json
{
    "id": 1234,
    "description": "Ivan's Developer Roadmap",
    "status": 1,
    "milestones-main": [
        {
            "id": 324,
            "description": "How to start",
            "course-link": "https:example.com",
            "status": 2,
            "steps": [
                {
                    "id": 12343,
                    "description": "How to start starting",
                    "step-link": "https:example.com/12312",
                    "status": 2
                },
                {
                    "id": 32423,
                    "description": "How to end starting",
                    "step-link": "https:example.com/8798",
                    "status": 2
                }
            ]
        },
        {
            "id": 678,
            "description": "How to finish",
            "course-link": "https:example.com",
            "status": 2,
            "steps": [
                {
                    "id": 2342,
                    "description": "How to start finishing",
                    "step-link": "https:example.com/78987",
                    "status": 2
                },
                {
                    "id": 43567,
                    "description": "How to end finishing",
                    "step-link": "https:example.com/76832",
                    "status": 0
                }
            ]
        }
    ],
    "milestones-other": //same structure as milestones-main
}
```

### api/badge/{id}

Request:

id: numerical roadmap id

Response:
```json
{
	"id": 1,
	"description": "Курс 'Основы программирования' пройден",
	"issue-date-time": "2019-06-09T15:27:43.63046+03:00"
}
```

### api/badge

Request:
```json
{
	"user-id": 1,
	"roadmap-id": 3,
	"milestone-id": 7
}
```

Response:
```json
{
	"id": 1,
	"description": "Курс 'Основы программирования' пройден",
	"issue-date-time": "2019-06-09T15:27:43.63046+03:00"
}
```

### api/certificate/{id}

Request:

id: numerical roadmap id

Response:
```json
{
	"id": 1,
	"description": "Специализация 'Путь программиста' получена",
	"issue-date-time": "2019-06-09T15:30:27.323818+03:00"
}
```

### api/certificate

Request:
```json
{
	"user-id": 1,
	"roadmap-id": 3
}
```

Response:
```json
{
	"id": 1,
	"description": "Специализация 'Путь программиста' получена",
	"issue-date-time": "2019-06-09T15:30:27.323818+03:00"
}
```