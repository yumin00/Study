# Algorithm
## Sort Alrgorithm
== 오름차순 내림차순 정렬

### 1. 버블
- 인접한 것들 끼리 비교해서 큰 값을 오른족으로 작은 값을 왼쪽으로
- 한 번을 반복을 하면 맨 오른쪽에 제일 큰 값이 남음
- 제일 큰 값을 제외하고 다시 반복
- 시간 : n(n-1)/2 => O(n^2)

### 2. 선택 (자리 선택)
- 가장 작은 값을 찾아서 맨 앞에 놓겠다!
- 첫 번째 값과 두 번째 값을 비교해서 작은 값을 min으로 놓고 idx 잡기
- min 값과 세 번째 값을 비교해서 작은 값은 min으로 놓고 idx 잡기
- 한 번 다 돌고 min 값을 맨 앞에 놓기
- 시간 : n(n-1)/2 => O(n^2)

### 3. 삽입
- 두 번째 idx부터 시작
- -1 idx와 비교해서 크면 오른쪽으로..
- 시간(worst) : 오름차순으로 하고 싶은데, 내림차순으로 되어있는 경우 - n(n-1)/2 => O(n^2)
- 시간(best) : O(n) - 뒤에를 먼저 만들어놓으면 끝까지 돌 필요가 없음
    
### 4. 병합


### 5. 힙


### 6. 퀵 == goland sort


### 7. 기수