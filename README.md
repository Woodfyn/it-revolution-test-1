# it-revolution-test-1

## Endpoints
All bodies must be with header `Content-Type: application/json`

### `POST /api/transform` 200
#### Body:
```json
{
    "original_link": "string"
}
```
#### Result:
```json
{
    "short link"
}
```

### `GET /api/original/:id` 200
### Result:
```json
"original link"
```

### `GET /api/statistics` 200
### Result:
```json
[
    {
        "created_at": "string",
        "count": "int"
    },
]
```

### `GET /api/statistics/:id` 200
### Result:
```json
{
    "created_at": "string",
    "count": "int"
}
```
