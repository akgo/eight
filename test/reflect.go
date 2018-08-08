package main

import (
	"fmt"
	"reflect"
	"gopkg.in/ini.v1"
	"os"
)

var (
	command string
	parameter string
	controllers string
	method string
	modelManage *ModelManage
)

type Model struct{
	Command string

	Controllers string
	Method string
}

type ModelManage struct{
	Models map[string] interface{}
}

func (manage *ModelManage) BindM() {

}

//动态加入新的Model
func (manage *ModelManage) Regist (command string,model *Model) {
	manage.Models[command] = model;
}

//删除Model
func (manage *ModelManage) UnRegist (command string) {
	delete (manage.Models,command)
}

func init() {
	modelManage = &ModelManage{make(map[string]interface{})}




	//注册UserM模块
	userModel := &User{"AkTest", 18, "00000001"}
	modelManage.Models["10000"] = userModel
	modelManage.Models["10001"] = userModel
	modelManage.Models["10002"] = userModel

	//注册Equip模块
	equipModel := &Equip{Id:"Item0000001"}
	modelManage.Models["20001"] = equipModel
	modelManage.Models["20002"] = equipModel

}

func main() {

	cfg, err := ini.Load("command.conf")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Please enter you commandId: ")
	for {

		fmt.Scanln(&command,&parameter)

		if(command == "bye") {
			os.Exit(1)
		}

		fmt.Println(command,parameter)

		controllers = cfg.Section(command).Key("controllers").String()
		method = cfg.Section(command).Key("method").String()

		if controllers =="" {
			fmt.Println(" Error !")
			os.Exit(1)
		}

		fmt.Println(controllers,method)

		testObj := modelManage.Models[command]
		object := reflect.ValueOf(testObj)



		v := object.MethodByName(method)
		v.Call([]reflect.Value{})
	}

}

type User struct {
	Name string
	Age  int
	Id   string
}

func (u *User) Login() {
	fmt.Println(" hello ",parameter)
}

func (u *User) Logout() {
	fmt.Println(" Bye !")
}

func (u *User) GetInfo() {
	fmt.Println("I'm " + u.Name + ", Id is " + u.Id + ". Nice to meet you! ")
}

type Equip struct {
	Id string
}

func (e *Equip) GetInfo () {

	fmt.Println(" it is Equip GetInfo! ",parameter)
}

func (e *Equip) List () {
	fmt.Println(" it is Equip List! ",parameter)
}