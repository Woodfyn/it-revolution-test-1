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
"short link"
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
        "short_link": "string",
        "count": "int"
    },
]
```
