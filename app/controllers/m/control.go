// Copyright 2017 Vckai Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package m

import (
	"github.com/astaxie/beego"

	"github.com/vckai/novel/app/controllers"
)

// 前台首页基类
type BaseController struct {
	controllers.BaseController
}

// 初始化执行
func (this *BaseController) Prepare() {
	this.BaseController.Prepare()

	theme := beego.AppConfig.String("website::mtheme")

	this.Module = "m/" + theme
	this.Layout = "m/" + theme + "/layout.tpl"

	// 模板参数
	mOut := make(map[string]interface{})
	mOut["ViewUrl"] = beego.AppConfig.String("static::viewurl") + this.Module + "/"
	this.Data["mOut"] = mOut

	this.Data["Title"] = "海量小说、图书、免费小说、漫画，免费畅读"
}
