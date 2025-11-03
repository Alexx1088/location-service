ALTER TABLE crossroads
ADD CONSTRAINT unique_crossroad UNIQUE (street_id, city_id);


