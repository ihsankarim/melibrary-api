# 📚 Melibrary API

Melibrary API adalah RESTful API backend untuk aplikasi **perpustakaan digital**. Dibangun dengan **Golang + Fiber**, menggunakan **PostgreSQL** sebagai database, dan mendukung dokumentasi API interaktif via Swagger.

---

## 🚀 Tech Stack

- [Golang](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [GORM](https://gorm.io/) (ORM)
- [PostgreSQL](https://www.postgresql.org/)
- [JWT](https://jwt.io/)
- [Swaggo](https://github.com/swaggo/swag) (Swagger auto-docs)

---

## 📦 Fitur Utama

- ✅ Autentikasi `Register` & `Login` dengan JWT
- 📚 CRUD Buku
- 🗂️ Manajemen Kategori Buku
- 🔐 Middleware proteksi JWT
- 📄 Swagger docs (`/swagger/index.html`)

---

## ⚙️ Requirements

- Go 1.18+
- PostgreSQL 12+
- [swag CLI](https://github.com/swaggo/swag) untuk generate dokumentasi API:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
