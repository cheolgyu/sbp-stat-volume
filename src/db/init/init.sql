-- DROP SCHEMA IF EXISTS "project_trading_volume" CASCADE;
-- CREATE SCHEMA "project_trading_volume";
DROP TABLE IF EXISTS project_trading_volume.tb_sum CASCADE;
CREATE TABLE project_trading_volume.tb_sum(
    row_pk varchar PRIMARY KEY,
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
    percent JSONB,
    CONSTRAINT tb_year_pk PRIMARY KEY (code_id, unit_type, yyyy)
);
DROP TABLE IF EXISTS project_trading_volume.tb_total CASCADE;
CREATE TABLE project_trading_volume.tb_total (
    code_id INTEGER NOT NULL,
    unit_type INTEGER NOT NULL,
    max_unit INTEGER,
    max_unit_arr INTEGER [ ],
    min_unit INTEGER,
    min_unit_arr INTEGER [ ],
    avg_vol bigint,
    percent JSONB,
    last_updated INTEGER,
    CONSTRAINT tb_total_pk PRIMARY KEY (code_id, unit_type)
);