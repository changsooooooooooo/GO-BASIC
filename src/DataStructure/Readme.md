```
m:=make(map[string]string, 5)

m["1"] = "seoul"
m["2"] = "seoul"
m["3"] = "seoul"
m["4"] = "seoul"
m["5"] = "seoul"
m["6"] = "seoul"
m["7"] = "seoul"
m["8"] = "seoul"
m["9"] = "seoul"
m["0"] = "seoul"

for k, v := range m{
    fmt.Println(k, v)
}

```

- 이렇게 해도 추가가되는 이유 : slice 로 작동하기 때문!


```

func formArr(arr [M]int){
	arr[2] = 10
}

arr:=[M]int{}
```

- 이렇게 const size 로 입력 받아도 동적배열로서 알아서 잡아주고 array 형식을 유지

    -> Because, 이렇게 해도 arr[2] value 가 바뀌지 않음
