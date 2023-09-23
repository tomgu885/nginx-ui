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

## ssl 证书路径

ssl_certificate /etc/nginx/ssl/ip-limit.cloudfy669.xyz/fullchain.cer;
ssl_certificate_key /etc/nginx/ssl/ip-limit.cloudfy669.xyz/private.key;

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

## acme / let's encrypt 限制

https://letsencrypt.org/docs/rate-limits/

185.244.208.50 - - [22/Sep/2023:22:18:56 +0800] "GET / HTTP/1.1" 304 0 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
ssl.cloudfy669.xyz 185.244.208.50(-, -, RO ) - - [22/Sep/2023:22:20:33 +0800] "http GET / HTTP/1.1" 304 0 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36" 80
ssl.cloudfy669.xyz 35.88.71.231(-, -, US ) - - [22/Sep/2023:22:59:42 +0800] "http GET /.well-known/acme-challenge/l-YMATzOopxC7y3SXp5rxR1qacetLXbZ8tewNswzu0Q HTTP/1.1" 200 87 "-" "Mozilla/5.0 (compatible; Let's Encrypt validation server; +https://www.letsencrypt.org)" 80
ssl.cloudfy669.xyz 18.222.24.83(-, -, US ) - - [22/Sep/2023:22:59:42 +0800] "http GET /.well-known/acme-challenge/l-YMATzOopxC7y3SXp5rxR1qacetLXbZ8tewNswzu0Q HTTP/1.1" 200 87 "-" "Mozilla/5.0 (compatible; Let's Encrypt validation server; +https://www.letsencrypt.org)" 80
ssl.cloudfy669.xyz 23.178.112.209(-, -, - ) - - [22/Sep/2023:22:59:42 +0800] "http GET /.well-known/acme-challenge/l-YMATzOopxC7y3SXp5rxR1qacetLXbZ8tewNswzu0Q HTTP/1.1" 200 87 "-" "Mozilla/5.0 (compatible; Let's Encrypt validation server; +https://www.letsencrypt.org)" 80
---
ssl2.cloudfy669.xyz 23.178.112.107(-, -, - ) - - [23/Sep/2023:09:49:16 +0800] "http GET /.well-known/acme-challenge/rN9hQ-hRMyeRKJrBFK54WGhW_yVkizhrmUfd_SOhmVM HTTP/1.1" 200 87 "-" "Mozilla/5.0 (compatible; Let's Encrypt validation server; +https://www.letsencrypt.org)" 80
ssl2.cloudfy669.xyz 35.85.41.235(-, -, US ) - - [23/Sep/2023:09:49:16 +0800] "http GET /.well-known/acme-challenge/rN9hQ-hRMyeRKJrBFK54WGhW_yVkizhrmUfd_SOhmVM HTTP/1.1" 200 87 "-" "Mozilla/5.0 (compatible; Let's Encrypt validation server; +https://www.letsencrypt.org)" 80
ssl2.cloudfy669.xyz 3.129.90.220(-, -, US ) - - [23/Sep/2023:09:49:16 +0800] "http GET /.well-known/acme-challenge/rN9hQ-hRMyeRKJrBFK54WGhW_yVkizhrmUfd_SOhmVM HTTP/1.1" 200 87 "-" "Mozilla/5.0 (compatible; Let's Encrypt validation server; +https://www.letsencrypt.org)" 80
