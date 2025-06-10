## Instalar migrations CLI

```bash
wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb
sudo dpkg -i migrate.linux-amd64.deb
```

## Pegar variáveis de ambiente

```bash
cp .env.example .env
```

## Criar migrations

```bash
make create_migration MIGRATION_NAME={NOME_DESEJADO_DA_MIGRATION (ex:add_user_phone)}
```

## Rodar migrations

```bash
make migrate_up
```

## Voltar migrations

```bash
make migrate_down
```

## Ajuda com migrations

```bash
make help
```
