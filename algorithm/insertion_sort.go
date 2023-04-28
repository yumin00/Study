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
- 두 번째 idx부터 시작
- -1 idx와 비교해서 크면 오른쪽으로..
- 시간(worst) : 오름차순으로 하고 싶은데, 내림차순으로 되어있는 경우 - n(n-1)/2 => O(n^2)
- 시간(best) : O(n) - 뒤에를 먼저 만들어놓으면 끝까지 돌 필요가 없음 : 이미 한 것은 신경쓰지 말자 / 이미 정렬됐으면 보장하고 break
*/

func main() {
	var N int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &arr[i])
	}

	for i := 0; i < N-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	for i := 0; i < N; i++ {
		fmt.Println(arr[i])
	}
}
