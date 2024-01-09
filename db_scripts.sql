DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS games_chances;
DROP TABLE IF EXISTS houses CASCADE;
DROP TABLE IF EXISTS houses_photos;
DROP TABLE IF EXISTS cities CASCADE;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS pickup_statuses CASCADE;
DROP TABLE IF EXISTS pickups;
DROP TABLE IF EXISTS reservations CASCADE;
DROP TABLE IF EXISTS blacklists;

CREATE TABLE users (
                       id INT UNIQUE NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       full_name VARCHAR(255) NOT NULL,
                       address VARCHAR(255),
                       city_id INT NOT NULL,
                       city_name VARCHAR (255),
                       role VARCHAR(10),
                       created_at TIMESTAMP DEFAULT Now(),
                       updated_at TIMESTAMP DEFAULT Now(),
                       deleted_at TIMESTAMP,
                       PRIMARY KEY(id)
);

CREATE TABLE wallets (
                         id INT UNIQUE NOT NULL,
                         user_id INT NOT NULL,
                         balance INT,
                         created_at TIMESTAMP DEFAULT Now(),
                         updated_at TIMESTAMP DEFAULT Now(),
                         deleted_at TIMESTAMP,
                         PRIMARY KEY (id),
                         CONSTRAINT users
                             FOREIGN KEY (user_id)
                                 REFERENCES users(id)
);

CREATE TABLE games_chances (
                               id INT UNIQUE NOT NULL,
                               user_id INT NOT NULL,
                               chance INT,
                               history INT,
                               created_at TIMESTAMP DEFAULT Now(),
                               updated_at TIMESTAMP DEFAULT Now(),
                               deleted_at TIMESTAMP,
                               PRIMARY KEY (id),
                               CONSTRAINT users
                                   FOREIGN KEY (user_id)
                                       REFERENCES users(id)
);

CREATE TABLE houses (
                        id INT UNIQUE NOT NULL,
                        house_name VARCHAR(255) NOT NULL,
                        user_id INT NOT NULL,
                        price_per_night INT,
                        description VARCHAR(255),
                        city_id INT NOT NULL,
                        max_guest INT,
                        created_at TIMESTAMP DEFAULT Now(),
                        updated_at TIMESTAMP DEFAULT Now(),
                        deleted_at TIMESTAMP
);

CREATE TABLE houses_photos (
                               id INT UNIQUE NOT NULL,
                               house_id INT NOT NULL,
                               photo_url VARCHAR(255),
                               created_at TIMESTAMP DEFAULT Now(),
                               updated_at TIMESTAMP DEFAULT Now(),
                               deleted_at TIMESTAMP
);

CREATE TABLE cities (
                        id INT UNIQUE NOT NULL,
                        name VARCHAR(255) NOT NULL,
                        created_at TIMESTAMP DEFAULT Now(),
                        updated_at TIMESTAMP DEFAULT Now(),
                        deleted_at TIMESTAMP
);

CREATE TABLE transactions (
                              id INT UNIQUE NOT NULL,
                              user_id INT NOT NULL,
                              house_id INT NOT NULL,
                              reservation_id INT NOT NULL,
                              created_at TIMESTAMP DEFAULT Now(),
                              updated_at TIMESTAMP DEFAULT Now(),
                              deleted_at TIMESTAMP
);

CREATE TABLE pickup_statuses (
                                 id INT UNIQUE NOT NULL,
                                 status VARCHAR(255) NOT NULL,
                                 created_at TIMESTAMP DEFAULT Now(),
                                 updated_at TIMESTAMP DEFAULT Now(),
                                 deleted_at TIMESTAMP
);

CREATE TABLE pickups (
                         id INT UNIQUE NOT NULL,
                         user_id INT NOT NULL,
                         reservation_id INT NOT NULL,
                         pick_up_status_id INT NOT NULL,
                         created_at TIMESTAMP DEFAULT Now(),
                         updated_at TIMESTAMP DEFAULT Now(),
                         deleted_at TIMESTAMP,
                         PRIMARY KEY (id),
                         CONSTRAINT pickup_statuses
                             FOREIGN KEY (pick_up_status_id)
                                 REFERENCES pickup_statuses(id)
);

CREATE TABLE reservations (
                              id INT UNIQUE NOT NULL,
                              house_id INT NOT NULL,
                              user_id INT NOT NULL,
                              check_in_date TIMESTAMP,
                              check_out_date TIMESTAMP,
                              total_price INT,
                              created_at TIMESTAMP DEFAULT Now(),
                              updated_at TIMESTAMP DEFAULT Now(),
                              deleted_at TIMESTAMP,
                              PRIMARY KEY(id)
);

CREATE TABLE blacklists (
    id INT UNIQUE NOT NULL,
    token TEXT,
    created_at TIMESTAMP DEFAULT Now(),
    updated_at TIMESTAMP DEFAULT Now(),
    deleted_at TIMESTAMP,
    PRIMARY KEY(id)
);
