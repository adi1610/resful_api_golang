# resful_api_golang

- Clone project :
```bash
git clone https://github.com/adi1610/resful_api_golang.git
```

- go mod terlebih dahulu :
```bash
go mod init resful-api-golang
```

- jalankan perintah dibawah ini untuk install library gin-gonic dan mysql:
```bash
go get -u github.com/gin-gonic/gin
go get -u grom.io/driver/mysql
go get -u gorm.io/gorm
```
- buat file  <b>setup.go</b>  duplicat file  <b>setup.go.example</b>
- buka file  <b>setup.go</b>  Rubah di bagian Connection database (sesuaikan dengan settingan mysql):

```
database, err := gorm.Open(mysql.Open("@user_name:@password@tcp(@host:@port)/@namaproject"))
```

- jangan lupa buat database

## Menjalankan Aplikasi

```bash
go run main.go
```