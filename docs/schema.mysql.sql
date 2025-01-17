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
    site_id int unsigned not null default '0' comment 'sites.id',
    name varchar(100) not null default '',
    domains text comment '',
    ssl_cert_state tinyint not null default '0' comment 'ssl 证书 状态 1:等待, 2: 申请开始, 3: 已完成, 4:失败',
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
drop table if exists sites;
create table sites (
    id int unsigned auto_increment primary key ,
    name varchar(200) not null default '' comment '',
    domains text comment '[]string, 逗号隔开',
    domain_count smallint unsigned default '0',
    state tinyint not null default '0' comment '1: 启用, 2:禁用',
    websocket_enable tinyint not null default '0' comment '1: 启用, 2:禁用',
    ssl_enable tinyint not null default '0' comment 'ssl是否启用 1:启用, 2:禁用',
    ssl_cert_state tinyint not null default '0' comment 'ssl 证书 状态 1:等待, 2: 申请开始, 3: 已完成, 4:失败',
    ssl_filename varchar(200) not null default '' comment 'ssl文件名',
    ssl_obtain_log text,
    ssl_private_key text,
    ssl_fullchain_cer text,
    ssl_expired_at int unsigned not null default '0' comment 'ssl 过期时间',
    http_ports varchar(300) not null default '' comment 'http端口',
    http_redirect tinyint not null default '0' comment 'http 转 https(443) 1: 转发, 2: 不转发',
    redirect varchar(200) not null default '' comment '跳转,',
    hsts_enable tinyint not null default '0' comment '1:启用, 2:禁用',
    https_ports varchar(300) not null default '' comment 'http',
    upstream_port_policy tinyint not null default '0' comment '1:同端口协议, 2: 回落到 80, 3: 回落到 443',
    upstream_rotate_policy tinyint not null default '0' comment '1: 轮询(robin round, 2: ip hash',
    upstream_ips varchar(100) not null default '' comment '上游ip,多个以逗号分隔,会轮询',
    upstream_host varchar(100) not null default '' comment '是否',
    created_at int unsigned not null default '0',
    updated_at int unsigned not null default '0',
    deleted_at int unsigned not null default '0' comment '删除标记',
    index(name),
    index(created_at),
    index(deleted_at)
) engine=innoDB default charset=utf8mb4 comment '域名';
-- CREATE INDEX `idx_sites_deleted_at` ON `sites`(`deleted_at`);
-- alter table sites add column redirect varchar(200) not null default '' comment '跳转' after http_redirect;

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
) engine=innoDB default charset=utf8mb4 comment '环境变量';

DROP TABLE IF EXISTS sys_log;
CREATE TABLE sys_log (
    id bigint unsigned auto_increment primary key ,
    trace_id varchar(100) not null default '' comment '跟踪码',
    node_id varchar(50) not null default '' comment '节点 master:主节点, node-xxx:子节点',
    restarted smallint not null default '0' comment '是否重启 nginx 1:重启, 2:没有',
    site_updated smallint not null default '0' comment '更新站点数量',
    content text,
    created_at int unsigned not null default '0',
    index(trace_id),
    index(created_at)
) engine=innoDB default charset=utf8mb4;

drop table if exists cdn_nodes;
CREATE TABLE cdn_nodes (
    id int unsigned auto_increment primary key ,
    title varchar(20) not null default '' comment '名称',
    api_url varchar(200) not null default '' comment '',
    state tinyint not null default '0' comment '1:正常 ,2: 暂停',
    created_at int unsigned not null default '0',
    updated_at int unsigned not null default '0',
    deleted_at int unsigned not null default '0' comment '删除标记',
    index(deleted_at)
) engine=innoDB default charset=utf8mb4 comment '节点';


drop table if exists settings;
CREATE TABLE settings (
    id int unsigned auto_increment primary key ,
    name varchar(30) not null default '',
    content varchar(1000) not null default '' comment '内容',
    data_type varchar(20) not null default '' comment 'json, array, int, string'
) engine=innoDB comment '配置' default charset=utf8mb4;
-- insert into settings (`name`, content, `data_type`) values ('cname', 'cname1.cloud699.xyz,cname2.cloud699.xyz', 'array');
--
