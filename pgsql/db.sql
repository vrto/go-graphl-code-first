CREATE TABLE pub(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE beer(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    maker VARCHAR(255) NOT NULL,
    type VARCHAR(20) NOT NULL
);

CREATE TABLE pubs_beers(
    pub_id INT NOT NULL,
    beer_id INT NOT NULL,
    PRIMARY KEY (pub_id, beer_id),
    CONSTRAINT fk_pub FOREIGN KEY(pub_id) REFERENCES pub(id),
    CONSTRAINT fk_beer FOREIGN KEY(beer_id) REFERENCES beer(id)
);

INSERT INTO pub(name) VALUES('Lokal');
INSERT INTO pub(name) VALUES('Poupe');

INSERT INTO beer(name, maker, type) VALUES ('Pilsner Urquell', 'Plzensky prazdroj', 'Lager');
INSERT INTO beer(name, maker, type) VALUES ('Raptor', 'Matuska', 'IPA');
INSERT INTO beer(name, maker, type) VALUES ('Budvar 33', 'Budvar', 'Lager');

-- serial starts with 1, safe to hardcode keys

INSERT INTO pubs_beers VALUES (1, 1);
INSERT INTO pubs_beers VALUES (1, 3);

INSERT INTO pubs_beers VALUES (2, 1);
INSERT INTO pubs_beers VALUES (2, 2);