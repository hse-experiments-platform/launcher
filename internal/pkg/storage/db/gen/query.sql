-- name: GetLaunches :many
select *,
       count(1) over () as count
from launches
where name like '%' || sqlc.arg(name) || '%'
  and (launch_type = $4 or $4 = '')
  and user_id = $5
order by created_at desc
limit $1 offset $2;

-- name: GetLaunch :one
select id,
       name,
       description,
       launch_status,
       launch_error,
       launch_type,
       input,
       output
from launches
where id = $1;

-- name: GetDatasetCreator :one
select creator_id
from datasets
where id = $1;

-- name: CreateLaunch :one
insert into launches (launch_type, user_id, name, description, launch_status, input)
values ($1, $2, $3, $4, $5, $6)
returning id;

-- name: UpdateLaunchStatus :exec
update launches
set launch_status = $2,
    updated_at    = now()
where id = $1;

-- name: FinishLaunch :exec
update launches
set launch_status = $2,
    updated_at    = now(),
    finished_at   = now(),
    launch_error  = $3,
    output        = $4
where id = $1;

-- name: CheckLaunchStatus :one
select launch_status = any (sqlc.arg(statuses)::text[])
from launches
where id = $1;

-- name: CheckModelExists :one
select class_name
from models
where id = $1;

-- name: GetModelHyperparameters :many
select h.id,
       h.name,
       h.description,
       h.type,
       h.default_value
from models m
         join hyperparameters h on m.id = h.model_id
where m.id = $1;

-- name: GetModelMetrics :many
select m2.metric_name
from models m
         join metrics m2 on m.id = m2.model_id
where m.id = $1;

-- name: GetLaunchType :one
select launch_type
from launches
where id = $1;

-- name: GetTrainedModelIDByLaunchID :one
select id
from trained_models
where launch_id = $1;

-- name: CreateTrainedModel :one
insert into trained_models (name, description, base_model_id, train_dataset_id, launch_id, target_col)
values (sqlc.arg(name), sqlc.arg(description), sqlc.arg(base_model_id), sqlc.arg(train_dataset_id), sqlc.arg(launch_id),
        sqlc.arg(target_column))
returning id;

-- name: CheckDatasetStatus :one
select status::text = any (sqlc.arg(statuses)::text[])
from datasets
where id = $1;

-- name: GetTrainedModelRunID :one
select l.output
from trained_models tm
         join launches l on tm.launch_id = l.id
where tm.id = $1;