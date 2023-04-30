# Challenge Backend: Customer Transaction Service

## Case Challenge

Buat relasi dari antara 2 tabel yaitu tabel `Customer` dan tabel `Transaction` dimana
terdapat relasi 1 to many dari `Customer` ke `Transaction`.
Artinya adalah setiap customer dapat melakukan banyak transaksi.

#### Buatlah API untuk:
1. Membuat Transaksi Baru (Post New Transaction)
2. Mendapatkan list transaksi diurutkan dari yang terbaru, dengan pilihan query sebagai berikut:
   - `Query`: untuk melakukan pencarian menu dan harganya pada transaksi
   - `Customer`: diurutkan dari nama customer

**Anda di bebaskan untuk menggunakan REST, gRPC, atau GraphQL.**

#### Hal yang patut anda perhatikan:
1. Bagaimana jika terdapat ribuan transaksi pada database?
2. Bagaimana jika terdapat banyak user yang mengakses API tersebut secara bersamaan?

Beri tahu dan jelaskan juga **Alasan Anda memakai tech stack** yang diambil dan
proses **Pengambilan solusi** ketika terdapat permasalahan dalam pembuatan **API**

## Syarat dan Ketentuan

- Tulis solusi yang Anda buat menggunakan `NodeJS` atau `Golang`
- Anda dapat menggunakan database, search engine atau hal lain yang populer
- Unit tests dan/atau integration test
- Menggunakan `docker-compose` untuk deployment yang lebih mudah
- **Tidak Diperbolehkan** menggunakan ORM (Mongoose, Sequelize, GORM, dll). Gunakan standard SQL