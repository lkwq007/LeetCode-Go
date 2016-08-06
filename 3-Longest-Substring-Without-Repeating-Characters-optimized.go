func lengthOfLongestSubstring(s string) int {
	log,front,max:=make(map[uint8]int),0,0
	var temp,i int
	var flag bool
	for i=0;i<len(s);i++ {
		temp,flag=log[s[i]]
		log[s[i]]=i
		if flag==true {
			if max<i-front {
				max=i-front
			}
			front=temp+1
			for i,v:=range log {
				if v<front {
					delete(log,i)
				}
			}
		}
	}
	if max<i-front {
		max=i-front
	}
	return max
}