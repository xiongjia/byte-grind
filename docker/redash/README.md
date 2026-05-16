# Redash Local Development Environment

Docker Compose setup for local Redash development with an additional PostgreSQL database for scratch work.

## Limitations for Production Use

1. **Credentials**: All passwords are hardcoded in `docker-compose.yml` for local development convenience. For production, move credentials to a `.env` file and update `docker-compose.yml` to reference them via `env_file` or `${VAR}` syntax.

2. **Memory Tuning**: PostgreSQL and Redis use conservative memory settings optimized for local development. Production deployments should adjust `shared_buffers`, `work_mem`, `maxmemory`, etc. based on available system resources.

## Services

| Service | Port | Description |
|---------|------|-------------|
| redash-server | 5501 | Redash web UI |
| redash-postgres | 6433 | PostgreSQL 18 (Redash db + scratch db) |
| redash-redis | 6479 | Redis 8 |
| redash-worker | - | RQ worker for background jobs |

## Databases

### Redash Database
- **Host**: localhost:6433
- **User**: redash
- **Database**: redash

### Scratch Database
Use this for importing local test data (Excel, CSV, etc.).

- **Host**: localhost:6433
- **User**: localdev
- **Password**: localdev123
- **Database**: scratch_db

## Quick Start

```bash
cd docker/redash
docker compose up -d
```

Wait ~10 seconds for initialization, then access Redash at http://localhost:5501

## Importing Test Data (Excel/CSV)

### 1. Connect to scratch_db

```bash
psql -h localhost -p 6433 -U localdev -d scratch_db
```

Password: `localdev123`

### 2. Create a table manually

```sql
CREATE TABLE my_data (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    value NUMERIC
);
```

### 3. Import from Excel/CSV

**Option A: Using psql with COPY**

Export Excel to CSV first, then:

```sql
\copy my_data(name, value) FROM '/path/to/data.csv' WITH (FORMAT csv, HEADER true);
```

**Option B: Using pandas (Python)**

```python
import pandas as pd
import psycopg2

df = pd.read_excel('/path/to/data.xlsx')
conn = psycopg2.connect(
    host='localhost',
    port=6433,
    database='scratch_db',
    user='localdev',
    password='localdev123'
)
df.to_sql('my_data', conn, if_exists='replace', index=False)
```

**Option C: Using pgAdmin or DBeaver**

Connect to `localhost:6433`, database `scratch_db`, and use the import wizard.

## Troubleshooting

### Reset scratch database
```bash
docker compose exec redash-postgres psql -U redash -c "DROP DATABASE scratch_db;"
docker compose exec redash-postgres psql -U redash -c "CREATE DATABASE scratch_db OWNER localdev;"
```

### Reset all data (including Redash data)
```bash
docker compose down -v
docker compose up -d
```

