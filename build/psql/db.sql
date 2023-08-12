CREATE TABLE accounts(
    id int NOT NULL AUTO_INCREMENT,
    user_id int
    amount int
    PRIMARY KEY (id)
)

CREATE TABLE transaction (
    id int NOT NULL AUTO_INCREMENT,
    type string
    amount int NOT NULL
    account_id int NOT NULL
    PRIMARY KEY (id)
    FOREIGN KEY (account_id) REFERENCES accounts(id)
)

