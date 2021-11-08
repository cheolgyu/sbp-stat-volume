# 피크월(거래량 기준)

## 목표
   이것의 목표는 손오공과 같이 특정 월에 몰리는 거래량을 가진 종목을 찾아서 매매할떄 참고하기 위함이다.   
   연도별 분기, 연도별 월, 연도별 주 단위로 구분하여 거래량이 몰리는 연도별 단위의 피크점를 찾은 후에 전체연도의 단위의 피크점과 연도별 단위의 피크점의 백분율을 구하여 백분율 순위 출력하기.

---

```
할것
 META.CONFIG에 코드값 추가 후 연결 짓기. 
   현재 1. 주 2.월 3. 분기 이지만
   CONFIG에서 가져와서 쓰기? : 프론트엔드는 CONFIG에서 가져와 쓰니깐 맞추면 좋지?
```

## 주요 용어
* 단위 구분   
  * 연도별 분기 quarter by year
  * 연도별 월 month by year
  * 연도별 주 weeks by year
* 단위의 피크점 unit peak

---



## 피크월 계산 방법

  1. 장열림 테이블에 거래일이 해당연도의 몇주, 몇 분기에 해당하는지 구한다.
  2. 종목의 같은 연도의 거래일을 (주/월/분기)단위로 묶어 거래량 합을 구한다.
     1. 종목의 해당연도의 일자에 해당하는 주가 몇번째 주인지 구하기
     2. 해당연도의 단위별 거래량 합 구하기
  3. 같은연도의 (주/월/분기)단위별 합의 값으로 같은연도의 단위별 평균값, 평균이하값들, 평균 이상값들, 최저, 최고 값을 구한다.
     1. 평균이하값들, 평균 이상값들에서 연결성의 존재 여부를 찾는다.

  4. 전체년도의 주단위 ***빈도분석*** : 전체연도로 놓고 연도별 최고거래량인 주, 최소거래량인 주, 평균거래량보다 큰 목록의 주, 평균거래량 보다 작은 목록의 주가 몇 퍼센트를 찾는다.
     1. 연도별 종목의 주단위 거래량의 빈도 분석

   추가 ? ) 특정연도의 단위별 최고가를 찾는다고 한다면 최고가 1등과 2등의 차이가 1밖에 없다면 ?? 그래서 배열이 필요함.
      전체연도에서 각자의 연도에서 뽑은 최고와 최저값이 각자의 연도별로 다르겠지? 그래서 각자연도별 최고저값이  전체 연도에서 몇퍼센트를 차지하는지 구하기?
   


---



  

## 실행계획

      빈도가 의미 있는 정보가 되기 위한 필요한 조건은 데이터들이 다 채워져야 한다.
      최소 단위의 데이터가 다 채워 지면 실행시키는게 맞다 .
      최소 단위의 데이터가 채워지기 전에 실행한다면 퍼센트값만 조금더 떨어질뿐.

   - 거래량단위 빈도 분석 티커 생성
   - 데이터 쓰기 티커 종료 후 실행
   - 금일이 단위(주/월)의 마지막인지 체크.
   - 마지막이라면 프로그램 실행.  stock-project-이름



