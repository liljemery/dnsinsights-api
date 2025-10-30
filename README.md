# DNS Insight API

API open-source en Go para diagnóstico de red y resolución de DNS/IP.

- Framework: Fiber
- ORM: GORM (PostgreSQL)
- Config: godotenv
- Logs: zerolog
- Docs: Swagger (`/docs`)

## Arquitectura

```
/cmd/dnsinsight-api/
/configs/
/internal/
  /adapters/http/
  /core/
  /middlewares/
  /modules/
    /dnscheck/
    /ipchecker/
    /network/
    /common/
  /repositories/
  /services/
/pkg/
/routes/
/docs/
/database/
/scripts/
/tests/
```

## Ejecutar

```bash
make deps      # instala swag
make tidy
make run
```

Docker:

```bash
docker-compose up --build
```

## Swagger

- Generar: `make docs`
- URL: `/docs`

## Endpoints

- GET `/api/public/dns/resolve?domain=example.com`
- GET `/api/public/ip/lookup?ip=8.8.8.8`
- GET `/api/public/ip/reverse?ip=8.8.8.8`
- GET `/api/public/network/ping?host=example.com&count=4`
- GET `/api/public/network/ttl?domain=example.com`

## Respuesta estándar

```json
{
  "status": "success",
  "message": "DNS records found",
  "data": { }
}
```

## Extender

- Agrega un servicio en `internal/services/<modulo>`
- Crea handlers HTTP en `internal/adapters/http/<modulo>`
- Cablea rutas en `routes/router.go`
- Documenta con comentarios Swagger sobre handlers

## Licencia

MIT
