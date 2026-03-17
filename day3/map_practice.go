package main

import "fmt"

func main() {
	///创建map
	person := map[string]string{
		"name":  "herb",
		"city":  "长沙",
		"hobby": "学go",
	}

	//访问单个值
	fmt.Println("我的名字是:", person["name"])
	fmt.Println("我住在:", person["city"])

	//如果key不存在，返回零值
	fmt.Println("不存在的key:", person["age"]) //输出空字符串

	//遍历 map
	fmt.Println("\n遍历个人信息:")
	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}

	//添加/修改
	person["age"] = "18"  //新增
	person["neme"] = "宝贝" //修改
	fmt.Println("\n更新后:", person)

	//删除 key
	delete(person, "hobby")
	fmt.Println("删除 hobby 后:", person)

	//判断key是否存在
	if age, ok := person["age"]; ok {
		fmt.Println("年龄存在:", age)
	} else {
		fmt.Println("年龄不存在")
	}
}