---
      

      meta.code 에서 코드목록과 rb_total을 조인 후 last_updated 칼럼을 가치 조회한 후 
      last_updated를 시작일 이후의 데이터를 조회 하는데 meta.opening을 조인하여 가격정보에 주,분기 칼럼을 붙여 조회한다.
         model: hist.price + meta.opening 

      조회한 가격데이터 목록을 반복문 돌려서 년도별 주, 월, 분기 배열에 각각 거래량의 합을 누적한다.
      
      거래량의 합을 구한 연도별 데이터는 집계테이블에 저장하기 위해 연도별 단위의 최고,최저,최저고의 목록 퍼센트등을 계산 한다.
      계산이 완료되면 집계테이블에 저장한다.
      
      집계테이블에서 데이터를 종목별로 가져온다.
      종목별로 가져온 데이터를 빈도에 따른 퍼센트를 계산한다.
      계산이 완료되면 결과테이블에 저장한다.
      
      
         방법: 
         //   코드에 해당하는 price의 min과 max 일자를 조회후 연도별 배열을 만들고
          //  연도 배열을 반복문돌려 연도에 해당하는 데이터를 일자순으로 소트하여 반환 받은 다음.

               
               
               집계공식 처리하기
                  집계 공식 넣기 전 처리
                     방법:
                        일자별 가격목록 데이터를 요구 단위(주,월,분기)에 맞게 끊어서 단위를 키 일자별 가격목록을 값으로 맵을 만든 후에.
                        반복문을 돌려 가격목록배열에 거래량의 합을 계산 후에 단위별 거래량의 합을  테이블에 저장 할까? 일단 저장 후에 삭제 하자 안쓰니깐
                        저장한다음. 

                  집계공식에 넣어서 처리 
                     방법:
                        코드값과 단위 값에 해당하는 주별 거래량의 합을 조회 해 온 후에 
                        연도별로 평균가, 최고가, 최소가, 최고가범위, 최소가범위를 계산 한 후에 저장한다.


               결과테이블에 저장하기
                  결과테이블 실행 전
                     방법:
                        연도별 집계처리가 완료된 코드값과 단위 값을 받아 전체 연도에 해당하는 집계 테이블을 조회한다.
                        조회한 결과에서 프로그램으로 값과 퍼센트를 구하고
                        저장한다.
            
            모드 결과 테이블저장이 끝나면 단위결 거래량의 합 테이블의 데이터를 지운다.
   
  ---
      dao
         전처리
            코드목록 조회
            처리기간 범위 만들기
               시작일 찾기 : 코드별 마지막 실행 일자 조회(rb_total 의 칼럼 )
               기간 만들기:
                  마지막 실행 일자 기준 이후의 일자 중 min과 max 값 조회
                  golang- main과 max 의 데이터 연도별로 나누기.
         
         가격합테이블 데이터 만들기
            연도별 가격 데이터 조회 루프
               golang- 가격일자를 단위구분으로 나누어 저장하기. map에다가
               golang-  나누어 저장된 거래량의 합 을 구하기. 
               구한 합을 tb_sum에 저장
         
         집계기능
            tb_sum에 데이터 조회
            golang- 연도별 평균가, 최고가,최소가, 최고저의 범위 계산
            계산 결과 저장를 tb_year에 저장

         결과
            tb_year에 종목의 전체연도 데이터 조회
            golang- 값과 퍼센트 만들기
            값과 퍼센트를 rb_total에 저장


