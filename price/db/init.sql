DROP TABLE IF EXISTS project.tb_52_weeks CASCADE;
CREATE TABLE IF NOT EXISTS project.tb_52_weeks (
    "code_id" INTEGER NOT NULL REFERENCES "meta"."code"(id),
    "price_type" INTEGER NOT NULL REFERENCES "meta"."config"(id),
    "p1x_left" INTEGER NOT NULL,
  
    "p1x" INTEGER,
    "p1y" numeric(20, 2),
    "p2x" INTEGER,
    "p2y" numeric(20, 2),
    "p3x" INTEGER,
    "p3y" numeric(20, 2),
    
    "p3_type" char(1) NOT NULL,
    "p32y_percent" numeric(10, 2)
);
COMMENT ON COLUMN "project"."tb_52_weeks"."code_id" IS '코드ID';
COMMENT ON COLUMN "project"."tb_52_weeks"."price_type" IS '가격종류 코드 값';
COMMENT ON COLUMN "project"."tb_52_weeks"."p1x_left" IS 'p1.x가 왼쪽으로 움직인 값, 일수';

COMMENT ON COLUMN "project"."tb_52_weeks"."p1x" IS '과거점 일자';
COMMENT ON COLUMN "project"."tb_52_weeks"."p1y" IS '과거점 값';
COMMENT ON COLUMN "project"."tb_52_weeks"."p2x" IS '현재점 일자';
COMMENT ON COLUMN "project"."tb_52_weeks"."p2y" IS '현재점 값';

COMMENT ON COLUMN "project"."tb_52_weeks"."p3x" IS '찾은점 일자';
COMMENT ON COLUMN "project"."tb_52_weeks"."p3y" IS '찾은점 값';

COMMENT ON COLUMN "project"."tb_52_weeks"."p3_type" IS 'p3 타입: 최고=H,최저점=L';
COMMENT ON COLUMN "project"."tb_52_weeks"."p32y_percent" IS '찾은점과 현재점의 차이 퍼센트';
----------------------------------------------