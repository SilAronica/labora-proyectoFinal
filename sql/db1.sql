create table logs (
    id serial,
    dni bigint NOT NULL,
    estado text,
    fecha_creación timestamp with time zone
);

create table wallets(
    id serial,
    dni bigint NOT NULL,
    estado bigint NOT NULL,
    fecha_creación timestamp with time zone
);