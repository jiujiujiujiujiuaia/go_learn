func minMeetingRooms(intervals [][]int) int {
	return merge(intervals)
 
 }
 
 type Interval struct {
	 Start int
	 End   int
 }
 
 type Intervals []Interval
 
 func (Interval Intervals) Len() int {
	 return len(Interval)
 }
 
 func (Interval Intervals) Swap(i, j int) {
	 Interval[i], Interval[j] = Interval[j], Interval[i]
 }
 
 func (Interval Intervals) Less(i, j int) bool {
	 return Interval[i].Start < Interval[j].Start
 }
 
 func merge(intervals [][]int) int {
 
	 room := make([]int,0)
 
	 //排序
	 intervals = tranformArrayToStruct(intervals)
 
	 for i := 0; i < len(intervals); i++ {
		 //TODO 这个地方代码太丑了，同时，你只关心最小的可以用堆来解决，降低
		 //时间复杂度
		 minEndIndex := 999999
		 for j:=0;j<len(room);j++{
			 //如果某个会议可以插入一个已有的房间，则选择最小的一个插入
			 if room[j] <= intervals[i][0]{
				 if minEndIndex == 999999{
					 minEndIndex = j
				 }else if  room[minEndIndex] > room[j]{
					 minEndIndex = j
				 }
			   
			 }
		 }
 
		 if minEndIndex == 999999{
			 room = append(room,intervals[i][1])
		 }else{
  room[minEndIndex] = intervals[i][1]
		 }
		
	 }
	 return len(room)
 }
 
 func max(a, b int) int {
	 if a > b {
		 return a
	 } else {
		 return b
	 }
 }
 
 //就是转个换排个序
 func tranformArrayToStruct(intervals [][]int) [][]int {
	 var Inter Intervals
 
	 for i := 0; i < len(intervals); i++ {
		 Inter = append(Inter, Interval{Start: intervals[i][0], End: intervals[i][1]})
	 }
 
	 sort.Sort(Inter)
 
	 res := make([][]int, 0)
	 for _, inter := range Inter {
		 res = append(res, []int{inter.Start, inter.End})
	 }
	 return res
 }