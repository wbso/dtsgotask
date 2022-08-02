create table tasks (
    id uuid primary key,
    detail text not null,
    is_done boolean not null default false,
    assignee text not null default '',
    deadline timestamp with time zone not null default CURRENT_TIMESTAMP,
    created_at timestamp with time zone not null default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone not null default CURRENT_TIMESTAMP
);