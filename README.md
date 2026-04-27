# 🏢 HR Go API

> Project belajar pertama menggunakan **Golang + Gin Framework** — sebuah REST API untuk sistem Human Resources sederhana.

## Tech Stack

| Teknologi | Keterangan |
|-----------|------------|
| **Go 1.25** | Bahasa utama |
| **Gin** | HTTP web framework |
| **GORM** | ORM untuk database |
| **PostgreSQL** | Database |
| **JWT** | Autentikasi token |
| **Swagger** | API documentation |

## Fitur

- 🔐 **Auth** — Register, Login, Get Profile (`/api/auth`)
- 🏗️ **Division** — CRUD divisi/departemen (`/api/division`)
- 👤 **Employee** — CRUD data karyawan (`/api/employee`)
- 📄 **Pagination & Search** — Semua list endpoint mendukung pagination dan pencarian
- 🔒 **JWT Middleware** — Proteksi endpoint dengan token
- 🌐 **CORS** — Siap diakses dari frontend
- 📖 **Swagger** — Dokumentasi API interaktif (`/swagger/index.html`)

## Struktur Project

```
hr-go-api/
├── cmd/api/main.go              # Entry point
├── docs/                        # Swagger generated docs
├── internal/
│   ├── config/                  # Konfigurasi (.env)
│   ├── database/                # Koneksi, migrasi, seeder
│   ├── dto/                     # Request & Response struct
│   ├── handler/                 # Controller / handler
│   ├── middleware/               # Auth, CORS, Logger
│   ├── models/                  # Model database (GORM)
│   ├── repository/              # Data access layer
│   └── service/                 # Business logic
├── mocks/                       # Mock untuk unit testing
├── pkg/
│   ├── hash/                    # Bcrypt helper
│   ├── jwt/                     # JWT generate & validate
│   ├── response/                # Standard API response
│   └── types/                   # Custom types (Date)
└── router/                      # Route definitions
```

## Cara Menjalankan

### 1. Clone repo

```bash
git clone https://github.com/rheadavin/hr-go-api.git
cd hr-go-api
```

### 2. Setup environment

Buat file `.env` di root project:

```env
APP_NAME=HRGoApi
APP_ENV=development
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=human_resources

JWT_SECRET=your_secret_key
JWT_EXPIRE_HOURS=1
```

### 3. Buat database PostgreSQL

```bash
createdb human_resources
```

### 4. Generate Swagger docs

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/api/main.go
```

### 5. Install dependencies & jalankan

```bash
go mod tidy
go run cmd/api/main.go
```

Server akan berjalan di `http://localhost:8080`. Migrasi dan seed data otomatis dijalankan saat start.

## 📖 Swagger Documentation

Akses Swagger UI di browser (hanya tersedia di mode non-production):

```
http://localhost:8080/swagger/index.html
```

Untuk regenerate docs setelah mengubah annotation:

```bash
swag init -g cmd/api/main.go
```

## API Endpoints

### Auth (Public)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/auth/register` | Register user baru |
| POST | `/api/auth/login` | Login, dapat JWT token |

### Auth (Protected 🔒)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/me` | Get profile user yang login |

### Division (Protected 🔒)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/division/` | List semua divisi (pagination) |
| GET | `/api/division/:id` | Get divisi by ID |
| POST | `/api/division/create` | Buat divisi baru |
| PUT | `/api/division/:id` | Update divisi |
| DELETE | `/api/division/:id` | Hapus divisi (soft delete) |

### Employee (Protected 🔒)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/employee/` | List semua karyawan (pagination) |
| GET | `/api/employee/:id` | Get karyawan by ID |
| POST | `/api/employee/create` | Tambah karyawan baru |
| PUT | `/api/employee/:id` | Update karyawan |
| DELETE | `/api/employee/:id` | Hapus karyawan (soft delete) |

## Health Check

```bash
curl http://localhost:8080/api/health
```

---

*Built with ☕ sambil belajar Golang*
