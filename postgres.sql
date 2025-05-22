
Create Table if not exists solicitudes (
    id varchar(255) not null default uuid_generate_v7(),
    user_id varchar(255) not null,
    activity_id varchar(255) not null,
    group varchar(255) not null,
    faculty varchar(255) not null,
    grade varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    status int not null,
    primary key (id)
    foreign key (user_id) references users(id),
    foreign key (activity_id) references activities(id)
);

Create Table if not exists users (
    id varchar(255) not null default uuid_generate_v7(),
    username varchar(255) not null unique,
    password varchar(255) not null,
    email varchar(255) not null unique,
    role varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key (id)
);

Create Table if not exists activities (
    id varchar(255) not null default uuid_generate_v7(),
    name varchar(255) not null,
    type varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    date varchar(255) not null,
    time varchar(255) not null,
    primary key (id)
);

