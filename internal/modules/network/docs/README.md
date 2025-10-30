# Módulo: network

## Propósito
Ejecutar comprobaciones de red (ping ICMP) y otras utilidades.

## Endpoints
- GET `/api/public/network/ping?host=example.com&count=4`

## Ejemplos
```bash
curl "http://localhost:8080/api/public/network/ping?host=1.1.1.1&count=3"
```

## Dependencias
- `github.com/go-ping/ping`

## Flujo
Handler HTTP -> Servicio `internal/services/network` -> Respuesta estándar.
