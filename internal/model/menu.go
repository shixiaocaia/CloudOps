/*
 * MIT License
 *
 * Copyright (c) 2024 Bamboo
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

package model

type Menu struct {
	ID         int     `json:"id" gorm:"primaryKey;column:id;comment:菜单ID"`
	Name       string  `json:"name" gorm:"column:name;type:varchar(50);not null;comment:菜单显示名称"`
	ParentID   int     `json:"parent_id" gorm:"column:parent_id;default:0;comment:上级菜单ID,0表示顶级菜单"`
	Path       string  `json:"path" gorm:"column:path;type:varchar(255);not null;comment:前端路由访问路径"`
	Component  string  `json:"component" gorm:"column:component;type:varchar(255);not null;comment:前端组件文件路径"`
	Icon       string  `json:"icon" gorm:"column:icon;type:varchar(50);default:'';comment:菜单显示图标"`
	SortOrder  int     `json:"sort_order" gorm:"column:sort_order;default:0;comment:菜单显示顺序,数值越小越靠前"`
	RouteName  string  `json:"route_name" gorm:"column:route_name;type:varchar(50);not null;comment:前端路由名称,需唯一"`
	Hidden     int     `json:"hidden" gorm:"column:hidden;type:tinyint(1);default:0;comment:菜单是否隐藏(0:显示 1:隐藏)"`
	CreateTime int64   `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:记录创建时间戳"`
	UpdateTime int64   `json:"update_time" gorm:"column:update_time;autoUpdateTime;comment:记录最后更新时间戳"`
	IsDeleted  int     `json:"is_deleted" gorm:"column:is_deleted;type:tinyint(1);default:0;comment:逻辑删除标记(0:未删除 1:已删除)"`
	Children   []*Menu `json:"children" gorm:"-"` // 子菜单列表,不映射到数据库
}

type CreateMenuRequest struct {
	Name      string `json:"name" binding:"required"`    // 菜单名称
	Path      string `json:"path" binding:"required"`    // 菜单路径
	ParentId  int    `json:"parent_id" binding:"gte=0"`  // 父菜单ID
	Component string `json:"component"`                  // 组件
	Icon      string `json:"icon"`                       // 图标
	SortOrder int    `json:"sort_order" binding:"gte=0"` // 排序
	RouteName string `json:"route_name"`                 // 路由名称
	Hidden    int    `json:"hidden" binding:"oneof=0 1"` // 是否隐藏
}

type GetMenuRequest struct {
	Id int `json:"id" binding:"required,gt=0"` // 菜单ID
}

type UpdateMenuRequest struct {
	Id        int    `json:"id" binding:"required,gt=0"` // 菜单ID
	Name      string `json:"name" binding:"required"`    // 菜单名称
	Path      string `json:"path" binding:"required"`    // 菜单路径
	ParentId  int    `json:"parent_id" binding:"gte=0"`  // 父菜单ID
	Component string `json:"component"`                  // 组件
	Icon      string `json:"icon"`                       // 图标
	SortOrder int    `json:"sort_order" binding:"gte=0"` // 排序
	RouteName string `json:"route_name"`                 // 路由名称
	Hidden    int    `json:"hidden" binding:"oneof=0 1"` // 是否隐藏
}

type DeleteMenuRequest struct {
	Id int `json:"id" binding:"required,gt=0"` // 菜单ID
}

type ListMenusRequest struct {
	PageNumber int  `json:"page_number" binding:"required,gt=0"` // 页码
	PageSize   int  `json:"page_size" binding:"required,gt=0"`   // 每页数量
	IsTree     bool `json:"is_tree"`                             // 是否树形结构
}
