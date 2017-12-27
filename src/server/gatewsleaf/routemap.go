// Copyright 2014 mqantleaf Author. All Rights Reserved.
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
package gatewsleaf
/**
重点函数
该函数里面注册 leaf msgid 与 mqant rpc函数名的映射关系
 */
func (this *CustomAgent) RouteMap(){
	this.routeMap[uint16(3001)]="Login/HD_Login" //  leaf msgid 3001 <--> moduleType/_func
	this.routeMap["Login/Success"]=uint16(3002)  // user.LoginSuccessResponse <--> Login/Success <--> leaf msgid 3002
}
