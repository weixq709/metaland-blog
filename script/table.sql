drop database if exists `metaland-blog`;
create database `metaland-blog` character set utf8 collate utf8_general_ci;
use `metaland-blog`;

drop table if exists user;
create table user(
    id bigint primary key AUTO_INCREMENT,
    username varchar(20) not null comment '用户名',
    password varchar(100) not null comment '密码',
    email varchar(100) comment '邮箱'
);

drop table if exists article;
create table article(
    id bigint primary key AUTO_INCREMENT,
    title varchar(100) not null comment '标题',
    user_id bigint comment '用户ID',
    content mediumtext comment '文章内容',
    create_time datetime comment '创建时间',
    update_time datetime comment '修改时间'
);

drop table if exists comment;
create table comment(
    id bigint primary key AUTO_INCREMENT,
    content tinytext comment '评论内容',
    user_id bigint comment '评论人用户ID',
    article_id bigint comment '评论文章ID',
    create_time datetime comment '创建时间'
);