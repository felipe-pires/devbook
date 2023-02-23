create if not exists database devbook;

use devbook;

drop table if exists usuarios;

create table usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criado_em timestamp default current_timestamp()
) engine=innodb;