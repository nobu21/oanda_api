package exchange

/*
func main() {
	//	d := GetData("USD_JPY", "2016-06-03 5:53:09", "2016-06-03 15:40:33")
	d := GetData("USD_JPY", "2014-06", "2017-06")

	var exportdata string
	exportdata = fmt.Sprintf("# DDL\n")
	db := "exchange"
	table := "USDJPY"
	exportdata += fmt.Sprintf("CREATE DATABASE %s\n\n", db)
	exportdata += fmt.Sprintf("# DML\n")
	exportdata += fmt.Sprintf("# CONTEXT-DATABASE: %s\n", db)
	for _, v := range d.Candles {
		exportdata += fmt.Sprintf("%s,bid=%f ask=%f %d\n", table, v.OpenBid, v.OpenAsk, v.Time.Unix())
	}
	fmt.Println(exportdata)

	/*
		d := new(OANDAStreamData)
		SetData(d, "USD_JPY")
		for {
			fmt.Println(<-d.c)
		}
*/
//}
