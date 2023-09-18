DROP TABLE IF EXISTS auth_tokens;
CREATE TABLE auth_tokens (
    id int unsigned auto_increment primary key ,
    token varchar(200) not null default ''
) engine=innoDB default charset=utf8mb4;

-- CREATE TABLE `certs` (`id` integer,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,
-- `name` text,`domains` text[],`filename` text,`ssl_certificate_path` text,
-- `ssl_certificate_key_path` text,
-- `auto_cert` integer,`challenge_method` text,`dns_credential_id` integer,`log` text,PRIMARY KEY (`id`));
DROP TABLE IF EXISTS certs;
CREATE TABLE certs (
    id int unsigned auto_increment primary key ,
    name varchar(100) not null default '',
    domains text comment '',
    ssl_key text comment 'ssl_certificate_key private.key;',
    ssl_cer text comment 'ssl_certificate fullchain.cer;',
    expired_at int unsigned not null default '0' comment '过期时间',
    log text,
    created_at int unsigned not null default '0',
    updated_at int unsigned not null default '0',
    deleted_at int unsigned not null default '0' comment '删除标记',
    index(name)
) engine=innoDB comment '证书表' default charset=utf8mb4;
-- CREATE INDEX `idx_certs_deleted_at` ON `certs`(`deleted_at`);


-- CREATE TABLE `auths` (
--    `id` integer,`created_at` datetime,`updated_at` datetime,
--    `deleted_at` datetime,`name` text,`password` text,PRIMARY KEY (`id`));
drop table if exists auths;
CREATE TABLE auths (
    id int unsigned auto_increment primary key ,
    username varchar(200) not null default '' comment '登陆名',
    encrypted_password varchar(200) not null default '' ,
    created_at int unsigned not null default '0',
    updated_at int unsigned not null default '0',
    deleted_at int unsigned not null default '0' comment '删除标记',
    index(username)
) engine =innoDB default charset=utf8mb4 comment '管理员表';

-- 123456
INSERT INTO auths (username, encrypted_password) values ('admin','$2a$12$o9/W9tgYusr3zKzuNZgtD.y4Sjn7mjVrAVNsKX/VyqOkYiNGfFeYu');
-- CREATE INDEX `idx_auths_deleted_at` ON `auths`(`deleted_at`);
-- CREATE TABLE `sites` (
 -- `id` integer,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,
-- `path` text,`advanced` numeric,PRIMARY KEY (`id`));
create table sites (
    id int unsigned auto_increment primary key ,
    name varchar(200) not null default '' comment '',
    domains json comment '[]string',
    state tinyint not null default '0' comment '1: 启用, 2:禁用',
    created_at int unsigned not null default '0',
    updated_at int unsigned not null default '0',
    deleted_at int unsigned not null default '0' comment '删除标记',
    index(name),
    index(created_at)
) engine=innoDB default charset=utf8mb4 comment '域名';
-- CREATE INDEX `idx_sites_deleted_at` ON `sites`(`deleted_at`);

--
-- CREATE TABLE `environments` (
-- `id` integer,`created_at` datetime,
-- `updated_at` datetime,`deleted_at` datetime,
-- `name` text,
-- `url` text,
-- `token` text,
-- `operation_sync` numeric,
-- `sync_api_regex` text,PRIMARY KEY (`id`));
-- CREATE INDEX `idx_environments_deleted_at` ON `environments`(`deleted_at`);
CREATE TABLE environments (
    id int unsigned auto_increment primary key ,
    name varchar(200) not null default '',
    url varchar(255) not null default '',
    token varchar(255) not null default '',
    operation_sync int not null default '0',
    sync_api_regex text,
    created_at int unsigned not null default '0',
    updated_at int unsigned not null default '0',
    deleted_at int unsigned not null default '0' comment '删除标记',
    index(deleted_at)
) engine=innoDB default charset=utf8mb4;
