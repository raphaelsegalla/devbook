insert into usuarios (nome, nick, email, senha)
values
("Raphael Fernando Pimentel Segalla", "rsegalla", "rsegalla@gmail.com", "$2a$10$8bRAEgXimffeaKtEuJk1HeS1N5z.5x.GYQaBK5P0ZBQeLiHPBCnp."),
("Thaina Liberato Segalla", "tsegalla", "tsegalla@gmail.com", "$2a$10$8bRAEgXimffeaKtEuJk1HeS1N5z.5x.GYQaBK5P0ZBQeLiHPBCnp."),
("Alessandra Liberato Segalla", "asegalla", "asegalla@gmail.com", "$2a$10$8bRAEgXimffeaKtEuJk1HeS1N5z.5x.GYQaBK5P0ZBQeLiHPBCnp."),
("Usuario 1", "usuario1", "usuario1@gmail.com", "$2a$10$8bRAEgXimffeaKtEuJk1HeS1N5z.5x.GYQaBK5P0ZBQeLiHPBCnp."),
("Usuario 2", "usuario2", "usuario2@gmail.com", "$2a$10$8bRAEgXimffeaKtEuJk1HeS1N5z.5x.GYQaBK5P0ZBQeLiHPBCnp."),
("Usuario 3", "usuario3", "usuario3@gmail.com", "$2a$10$8bRAEgXimffeaKtEuJk1HeS1N5z.5x.GYQaBK5P0ZBQeLiHPBCnp.");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3),
(1, 4),
(1, 5),
(2, 1),
(2, 6);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicação do Raphael", "Essa é a publicação do Raphael! Oba!", 1),
("Publicação da Thaina", "Essa é a publicação da Thaina! Oba!", 2),
("Publicação da Alessandra", "Essa é a publicação da Alessandra! Oba!", 3),
("Publicação do Usuario 1", "Essa é a publicação do Usuario 1! Oba!", 4),
("Publicação do Usuario 2", "Essa é a publicação do Usuario 2! Oba!", 5),
("Publicação do Usuario 3", "Essa é a publicação do Usuario 3! Oba!", 6);