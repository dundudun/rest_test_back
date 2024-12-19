-- name: CreateWasteStorage :exec
insert into waste_storage (
    name, 
    plastic_limit, 
    glass_limit, 
    biowaste_limit,
    stored_plastic,
    stored_glass,
    stored_biowaste
) values (
    $1, $2, $3, $4, $5, $6, $7
);

-- name: ListWasteStorage :many
select * from waste_storage 
order by name;

-- name: GetWasteStorage :one
select * from waste_storage 
where id = $1 limit 1;

-- name: UpdateWasteStorage :exec
update waste_storage set
    name = $2, 
    plastic_limit = $3, 
    glass_limit = $4, 
    biowaste_limit = $5,
    stored_plastic = $6,
    stored_glass = $7,
    stored_biowaste = $8
where id = $1;

-- name: PartlyUpdateWasteStorage :one
update waste_storage set
    name = coalesce($2, name),
    plastic_limit = coalesce($3, plastic_limit),
    glass_limit = coalesce($4, glass_limit),
    biowaste_limit = coalesce($5, biowaste_limit),
    stored_plastic = coalesce($6, stored_plastic),
    stored_glass = coalesce($7, stored_glass),
    stored_biowaste = coalesce($8, stored_biowaste)
where id  = $1
returning *;

-- name: DeleteWasteStorage :exec
delete from waste_storage
where id = $1;