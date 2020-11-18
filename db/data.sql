-- Drop table

-- DROP TABLE public."data";

CREATE TABLE public."data" (
	id serial NOT NULL,
	"timestamp" timestamp NOT NULL,
	duration int2 NULL, -- seconds
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