DROP TABLE IF EXISTS playlist;
CREATE TABLE playlist
(
	track_id bigserial PRIMARY KEY NOT NULL,
	track_name varchar(256) NOT NULL,
	track_duration integer NOT NULL
);

INSERT INTO playlist (track_name, track_duration)
VALUES
('Queen - «Bohemian Rhapsody»', '12'),
('Joy Division - «Love Will Tear Us Apart»', '15'),
('The Beatles - «A Day in the Life»', '4');
-- ('Queen - «Bohemian Rhapsody»', '355'),
-- ('Joy Division - «Love Will Tear Us Apart»', '205'),
-- ('The Beatles - «A Day in the Life»', '333');