# Módulo: dnscheck

## Propósito
Resolver dominios (A, AAAA, MX, TXT, CNAME), reverse IP y TTL.

## Endpoints
- GET `/api/public/dns/resolve?domain=example.com`
- GET `/api/public/ip/reverse?ip=8.8.8.8`
- GET `/api/public/network/ttl?domain=example.com`

## Ejemplos
```bash
curl "http://localhost:8080/api/public/dns/resolve?domain=example.com"
```

## Dependencias
- `net` (stdlib)
- `github.com/miekg/dns`

## Flujo
Handler HTTP -> Servicio `internal/services/dnscheck` -> Resolver -> Respuesta estándar.
