func lengthOfLongestSubstring(s string) int {
	log,count,max:=make(map[uint8]int),0,0
	var temp,i int
	var flag bool
	for i=0;i<len(s);i++ {
		temp,flag=log[s[i]]
		log[s[i]]=i
		if flag==true {
			if max<count {
				max=count
			}
			count=0
			log=make(map[uint8]int)
			i=temp
		}
		else {
			count++
		}
	}
	if max<count {
		max=count
	}
	return max
}