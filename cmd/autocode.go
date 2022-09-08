package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/owenzhou/artisan/app"
	"github.com/owenzhou/artisan/config"
	"github.com/owenzhou/artisan/controller"
	"github.com/owenzhou/artisan/event"
	"github.com/owenzhou/artisan/listener"
	"github.com/owenzhou/artisan/model"
)

type tableField struct {
	Field      string `json:"Field"`
	Type       string `json:"Type"`
	Collation  string `json:"Collation"`
	Null       string `json:"Null"`
	Key        string `json:"Key"`
	Default    string `json:"Default"`
	Extra      string `json:"Extra"`
	Privileges string `json:"Privileges"`
	Comment    string `json:"Comment"`
}

//创建文件夹及文件
func makeFile(filename string) *os.File {
	//文件夹不存在则创建文件夹
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); err != nil {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("create dir fail.")
			return nil
		}
	}
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	return f
}

//创建模板文件
func makeTplFile(name string, tplStr string, data map[string]interface{}, funcs ...template.FuncMap) (string, error) {
	var f template.FuncMap
	if len(funcs) > 0 {
		f = funcs[0]
	} else {
		f = template.FuncMap{}
	}
	tmpl, _ := template.New(name).Funcs(f).Parse(tplStr)
	file := makeFile(name)
	defer file.Close()

	return name, tmpl.Execute(file, data)
}

//等待用户输入
func waitEnter(name string) string {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print(name + " file or directory already exist, continue?(Y/N):")
	input.Scan()
	return input.Text()
}

//创建控制器
func makeController(name string, resource ...bool) string {
	var tplStr string
	if len(resource) > 0 && resource[0] {
		tplStr = controller.ResourceTemplate
	} else {
		tplStr = controller.Template
	}
	if config.Config == nil {
		return "Make Controller error: config file is not exists."
	}

	filePath := config.Config.ControllerPath + "/" + name + ".go"
	if _, err := os.Stat(filePath); err == nil {
		s := waitEnter(filePath)
		if s == "N" {
			return "exit."
		}
	}

	tmplArr := map[string]interface{}{
		"packageName":    filepath.Base(filepath.Dir(filePath)),
		"controllerName": filepath.Base(name),
	}
	createdName, err := makeTplFile(filePath, tplStr, tmplArr)
	if err != nil {
		return createdName + " controller create failed."
	}
	return createdName + " controller created."
}

