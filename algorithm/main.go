package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
N개의 수가 주어졌을 때, 이를 오름차순으로 정렬하는 프로그램을 작성하시오.

입력
첫째 줄에 수의 개수 N(1 ≤ N ≤ 1,000)이 주어진다. 둘째 줄부터 N개의 줄에는 수가 주어진다. 이 수는 절댓값이 1,000보다 작거나 같은 정수이다. 수는 중복되지 않는다.

출력
첫째 줄부터 N개의 줄에 오름차순으로 정렬한 결과를 한 줄에 하나씩 출력한다.
*/

/*
<우선순위 큐>
- min heap
  - 우선순위가 높으면 오름차순
- max heap
  - 우선수위가 높으면 내림차순

- head를 한 번 빼는 것 : O(1)
- 이진트리에서 숫자 정렬(head 채우기) : O(logN) -> N(노드)만[큼 반복

-> O(NlogN)
*/

func main() {
	var N int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &arr[i])
	}
}

/*
예제 입력
5
5
2
3
4
1
*/
