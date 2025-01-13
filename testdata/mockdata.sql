CREATE TABLE organization (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    plastic_limit integer,
    glass_limit integer,
    biowaste_limit integer,
    produced_plastic integer,
    produced_glass integer,
    produced_biowaste integer
);

CREATE TABLE waste_storage (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    plastic_limit integer,
    glass_limit integer,
    biowaste_limit integer,
    stored_plastic integer,
    stored_glass integer,
    stored_biowaste integer
);

CREATE TABLE org_to_stor (
    id bigserial PRIMARY KEY,
    organization_id bigint REFERENCES organization(id),
    waste_storage_id bigint REFERENCES waste_storage(id),
    distance_meters integer NOT NULL
);

CREATE TABLE stor_to_stor (
    id bigserial PRIMARY KEY,
    prev_waste_storage_id bigint REFERENCES waste_storage(id),
    next_waste_storage_id bigint REFERENCES waste_storage(id),
    distance_meters integer NOT NULL
);

--check for creation and further adding, id specifically
insert into organization values (1, 'test_org', 100, 50, 0, 15, 0, 0)