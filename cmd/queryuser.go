// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/cyulei/agenda/datarw"
	"github.com/modood/table"
	"github.com/spf13/cobra"
)

// queryuserCmd represents the queryuser command
var queryuserCmd = &cobra.Command{
	Use:   "queryuser",
	Short: "Show name,email,phone of users",
	Long: `queryuser:Show name,email,phone of users
	you must login before query
	For example:
	agenda queryuser             :show all registered users' information
	agenda queryuser -n user1    :show user1' information if registered
	`,
	Run: func(cmd *cobra.Command, args []string) {

		queryuser()

	},
}

func init() {
	rootCmd.AddCommand(queryuserCmd)
	queryuserCmd.Flags().StringVarP(&queryuserName, "name", "n", "", "user's name")
}

var queryuserName string

func queryuser() {
	logInit()
	defer logFile.Close()

	curUser = datarw.GetCurUser()

	if curUser == nil { //是否已登陆
		logSave("isn't login,please use command login", "[Error]")
		return
	} else {
		logSave("cmd: queryuser called", "[Info]")
	}

	//获取所有用户
	users := datarw.GetUsers()

	if queryuserName == "" { //查询所有用户（因为已登录，所以不可能没有用户）

		for i := range users { //掩盖密码
			users[i].Password = "********"

		}

		fmt.Println(table.Table(users))

	} else { //查询单个用户
		for _, user := range users {
			if user.Name == queryuserName {
				fmt.Println("\t", user.Name, "\t", user.Email, "\t", user.Phone)
				return //查询成功
			}
		}
		logSave(queryuserName+"doesn't registered", "[Warning]") //查询失败

	}

	logSave("cmd: queryuser success", "[Info]")
}
