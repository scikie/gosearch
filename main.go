package main
import (
	"gosearch/model"
	"github.com/sirupsen/logrus"
	"gosearch/routes"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ // 输出格式 logfmt 风格
		DisableColors: true,
		FullTimestamp: false,
	})
	// 设置日志等级为 Debug
	// Debug 级别以及更高级别的日志（如 Info、Warning、Error、Fatal 和 Panic）都会被记录。
	logrus.SetLevel(logrus.DebugLevel)
	// 引用数据库
	model.InitDb()
	// 引入路由组件
	routes.InitRouter()

}
