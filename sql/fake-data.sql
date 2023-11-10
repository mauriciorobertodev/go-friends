-- Active: 1699462180782@@127.0.0.1@3306@gofriends
INSERT INTO users (name, nick, email, password) VALUES
("User 1", "user_1", "test1@email.com", "$2a$10$3cQ9NNp0CZKEUIdAK5gpPealq989NwV/k6P91KX.EUKhelSSNlnQm"),
("User 2", "user_2", "test2@email.com", "$2a$10$3cQ9NNp0CZKEUIdAK5gpPealq989NwV/k6P91KX.EUKhelSSNlnQm"),
("User 3", "user_3", "test3@email.com", "$2a$10$3cQ9NNp0CZKEUIdAK5gpPealq989NwV/k6P91KX.EUKhelSSNlnQm"),
("User 4", "user_4", "test4@email.com", "$2a$10$3cQ9NNp0CZKEUIdAK5gpPealq989NwV/k6P91KX.EUKhelSSNlnQm"),
("User 5", "user_5", "test5@email.com", "$2a$10$3cQ9NNp0CZKEUIdAK5gpPealq989NwV/k6P91KX.EUKhelSSNlnQm")

INSERT INTO followers (user_id, follower_id) VALUES
(1, 2),
(2, 1),
(2, 3),
(3, 4),
(4, 5),
(5, 1),
