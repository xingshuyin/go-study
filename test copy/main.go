package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type People struct {
	Name string `json:name`
	Age  int
}

func (p People) GetName() string {
	return p.Name
}

// func (p *People) SetMessage(name string, age int) {
func (p *People) SetMessage(people People) {
	p.Name = people.Name
	p.Age = people.Age
}
func get_value(v interface{}) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	v_, ok := v.(int)
	if ok {
		fmt.Println("类型断言", v_)
	}
	value := reflect.ValueOf(v)
	value_type := reflect.TypeOf(v)

	switch value.Kind() {
	case reflect.Int:
		fmt.Println("int", value.Int())
	case reflect.String:
		fmt.Println("string", value.String())
	case reflect.Struct:
		fmt.Println("struct", value)
		fmt.Println("struct_field0_tag_json", value_type.Field(0).Tag.Get("json"))
	case reflect.Pointer: // 指针类型的变量
		fmt.Println("pointer", value.Pointer())
		switch value.Elem().Kind() { //获取指针类型的值
		case reflect.Int:
			value.Elem().SetInt(55) //指针类型设置值
		case reflect.String:
			value.Elem().SetString("abcde")
		case reflect.Struct:
			fmt.Println("指针结构体的值", value.Elem())       //获取指针的值
			f0 := value.Elem().Field(0)                //获取结构体的第一个字段的值
			f_num := value.Elem().NumField()           //获取字段数量
			f_Name := value.Elem().FieldByName("Name") //获取结构体的Name字段的值
			f_type_Name, _ := value_type.Elem().FieldByName("Name")
			f_Name_tag := f_type_Name.Tag.Get("json")                                                                                  // 获取字段的tag,FieldByName的字段无法直接连接tag,因为他返回两个值
			f_Name_tag = value_type.Elem().Field(0).Tag.Get("json")                                                                    // 获取字段的tag
			fmt.Println("字段数", f_num, "第0个字段及类型 ", f0, f0.Type().Name(), f0.Type().Kind(), "Name字段", f_Name, "f_Name_tag", f_Name_tag) //获取结构体的字段的值名称及类型

			m0 := value.Elem().Method(0)                      //获取结构体的第一个方法
			m_GetName := value.Elem().MethodByName("GetName") //获取结构体的GetName方法
			m_GetName_r := m_GetName.Call(nil)                //执行结构体的GetName方法, nil代表无参数, 返回结果切片

			//因为是SetMessage指针的方法,所以这不用Elem(), 参数是 []reflect.Value{reflect.ValueOf("mike")}   /  reflect.ValueOf(People{Name: "mla", Age: 26})
			value.MethodByName("SetMessage").Call([]reflect.Value{reflect.ValueOf(People{Name: "mla", Age: 26})}) //执行结构体的GetName方法, 并传入参数, 返回结果切片
			fmt.Println("第0个方法哈希值排序", m0, "GetName方法及调用", m_GetName, m_GetName_r)
		}
	}

}

func file() {
	file, err := os.Open("E:/go/test copy/a.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("read file err", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Print(line)
	}
}

func byte_file() {
	file, err := os.Open("E:/go/test copy/main.exe")
	defer file.Close()
	if err != nil {
		fmt.Println("read file err", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Print(line)
	}
}

func ioutil_file() {
	file_byte, err := os.ReadFile("E:/go/test copy/a.txt")
	if err != nil {
		return
	}
	fmt.Print(string(file_byte))
}

func copy_file(file_old string, file_new string) {
	f, _ := os.ReadFile(file_old)
	os.WriteFile(file_new, f, 066)
}

func copy_file_(file_old string, file_new string) {
	f_old, _ := os.Open(file_old)
	defer f_old.Close()
	f_new, _ := os.OpenFile(file_new, os.O_CREATE|os.O_WRONLY, 066)
	defer f_new.Close()
	temp := make([]byte, 1280)
	for {
		n, err := f_old.Read(temp)
		fmt.Println(n)
		if err != nil {
			break
		}
		f_new.Write(temp[:n])
	}

}

func main() {
	copy_file_("D:/下载/bandicam 2023-04-04 19-08-42-408.mp4", "E:/go/test copy/main2.mp4")
}
