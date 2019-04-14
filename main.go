package main

func twoSun(nums []int,target int) []int{
	makars:=make(map[int]int,len(nums))
	for i:=0;i<len(nums) ;i++  {
		num:=nums[i]
		numTarget:=target-num
		index ,ok:=makars[numTarget]
		if ok{
			return  []int{i ,index}
		}
		makars[num]=i;
	}
	return  nil;
}
func main()  {

}
