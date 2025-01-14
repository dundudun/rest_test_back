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
) returning *;

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
    name = coalesce($2, name),
    plastic_limit = coalesce($3, plastic_limit),
    glass_limit = coalesce($4, glass_limit),
    biowaste_limit = coalesce($5, biowaste_limit),
    produced_plastic = coalesce($6, produced_plastic),
    produced_glass = coalesce($7, produced_glass),
    produced_biowaste = coalesce($8, produced_biowaste)
where id = $1
returning *;

-- name: DeleteOrganization :one
delete from organization
where id = $1
returning id;