## 실행계획2 - 데이터가 매일 들어오는것 처리
      일자별 가격 데이터들이 모여 tb_sum의 합 데이터가 되는데
      2주차에 4개 일자가 있다고 한다면
      2번쨰 일짜까지 처리한 후 다음날 장마감 후 3번째 일자를 처리한다고 가정해보면
      해당주에 해당하는 tb_sum 데이터를 삭제한다. 
            현재 upsert 때문에 pk를 여러개 설치했더니 25mb 테이블에 23mb 인텍스 사이즈가 잡혀 있다. pk 줄이고 
      특정일자가 포함된 단위의 데이터들을 삭제한다. 
            주단위는 해당 주이고
            월 단위는 해당 월 이다.
            분기단위는 해당 분기이다
         그러면 틀정일 이전의 단위에 포함되는 가격 데이터도 추가로 조회해야한다. 먼가 더 복잡하게 하는것 같다.
            
            *** tb_sum 데이터 처리 ***
            1.
            tb_sum에 몇주차 데이터가 까지 저장 했나 lasted 칼럼을 만든 후에 tb_sum에서 함수를 태워 더하기? 아님
            해당 주차의 데이터가 있다면 누적으로 더하기를 하고 없다면 새로운 줄을 추가한다. 
               => 이건 같은날 2번 실행시 일자의 거래량이 값이 업데이트 된 경우 반영이 안된다.
            
            2.
               일자을 배열로 가지고 있을까?
                  => 이건 같은날 2번 실행시 일자의 거래량이 값이 업데이트 된 경우 반영이 안된다.

            3.   sql 함수로 처리
               가격목록을 조회해 오지마 말고 특정일을 주면 포함된 데이터를 다시 더하는 sql 함수로 만들자.
               코드의 마지막 계산일을 tb_total에 있으니 코드를 주면 
               tb_total의 last_updated 보다 +1 큰  일자를 조회 후에 일자가 포함된 주, 월, 분기 를 
               hist.price에서 조회해서 합을구 하여 update나 insert 해야하는데 
               => 가격목록이 1개라면 상관 없지만 여러개라면   
                  시작일, 마지막일 에 포함된 주,월,분기 데이터를 upsert 해야한다.
                  hist.price와 meta.opening을 조인후 meta.opening의 
                  yyy, week 별 합 그룹
                  yy,, month 별 합 그룹 
                  yy,, q 별 합 그룹 
                  테이블을 만들어 upsert 시켜야된다. 
                  월까지는 데이터량이 작지만 분기 부터는 달라질것. 
               ==> sql 함수로 처리는 코드가 더 복잡해진다.
            
            4. 아짱난다. 삭제 후 재 등록?
            5. 코드별 년도 주차,월,분기 값을 넘겨주면 해당 데이터 삭제하고 새로 select 후 집어 넣기
            ==> 6. upsert로 기존에 있는 항목은 sum_val 누적으로 더해서 처리하기.
               ==> 테스트 사항.
                  가격데이터가 0-10까지 저장한 상태에서 11~ 15일 데이터가 들올경우 
                  11~15 에 해당 되는 단위들의 sum 값이 반영되었나 확인.
            
            

               


               

      tb_sum 다음이 tb_year 이다.
      tb_year에서 새로온 데이터를 반영하려면?
      
      
      *** tb_year 데이터 처리 ***
      마지막 계산일 이 포함된 년도와 이후 년도에 해당하는  tb_sum을 조회한다.
      계산한다. upsert한다.

      *** tb_total 데이터 처리 ***
      코드의 전체년도에 해당하는 데이터를 tb_year에서 조회한다.
      계산한다. 저장한다.
      



      특정일이 포함된 데이터를 삭제안한다 왜냐면 3줄 뿐이니 upsert로 처리한다.
         해당 연도의 주, 월, 분기 줄이다. 
      tb_sum에서 특정 연도의 주,월, 분기 데이터를 조회한다.
      다시 계산한다. 저장한다.


      tb_total 에서는 
      특정일이 포함된 데이터를 삭제안한다 왜냐면 3줄 뿐이니 upsert로 처리한다.

         

            





         
  



## db

      스키마 
         project_trading_volume
      테이블
         rb_total
         tb_year
----

   1. 빈도 테이블  rb_total
      | 종목id | 최고거래량인 주 | 최고거래량인 주의 퍼센트 | 최소거래량인 주 | 최소거래량인 주의 퍼센트 | 평균 최소 범위 | 평균 최소 범위의 퍼센트 | 평균 최대 범위 | 평균 최대 범위의 퍼센트 | 마지막 계산일 |
      | ------ | --------------- | ------------------------ | --------------- | ------------------------ | -------------- | ----------------------- | -------------- | ----------------------- | ------------- |

   2. 연도별 집계 테이블 tb_year
      | 종목id | 단위(주/분기/월) | 연도 | 최고거래량인 단위값(주/분기/월) | 최소거래량인 단위값(주/분기/월) | 평균 이상의 단위값(주/분기/월) 목록 | 평균 이하의 단위값(주/분기/월) 목록 |
      | ------ | ---------------- | ---- | ------------------------------- | ------------------------------- | ----------------------------------- | ----------------------------------- |

   3. 주별 테이블  tb_sum
      1. 저장소 낭비
         1. 2000(종목수) * (종목별)전체년도 수 * 52(주) 의 줄 발생
         2. rdb말고 nosql db로 한다면?
         
      2. 연도 파티션 테이블
      
      | 종목id | 연도 | 주  | 거래량의 합 |
      | ------ | ---- | --- | ----------- |

meta.opening 에 주, 분기 칼럼 추가 및 계산.

---







## 계산 범위
현재 입력된 달은 매일 갱신.   
매일 테이블을 갱신 시켜줌. 해당년도의 달만.    
하루에 2번 실행 될 되도 문제없게. 
   

---


## 실행시작
가격다 넣고 반등 계산하고.   
종목별 마지막 가격일 조회 해서 가격일 갱신하면 이전 피크월은 계산이 안됨.   

피크월 테이블 기준으로 시작하기 vs 가격 테이블 기준 시작하기   
종목의 마지막 년도 피크월이 있으면? ( 년도포함월 칼럼추가. )    

---