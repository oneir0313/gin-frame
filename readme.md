# Gin Frame

Architecture template with gin framework, cronjob

## Environment Variables

- Run directly

    `configs/config.yml`
- Run with docker-compose (Some values are replaced at `lib/config/config.go`)
    
    `.env` 


</details>

## Running the project 

- Make sure you have docker installed and mysql is running. 
- Configure `configs/config.yml` mysql host
- Run `docker-compose up -d` 
- Check `localhost:9220`. 

## JWT API Demo

### login request
```bash
curl --location --request POST 'http://localhost:9220/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account": "admin",
    "password": "admin"
}'
```
### login response example
```json
{
    "code": 200,
    "expire": "2022-11-11T14:18:54+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjgxNDc1MzQsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2ODE0MzkzNH0.ZpjQ55Zf9eGKgv_EzlTQDaiK6MlbYjrmHzY3mRi08N8"
}
```




### GET Picture API with Authorization
use `Authorization: Bearer ${token}`
```bash
curl --location --request GET 'http://localhost:9220/auth/picture' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjgxNDc1MzQsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTY2ODE0MzkzNH0.ZpjQ55Zf9eGKgv_EzlTQDaiK6MlbYjrmHzY3mRi08N8' \
--header 'Content-Type: application/json'
```
### Picture API Response
![](https://www.taiwan.net.tw/userfiles/image/Wallpaper_en/1920x1080_03.jpg)
