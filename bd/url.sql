create table Url (
    id serial primary key ,
    name text not null,
    method text not null,
     is_active boolean default true,
                created_at timestamp default now(),
                updated_at timestamp default now ()
);

