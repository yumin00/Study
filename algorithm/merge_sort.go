package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
N개의 수가 주어졌을 때, 이를 오름차순으로 정렬하는 프로그램을 작성하시오.

입력
첫째 줄에 수의 개수 N(1 ≤ N ≤ 1,000)이 주어진다. 둘째 줄부터 N개의 줄에는 수가 주어진다. 이 수는 절댓값이 1,000보다 작거나 같은 정수이다. 수는 중복되지 않는다.

출력
첫째 줄부터 N개의 줄에 오름차순으로 정렬한 결과를 한 줄에 하나씩 출력한다.
*/

/*
[분할]
- 균등하게 잘라줌
- 배열의 중간을 기준으로 잡아서 두 가지 그룹으로 나눔
- 두 가지 그룹에서 각각 중간을 기준으로 잡아서 다시 두 가지 그룹으로 나눔
- 그룹이 없어질 때까지 분할하기
- 시간 : logN

[정복]
- 합병을 하면서 정렬
- 시간 : N
*/

func main() {
	var N int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &arr[i])
	}
	result := divide(arr)

	for i := 0; i < N; i++ {
		fmt.Println(result[i])
	}
}

func divide(arr []int) []int { // 재귀를 통해서 좌우배열을 자른다.
	if len(arr) == 1 {
		return arr
	}

	middle := len(arr) / 2
	return merge(divide(arr[middle:]), divide(arr[:middle]))
}

func merge(arr1, arr2 []int) []int { // 잘린 배열을 통합하면서 정렬한다.
	arr := append(arr1, arr2...)

	sort.Ints(arr)
	return arr

}
