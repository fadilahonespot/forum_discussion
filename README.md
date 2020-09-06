# Untuk Jawaban Soal Nomer 1

## Library
- Gin (router)
- JWT (security)
- GORM (ORM Database)
- Viper (envirotment variabel)
- Crypto (hashing)
- Gomod (depedensi)

## Setting
lakukan penyesuaian setting database dan sesuikan destinasi folder pada komputer anda pada folder kaskus/config/config.json
lalu run program
```sh
go run main.go
```

Maka otomatis akan tergenerat table pada database
- lakukan registrasi
- lalu setelah berhasil registrasi lakukan login, maka bila berhasil akan mendapatkan kode token
- Maka masukan token tersebut pada bagian authorizer dan pilih bearer token untuk mengakses end poin- end poin lainnya
- Untuk membuat catagory dan menghapus catagory hanya admin yang bisa melakukannya. Bila ada user mendaftar secara default tingkatannya adalah user. Untuk mengubahnya menjadi admin bisa di lakukan lewat databasenya langsung pada table user.
- Untuk keperluan testing saya sudah sediakan file json yang berisi end poin - end poin untuk di import ke postman bersamaan dengan file ini.
