# Computer Architecture
<메모리 계층 구조>

![img_3.png](https://user-images.githubusercontent.com/130362583/233977029-5697b627-dd54-4f69-97d5-7b4de601b4da.png)

1. remote secondary stroage
- 네트워크를 통해서
- 속도가 느리고 공간이 커짐

2. tradtional disk
- hard disk

3. flash disk
- usb

4. main memory
- 우리가 접근할 수 있는 최대

5. caches
- main memory가 cpu로 가기 위한 통로(?)
- l1 : 연산자 (많이 쓰는 것들을 먼저 / 속도가 빠른 것들을 먼저)
- l2 : 피연산자
- l3 :


*** buffer = caches와 똑같은 개념
- buffer : 대기실 같은 존재
- fScan에는 buffer가 있어서 속도가 훨씬 빠름

6. registers
- 속도가 빠르고 공간이 작아짐

### cache 알고리즘

[hardware]

cpu - cache - main memory

[software]

cpu - memory - disk

--> cache 알고리즘 = 캐싱한다


- 네트워크와의 입출력(remote secondary storage) 없이 내 메모리(main memory에서 꺼내서 준다 => 속도가 훨씬 빠름
- 같은 걸 물어보면 또 찾아서 보여주는 게 아니라 메모리에 저장해놓고 물어보면 그걸 그대로 뱉어준다.

### data 지역성
시간 / 공간 / 순차

참조한다 == 접근하다 == 가져간다
1. 시간 지역성 (빈도)
- 한번 참조한 메모리는 또 참조할 가능성이 있다.

2. 공간 지역성 
- 참조한 메모리의 주변을 참조할 가능성이 있다.

3. 순차 지역성
- 메모리는 순차적으로 넣는다.
- 참조한 메모리의 다음을 실행할 가능성이 있다.

-> 이 중 1개, 2개를 가지고 알고리즘을 짜면 캐싱 알고리즘을 만들 수 있다.

-> redis : cache할 때 많이 사용함 (why? 미리 저장해서 뱉어주는 개념)

### cache 알고리즘 구현

1. LRU(시간 기준) : 가장 오랫동안 참조되지 않은 것을 뺀다 / 가장 최근에 참조한 것을 냅둔다

memory - cache

miss(아주 느림) : cache에 없는 것을 부르는 것 == cache에 접근하는 것이 아니라 memory에 접근하는 것과 같다

hit(아주 빠름) : cache에 있는 것을 뱉어주는 것

가장 적은 용량을 가지고 최대한 많은 hit를 치는 것이 목표

빠른 퍼포먼스를 위해서 어떤 곳에 저장해서 뱉어주는 것 == 캐싱한다

강력 캐쉬 비우기 : flush

2. (빈도수 기준) : 가장 조금 참조된 것을 뺀다


