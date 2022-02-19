insert into usuarios (nome, nick, email, senha)
values
("Usuario1","userNick@1","usuario01@gmail.com","$2a$10$qGkud74BEbgDdShaBftaKOLL77iJwXBFkw7pdqfLYGXy4DOjxmd8a"), --
("Usuario2","userNick@2","usuario02@gmail.com","$2a$10$qGkud74BEbgDdShaBftaKOLL77iJwXBFkw7pdqfLYGXy4DOjxmd8a"), --
("Usuario3","userNick@3","usuario03@gmail.com","$2a$10$qGkud74BEbgDdShaBftaKOLL77iJwXBFkw7pdqfLYGXy4DOjxmd8a"); --

insert into seguidores (usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes (titulo, conteudo, autor_id)
values
("Publicação do usuario 1","Essa é a publicação do usuário 1, Oba!", 1),
("Publicação do usuario 2","Essa é a publicação do usuário 2, Oba!", 2),
("Publicação do usuario 3","Essa é a publicação do usuário 3, Oba!", 3);
