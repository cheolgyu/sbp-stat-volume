/*

파일명: 단타
단타용 종목 찾기 
pubapi query 참고.

*/

/*
	example: 
	=====================================================================================
	=====================================================================================
	=====================================================================================
	v.0.0.1
	select  	
		t.id 
		,t.code
		,t.name
		,t.market_type 
		,(select name from meta.config where id = t.market_type ) as market_name
		,t.stop
		,t.avg_l2h
		,t.avg_o2c
		,t.std_l2h
		,t.std_o2c
	from
	(
		SELECT  
			pc.id
			, pc.code
			, pc.market_type
				
				,pc.name
			,pc.stop
			, round(AVG(hp.L2H),2) AS AVG_L2H
			, round(STDDEV(hp.L2H),2) AS STD_L2H
			, round(AVG(hp.o2c),2) AS AVG_o2c
			, round(STDDEV(hp.o2c),2) AS STD_o2c
		
			, count(hp.L2H) as cnt_l2h
			, count(hp.o2c) as cnt_o2c
			, array_agg(hp.L2H) 
			, array_agg(hp.lp) 
			, array_agg(hp.hp) 
		from 
			public.tb_code pc
			left join hist.PRICE hp  on hp.code_id = pc.id 
			
		where 1=1
		and pc.code_type = 1
		and hp. dt >= (select min(dt) from (select dt from meta.opening  order by dt desc limit 10 )t )
		and hp.vol != 0
		group by pc.id ,pc.code,pc.market_type,pc.name,pc.stop
		having count(hp.L2H) = 10
	)t 

	=====================================================================================
	=====================================================================================
	=====================================================================================
	v.0.0.2

	with tbo as (
		select min(dt) from (select dt from meta.opening  order by dt desc limit 100 )t
	), tb as (
		SELECT  
			pc.id
			, round(AVG(hp.L2H),2) AS AVG_L2H
			, round(STDDEV(hp.L2H),2) AS STD_L2H
			, round(AVG(hp.o2c),2) AS AVG_o2c
			, round(STDDEV(hp.o2c),2) AS STD_o2c
		
			, count(hp.L2H) as cnt_l2h
			, count(hp.o2c) as cnt_o2c
			, array_agg(hp.L2H) 
			, array_agg(hp.lp) 
			, array_agg(hp.hp) 
		from 
				tbo, public.tb_code pc 
				--left join meta.config mc on pc.market_type = mc.id
				left join hist.PRICE hp  on hp.code_id = pc.id 
			
		where 1=1
		and pc.code_type = 1
		and hp. dt >= tbo.min --(select min(dt) from (select dt from meta.opening  order by dt desc limit 10 )t )
		and hp.vol != 0
		group by pc.id --,pc.code
		having count(hp.L2H) = 100
	)


	select  	
		tb.id 
		,pc.code
		,pc.name
		,pc.market_type 
		,mc.name as market_name
		,pc.stop
		,tb.avg_l2h
		,tb.avg_o2c
		,tb.std_l2h
		,tb.std_o2c
	from tb 
	left join  public.tb_code pc on tb.id = pc.id 
	left join meta.config mc on pc.market_type = mc.id

*/

CREATE OR REPLACE FUNCTION  project.func_day_trading(
	inp_term 		integer 
	,inp_limit 		integer 
	,inp_offset 	integer
	,inp_sort 		VARCHAR 
	,inp_desc 		VARCHAR
	,inp_market_arr integer[]
	)
 RETURNS TABLE (
	 code_id integer  
	 ,code VARCHAR  
	 ,name VARCHAR 
	 ,market_type integer 
	 ,market_name  VARCHAR 
	 ,stop boolean
	 ,avg_l2h numeric
	 ,avg_o2c numeric
	 ,std_l2h numeric
	 ,std_o2c numeric
	 ) AS
 $$
declare
    r record;
	x1 numeric;
	x2 numeric;
	y1 numeric;
	y2 numeric;
	count integer := 0;
BEGIN

RETURN QUERY 
 -- code  
--------------------
EXECUTE format('
with tbo as (
	select min(dt) from (select dt from meta.opening  order by dt desc limit $1 )t
), tb as (
	SELECT  
		  pc.code_id
		  , round(AVG(hp.L2H),2) AS AVG_L2H
		  , round(STDDEV(hp.L2H),2) AS STD_L2H
		  , round(AVG(hp.o2c),2) AS AVG_o2c
		  , round(STDDEV(hp.o2c),2) AS STD_o2c
		  , count(hp.L2H) as cnt_l2h
		  , count(hp.o2c) as cnt_o2c
		  , array_agg(hp.L2H) 
		  , array_agg(hp.lp) 
		  , array_agg(hp.hp) 
	   from 
		  	tbo, only public.company pc 
		  	left join hist.PRICE hp  on hp.code_id = pc.code_id 
		  
	   where 1=1
       and pc.code_type = 1
	   and hp. dt >= tbo.min 
	   and hp.vol != 0
	   group by pc.code_id 
	   having count(hp.L2H) = $1
)


select  	
	  tb.code_id 
	  ,pc.code
	  ,pc.name
	  ,pc.market_type 
	  ,mc.name as market_name
	  ,pc.stop
	  ,tb.avg_l2h
	  ,tb.avg_o2c
	  ,tb.std_l2h
	  ,tb.std_o2c
  from tb 
  left join  only public.company pc on tb.code_id = pc.code_id 
  left join meta.config mc on pc.market_type = mc.id
  where 1=1
  and market_type =any ($6)
  order by %s %s
  limit $2  offset $3
  ', inp_sort, inp_desc )using 
  inp_term 		
,inp_limit 		
,inp_offset 	
,inp_sort 		
,inp_desc 		
,inp_market_arr  
  ;
--------------------


-------------
END;
$$
LANGUAGE plpgsql;

	-- inp_term 		integer 
	-- ,inp_limit 		integer 
	-- ,inp_offset 	integer
	-- ,inp_sort 		VARCHAR 
	-- ,inp_desc 		VARCHAR
	-- ,inp_market_arr integer[]
select * from project.func_day_trading(10,10,1,'avg_l2h','desc',ARRAY[	7,8,9]);

