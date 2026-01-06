# Employee Backend Service (Golang + Gin)

Backend service ini dibuat menggunakan **Golang**, **Gin Framework**, dan **PostgreSQL**, dengan pendekatan **Clean Architecture**. Project ini mendukung migrasi database, seeding data, serta beberapa endpoint query berbasis requirement bisnis.

---

## ✨ Fitur Utama

- REST API menggunakan Gin
- Clean Architecture (handler, service, repository)
- PostgreSQL sebagai database
- Database migration menggunakan `golang-migrate`
- Seeder untuk dummy data
- Export response JSON ke file `.txt`
- Download & read file JSON
- Docker & Docker Compose ready

---

## 🗄️ Struktur Database

### Entity Relationship Diagram (ERD)

```mermaid
erDiagram
    EMPLOYEES {
        BIGINT id PK
        VARCHAR first_name
        VARCHAR last_name
        DATE hire_date
        DATE termination_date
        INTEGER salary
        TIMESTAMPTZ created_at
        TIMESTAMPTZ updated_at
    }

    ANNUAL_REVIEWS {
        BIGINT id PK
        BIGINT employee_id FK
        DATE review_date
        TIMESTAMPTZ created_at
        TIMESTAMPTZ updated_at
    }

    EMPLOYEES ||--o{ ANNUAL_REVIEWS : has
```

### Employees

```sql
CREATE TABLE employees (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    hire_date DATE NOT NULL,
    termination_date DATE,
    salary INTEGER NOT NULL CHECK (salary > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

### Annual Reviews

```sql
CREATE TABLE annual_reviews (
    id BIGSERIAL PRIMARY KEY,
    employee_id BIGINT NOT NULL,
    review_date DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## ⚙️ Environment Variable

```env
APP_PORT=<your_app_port>
DATABASE_URL=postgres://postgres:<password>@<host>:<port>/<db_name>?sslmode=disable
```

---

## 🐳 Menjalankan dengan Docker

### 1️⃣ Build & Run

```bash
docker compose up --build
```

### 2️⃣ Jalankan Migrasi

```bash
docker compose exec backend migrate -path ./migrations -database "postgres://postgres:postgres@postgres:5432/employee_db?sslmode=disable" up
```

### 3️⃣ Jalankan Seeder

```bash
docker compose exec backend ./seeder
```

---

## 🔌 Daftar Endpoint

### 1️⃣ Karyawan Aktif dengan Nama Belakang "Smith"

```http
GET /api/v1/employees/active-smith
```

### 2️⃣ Karyawan Tanpa Review

```http
GET /api/v1/employees/no-reviews
```

### 3️⃣ Selisih Hari Hire Date Awal & Akhir

```http
GET /api/v1/employees/hire-date-diff
```

### 4️⃣ Estimasi Gaji Tahun 2016 + Total Review

```http
GET /api/v1/employees/salary-estimation
```

---

## 📁 Export & Read File JSON

### Simpan Response ke File TXT

```http
POST /api/v1/employees/save/contoh2
POST /api/v1/employees/save/contoh3
POST /api/v1/employees/save/contoh4
POST /api/v1/employees/save/contoh5
```

File akan tersimpan di:

```
storage/reports/contoh2.txt
storage/reports/contoh3.txt
storage/reports/contoh4.txt
storage/reports/contoh5.txt
```

### Read File

```http
GET /api/v1/employees/read/file
```

---

## 🧪 Reset Project (Fresh Start)

```bash
docker compose down -v
docker compose up --build
```

Lalu jalankan migrasi & seeder ulang.

---

## 🏗️ Teknologi

- Golang 1.22+
- Gin
- PostgreSQL 16
- pgx
- golang-migrate
- Docker & Docker Compose
