create table logs (
    id serial,
    dni bigint NOT NULL,
    estado text,
    fecha_creacion timestamp with time zone
);

create table wallets(
    id serial,
    dni bigint NOT NULL,
    pais_id text DEFAULT 'PE' ::text NOT NULL,
    fecha_creacion timestamp with time zone
);


insert into logs (dni, estado, fecha_creación) values (9483626273, 'PENDIENTE', now());
insert into logs (dni, estado, fecha_creación) values (35426251738636, 'RECHAZADO', now();)
select * from logs;