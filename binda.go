package main

import (
	"fmt"
	"reflect"
	"os"
	"strings"
	"github.com/labstack/gommon/log"
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

type IObject interface {
}

type ModelManage struct{
	handlers map[string]*handlerEntity
}

type handlerEntity struct {
	object IObject
	method reflect.Value
	argType reflect.Type
	argIsRaw bool
}

func isHandlerMethod(method reflect.Method) bool{
	mt := method.Type

	if mt.NumIn() != 3{
		return true
	}

	return true
}

func (manage *ModelManage) BindM(obj IObject) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	name := reflect.Indirect(v).Type().Name()

	//if manage.handlers == nil{
	//	manage.handlers = make(map[string]*handlerEntity)
	//}
	fmt.Println(t,v,name)
	for m := 0; m < t.NumMethod(); m++ {
		method := t.Method(m)
		mt := method.Type
		mn := method.Name
		fmt.Println(method,mt,mn)
		if isHandlerMethod(method){
			raw := false
			manage.handlers[strings.ToLower(fmt.Sprintf("%s.%s", name, mn))] = &handlerEntity{
				object: obj,
				method: v.Method(m),
				argType: mt.In(0),
				argIsRaw: raw,
			}
		}else{
			log.Printf("%s.%s register failed, argc=%d\n", name, mn, mt.NumIn())
		}
	}
}

func (manage *ModelManage) Exec(name string, data interface{}) {
	fmt.Println(strings.ToLower(name))

	h, ok := manage.handlers[strings.ToLower(name)]
	if !ok {
		fmt.Println("not found handle by", name)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("dispatch error", name, err)
		}
	}()

	//var argv reflect.Value
	//argv = reflect.ValueOf(data)

	args := []reflect.Value{}
	h.method.Call(args)
}


func init() {
	modelManage = &ModelManage{}
	modelManage.handlers = make(map[string]*handlerEntity);

	//绑定人物模块
	modelManage.BindM(&User{"ak",18,"0000001"})

	//绑定装备模块
	//modelManage.BindM(Equip{"item_0000001"})

}

func main() {

	fmt.Printf("Please enter you commandId: ")
	for {

		fmt.Scanln(&command,&parameter)

		if(command == "bye") {
			os.Exit(1)
		}

		fmt.Println(command,parameter)

		modelManage.Exec(command,parameter)

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
