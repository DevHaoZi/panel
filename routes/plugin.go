package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"panel/app/http/controllers/plugins/fail2ban"
	"panel/app/http/controllers/plugins/mysql57"
	"panel/app/http/controllers/plugins/mysql80"
	"panel/app/http/controllers/plugins/openresty"
	"panel/app/http/controllers/plugins/php74"
	"panel/app/http/controllers/plugins/php80"
	"panel/app/http/controllers/plugins/php81"
	"panel/app/http/controllers/plugins/php82"
	"panel/app/http/controllers/plugins/phpmyadmin"
	"panel/app/http/controllers/plugins/postgresql15"
	"panel/app/http/controllers/plugins/postgresql16"
	"panel/app/http/controllers/plugins/pureftpd"
	"panel/app/http/controllers/plugins/redis"
	"panel/app/http/controllers/plugins/s3fs"
	"panel/app/http/controllers/plugins/supervisor"
	"panel/app/http/controllers/plugins/toolbox"
	"panel/app/http/middleware"
)

// Plugin 加载插件路由
func Plugin() {
	facades.Route().Prefix("api/plugins/openresty").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		openRestyController := openresty.NewOpenrestyController()
		route.Get("status", openRestyController.Status)
		route.Post("reload", openRestyController.Reload)
		route.Post("start", openRestyController.Start)
		route.Post("stop", openRestyController.Stop)
		route.Post("restart", openRestyController.Restart)
		route.Get("load", openRestyController.Load)
		route.Get("config", openRestyController.GetConfig)
		route.Post("config", openRestyController.SaveConfig)
		route.Get("errorLog", openRestyController.ErrorLog)
		route.Post("clearErrorLog", openRestyController.ClearErrorLog)
	})
	facades.Route().Prefix("api/plugins/mysql57").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		mysql57Controller := mysql57.NewMysql57Controller()
		route.Get("status", mysql57Controller.Status)
		route.Post("reload", mysql57Controller.Reload)
		route.Post("start", mysql57Controller.Start)
		route.Post("stop", mysql57Controller.Stop)
		route.Post("restart", mysql57Controller.Restart)
		route.Get("load", mysql57Controller.Load)
		route.Get("config", mysql57Controller.GetConfig)
		route.Post("config", mysql57Controller.SaveConfig)
		route.Get("errorLog", mysql57Controller.ErrorLog)
		route.Post("clearErrorLog", mysql57Controller.ClearErrorLog)
		route.Get("slowLog", mysql57Controller.SlowLog)
		route.Post("clearSlowLog", mysql57Controller.ClearSlowLog)
		route.Get("rootPassword", mysql57Controller.GetRootPassword)
		route.Post("rootPassword", mysql57Controller.SetRootPassword)
		route.Get("databases", mysql57Controller.DatabaseList)
		route.Post("databases", mysql57Controller.AddDatabase)
		route.Delete("databases/{database}", mysql57Controller.DeleteDatabase)
		route.Get("backups", mysql57Controller.BackupList)
		route.Post("backups", mysql57Controller.CreateBackup)
		route.Put("backups", mysql57Controller.UploadBackup)
		route.Delete("backups/{name}", mysql57Controller.DeleteBackup)
		route.Post("backups/restore", mysql57Controller.RestoreBackup)
		route.Get("users", mysql57Controller.UserList)
		route.Post("users", mysql57Controller.AddUser)
		route.Delete("users/{user}", mysql57Controller.DeleteUser)
		route.Post("users/password", mysql57Controller.SetUserPassword)
		route.Post("users/privileges", mysql57Controller.SetUserPrivileges)
	})
	facades.Route().Prefix("api/plugins/mysql80").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		mysql80Controller := mysql80.NewMysql80Controller()
		route.Get("status", mysql80Controller.Status)
		route.Post("reload", mysql80Controller.Reload)
		route.Post("start", mysql80Controller.Start)
		route.Post("stop", mysql80Controller.Stop)
		route.Post("restart", mysql80Controller.Restart)
		route.Get("load", mysql80Controller.Load)
		route.Get("config", mysql80Controller.GetConfig)
		route.Post("config", mysql80Controller.SaveConfig)
		route.Get("errorLog", mysql80Controller.ErrorLog)
		route.Post("clearErrorLog", mysql80Controller.ClearErrorLog)
		route.Get("slowLog", mysql80Controller.SlowLog)
		route.Post("clearSlowLog", mysql80Controller.ClearSlowLog)
		route.Get("rootPassword", mysql80Controller.GetRootPassword)
		route.Post("rootPassword", mysql80Controller.SetRootPassword)
		route.Get("databases", mysql80Controller.DatabaseList)
		route.Post("databases", mysql80Controller.AddDatabase)
		route.Delete("databases/{database}", mysql80Controller.DeleteDatabase)
		route.Get("backups", mysql80Controller.BackupList)
		route.Post("backups", mysql80Controller.CreateBackup)
		route.Put("backups", mysql80Controller.UploadBackup)
		route.Delete("backups/{name}", mysql80Controller.DeleteBackup)
		route.Post("backups/restore", mysql80Controller.RestoreBackup)
		route.Get("users", mysql80Controller.UserList)
		route.Post("users", mysql80Controller.AddUser)
		route.Delete("users/{user}", mysql80Controller.DeleteUser)
		route.Post("users/password", mysql80Controller.SetUserPassword)
		route.Post("users/privileges", mysql80Controller.SetUserPrivileges)
	})
	facades.Route().Prefix("api/plugins/postgresql15").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		postgresql15Controller := postgresql15.NewPostgresql15Controller()
		route.Get("status", postgresql15Controller.Status)
		route.Post("reload", postgresql15Controller.Reload)
		route.Post("start", postgresql15Controller.Start)
		route.Post("stop", postgresql15Controller.Stop)
		route.Post("restart", postgresql15Controller.Restart)
		route.Get("load", postgresql15Controller.Load)
		route.Get("config", postgresql15Controller.GetConfig)
		route.Post("config", postgresql15Controller.SaveConfig)
		route.Get("userConfig", postgresql15Controller.GetUserConfig)
		route.Post("userConfig", postgresql15Controller.SaveUserConfig)
		route.Get("log", postgresql15Controller.Log)
		route.Post("clearLog", postgresql15Controller.ClearLog)
		route.Get("database", postgresql15Controller.DatabaseList)
		route.Post("addDatabase", postgresql15Controller.AddDatabase)
		route.Post("deleteDatabase", postgresql15Controller.DeleteDatabase)
		route.Get("backup", postgresql15Controller.BackupList)
		route.Post("createBackup", postgresql15Controller.CreateBackup)
		route.Post("uploadBackup", postgresql15Controller.UploadBackup)
		route.Post("deleteBackup", postgresql15Controller.DeleteBackup)
		route.Post("restoreBackup", postgresql15Controller.RestoreBackup)
		route.Get("user", postgresql15Controller.UserList)
		route.Post("addUser", postgresql15Controller.AddUser)
		route.Post("deleteUser", postgresql15Controller.DeleteUser)
		route.Post("userPassword", postgresql15Controller.SetUserPassword)
	})
	facades.Route().Prefix("api/plugins/postgresql16").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		postgresql16Controller := postgresql16.NewPostgresql16Controller()
		route.Get("status", postgresql16Controller.Status)
		route.Post("reload", postgresql16Controller.Reload)
		route.Post("start", postgresql16Controller.Start)
		route.Post("stop", postgresql16Controller.Stop)
		route.Post("restart", postgresql16Controller.Restart)
		route.Get("load", postgresql16Controller.Load)
		route.Get("config", postgresql16Controller.GetConfig)
		route.Post("config", postgresql16Controller.SaveConfig)
		route.Get("userConfig", postgresql16Controller.GetUserConfig)
		route.Post("userConfig", postgresql16Controller.SaveUserConfig)
		route.Get("log", postgresql16Controller.Log)
		route.Post("clearLog", postgresql16Controller.ClearLog)
		route.Get("database", postgresql16Controller.DatabaseList)
		route.Post("addDatabase", postgresql16Controller.AddDatabase)
		route.Post("deleteDatabase", postgresql16Controller.DeleteDatabase)
		route.Get("backup", postgresql16Controller.BackupList)
		route.Post("createBackup", postgresql16Controller.CreateBackup)
		route.Post("uploadBackup", postgresql16Controller.UploadBackup)
		route.Post("deleteBackup", postgresql16Controller.DeleteBackup)
		route.Post("restoreBackup", postgresql16Controller.RestoreBackup)
		route.Get("user", postgresql16Controller.UserList)
		route.Post("addUser", postgresql16Controller.AddUser)
		route.Post("deleteUser", postgresql16Controller.DeleteUser)
		route.Post("userPassword", postgresql16Controller.SetUserPassword)
	})
	facades.Route().Prefix("api/plugins/php74").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		php74Controller := php74.NewPhp74Controller()
		route.Get("status", php74Controller.Status)
		route.Post("reload", php74Controller.Reload)
		route.Post("start", php74Controller.Start)
		route.Post("stop", php74Controller.Stop)
		route.Post("restart", php74Controller.Restart)
		route.Get("load", php74Controller.Load)
		route.Get("config", php74Controller.GetConfig)
		route.Post("config", php74Controller.SaveConfig)
		route.Get("errorLog", php74Controller.ErrorLog)
		route.Get("slowLog", php74Controller.SlowLog)
		route.Post("clearErrorLog", php74Controller.ClearErrorLog)
		route.Post("clearSlowLog", php74Controller.ClearSlowLog)
		route.Get("extensions", php74Controller.GetExtensionList)
		route.Post("installExtension", php74Controller.InstallExtension)
		route.Post("uninstallExtension", php74Controller.UninstallExtension)
	})
	facades.Route().Prefix("api/plugins/php80").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		php80Controller := php80.NewPhp80Controller()
		route.Get("status", php80Controller.Status)
		route.Post("reload", php80Controller.Reload)
		route.Post("start", php80Controller.Start)
		route.Post("stop", php80Controller.Stop)
		route.Post("restart", php80Controller.Restart)
		route.Get("load", php80Controller.Load)
		route.Get("config", php80Controller.GetConfig)
		route.Post("config", php80Controller.SaveConfig)
		route.Get("errorLog", php80Controller.ErrorLog)
		route.Get("slowLog", php80Controller.SlowLog)
		route.Post("clearErrorLog", php80Controller.ClearErrorLog)
		route.Post("clearSlowLog", php80Controller.ClearSlowLog)
		route.Get("extensions", php80Controller.GetExtensionList)
		route.Post("installExtension", php80Controller.InstallExtension)
		route.Post("uninstallExtension", php80Controller.UninstallExtension)
	})
	facades.Route().Prefix("api/plugins/php81").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		php81Controller := php81.NewPhp81Controller()
		route.Get("status", php81Controller.Status)
		route.Post("reload", php81Controller.Reload)
		route.Post("start", php81Controller.Start)
		route.Post("stop", php81Controller.Stop)
		route.Post("restart", php81Controller.Restart)
		route.Get("load", php81Controller.Load)
		route.Get("config", php81Controller.GetConfig)
		route.Post("config", php81Controller.SaveConfig)
		route.Get("errorLog", php81Controller.ErrorLog)
		route.Get("slowLog", php81Controller.SlowLog)
		route.Post("clearErrorLog", php81Controller.ClearErrorLog)
		route.Post("clearSlowLog", php81Controller.ClearSlowLog)
		route.Get("extensions", php81Controller.GetExtensionList)
		route.Post("installExtension", php81Controller.InstallExtension)
		route.Post("uninstallExtension", php81Controller.UninstallExtension)
	})
	facades.Route().Prefix("api/plugins/php82").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		php82Controller := php82.NewPhp82Controller()
		route.Get("status", php82Controller.Status)
		route.Post("reload", php82Controller.Reload)
		route.Post("start", php82Controller.Start)
		route.Post("stop", php82Controller.Stop)
		route.Post("restart", php82Controller.Restart)
		route.Get("load", php82Controller.Load)
		route.Get("config", php82Controller.GetConfig)
		route.Post("config", php82Controller.SaveConfig)
		route.Get("errorLog", php82Controller.ErrorLog)
		route.Get("slowLog", php82Controller.SlowLog)
		route.Post("clearErrorLog", php82Controller.ClearErrorLog)
		route.Post("clearSlowLog", php82Controller.ClearSlowLog)
		route.Get("extensions", php82Controller.GetExtensionList)
		route.Post("installExtension", php82Controller.InstallExtension)
		route.Post("uninstallExtension", php82Controller.UninstallExtension)
	})
	facades.Route().Prefix("api/plugins/phpmyadmin").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		phpMyAdminController := phpmyadmin.NewPhpMyAdminController()
		route.Get("info", phpMyAdminController.Info)
		route.Post("port", phpMyAdminController.SetPort)
	})
	facades.Route().Prefix("api/plugins/pureftpd").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		pureFtpdController := pureftpd.NewPureFtpdController()
		route.Get("status", pureFtpdController.Status)
		route.Post("reload", pureFtpdController.Reload)
		route.Post("start", pureFtpdController.Start)
		route.Post("stop", pureFtpdController.Stop)
		route.Post("restart", pureFtpdController.Restart)
		route.Get("list", pureFtpdController.List)
		route.Post("add", pureFtpdController.Add)
		route.Post("delete", pureFtpdController.Delete)
		route.Post("changePassword", pureFtpdController.ChangePassword)
		route.Get("port", pureFtpdController.GetPort)
		route.Post("port", pureFtpdController.SetPort)
	})
	facades.Route().Prefix("api/plugins/redis").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		redisController := redis.NewRedisController()
		route.Get("status", redisController.Status)
		route.Post("reload", redisController.Reload)
		route.Post("start", redisController.Start)
		route.Post("stop", redisController.Stop)
		route.Post("restart", redisController.Restart)
		route.Get("load", redisController.Load)
		route.Get("config", redisController.GetConfig)
		route.Post("config", redisController.SaveConfig)
	})
	facades.Route().Prefix("api/plugins/s3fs").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		s3fsController := s3fs.NewS3fsController()
		route.Get("list", s3fsController.List)
		route.Post("add", s3fsController.Add)
		route.Post("delete", s3fsController.Delete)
	})
	facades.Route().Prefix("api/plugins/supervisor").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		supervisorController := supervisor.NewSupervisorController()
		route.Get("status", supervisorController.Status)
		route.Post("start", supervisorController.Start)
		route.Post("stop", supervisorController.Stop)
		route.Post("restart", supervisorController.Restart)
		route.Post("reload", supervisorController.Reload)
		route.Get("log", supervisorController.Log)
		route.Post("clearLog", supervisorController.ClearLog)
		route.Get("config", supervisorController.Config)
		route.Post("config", supervisorController.SaveConfig)
		route.Get("processes", supervisorController.Processes)
		route.Post("startProcess", supervisorController.StartProcess)
		route.Post("stopProcess", supervisorController.StopProcess)
		route.Post("restartProcess", supervisorController.RestartProcess)
		route.Get("processLog", supervisorController.ProcessLog)
		route.Post("clearProcessLog", supervisorController.ClearProcessLog)
		route.Get("processConfig", supervisorController.ProcessConfig)
		route.Post("processConfig", supervisorController.SaveProcessConfig)
		route.Post("deleteProcess", supervisorController.DeleteProcess)
		route.Post("addProcess", supervisorController.AddProcess)

	})
	facades.Route().Prefix("api/plugins/fail2ban").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		fail2banController := fail2ban.NewFail2banController()
		route.Get("status", fail2banController.Status)
		route.Post("start", fail2banController.Start)
		route.Post("stop", fail2banController.Stop)
		route.Post("restart", fail2banController.Restart)
		route.Post("reload", fail2banController.Reload)
		route.Get("list", fail2banController.List)
		route.Post("add", fail2banController.Add)
		route.Post("delete", fail2banController.Delete)
		route.Get("ban", fail2banController.BanList)
		route.Post("unban", fail2banController.Unban)
		route.Post("whiteList", fail2banController.SetWhiteList)
		route.Get("whiteList", fail2banController.GetWhiteList)
	})
	facades.Route().Prefix("api/plugins/toolbox").Middleware(middleware.Jwt()).Group(func(route route.Router) {
		toolboxController := toolbox.NewToolBoxController()
		route.Get("dns", toolboxController.GetDNS)
		route.Post("dns", toolboxController.SetDNS)
		route.Get("swap", toolboxController.GetSWAP)
		route.Post("swap", toolboxController.SetSWAP)
		route.Get("timezone", toolboxController.GetTimezone)
		route.Post("timezone", toolboxController.SetTimezone)
		route.Get("hosts", toolboxController.GetHosts)
		route.Post("hosts", toolboxController.SetHosts)
		route.Post("rootPassword", toolboxController.SetRootPassword)
	})
}
