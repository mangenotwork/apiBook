package define

import "fmt"

var (
	DirNameExistErr         = fmt.Errorf("目录名已存在")
	ProjectExistErr         = fmt.Errorf("项目名已存在")
	NoPermission            = fmt.Errorf("没有权限")
	ProjectPublicNotAddUser = fmt.Errorf("公有项目无需添加协助者")
	NotDelProjectOwner      = fmt.Errorf("不能移除项目创建者")
)
