package main

import "fmt"

type Receiver struct {
	PhoneNum string
}

var AlertPhoneNumberMapping = map[string][]string{
	"小学数学": {"18501365488", "13476270040"},
	"初中数学": {"13011134178", "18162663189"},
	"高中数学": {"17610066831", "13476981862"},
	"小学语文": {"13477090039", "18627804473"},
	"初中语文": {"13477090039", "18627919470"},
	"高中语文": {"13477090039", "18627913420"},
	"小学英语": {"17771841962", "18627960465"},
	"初中英语": {"17771841962", "18627943421"},
	"高中英语": {"17771841962", "18627961426"},
	"初中物理": {"13011134178", "15071228040"},
	"高中物理": {"17610066831", "18627851467"},
	"初中化学": {"13011134178", "18672958276"},
	"高中化学": {"17610066831", "18271475825"},
}
var demo map[string][]string

var a string

func main() {
	//data, _ := json.Marshal(AlertPhoneNumberMapping)
	//fmt.Println(string(data))
	//json.Unmarshal(data, &demo)
	//fmt.Println(demo)
	//data,_ := json.Marshal(AlertPhoneNumberMapping)
	//json.Unmarshal(data,demo)

	//maps := make(map[int]string)
	//target := map[int]string{
	//	1:"aa",6:"cc",
	//}
	//a := []string{"a","b","c","d"}
	//for index,v := range a{
	//	maps[index] = v
	//}
	//var b []string
	//for index,_ := range a{
	//	if _,ok := target[index];ok  {
	//		b = append(b,a[index])
	//	}
	//}
	//fmt.Println(b)

	//maps := map[string]string{"4":"a","2":"b","3":"c"}
	//for k,v := range maps{
	//	fmt.Println(k+":"+v)
	//}
	//a := make([]int,0)
	//for i:=0;i<10;i++{
	//	a = append(a,i)
	//}
	//fmt.Println(a)

	//a := "a"
	//b := strings.Split(a,",")
	//fmt.Println(b[0])

	//var a map[int]string
	//a[1] = "a"
	//fmt.Println(a[1])
	//
	//a := []int{1,2,3,4,5}
	//b := []int{6,7,8,9,10}
	//
	//a = append(a,b...)
	//fmt.Println(a)
	//
	//a := []int{1,2,3}
	//fmt.Println(a[0:3])
	//teacherList := "[1111,1111,1111,111,111]"
	//var teacherIdList []int
	//err := json.Unmarshal([]byte(teacherList), &teacherIdList)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(teacherIdList)
	//for a := range os.Args {
	//	fmt.Println(a)
	//}
	//fmt.Println(len("[error_code] = %d   "))
	//maps := make(map[int]int)
	//for i := 0; i < 10; i++ {
	//	maps[i] = i
	//}
	//
	//for key, value := range maps {
	//	fmt.Printf("%v:%v \n ", key, value)
	//}
	//
	//fmt.Println("a \n  b \n")
	//a := fmt.Sprintf("http://10.56.230.122:8090/error/apply?server_name=go_%s&user_name=%s", "bingo_test", "username")
	//fmt.Println(a)

	//a := "aaa   aaa bb"
	//for _, v := range strings.Split(a, " ") {
	//	fmt.Println(v)
	//}

	//handled := `spp_k12_logic_status_svr spp_k12_online_logic_svr spp_k12_txcloud_hls_auth_svr k12_edu_msg_center k12_edu_online_userlist k12_edu_info_write_svr k12_edu_cs_class_interaction k12_edu_info_read_svr k12_edu_enter_live_auth`
	//handled := `spp_k12_txcloud_proxy_svr k12_video_audit spp_k12_logic_status_svr spp_k12_online_logic_svr spp_k12_txcloud_hls_auth_svr k12_edu_msg_center k12_edu_online_userlist k12_edu_info_write_svr k12_edu_cs_class_interaction k12_edu_info_read_svr k12_edu_enter_live_auth`
	//target := `k12_edu_all_agent k12_edu_course_write_svr k12_edu_info_read_svr k12_edu_info_write_svr k12_edu_msg_center k12_edu_online_userlist k12_video_audit spp_edu_cs_interface spp_k12_logic_status_svr spp_k12_online_logic_svr spp_k12_txcloud_hls_auth_svr spp_k12_txcloud_proxy_svr`
	//a := strings.Split(handled, " ")
	//b := strings.Split(target, " ")
	//fmt.Println("handle = ", len(a))
	//fmt.Println("target = ", len(b))
	//handle_map := make(map[string]string)
	//target_map := make(map[string]string)
	//for _, v := range a {
	//	handle_map[v] = v
	//}
	//for _, v := range b {
	//	target_map[v] = v
	//}
	//
	//for k, _ := range target_map {
	//	if _, ok := handle_map[k]; !ok {
	//		fmt.Println(k)
	//	}
	//}
	a := []int{361,
		355,
		353,
		355,
		336,
		336,
		315,
		373,
		361,
		352,
		342,
		337,
		334,
		351,
		314,
		337,
		377,
		335,
		352,
		353,
		341,
		312,
		343,
		329,
		342,
		392,
		334,
		353,
		351}

	var total int
	var max int
	var min int
	min = a[0]
	for _, i := range a {
		total += i
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	ave := total / len(a)
	fmt.Println(len(a), ave, min, max)
}
