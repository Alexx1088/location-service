# location-service
The service provides information about **cities**, **streets**, and **crossroads**.

## Entity schema
### City
```
- id
- name
```
### Street
```
- id
- name
- city_id
```
### Crossroad
```
- id
- street_id
```

## Requirements:
The location-service must implement simple **CRUD** operations (```controller```, ```service``` and ```repository``` layers) for all entities.
