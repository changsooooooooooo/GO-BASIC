```
func CaptureLoop(){
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i:=0; i<3; i++{
		f[i] = func(){
			fmt.Println(i)
		}
	}

	for i:=0; i<3; i++{
		f[i]()
	}
}

func CaptureLoop2(){
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i:=0; i<3; i++{
		v:=i
		f[v] = func(){
			fmt.Println(v)
		}
	}

	for i:=0; i<3; i++{
		f[i]()
	}
}
```

- 위, 아래 코드의 차이는 포인터 값 복사 인지 아니면 단순히 포인터 자체의 값을 계속 활용하는지의 차이

    -> 함수 literal 에서의 특징!
