package mysql57

var (
	Name        = "MySQL-5.7"
	Description = "MySQL 是最流行的关系型数据库管理系统之一，Oracle 旗下产品。"
	Slug        = "mysql57"
	Version     = "5.7.44"
	Requires    = []string{}
	Excludes    = []string{"mysql80"}
	Install     = `bash /www/panel/scripts/mysql/install.sh 57`
	Uninstall   = `bash /www/panel/scripts/mysql/uninstall.sh 57`
	Update      = `bash /www/panel/scripts/mysql/update.sh 57`
)
