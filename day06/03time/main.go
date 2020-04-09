package main

import (
	"fmt"
	"time"
)

//时间
//把一个字符串的时间转化为时间戳


//时区
func f1()  {
	now := time.Now() //获取本地的时间
	fmt.Println(now)
	time.Parse("2006-01-02 15:04:05","2019-08-04 14:41:50")
	//按照东八区的时区去解析时间
	//根据字符串加载一个时区
	loc, err :=time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println("load location failed")
	}
	//按照指定时区解析时间
	timeobj, err:= time.ParseInLocation("2006-01-02 15:04:05","2019-08-04 14:41:50",loc)
	fmt.Println(timeobj)
}

func main()  {
	now := time.Now()
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	timeStamp := now.Unix()
	ret := time.Unix(timeStamp,0)
	fmt.Println(ret)
	// now + 1hour
	fmt.Println(now.Add(24*time.Hour))
	//Sub计算两个时间的差值
	//Equal 判断两个时间是否相等
	//Befor after 判断时间先后


	//定时器
	//timer := time.Tick(time.Second) //1s执行一次
	//for t := range timer{
	//	fmt.Println(t)
	//}
	//格式化时间，将语言对象转化成字符串类型的时间
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 03:04:05.000 PM"))
	//按照对应的格式 解析字符串类型的时间

	timeObj, _ := time.Parse("2006/01/02 15:04:05","2000/05/12 15:04:05")
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	//n := 100
	//time.Sleep(time.Duration(n))//强制转换为Duration类型
	//睡眠 和python一样
	//time.Sleep(3*time.Minute)
	//时间相减
	now.UTC() //转换为UTC时间
	t2 := now.Add(24*time.Hour+10*time.Minute)
	fmt.Println(t2.Sub(now))
	f1()
}


