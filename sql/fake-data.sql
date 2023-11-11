-- Active: 1699462180782@@127.0.0.1@3306@gofriends
INSERT INTO users (name, nick, email, password) VALUES
("User 1", "user_1", "test1@email.com", "$2a$10$3cQ9NNp0CZKEUIdAK5gpPealq989NwV/k6P91KX.EUKhelSSNlnQm"),
("User 2", "user_2", "test2@email.com", "$2a$10$3cQ9NNp0CZKEUIdAK5gpPealq989NwV/k6P91KX.EUKhelSSNlnQm"),
("User 3", "user_3", "test3@email.com", "$2a$10$3cQ9NNp0CZKEUIdAK5gpPealq989NwV/k6P91KX.EUKhelSSNlnQm")

INSERT INTO followers (user_id, follower_id) VALUES
(1, 2),
(2, 1),
(2, 3)

INSERT INTO posts (title, content, author_id) VALUES
("Post do 1", "Conteúdo", 1),
("Post do 2", "Conteúdo", 2),
("Post do 3", "Conteúdo", 3)
