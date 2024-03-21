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
select $1, unnest(sqlc.arg(hyperparameter_ids)::bigint[]), unnest(sqlc.arg(values)::jsonb[]);

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
