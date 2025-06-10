## Instalar migrations CLI

```bash
wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb
sudo dpkg -i migrate.linux-amd64.deb
```

## Criar migrations

```bash
migrate create -ext=sql -dir=internal/database/migrations -seq {NOME_DA_MIGRAION}
```
