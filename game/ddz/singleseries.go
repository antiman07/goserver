package ddz
/*
从数组中找出最长连续的数字 分析是否构成单顺
*/

func MaxSerialArr(arr []int) []int {
	//总长
	var max = 0
	//集合点
	var jointIndex = -1
	//结果集
	var rs = make([]int, 0, len(arr))
	//集合点左边(比集合点小的部分长度)
	leftLen:= 0
	for i, v := range arr {
		lengthTmp,leftLenTmp,_:= depth(arr,v)
		if lengthTmp > max {
			max = lengthTmp
			jointIndex = i
			leftLen = leftLenTmp
		}
	}
	min := arr[jointIndex] - leftLen
	for j:=0;j<max;j++{
		rs = append(rs,min+j)
	}
	return rs
}

//总深度，左深度，右深度
func depth(arr []int, elem int) (int,int,int) {
	return  1 + containLeftLen(arr, elem-1) + containRightLen(arr, elem+1),containLeftLen(arr,elem-1), containRightLen(arr,elem+1)
}
//arr是否包含elem
func contain(arr []int, elem int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
//左深度
func containLeftLen(arr []int, elem int) int {
	if contain(arr, elem) {
		return containLeftLen(arr, elem-1)
	}
	return 0
}
//右深度
func containRightLen(arr []int, elem int) int {
	if contain(arr, elem) {
		return containRightLen(arr, elem+1) + 1
	}
	return 0
}