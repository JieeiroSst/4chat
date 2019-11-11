package testing

import (
	"4chat/json"
	"testing"
)

func AccountMovieRead(b *testing.B) {
	for n := 0; n < b.N; n++ {
	 	json.ReadJson("../json/account.json", func(data map[string]interface{}) bool {
			return data["Account"].(float64) >= 2010
		})
	}
}

func AccountMovieReadToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.ReadJsonToken("../json/account.json", func(data map[string]interface{}) bool {
			return data["Account"].(float64) >= 2010
		})
	}
}

func BenchmarkQRead(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.ReadJson("../json/account.json", func(data map[string]interface{}) bool {
			return data["Account"].(string) == "4680"
		})
	}
}

func BenchmarkQReadToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.ReadJsonToken("../json/account.json", func(data map[string]interface{}) bool {
			return data["Account"].(string) == "4680"
		})
	}
}

