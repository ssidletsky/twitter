CREATE DATABASE IF NOT EXISTS twitter;
USE twitter;

CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED AUTO_INCREMENT,
    username CHAR(15) NOT NULL,
    password CHAR(60) NOT NULL,
    email VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL DEFAULT "",
    last_name VARCHAR(100) NOT NULL DEFAULT "",
    age TINYINT UNSIGNED NOT NULL DEFAULT 0,
    
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS tweets (
	id INT UNSIGNED AUTO_INCREMENT,
	author_user_id INT UNSIGNED NOT NULL,
	text VARCHAR(280) NOT NULL,
	publication_date DATETIME NOT NULL,
	
	PRIMARY KEY (id),
	FOREIGN KEY (author_user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS followers (
	followed_user_id INT UNSIGNED NOT NULL,
	follower_user_id INT UNSIGNED NOT NULL,
	
	PRIMARY KEY (followed_user_id, follower_user_id),
	FOREIGN KEY (followed_user_id) REFERENCES users(id),
	FOREIGN KEY (follower_user_id) REFERENCES users(id)
) ENGINE=InnoDB;
