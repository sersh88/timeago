package timeago

func thLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"เมื่อสักครู่นี้", "อีกสักครู่"},
		{"%d วินาทีที่แล้ว", "ใน %d วินาที"},
		{"1 นาทีที่แล้ว", "ใน 1 นาที"},
		{"%d นาทีที่แล้ว", "ใน %d นาที"},
		{"1 ชั่วโมงที่แล้ว", "ใน 1 ชั่วโมง"},
		{"%d ชั่วโมงที่แล้ว", "ใน %d ชั่วโมง"},
		{"1 วันที่แล้ว", "ใน 1 วัน"},
		{"%d วันที่แล้ว", "ใน %d วัน"},
		{"1 อาทิตย์ที่แล้ว", "ใน 1 อาทิตย์"},
		{"%d อาทิตย์ที่แล้ว", "ใน %d อาทิตย์"},
		{"1 เดือนที่แล้ว", "ใน 1 เดือน"},
		{"%d เดือนที่แล้ว", "ใน %d เดือน"},
		{"1 ปีที่แล้ว", "ใน 1 ปี"},
		{"%d ปีที่แล้ว", "ใน %d ปี"},
	}[index]
	return res[0], res[1]
}
