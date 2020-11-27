-- Drop table
-- DROP TABLE public."data";
CREATE TABLE public."data" (
	id serial NOT NULL,
	"timestamp" timestamp NOT NULL,
	duration int2 NULL,
	-- seconds
	lat int NULL,
	lon int NULL,
	meters int2 NULL,
	operator_description varchar(64),
	operator_id int,
	site varchar(128),
	event_type varchar(8),
	vehicle_id varchar(8),
	zone varchar(128)
);

-- Column comments
COMMENT ON COLUMN public."data".duration IS 'seconds';

drop table public.resumen_mensual;

CREATE TABLE public.resumen_mensual (
	row_id serial NOT NULL,
	period varchar(10) not null,
	accCent int2 not null,
	accelerations int2 not null,
	braking int2 not null,
	identificador TEXT not null,
	infHigh int2 not null,
	infLight int2 not null,
	infMedium int2 not null,
	km numeric(8,2) not null,
	score numeric(6,2) not null,
	scoreAP numeric(6,2) not null,
	scoreRR numeric(6,2) not null,
	scoreZU numeric(6,2) not null
);

alter table public.resumen_mensual owner to leer_insertar_user;