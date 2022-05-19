# url-shortener

- Small app for study

## Dev

```bash
go run cmd/main.go
```

## Prod

- https://url-shortener.up.railway.app↗

### Create Short URL Path

- `POST /api/v1/shorten`
```json
{
  "orgUrl": "https://www.naver.com"
}
```

### Get Original URL

- `GET /api/v1/:shortUrlPath`
