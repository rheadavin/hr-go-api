# 🏢 HR Go API

> Project belajar pertama gue pakai **Golang + Gin Framework** — sebuah REST API untuk sistem Human Resources sederhana.

## Tech Stack

| Teknologi | Keterangan |
|-----------|------------|
| **Go 1.25** | Bahasa utama |
| **Gin** | HTTP web framework |
| **GORM** | ORM untuk database |
| **PostgreSQL** | Database |
| **JWT** | Autentikasi token |

## Fitur

- 🔐 **Auth** — Register, Login, Get Profile (`/api/auth`)
- 🏗️ **Division** — CRUD divisi/departemen (`/api/division`)
- 👤 **Employee** — CRUD data karyawan (`/api/employee`)
- 📄 **Pagination & Search** — Semua list endpoint mendukung pagination dan pencarian
- 🔒 **JWT Middleware** — Proteksi endpoint dengan token
- 🌐 **CORS** — Siap diakses dari frontend

## Struktur Project

```
hr-go-api/
├── cmd/api/main.go              # Entry point
├── internal/
│   ├── config/                  # Konfigurasi (.env)
│   ├── database/                # Koneksi, migrasi, seeder
│   ├── dto/                     # Request & Response struct
│   ├── handler/                 # Controller / handler
│   ├── middleware/               # Auth, CORS, Logger
│   ├── models/                  # Model database (GORM)
│   ├── repository/              # Data access layer
│   └── service/                 # Business logic
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

### 4. Install dependencies & jalankan

```bash
go mod tidy
go run cmd/api/main.go
```

Server akan berjalan di `http://localhost:8080`. Migrasi dan seed data otomatis dijalankan saat start.

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

## Contoh Request

### Register

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name": "Rhea Davin", "email": "rheadavin@yopmail.com", "password": "password123"}'
```

### Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "rheadavin@yopmail.com", "password": "password123"}'
```

### Create Employee (pakai token)

```bash
curl -X POST http://localhost:8080/api/employee/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{
    "nik": "EMP001",
    "full_name": "John Doe",
    "email": "john@example.com",
    "phone": "08123456789",
    "position": "Backend Developer",
    "salary": 15000000,
    "join_date": "2026-01-15",
    "division_id": 1
  }'
```

### List Employee (pagination & search)

```bash
curl -X POST http://localhost:8080/api/employee/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{"page": 1, "limit": 10}'
```

Dengan search:

```bash
curl -X POST http://localhost:8080/api/employee/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{"page": 1, "limit": 10, "search": "John"}'
```

## Health Check

```bash
curl http://localhost:8080/api/health
```

---

*Built with ☕ sambil belajar Golang*
