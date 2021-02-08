package timeago

var bnTimeTypes = [][]string{
	{"এইমাত্র", "একটা সময়"},
	{"%d সেকেন্ড আগে", "%d এর সেকেন্ডের মধ্যে"},
	{"1 মিনিট আগে", "1 মিনিটে"},
	{"%d এর মিনিট আগে", "%d এর মিনিটের মধ্যে"},
	{"1 ঘন্টা আগে", "1 ঘন্টা"},
	{"%d ঘণ্টা আগে", "%d এর ঘন্টার মধ্যে"},
	{"1 দিন আগে", "1 দিনের মধ্যে"},
	{"%d এর দিন আগে", "%d এর দিন"},
	{"1 সপ্তাহ আগে", "1 সপ্তাহের মধ্যে"},
	{"%d এর সপ্তাহ আগে", "%d সপ্তাহের মধ্যে"},
	{"1 মাস আগে", "1 মাসে"},
	{"%d মাস আগে", "%d মাসে"},
	{"1 বছর আগে", "1 বছরের মধ্যে"},
	{"%d বছর আগে", "%d বছরে"},
}

func bnLocale(_ float64, idx int) (ago string, in string) {
	var res = bnTimeTypes[idx]
	return res[0], res[1]
}