//创建模型
func makeModel(name string) (result string) {

	var table tableField
	var modelFields = make([]map[string]string, 0)
	rows, err := db.Model(&tableField{}).Raw("show full columns from " + filepath.Base(name)).Rows()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	if config.Config == nil {
		return "Make Model error: config file is not exists."
	}

	filePath := config.Config.ModelPath + "/" + name + ".go"
	if _, err := os.Stat(filePath); err == nil {
		s := waitEnter(filePath)
		if s == "N" {
			return "exit."
		}
	}

	//是否导入time包
	importTime := false
	//是否导入sql包
	importSql := false
	for rows.Next() {
		db.ScanRows(rows, &table)
		//处理table字段
		var tableFields = make(map[string]string)
		var gormTagStr = ""
		var bindingStr = ""
		var formStr = ""
		var labelStr = ""
		tableFields["field"] = ""
		tableFields["type"] = ""
		tableFields["tag"] = ""

		//默认添加clumn标识
		gormTagStr += "column:" + table.Field + ";"
		//处理模型的 field 字段
		if table.Field == "id" {
			tableFields["field"] = "ID"
		} else if strings.Contains(table.Field, "_") {
			tableFields["field"] = strings.ReplaceAll(strings.Title(strings.ReplaceAll(table.Field, "_", " ")), " ", "")
		} else {
			tableFields["field"] = strings.Title(table.Field)
		}

		//处理模型 type 字段
		if strings.Contains(table.Type, "tinyint") {
			tableFields["type"] = "int8"
			if strings.Contains(table.Type, "unsigned") {
				tableFields["type"] = "uint8"
			}
		} else if strings.Contains(table.Type, "mediumint") || strings.Contains(table.Type, "smallint") {
			tableFields["type"] = "int16"
			if strings.Contains(table.Type, "unsigned") {
				tableFields["type"] = "uint16"
			}
		} else if strings.Contains(table.Type, "bigint") {
			tableFields["type"] = "int64"
			if strings.Contains(table.Type, "unsigned") {
				tableFields["type"] = "uint64"
			}
		} else if strings.Contains(table.Type, "int") {
			tableFields["type"] = "int"
			if strings.Contains(table.Type, "unsigned") {
				tableFields["type"] = "uint"
			}
		} else if strings.Contains(table.Type, "varchar") || strings.Contains(table.Type, "char") {
			tableFields["type"] = "string"
			reg := regexp.MustCompile(`\d+`)
			cd := reg.FindString(table.Type)
			if cd != "" {
				bindingStr += "min=0,max=" + cd + ","
			}
		} else if strings.Contains(table.Type, "timestamp") ||
			strings.Contains(table.Type, "datetime") ||
			strings.Contains(table.Type, "date") ||
			strings.Contains(table.Type, "time") ||
			strings.Contains(table.Type, "year") {
			importTime = true
			tableFields["type"] = "time.Time"
		} else if strings.Contains(table.Type, "text") {
			tableFields["type"] = "string"
			gormTagStr += "type:longText;"
		} else if strings.Contains(table.Type, "float") ||
			strings.Contains(table.Type, "double") ||
			strings.Contains(table.Type, "decimal") {
			tableFields["type"] = "float64"
		} else {
			tableFields["type"] = table.Type
		}

		//当字段默认值不为空时，再次处理Type
		if table.Default != "" {
			defaultValue := table.Default
			tableFields["type"] = "*" + tableFields["type"]

			/*不使用sql.Null...类型，gin shouldBind不支持
			importSql = true		
			if tableFields["type"] == "bool" {
				if table.Default == "0" {
					defaultValue = "false"
				}else{
					defaultValue = "true"
				}
				
				tableFields["type"] = "*bool"
			}
			
			if tableFields["type"] == "int16" || tableFields["type"] == "uint16" {
				tableFields["type"] = "sql.NullInt16"
			}
			if tableFields["type"] == "int" || tableFields["type"] == "uint" {
				tableFields["type"] = "sql.NullInt32"
			}
			if tableFields["type"] == "int64" || tableFields["type"] == "uint64" {
				tableFields["type"] = "sql.NullInt64"
			}
			if tableFields["type"] == "float64" {
				tableFields["type"] = "sql.NullFloat64"
			}
			*/
			gormTagStr += "default:" + defaultValue + ";"
		}

		//处理模型 tag 字段
		if table.Key == "PRI" {
			gormTagStr += "primaryKey;"
		}
		if table.Key == "UNI" {
			gormTagStr += "unique;"
		}
		
		if table.Extra == "auto_increment" {
			gormTagStr += "autoIncrement;"
		}

		//字段默认加上form,label标签
		formStr += table.Field
		labelStr += table.Comment

		//字段不为空，则设置binding标签
		if table.Null == "NO" {
			gormTagStr += "not null;"
			//创建binding，主键不用创建binding
			if table.Key != "PRI" {
				bindingStr += "required"
			}
		}else if strings.Contains(tableFields["type"], "string"){	//字段为空不更新问题，加入指针
			tableFields["type"] = "*string"
		}

		if table.Comment != "" {
			gormTagStr += "comment:'" + table.Comment + "';"
		}

		//最后将tag合并
		if gormTagStr != "" {
			gormTagStr = ` gorm:"` + gormTagStr + `"`
		}

		if bindingStr != "" {
			bindingStr = ` binding:"` + bindingStr + `"`
		}

		if formStr != "" {
			formStr = ` form:"` + formStr + `"`
		}

		if labelStr != "" {
			labelStr = ` label:"` + labelStr + `"`
		}

		tableFields["tag"] = fmt.Sprintf("`json:\"%s\"%s%s%s%s`", table.Field, gormTagStr, bindingStr, formStr, labelStr)

		modelFields = append(modelFields, tableFields)
	}

	tmplArr := map[string]interface{}{
		"packageName": filepath.Base(filepath.Dir(filePath)),
		"importTime":  importTime,
		"importSql":   importSql,
		"modelName":   strings.ReplaceAll(strings.Title(strings.ReplaceAll(filepath.Base(name), "_", " ")), " ", ""),
		"fields":      modelFields,
	}
	createdName, err := makeTplFile(filePath, model.Template, tmplArr, template.FuncMap{"ucfirst": strings.Title})
	if err != nil {
		return createdName + " model create failed."
	}
	return createdName + " model created."
}

//创建事件
func makeEvent(name string) string {
	if config.Config == nil {
		return "Make Event error: config file is not exists."
	}
	filePath := config.Config.EventPath + "/" + name + ".go"
	if _, err := os.Stat(filePath); err == nil {
		s := waitEnter(filePath)
		if s == "N" {
			return "exit."
		}
	}

	tmplArr := map[string]interface{}{
		"packageName": filepath.Base(filepath.Dir(filePath)),
		"eventName":   filepath.Base(name),
	}
	createdName, err := makeTplFile(filePath, event.Template, tmplArr)
	if err != nil {
		return createdName + " event create failed."
	}
	return createdName + " event created."
}

