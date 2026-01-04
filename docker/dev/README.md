[#](#) Local Development Containers

## Services

- PostgreSQL: localhost:5432

## Start Services
```bash
docker-compose up -d
```

## Stop Serivces

```bash
docker-compose down
```

## Restart Data

```bash
docker-compose down -v
rm -rf volumes/pg
```

## ENV

POSTGRES_USER=dev
POSTGRES_PASSWORD=dev
POSTGRES_DB=devdb

