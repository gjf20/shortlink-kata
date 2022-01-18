-- convenient queries for modification
INSERT INTO links VALUES ('test-slug', 'https://google.com');

SELECT * FROM link_visits WHERE slug = 'test-slug';

ALTER TRIGGER trigger_new_link
ON link 
RENAME TO new_trigger_name_1;