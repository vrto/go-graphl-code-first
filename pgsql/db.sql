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

INSERT INTO pub VALUES(1, 'Lokal');
INSERT INTO pub VALUES(2, 'Poupe');

INSERT INTO beer VALUES (1, 'Pilsner Urquell', 'Plzensky prazdroj', 'Lager');
INSERT INTO beer VALUES (2, 'Raptor', 'Matuska', 'IPA');
INSERT INTO beer VALUES (3, 'Budvar 33', 'Budvar', 'Lager');

INSERT INTO pubs_beers VALUES (1, 1);
INSERT INTO pubs_beers VALUES (1, 3);

INSERT INTO pubs_beers VALUES (2, 1);
INSERT INTO pubs_beers VALUES (2, 2);