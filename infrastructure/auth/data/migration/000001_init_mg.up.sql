create table if not exists public.users (
    id serial primary key,
    email text not null,
    password text not null
)






