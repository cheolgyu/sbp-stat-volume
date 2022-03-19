-- DROP SCHEMA IF EXISTS "stat" CASCADE;
-- CREATE SCHEMA "stat";
DROP TABLE IF EXISTS stat.volume_sum CASCADE;
CREATE TABLE stat.volume_sum(
    row_pk VARCHAR PRIMARY KEY,
    code_id INTEGER NOT NULL,
    unit_type INTEGER NOT NULL,
    yyyy INTEGER,
    unit INTEGER,
    sum_val bigint
);
DROP TABLE IF EXISTS stat.volume_year CASCADE;
-- // 1 :week ,2 : month, 3: q
CREATE TABLE stat.volume_year (
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
DROP TABLE IF EXISTS stat.volume_total CASCADE;
CREATE TABLE stat.volume_total (
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
    CONSTRAINT volume_total_pk PRIMARY KEY (code_id, unit_type)
);


DROP VIEW IF EXISTS "public.view_volume";
CREATE VIEW PUBLIC.view_volume AS
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
	left join stat.volume_total tt on tt.code_id = pc.code_id 
where  tt.max_unit is not null
order by tt.max_percent desc 


DROP VIEW IF EXISTS "view_volume";


CREATE VIEW PUBLIC.view_volume AS
SELECT CMP.CODE_ID,
	CMP.CODE,
	CMP.NAME,
	CMP.CODE_TYPE,
	CMP.MARKET_TYPE,
	CMP.STOP,

	(SELECT NAME
		FROM META.CONFIG
		WHERE ID = CMP.CODE_ID ) AS CODE_TYPE_NAME,

	(SELECT NAME
		FROM META.CONFIG
		WHERE ID = CMP.MARKET_TYPE ) AS MARKET_TYPE_NAME,
	VTB.UNIT_TYPE,
	VTB.YYYY_CNT,
	VTB.MAX_UNIT,
	VTB.MAX_PERCENT,
	VTB.MIN_UNIT,
	VTB.MIN_PERCENT,
	VTB.MAX_RATE,
	VTB.MIN_RATE,
	VTB.MAX_ARR_RATE,
	VTB.MIN_ARR_RATE,
	VTB.AVG_VOL,
	VTB.LAST_UPDATED
FROM stat.volume_total VTB
LEFT JOIN ONLY COMPANY CMP ON VTB.CODE_ID = CMP.CODE_ID
ORDER BY VTB.MAX_PERCENT DESC,
	VTB.YYYY_CNT DESC ;