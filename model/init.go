package model

import "github.com/satori/go.uuid"

func init() {
	modelManage = &ModelManage{}
	modelManage.handlers = make(map[string]*handlerEntity);

	//绑定人物模块
	user := NewUser("ak",18,uuid.Must(uuid.NewV4()).String())
	modelManage.BindM(user)

	//绑定装备模块
	//modelManage.BindM(Equip{"item_0000001"})
}

