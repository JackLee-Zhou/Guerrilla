/*
 *  Copyright(c)
 *    ProjectName:Guerrilla
 *    FileName:eventstruct.go
 *    Date:2022/9/12 下午3:20
 *    Author:JackLee
 *   All rights reserved.
 */

package event

import "context"

type EventHandler func()

// Date 事件结构
type Date struct {
	Typ  string
	Hand EventHandler
	Ctx  context.Context
}

type Events map[string]Date
