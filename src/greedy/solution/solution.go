package solution

func Solution(n int, lost []int, reserve []int) int {

	status := make([]int, n+2)

	for _, v := range reserve{
		status[v]=1
	}

	for _, v := range lost{
		status[v]-=1
	}

	for i, v := range status{
		if i==0{
			continue
		}
		if v<0 && 0<status[i-1]{
			status[i] = 0
			status[i-1]-=1
			continue
		}
		if v<0 && 0<status[i+1]{
			status[i] = 0
			status[i+1]-=1
			continue
		}
	}

	answer:=0
	for i:=1; i<n+1; i++{
		if -1<status[i]{
			answer += 1
		}
	}
	return answer
}
