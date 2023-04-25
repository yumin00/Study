# Go Lang

## 1. 입출력

(1) 여러 줄을 읽을 때

**fmt.Scan**

(2) 대량의 입력을 받을 때

많은 줄을 읽어야한다면, fmt.Scan은 버퍼링을 하지 않아 느리다.

bufio 패키지와 **fmt.Fscan**을 활용

```go
r := bufio.NewReader(os.Stdin)
fmt.Fscan(r, &v)
```

