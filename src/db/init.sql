-- DROP SCHEMA IF EXISTS "project_trading_volume" CASCADE;
-- CREATE SCHEMA "project_trading_volume";
DROP TABLE IF EXISTS project_trading_volume.tb_sum CASCADE;
CREATE TABLE project_trading_volume.tb_sum(
    row_pk VARCHAR PRIMARY KEY,
    code_id INTEGER NOT NULL,
    unit_type INTEGER NOT NULL,
    yyyy INTEGER,
    unit INTEGER,
    sum_val bigint
);
DROP TABLE IF EXISTS project_trading_volume.tb_year CASCADE;
-- // 1 :week ,2 : month, 3: q
CREATE TABLE project_trading_volume.tb_year (
    code_id INTEGER NOT NULL,
    unit_type INTEGER NOT NULL,
    yyyy INTEGER,
    max_unit INTEGER,
    max_unit_arr INTEGER [ ],
    min_unit INTEGER,
    min_unit_arr INTEGER [ ],
    avg_vol bigint,
    rate JSONB,
    CONSTRAINT tb_year_pk PRIMARY KEY (code_id, unit_type, yyyy)
);
DROP TABLE IF EXISTS project_trading_volume.tb_total CASCADE;
CREATE TABLE project_trading_volume.tb_total (
    code_id INTEGER NOT NULL,
    unit_type INTEGER NOT NULL,
    yyyy_cnt INTEGER,
    max_unit INTEGER,
    max_percent DOUBLE PRECISION,
    min_unit INTEGER,
    min_percent DOUBLE PRECISION,
    max_rate JSONB,
    min_rate JSONB,
    max_arr_rate JSONB,
    min_arr_rate JSONB,
    avg_vol bigint,
    last_updated INT,
    CONSTRAINT tb_total_pk PRIMARY KEY (code_id, unit_type)
);


DROP VIEW IF EXISTS "public.view_trading_volume";
CREATE VIEW PUBLIC.view_trading_volume AS
SELECT tt.CODE_id as mp_code,
    tt.unit_type ,
    --(select name from meta.config where id = tt.unit_type ) as unit_type_name,
	pc.code_id,
	pc.code,
	pc.name,
	pc.code_type,
	(select name from meta.config where id = pc.code_type ) as code_type_name,
	pc.market_type,
	(select name from meta.config where id = pc.market_type ) as market_type_name,
	tt.max_unit,
	tt.max_percent,
	tt.min_unit,
	tt.min_percent,
    tt.max_rate,
    tt.min_rate,
    tt.avg_vol,
    tt.last_updated
FROM only public.company pc
	left join project_trading_volume.tb_total tt on tt.code_id = pc.code_id 
where  tt.max_unit is not null
order by tt.max_percent desc 