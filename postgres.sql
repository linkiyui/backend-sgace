
Create Table if not exists users (
    id uuid not null default gen_random_uuid(),
    username varchar(255) not null unique,
    password varchar(255) not null,
    email varchar(255) not null unique,
    role varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key (id)
);


Create Table if not exists activities (
    id uuid not null default gen_random_uuid(),
    name varchar(255) not null,
    type varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    date varchar(255) not null,
    time varchar(255) not null,
    primary key (id)
);


Create Table if not exists solicitudes (
    id uuid not null default gen_random_uuid(),
    user_id uuid not null,
    activity_id uuid not null,
    group_id varchar(255) not null,
    faculty varchar(255) not null,
    grade varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    status int not null,
    primary key (id),
    foreign key (user_id) references users(id),
    foreign key (activity_id) references activities(id)
);



