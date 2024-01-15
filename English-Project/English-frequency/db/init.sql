
CREATE DATABASE english_frequency;
USE english_frequency;


CREATE TABLE word (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    word VARCHAR(45),
    word_type TINYINT
);

CREATE TABLE mst_wordtype(
    word_type_id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    word_type_name VARCHAR(20)
);

CREATE TABLE translation(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    word_id INT,
    word_jp VARCHAR(100),
    word_type_id INT,
    FOREIGN KEY fk_word_id (word_id) REFERENCES word (id),
    FOREIGN KEY fk_word_type (word_type_id) REFERENCES mst_wordtype (word_type_id)
);

CREATE TABLE mst_provider(
    id VARCHAR(20) PRIMARY KEY NOT NULL,
    site_name VARCHAR(50),
    url VARCHAR(100)
);

CREATE TABLE frequency(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    provider_id VARCHAR(20),
    word_id INT,
    count INT,
    date DATE,
    FOREIGN KEY fk_provider_id (provider_id) REFERENCES mst_provider (id),
    FOREIGN KEY fk_word_id (word_id) REFERENCES word (id)
);