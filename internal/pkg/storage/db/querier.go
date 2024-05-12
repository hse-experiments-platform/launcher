// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CheckLaunchStatus(ctx context.Context, arg CheckLaunchStatusParams) (bool, error)
	CreateLaunch(ctx context.Context, arg CreateLaunchParams) (int64, error)
	FinishLaunch(ctx context.Context, arg FinishLaunchParams) error
	GetDatasetCreator(ctx context.Context, id int64) (int64, error)
	GetLaunch(ctx context.Context, id int64) (GetLaunchRow, error)
	// launches schema
	// create table if not exists launches
	// (
	//     id            bigint primary key generated by default as identity,
	//     launch_type   text                     not null,
	//     user_id       bigint                   not null references users (id),
	//     name          text                     not null,
	//     description   text                     not null,
	//     created_at    timestamp with time zone not null default now(),
	//     updated_at    timestamp with time zone not null default now(),
	//     finished_at   timestamp with time zone,
	//     launch_status text                     not null,
	//     launch_error  text
	// );
	GetLaunches(ctx context.Context, arg GetLaunchesParams) ([]GetLaunchesRow, error)
	UpdateLaunchStatus(ctx context.Context, arg UpdateLaunchStatusParams) error
}

var _ Querier = (*Queries)(nil)
