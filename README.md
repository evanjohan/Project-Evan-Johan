# Project-Evan-Johan
Project penerapan:
- RabbitMQ (Messages broker)
- Connection database
- Rest API
# Producer 
COMMAND : go run main.go  <br/>
POST : /news ( kalau di local :localhost:8080/News)<br/>
JSON BODY example : {<br/>
	"author" : "Berita IT",<br/>
	"body" : "bahwa IT sekarang sangat berkembang pesat"<br/>
} <br/>
GET : /news ( kalau di local :localhost:8080/News) karena pakai pagination, sehingga urlnya :localhost:8080/News?page=1<br/>
kita pilih page berapa yang mau kita tampilkan, contoh response:<br/>
>>{<br/>
  >>  "status": 1,<br/>
  >> "message": "Success Get Data News",<br/>
  >> "totalData": 60,<br/>
  >>  "totalPage": 6,<br/>
  >>  "Data": [
  >     {
  >          "id": "10",
  >          "author": "Test",
  >          "body": "Aja"
        }]<br/>
}<br/>

data akan selalu masuk, dan tersimpan di queue, disini kita menggunakan Message broker Rabbit MQ<br/>

# Rabbit MQ info:
- Queue diberi nama : news
- Exchanges diberi nama : notifExchange
- type : headers,
- durable : true

# consumer
Command : go run notification.go (file ini ada  di folder main/consumer) bisa dijalankan secara bersamaan dengan berbeda terminal
