```
for n:=range ch{
		fmt.Println(n)
		time.Sleep(time.Second)
	}
```

- 단순히 channel close 없이 이렇게 짜면 무한루프 돌면서 데이터가 들어오길 계속 기다리게 된다.


```
ctx, cancel := context.WithCancel(context.Background())
ctx = context.WithValue(ctx, "Name", "Thomas")
```

- 이런식으로 상위 context 에 할당해서 여러 특성을 줄 수 있다.
