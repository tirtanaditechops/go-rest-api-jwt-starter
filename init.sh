#!/bin/bash

echo "📁 Membuat struktur direktori loketnadi-be-go..."

mkdir -p cmd/api
mkdir -p internal/{db,handler,model,service}

touch .env
touch go.mod
touch go.sum

touch cmd/api/main.go
touch internal/db/mssql.go
touch internal/handler/customer.go
touch internal/model/customer.go
touch internal/service/customer_service.go

echo "✅ Struktur proyek selesai dibuat."
