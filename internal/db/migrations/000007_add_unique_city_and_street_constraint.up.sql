ALTER TABLE cities
ADD CONSTRAINT unique_city_name UNIQUE (name);

ALTER TABLE streets
ADD CONSTRAINT unique_street_name UNIQUE (name);


