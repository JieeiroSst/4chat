package json

func init() {
	_ = ReadFileJson("account.json")
	_ = WriteJson("account.json")
}
