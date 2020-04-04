package main

//2020/2/27

//这道题有两个感悟
//1.一开始对于这类，类似于要输出所有可能结果的题目，有点无从下手，后来发现都是类似于回溯+递归来完成。
//2.完全回溯就是暴力法了，因此一边回溯还要一边剪枝。
//3.回到这道题目上，后来我一分析，发现就是考组合的题目，假设n为2，你只需要算出4个坑中，如何组合两个"("，因为右括号")"填进来不影响结果。
//然后组合的结果中你要过滤出不合规的组合。
//4.这里的剪枝操作就是，无论何时，优先放左括号，只要在右括号多于左括号时才放右括号。
var res []string

func generateParenthesis1(n int) []string {
	res = []string{}
	recursive(n, n, "")
	return res
}

func recursive(leftNum, rightNum int, str string) {
	left, right := str, str
	left += "("
	right += ")"
	if leftNum == 0 && rightNum == 0 {
		res = append(res, str)
		return
	}

	if leftNum != 0 {
		recursive(leftNum-1, rightNum, left)
	}
	if rightNum > leftNum && rightNum != 0 {
		recursive(leftNum, rightNum-1, right)
	}

}

//2020/2/27
func generateParenthesis(n int) []string {
	res = make([]string, 0)
	helpGenerateParenthesis("", n, n)
	return res
}

func helpGenerateParenthesis(str string, leftNum, rightNum int) {
	if leftNum == 0 && rightNum == 0 {
		res = append(res, str)
		return
	}

	if leftNum != 0 {
		helpGenerateParenthesis(str+"(", leftNum-1, rightNum)
	}

	if rightNum > leftNum {
		helpGenerateParenthesis(str+")", leftNum, rightNum-1)
	}
	return
}
