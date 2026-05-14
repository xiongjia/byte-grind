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
| redash-mcp | - | MCP Server for AI integration |

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

## Redash MCP Server (AI Integration)

The `redash-mcp` service exposes a [Model Context Protocol](https://modelcontextprotocol.io/) server, allowing AI assistants to query Redash data sources.

### Get API Key

1. Visit http://localhost:5501 and log in
2. Go to **Settings > API Key** and generate one
3. Set the API key in `.env` file:
   ```bash
   REDASH_API_KEY=your_api_key_here
   ```

### Run with Claude Desktop

Add to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "redash": {
      "command": "npx",
      "args": ["-y", "@suthio/redash-mcp"],
      "env": {
        "REDASH_API_KEY": "your_api_key_here",
        "REDASH_URL": "http://localhost:5501"
      }
    }
  }
}
```

### Available MCP Tools
- `list-queries` / `get-query` — query management
- `execute-query` / `execute-adhoc-query` — run queries
- `list-data-sources` — list available data sources
- `list-dashboards` / `get-dashboard` — dashboard management