package main

import "strconv"

/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */

// @lc code=start

//边界条件：（1）确定题目给的数据没有除数字冒号逗号外的任何数字
//对于ipv4而言：1.是否被分成了4段 2.是否被逗号分隔 3.是否小于256 4.长度大于1但是首位不能为0
//对于ipv6而言：1.是否被分成了8段 2.是否被逗号分隔 3.每一段长度是否在1-4之间 4.可以0打头可以不是0打头，可以只有一位 但是不能没有
//f:f:f:f:f:f 0:0:0:0:0:0:0  0213:1233.... 213:1233..... 等等都是合法的

func validIPAddress(IP string) string {
	//(1)区分两类ip 一类是4 一类是6
	if judgeIpIsIpv6(IP) {
		return "IPv6"
	} else if judgeIpIsIpv4(IP) {
		return "IPv4"
	} else {
		return "Neither"
	}
}

func judgeIpIsIpv6(ip string) bool {
	ip += ":"
	splitTimes := 0
	l := 0
	for r := 0; r < len(ip); {
		if ip[r] != ':' {
			r++
			continue
		}

		//(1)先判断位数 必须是小于4位大于1位，再判断里面的字符是否是合法的
		//如:::::: 02001:都是不合法的
		if len(ip[l:r]) <= 4 && len(ip[l:r]) >= 1 {
			for l < r {
				if (ip[l] <= '9' && ip[l] >= '0') ||
					(ip[l] <= 'f' && ip[l] >= 'a') ||
					(ip[l] <= 'F' && ip[l] >= 'A') {
					l++
					continue
				} else {
					return false
				}
			}
			splitTimes++
			r++
			l = r
		} else {
			return false
		}
	}
	return splitTimes == 8
}

func judgeIpIsIpv4(ip string) bool {
	//(1)注意这里要加一个逗号 帮助下代码
	ip += "."
	splitTimes := 0
	l := 0
	for r := 0; r < len(ip); {
		if ip[r] != '.' {
			r++
			continue
		}

		if val, err := strconv.Atoi(ip[l:r]); err == nil && val <= 255 {
			//001.123.121.1 这种情况
			if len(ip[l:r]) > 1 && ip[l] == '0' {
				return false
			}

			for l < r {
				if ip[l] <= '9' && ip[l] >= '0' {
					l++
					continue
				} else {
					return false
				}

			}
			splitTimes++
			r++
			l = r
		} else {
			return false
		}
	}

	return splitTimes == 4
}

//测试用例：1.功能测试 只包含数字和逗号情况 看能不能正确分辨 2.异常测试 有冒号和字母呢
//3.有冒号逗号字母呢
// func main() {
// 	// fmt.Println(judgeIpIsIpv4("172.16.254.1"))
// 	// fmt.Println(judgeIpIsIpv4("172.16.254.001"))
// 	// fmt.Println(judgeIpIsIpv4("2001:0db8:85a3:0:0:8A2E:0370:7334"))
// 	//
// 	// fmt.Println(judgeIpIsIpv6("2001:0db8:85a3::8A2E:0370:7334"))
// 	// fmt.Println(judgeIpIsIpv6("02001:0db8:85a3:0000:0000:8a2e:0370:7334"))
// 	// fmt.Println(judgeIpIsIpv6("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
// 	// fmt.Println(judgeIpIsIpv6("2001:db8:85a3:0:0:8A2E:0370:7334"))

// }

// @lc code=end
