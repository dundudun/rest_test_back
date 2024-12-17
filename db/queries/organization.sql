-- name: CreateOrganization :one
insert into organization (
    name, 
    plastic_limit, 
    glass_limit, 
    biowaste_limit,
    produced_plastic,
    produced_glass,
    produced_biowaste
) values (
    $1, $2, $3, $4, $5, $6, $7
) 
returning *;

-- name: ListOrganizations :many
select * from organization 
order by name;

-- name: GetOrganization :one
select * from organization 
where id = $1 limit 1;

-- name: UpdateOrganization :one
update organization set
    name = $2, 
    plastic_limit = $3, 
    glass_limit = $4, 
    biowaste_limit = $5,
    produced_plastic = $6,
    produced_glass = $7,
    produced_biowaste = $8
where id = $1
returning *;

-- name: PartlyUpdateOrganization :one
update organization set
    name = $2, 
    plastic_limit = $3, 
    glass_limit = $4, 
    biowaste_limit = $5,
    produced_plastic = $6,
    produced_glass = $7,
    produced_biowaste = $8
where id = $1
returning *;

-- name: DeleteOrganization :exec
delete from organization
where id = $1;