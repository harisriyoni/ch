# Deployment Aplikasi Frontend

Kita akan membuat frontend yang sudah kita buat berkomunikasi dengan backend yang sudah dibuat sebelumnya. Contoh :
https://gocroot.herokuapp.com/presensi
Adapun langkah yang akan kita lakukan :
1. Membuat tabel di dalam kontainer dari tampilan yang sudah kita buat
2. Menambahkan jscroot di dalam file html yang kita buat
3. Membuat dan mengisi kerangka jscroot

# Kontainer Tabel

Pertama pakai frontend yang sudah kemaren kita buat, deploy di repo kampus. Kita buat di dalam kontainer sebuah tabel untuk mengkonsumsi API.



![image](https://user-images.githubusercontent.com/11188109/214906570-fdf8403b-e82d-41ec-a61a-3a8918eb97fb.png)

go mod init chapter01

go get github.com/gofiber/fiber/v2

github.com/gofiber/websocket/v2

ws://127.0.0.1:3000/ws