//创建监听器
func makeListener(name string) string {
	if config.Config == nil {
		return "Make Listener error: config file is not exists."
	}

	filePath := config.Config.ListenerPath + "/" + name + ".go"
	if _, err := os.Stat(filePath); err == nil {
		s := waitEnter(filePath)
		if s == "N" {
			return "exit."
		}
	}

	tmplArr := map[string]interface{}{
		"packageName":  filepath.Base(filepath.Dir(filePath)),
		"listenerName": filepath.Base(name),
	}
	createdName, err := makeTplFile(filePath, listener.Template, tmplArr)
	if err != nil {
		return createdName + " listener created failed."
	}
	return createdName + " listener created."
}

//创建app
func newApp(name string) string {
	currentDir, _ := os.ReadDir(".")
	if len(currentDir) > 0 {
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("The current dir is not empty, continue?(Y/N):")
		input.Scan()
		if input.Text() == "N" {
			return "exit."
		}
	}
	//创建配置文件
	configName, err := makeTplFile("config.yaml", app.ConfigTemplate, map[string]interface{}{})
	if err != nil {
		fmt.Println("./" + configName + " create failed.")
	} else {
		fmt.Println("./" + configName + " created.")
	}

	//创建go.mod文件
	gomodName, err := makeTplFile("go.mod", app.GoModTemplate, map[string]interface{}{
		"module": name,
	})
	if err != nil {
		fmt.Println("./" + gomodName + " create failed.")
	} else {
		fmt.Println("./" + gomodName + " created.")
	}
	//创建一个默认的控制器
	ctrlPath := "app/http/controllers/HomeController.go"
	ctrlName, err := makeTplFile(ctrlPath, app.CtrlTemplate, map[string]interface{}{"packageName": "controllers"})
	if err != nil {
		fmt.Println("./" + ctrlName + " create failed.")
	} else {
		fmt.Println("./" + ctrlName + " created.")
	}
	//创建控制器服务
	ctrlProviderName, err := makeTplFile("app/providers/ControllerServiceProvider.go", app.CtrlProviderTemplate, map[string]interface{}{"moduleName": name})
	if err != nil {
		fmt.Println("./" + ctrlProviderName + " create failed.")
	} else {
		fmt.Println("./" + ctrlProviderName + " created.")
	}
	//创建控制器facades
	ctrlFacadeName, err := makeTplFile("app/facades/controller.go", app.CtrlFacadeTemplate, map[string]interface{}{"moduleName": name})
	if err != nil {
		fmt.Println("./" + ctrlFacadeName + " create failed.")
	} else {
		fmt.Println("./" + ctrlFacadeName + " created.")
	}
	//创建控制器的实现
	ctrlConcreteName, err := makeTplFile("app/concretes/controller.go", app.CtrlConcreteTemplate, map[string]interface{}{"moduleName": name})
	if err != nil {
		fmt.Println("./" + ctrlConcreteName + " create failed.")
	} else {
		fmt.Println("./" + ctrlConcreteName + " created.")
	}
	//创建config文件
	appConfigName, err := makeTplFile("config/app.go", app.AppConfigTemplate, map[string]interface{}{"moduleName": name})
	if err != nil {
		fmt.Println("./" + appConfigName + " create failed.")
	} else {
		fmt.Println("./" + appConfigName + " created.")
	}
	//创建logs目录
	os.MkdirAll("logs", 0755)
	//创建routes目录
	webRouteName, err := makeTplFile("routes/web.go", app.WebRouteTemplate, map[string]interface{}{"moduleName": name})
	if err != nil {
		fmt.Println("./" + webRouteName + " create failed.")
	} else {
		fmt.Println("./" + webRouteName + " created.")
	}
	//创建utils目录
	os.MkdirAll("utils", 0755)
	//创建views目录
	viewLayoutName, err := makeTplFile("views/layouts/layout.html", app.ViewLayoutTemplate, map[string]interface{}{})
	if err != nil {
		fmt.Println("./" + viewLayoutName + " create faild.")
	} else {
		fmt.Println("./" + viewLayoutName + " created.")
	}
	viewContentName, err := makeTplFile("views/home/index.html", app.ViewContentTemplate, map[string]interface{}{})
	if err != nil {
		fmt.Println("./" + viewContentName + " create faild.")
	} else {
		fmt.Println("./" + viewContentName + " created.")
	}
	//创建main.go文件
	mainName, err := makeTplFile("main.go", app.MainTemplate, map[string]interface{}{"moduleName": name})
	if err != nil {
		fmt.Println("./" + mainName + " create failed.")
	} else {
		fmt.Println("./" + mainName + " created.")
	}
	return name + " application created."
}
