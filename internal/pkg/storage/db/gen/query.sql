-- name: CreateLaunch :one
insert into launches (launch_type, name, description)
values ($1, $2, $3)
returning id;

-- name: CreateTrainedModel :one
insert into trained_models (name, description, model_id, model_training_status, training_dataset_id, target_column,
                            launch_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
returning id;

-- name: CreateTrainingHyperparameters :exec
insert into train_hyperparameters (train_model_id, hyperparameter_id, value)
select $1, unnest(sqlc.arg(hyperparameter_ids)::bigint[]), unnest(sqlc.arg(values)::text[]);

-- name: UpdateTrainedModelStatus :exec
update trained_models
set train_error           = $2,
    model_training_status = $3
where id = $1;

-- name: UpdateLaunchStatus :exec
update launches
set launch_error = $2,
    updated_at   = default,
    finished_at  = default
where id = $1;

-- name: GetDataset :one
select id,
       status,
       creator_id,
       rows_count
from datasets
where id = $1;

-- name: GetDatasetSchema :many
select column_name, column_type
from dataset_schemas
where dataset_id = $1
order by column_number;

-- name: InsertTrainMetrics :exec
insert into train_metrics (launch_id, trained_model_id, metric_id, value)
VALUES ($1, $2, $3, $4);

-- name: GetTrainedModelMetrics :many
select tm.*
from train_metrics tm
         join trained_models tm on tm.launch_id = $1;