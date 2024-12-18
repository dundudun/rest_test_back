-- name: ConnectStorAndOrg :exec
insert into org_to_stor(
    organization_id,
    waste_storage_id,
    distance_meters
) values (
    $1, $2, $3
);

-- name: DelConnStorAndOrg :exec
delete from org_to_stor
where organization_id = $1 and waste_storage_id = $2;

-- name: ConnectTwoStors :exec
insert into stor_to_stor(
    prev_waste_storage_id,
    next_waste_storage_id,
    distance_meters
) values (
    $1, $2, $3
);

-- name: DelConnBetweenStors :exec
delete from stor_to_stor
where prev_waste_storage_id = $1 and next_waste_storage_id = $2;





-- name: FromOrgPlasticStors :many
select 
    ws.id, 
    ws.name, 
    ws.plastic_limit, 
    ws.stored_plastic,
    rel.distance_meters
from org_to_stor rel
join waste_storage ws 
    on rel.waste_storage_id = ws.id
where rel.organization_id = $1
    and plastic_limit is not NULL
    and plastic_limit <> 0
order by distance_meters asc;

-- name: FromStorsPlasticStors :many
select 
    ws.id, 
    ws.name, 
    ws.plastic_limit, 
    ws.stored_plastic,
    rel.distance_meters
from stor_to_stor rel
join waste_storage ws 
    on rel.next_waste_storage_id = ws.id
where rel.prev_waste_storage_id = $1
    and ws.plastic_limit is not NULL
    and ws.plastic_limit <> 0
order by distance_meters asc;





-- name: FromOrgGlassStors :many
select 
    ws.id, 
    ws.name, 
    ws.glass_limit, 
    ws.stored_glass,
    rel.distance_meters
from org_to_stor rel
join waste_storage ws 
    on rel.waste_storage_id = ws.id
where rel.organization_id = $1
    and glass_limit is not NULL
    and glass_limit <> 0
order by distance_meters asc;





-- name: FromOrgBiowasteStors :many
select 
    ws.id, 
    ws.name, 
    ws.biowaste_limit, 
    ws.stored_biowaste,
    rel.distance_meters
from org_to_stor rel
join waste_storage ws 
    on rel.waste_storage_id = ws.id
where rel.organization_id = $1
    and biowaste_limit is not NULL
    and biowaste_limit <> 0
order by distance_meters asc;