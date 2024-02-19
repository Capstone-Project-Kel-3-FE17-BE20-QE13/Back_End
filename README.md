# JobHuntz

JobHuntz is a web application that allows company to post job openings and job seekers to search for and apply to jobs.


# ERD 

![Untitled Diagram drawio](https://github.com/Capstone-Project-Kel-3-FE17-BE20-QE13/Back_End/assets/52233444/a1fe1ec3-776b-4f64-897c-a0fce48ccf8e)


## OPENAPI Swaggerhub
https://app.swaggerhub.com/apis/zidan70/Job-Huntz-Swaggerhub/1.0.0

## Fitur Utama
- Autentikasi
- Manajemen Lowongan
- Manajemen Lamaran
- Proses Penerimaan Pelamar
- Payment Gateway Premium

### Instalasi
1. Clone repositori ini menggunakan Git:
   ```bash
   git clone https://github.com/Capstone-Project-Kel-3-FE17-BE20-QE13/Back_End.git

2. Masuk ke direktori aplikasi:
    ```
    cd Back_End

3. Instal dependensi yang diperlukan:
    ```
    go mod tidy


### Menjalankan Server
Untuk menjalankan server, gunakan perintah:
    ```
    go run main.go

Server akan berjalan di http://localhost:8080 atau port yang telah kamu tentukan.

### ENV

- export DBUSER='......'
- export DBPASS='......'
- export DBHOST='.......'
- export DBPORT='.......'
- export DBNAME='.......'
- export JWTSECRET='...........'

- export CLOUDINARY_KEY='.....'
- export CLOUDINARY_SECRET='......'
- export CLOUDINARY_CLOUD_NAME='....'
- export MIDTRANS_SERVERKEY='......'
