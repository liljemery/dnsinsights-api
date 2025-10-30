# Módulo: ipchecker

## Propósito
Información básica de una IP (versión, privada, loopback, etc.).

## Endpoints
- GET `/api/public/ip/lookup?ip=8.8.8.8`

## Ejemplos
```bash
curl "http://localhost:8080/api/public/ip/lookup?ip=8.8.8.8"
```

## Dependencias
- `net` (stdlib)

## Flujo
Handler HTTP -> Servicio `internal/services/ipchecker` -> Respuesta estándar.
