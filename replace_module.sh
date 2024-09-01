#!/bin/bash

# Новый модуль, переданный как аргумент
NEW_MODULE=$1

# Старый модуль (из go.mod вашего шаблона)
OLD_MODULE="github.com/mixdjoker/server-template"

# Замена всех упоминаний старого модуля на новый
find . -type f \( -name "*.go" -o -name "go.mod" \) -exec sed -i "s|$OLD_MODULE|$NEW_MODULE|g" {} +

# Обновляем go.mod
go mod tidy
