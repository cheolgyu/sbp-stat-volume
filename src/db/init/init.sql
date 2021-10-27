-- DROP SCHEMA IF EXISTS "project_trading_volume" CASCADE;
-- CREATE SCHEMA "project_trading_volume";
DROP TABLE IF EXISTS project_trading_volume.tb_sum_by_unit CASCADE;
CREATE TABLE project_trading_volume.tb_sum_by_unit (
    code_id INTEGER NOT NULL,
    unit_type INTEGER NOT NULL,
    unit INTEGER,
    sum_val INTEGER,
    CONSTRAINT tb_sum_by_unit_pk PRIMARY KEY (code_id, unit_type)
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
DROP TABLE IF EXISTS project_trading_volume.tb_result CASCADE;
CREATE TABLE project_trading_volume.tb_result (
    code_id INTEGER NOT NULL,
    unit_type INTEGER NOT NULL,
    high_val INTEGER,
    high_percent INTEGER,
    low_val INTEGER,
    low_percent INTEGER,
    val_range json,
    etc json,
    last_updated INTEGER,
    CONSTRAINT tb_result_pk PRIMARY KEY (code_id, unit_type)
);