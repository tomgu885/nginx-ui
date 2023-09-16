学习


https://nginxui.com/

## 开发步骤

```bash
$cd frontend
$yarn install
$yarn build
$cd ..
$go run main.go
```

## 启动

1. gocron [cert.AutoObtain]
1. http.Server( gin)

## 用户认证

jwt: Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4iLCJleHAiOjE2OTQ2NDY5NDd9.hTOOTQjjWTMzXCDAEV34v9OJkRRG7wOiggEBkMBqFOM

## 接口/路由

1. GET / 65
2. POST / 29
3. DELETE / 11

## 数据库-表

sqlite3

```text
auth_tokens      certs            config_backups   environments
auths            chat_gpt_logs    dns_credentials  sites
```

```sqlite
-- auths
CREATE TABLE `auths` (
    `id` integer,
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime,`name` text,
    `password` text,PRIMARY KEY (`id`));
CREATE INDEX `idx_auths_deleted_at` ON `auths`(`deleted_at`);
-- tokens
CREATE TABLE `auth_tokens` (`token` text);


CREATE TABLE `config_backups` (
    `id` integer,
    `created_at` datetime,
    `updated_at` datetime,`deleted_at` datetime,
    `name` text,`file_path` text,`content` text,PRIMARY KEY (`id`));
CREATE INDEX `idx_config_backups_deleted_at` ON `config_backups`(`deleted_at`);


CREATE TABLE `certs` (
    `id` integer,
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime,
    `name` text,`domains` text[],
    `filename` text,
    `ssl_certificate_path` text,
    `ssl_certificate_key_path` text,
    `auto_cert` integer,`challenge_method` text,
    `dns_credential_id` integer,`log` text,
    PRIMARY KEY (`id`)
);
CREATE INDEX `idx_certs_deleted_at` ON `certs`(`deleted_at`);
```

## go:embed

1. https://colobu.com/2021/01/17/go-embed-tutorial/
2. https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/

go 1.6 开始


## sqlite

1. https://sqlite.org/cli.html
