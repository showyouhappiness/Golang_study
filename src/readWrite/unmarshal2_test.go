package readWrite

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Opus 作品
type Opus struct {
	Date  string
	Title string
}
type Opus1 struct {
	Type  string
	Title string
}

type Actress1 struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       []string
}

type Actress2 struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       Opus
}

type Actress3 struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       []Opus
}

type Actress4 struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       map[string]Opus1
}

// 因为json.UnMarshal() 函数接收的参数是字节切片，所以需要把JSON字符串转换成字节切片。

// 普通JSON
func TestUnmarshal1(t *testing.T) {
	jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus":[
         "《阿娜尔罕》",
         "《逆光之恋》",
         "《克拉恋人》"
      ]
   }`)

	var actress Actress1
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	for _, val := range actress.Opus {
		fmt.Println("\t", val)
	}
}

//JSON内嵌普通JSON
func TestUnmarshal2(t *testing.T) {
	// JSON嵌套普通JSON
	jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus": {
         "Date":"2013",
         "Title":"《阿娜尔罕》"
      }
   }`)
	var actress Actress2
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	fmt.Printf("\t%s:%s", actress.Opus.Date, actress.Opus.Title)
}

//JSON内嵌数组JSON
func TestUnmarshal3(t *testing.T) {
	// JSON嵌套数组JSON
	jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus":[
         {
            "date":"2013",
            "title":"《阿娜尔罕》"
         },
         {
            "date":"2014",
            "title":"《逆光之恋》"
         },
         {
            "date":"2015",
            "title":"《克拉恋人》"
         }
      ]
   }`)
	var actress Actress3
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	for _, val := range actress.Opus {
		fmt.Printf("\t%s - %s\n", val.Date, val.Title)
	}
}

//JSON内嵌具有动态Key的JSON
func TestUnmarshal4(t *testing.T) {
	jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus":{
         "2013":{
            "Type":"近代革命剧",
            "Title":"《阿娜尔罕》"
         },
         "2014":{
            "Type":"奇幻剧",
            "Title":"《逆光之恋》"
         },
         "2015":{
            "Type":"爱情剧",
            "Title":"《克拉恋人》"
         }
      }
   }`)
	var actress Actress4
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	for index, value := range actress.Opus {
		fmt.Printf("\t日期：%s\n", index)
		fmt.Printf("\t\t分类：%s\n", value.Type)
		fmt.Printf("\t\t标题：%s\n", value.Title)
	}
}
