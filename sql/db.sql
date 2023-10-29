CREATE TABLE IF NOT EXISTS products (
    id INT,
    PRIMARY KEY(id)
);

INSERT INTO products
VALUES (1);

CREATE TABLE IF NOT EXISTS packsizes(
  pid INT, 
  size INT,
  PRIMARY KEY(pid,size),
  FOREIGN KEY(pid) REFERENCES products(id)
);

INSERT INTO packsizes
VALUES (1,250),(1,500),(1,1000),(1,2000),(1,5000);


