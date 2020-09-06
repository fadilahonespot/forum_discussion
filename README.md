# Forum Diskusi
Ini adalah projek untuk jawaban soal nomer 1, menggunakan golang untuk membuat Rest API dan MYSQL sebagai databasenya

## Skema Database
![SkemaDatabase](https://github.com/fadilahonespot/forum_discussion/blob/master/skema_kaskus.png)

## Library
- Gin (router)
- JWT (security)
- GORM (ORM Database)
- Viper (environment variabel)
- Crypto (hashing)
- Gomod (depedensi)

## Setting
lakukan penyesuaian setting database pada bagian database dan setting destinasi file pada bagian asset di 
```destinasi
kaskus/config/config.json
```
lalu run program
```sh
go run main.go
```

Maka otomatis akan tergenerate table pada database
- lakukan registrasi
- lalu setelah berhasil registrasi lakukan login, maka bila berhasil akan mendapatkan kode token
- Maka masukan token tersebut pada bagian authorrization dan pilih bearer token untuk mengakses endpoin-endpoin lainnya
![inputToken](https://github.com/fadilahonespot/forum_discussion/blob/master/authorization.PNG)
- Untuk membuat catagory dan menghapus catagory hanya admin yang bisa melakukannya. Bila ada user mendaftar secara default tingkatannya adalah user. Untuk mengubahnya menjadi tingkatan admin bisa di ubah lewat databasenya langsung pada field role di table user.
- Untuk keperluan testing saya sudah sertakan file json yang berisi setting endpoin-endpoin untuk di import ke postman bersamaan dengan file ini.

## Endpoin

Not Authorization
```endpoint
localhost:7788/register #Post
localhost:7788/login    #Post
```
User Authorization
```endpoint
localhost:7788/user                   #Get All user
localhost:7788/user/profile           #Get profile user
localhost:7788/catagory               #Get all catagory
localhost:7788/discussion             #Post Discussion
localhost:7788/discussion/answerf/1   #Post Discussion untuk membalas diskusi tingkat pertama, angka 1 adalah id diskusi
localhost:7788/discussion/answers/1   #Post Discussion untuk membalas diskusi tingkat kedua, angka 1 adalah id diskusi tingkat pertama
localhost:7788/discussion             #Get all discussion
localhost:7788/discussion/1           #Get detail discussion, angka satu adalah id diskusi
localhost:7788/discussion/1           #Put discussion, angka 1 adalah id diskusi
localhost:7788/discussion/1           #Delete discussion, angka 1 adalah id diskusi 
```
Admin Authorization
```endpoint
localhost:7788/catagory     #Post catagory, menambahkan katagory
localhost:7788/catagory/4   #Delete catagory, angka 1 adalah id catagory
```
