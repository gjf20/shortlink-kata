-- size of slug is linked to max output of the hashing function used in generateHash()
-- size of link is based off of Chrome's max url length
 CREATE TABLE links(
    slug VARCHAR(20) PRIMARY KEY,
    link VARCHAR(2048));


CREATE TABLE link_visits(
    slug VARCHAR(20) PRIMARY KEY,
    visits INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_slug
        FOREIGN KEY(slug)
            REFERENCES links(slug));


CREATE OR REPLACE FUNCTION func_new_link() RETURNS TRIGGER AS
$BODY$
BEGIN
    INSERT INTO
        link_visits
        VALUES (new.slug);
           RETURN new;
END;
$BODY$
language plpgsql;

CREATE TRIGGER trigger_new_link
     AFTER INSERT ON links
     FOR EACH ROW
     EXECUTE FUNCTION func_new_link();


INSERT INTO links VALUES ('test-slug', 'https://google.com');

SELECT * FROM link_visits WHERE slug = 'test-slug';



ALTER TRIGGER trigger_new_link
ON link 
RENAME TO new_trigger_name_1;