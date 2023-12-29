// @Author liuzhen
// @Date 2023/12/19 20:20:00
// @Desc
package main

import (
	"fmt"
	"time"
)

/**
time包
时间和日期是我们编程中经常会用到的，本文主要介绍了 Go 语言内置的 time 包的基本用法。time 包提供了一些关于时间显示和测量用的函数。time 包中日历的计算采用的是公历，不考虑润秒。

1. 时间类型: Go 语言中使用 time.Time 类型表示时间。
	1.1. time.Now()  获取当前的时间对象time.Time	eg: now = time.Now()
		time.Time对象常用api:
			1.1.1. now.Year()  获取当前时间年
			1.1.2. now.Month()  获取当前时间月
			1.1.3. now.Day()  获取当前时间日
			1.1.4. now.Hour()  获取当前时间小时
			1.1.5. now.Minute()  获取当前时间分钟
			1.1.6. now.Second()  获取当前时间秒
			1.1.7. now.Unix()  秒级时间戳
			1.1.8. now.UnixMilli()  毫秒时间戳 Go1.17+
			1.1.9. now.UnixMicro()  微秒时间戳 Go1.17+
			1.1.10. now.UnixNano()  纳秒时间戳
	1.2.

2. Location和time zone
Go 语言中使用 location 来映射具体的时区。时区（Time Zone）是根据世界各国家与地区不同的经度而划分的时间定义，全球共分为24个时区。
中国差不多跨5个时区，但为了使用方便只用东八时区的标准时即北京时间为准。
	2.1. time.FixedZone()  返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	2.2. time.LoadLocation()  如果当前系统有时区数据库，则可以加载一个位置得到对应的时区，返回Location。
	2.3. time.Date() time.Time  创建时间对象time.Time，需要指定Location，常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）
	2.3. time.UTC  UTC时间
	2.4. .Local  当地时间
	2.5. time.Equal(sameTimeInBeijing)  比较时间

3. 时间间隔: time.Duration 是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。time.Duration 表示一段时间间隔，可表示的最长时间段大约290年。
	time 包中定义的时间间隔类型的常量如下: eg: time.Duration 表示1纳秒，time.Second 表示1秒。
		const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)

4. 时间操作
	4.1. 增加时间: now.Add(d Duration) Time
	4.2. 两个时间之间的差值: now.Sub(u Time) Duration
			返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。
	4.3. 判断两个时间是否相同: now.Equal(u Time) bool
			会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。
	4.4. t代表的时间点在u之前，返回真；否则返回假: now.Before(u Time) bool
	4.5. t代表的时间点在u之后，返回真；否则返回假: now.After(u Time) bool

5. 定时器: 使用 time.Tick(时间间隔) 来设置定时器，定时器的本质上是一个通道（channel）。
	eg: tickDemo() 定义一个1秒间隔的定时器 ticker := time.Tick(time.Second)

6. 时间格式化: time.Time => 字符串
	now.Format() 函数能够将一个时间对象格式化输出为指定布局的文本表示形式，需要注意的是 Go 语言中时间格式化的布局不是常见的 Y-m-d H:M:S。
		格式模板(记忆口诀为: 2006 1 2 3 4 5): 2006-01-02 15:04:05.000
			其中: 2006: 年(Y)   01: 月(m)   02: 日(d)   15: 时(H)   04: 分(M)   05: 秒(S)  AM/PM: 上午/下午
		补充: 如果想格式化为12小时格式，需在格式化布局中添加PM。小数部分想保留指定位数就写0，如果想省略末尾可能的0就写9。

7. 解析字符串格式的时间: 字符串 => time.Time
	格式模板(记忆口诀为: 2006 1 2 3 4 5): 2006-01-02 15:04:05.000
	7.1. time.Parse() 解析字符串格式时间。
	7.2. time.ParseInLocation()  解析字符串格式时间。注: 需要在解析时额外指定时区信息。
*/

func main() {
	// 1.
	timeDemo()
	// 2.
	timezoneDemo()
	// 4.
	timeOpDemo()
	// 5.
	tickDemo()
	// 6.
	formatDemo()
	// 7.1.
	parseDemo()
	// 7.2.
	parseInLocation()
}

// 1. timeDemo 时间对象的年月日时分秒
func timeDemo() {
	fmt.Println("1. timeDemo() -------------------->")

	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
}

// 2. timezoneDemo 时区示例。使用beijing来表示东八区8小时的偏移量，其中 time.FixedZone() 和 time.LoadLocation() 这两个函数则是用来获取location信息。
func timezoneDemo() {
	fmt.Println("2. timezoneDemo() -------------------->")

	// 2.1.
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 2.2.
	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
	if err != nil {
		fmt.Println("load America/New_York location failed", err)
		return
	}
	fmt.Println()
	// 加载上海所在的时区
	// shanghai, err := time.LoadLocation("Asia/Shanghai") // UTC+08:00
	// 加载东京所在的时区
	// tokyo, err := time.LoadLocation("Asia/Tokyo") // UTC+09:00

	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	// timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)  // 系统本地时间
	timeInUTC := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2024, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2024, 1, 1, 7, 0, 0, 0, newYork)
	fmt.Println("timeInUTC: ", timeInUTC)
	fmt.Println("sameTimeInBeijing: ", sameTimeInBeijing)
	fmt.Println("sameTimeInNewYork: ", sameTimeInNewYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	fmt.Println(timesAreEqual)
}

// 4. 时间操作
func timeOpDemo() {
	fmt.Println("4. timeOpDemo() -------------------->")
	now := time.Now()
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)

}

// 5. 定时器
func tickDemo() {
	fmt.Println("5. tickDemo() -------------------->")

	// 定义一个1秒间隔的定时器
	ticker := time.Tick(time.Second)
	count := 0
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务

		if count > 3 {
			break
		}
		count++
	}
}

// 6. 时间格式化
func formatDemo() {
	fmt.Println("6. formatDemo() -------------------->")

	now := time.Now()
	// 格式化的模板为 2006-01-02 15:04:05

	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96

	// 只格式化时分秒部分
	fmt.Println(now.Format("15:04:05"))
	// 只格式化日期部分
	fmt.Println(now.Format("2006.01.02"))
}

// 7.1. time.Parse() 解析时间
func parseDemo() {
	fmt.Println("7.1. parseDemo() -------------------->")

	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", "2022/10/05 11:25:20")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0000 UTC

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	timeObj, err = time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0800 CST
}

// 7.2.  time.ParseInLocation() 解析时间
func parseInLocation() {
	fmt.Println("7.2. ParseInLocation() -------------------->")

	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/10/05 11:25:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
