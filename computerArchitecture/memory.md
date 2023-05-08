# Memory

## 1. 메모리 구조
lower address

higher address

### (1) text == code
- exe와 같은 파일에 담겨있는 코드는 text에 저장됨

### (2) Data
- 상수 (== const)
- 전역변수


### * 참고 *
- 정적 할당 : 값이 중요한 변수 (ex. a := 2) - 코드가 써지는 순간, 메모리가 할당이 됨
    - array
- 동적 할당 : 메모리 할당이 안됨
    - slice

- 정적 할당이냐 동적 할당이냐에 따라서 heap과 stack으로 나뉘어짐

### (3) heap
- {} 안에 들어있는 동적 할당
- 메모리를 아래로 채움
- 사용자가 선언도 하고 해제도 해야됨
- 사용자가 해제하지 않으면 메모리가 쌓여있음 == 메모리 누수(memory leak)
- 주소를 쓰면 대부분 heap을 사용한다
- 여러 함수에서 사용된다 ? heap

### * Garbage Colletor(GC)
- 메모리 누수를 막기 위해서 쓰지 않는 메모리를 지워줌
- 최근에 사용한 흔적이 있으면 삭제하지 않음

### (4) Stack
- {} 안에 들어있는 정적 할당
- 메모리를 위로 채움
- 선언됐을 때, 메모리가 할당이 됨
- 절({})이 끝났을 때 할당된 메모리가 해제가 됨

### * GO에서의 동적 할당
- 동적 할당 생성 : make(type, length) 사용
- 동적 할당 해제 : 없애지 않음 - garbage collector가 알아서 해줌
- 동적 할당 해제(최선의 방법)
  - ex. arr = make([]int, 4)
  - arr = nil
- garbage collector가 돌면 프로그램이 엄청 느려짐
- garbage collector가 자주 돈다는 것은 좋지 않은 것


### * 최신 GO에서 Heap, Stack 어떻게 할당하는가?
- 사용자가 결정하지만 마지막에는 compiler가 heap, stack을 할당한다
```go
package main

import "fmt"

func main() {
	size := 8
	a := make(map[string]int, size)
	
	fmt.Print(a)
	//a = map[]
}
```
- a는 make를 사용했으니까 heap으로 할당한다고 생각하는 것이 기본
- 하지만 go compiler는 a가 func temp()에서만 사용하기 때문에 stack으로 할당
- 때문에 go compiler는 밖에서도 사용하면 heap에 할당하고, 안에서만 사용하면 stack으로 할당


- heap에는 포인터만 들어올 수 있음
- 포인터는 heap도 될 수 있고, stack도 될 수 있다

- garbage issue 때문
  - stack address가 높아서 금방 금방 빠짐
  - heap은 누군가 해제해줄 때까지 기다림
  - 최대한 stack에 쌓는 것이 성능에 좋음

![](https://user-images.githubusercontent.com/116709146/235143286-af1f6f02-10b6-4362-9858-6380ede0e542.png)
- 하위함수에게 메모리를 전달할 때는 stack으로 메모리의 비용(==할당비용)을 보존함
- 상위함수에게 메모리를 전달할 때는 heap으로 할당