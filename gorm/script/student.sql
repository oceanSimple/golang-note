use project;
drop table if exists students;
create table students
(
    id           int unsigned primary key auto_increment,
    name         varchar(255),
    age          int,
    is_deleted   tinyint unsigned not null default 0 comment '是否删除, 0: 未删除, 1: 已删除',
    gmt_create   timestamp        not null comment '创建时间',
    gmt_modified timestamp         not null comment '修改时间'
) comment 'student' charset = utf8
                    engine = InnoDB;