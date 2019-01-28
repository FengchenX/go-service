package jwt

import (
	"agfun/auth/entity"
	"agfun/dbcentral/etcddb"
	"agfun/dbcentral/mysqldb"
	"agfun/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
)

func AuthMiddleWare(authDB *gorm.DB, cli *etcddb.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开关
		if true {
			c.Next()
			return
		}
		session := c.GetHeader("session")
		var ID string
		e := cli.Get(session, &ID)
		if e != nil {
			util.Fail(c, e)
			c.Abort()
			return
		}
		var userRoles []*entity.UserRole
		db := authDB.Find(&userRoles, "user_id = ?", ID)
		if db.Error != nil {
			util.Fail(c, db.Error)
			c.Abort()
			return
		}
		verb := entity.Verb{
			ID:   "",
			Name: "GET",
		}
		switch c.Request.Method {
		case "GET":
			verb.Name = "GET"
		case "POST":
			verb.Name = "POST"
		case "UPDATE":
			verb.Name = "UPDATE"
		case "DELETE":
			verb.Name = "DELETE"
		}
		db = authDB.Where("name=?", verb.Name).Find(&verb)
		if db.Error != nil {
			util.Fail(c, db.Error)
			c.Abort()
			return
		}

		for _, userRole := range userRoles {
			rule := entity.Rule{
				RoleID: userRole.RoleID,
				VerbID: verb.ID,
			}
			var resources []*entity.Resource
			db := authDB.Select("resources.name, resources.layer").
				Table("rules").
				Joins("INNER JOIN resources ON rules.resource_id = resources.id").
				Where("rules.role_id =? AND rules.verb_id=?", rule.RoleID, rule.VerbID).
				Find(&resources)
			if db.Error != nil {
				continue
			}

			b, e := IsResursIncludePath(resources, c.Request.RequestURI)
			if e != nil {
				util.Fail(c, e)
				return
			}
			if b {
				c.Next()
				return
			}
		}
		b, e := IsResources(c.Request.RequestURI)
		if e != nil {
			util.Fail(c, e)
			return
		}
		if !b {
			c.Next()
			return
		}
		util.Fail(c, fmt.Errorf("no auth"))
		c.Abort()
	}
}

func GetResourceParts(parentID string) (string, error) {
	if parentID == "0" {
		return "", nil
	}
	resource := entity.Resource{}
	db := mysqldb.GetAuthDB().Where("id = ?", parentID).First(&resource)
	if db.Error != nil {
		return "", db.Error
	}
	url := ""
	if len(resource.Name) > 0 {
		url = fmt.Sprintf("%s/%s", resource.Name, url)
	}
	if len(resource.Type) > 0 {
		url = fmt.Sprintf("%s/%s", resource.Type, url)
	}
	s, e := GetResourceParts(resource.ParentID)
	if e != nil {
		return "", e
	}
	if len(s) > 0 {
		url = fmt.Sprintf("%s/%s", s, url)
	}
	return url, nil
}

func IsResources(path string) (bool, error) {
	var resources []*entity.Resource
	mysqldb.GetAuthDB().Where("name=?", "").Find(&resources)
	b, e := IsResursIncludePath(resources, path)
	return b, e
}

func IsResursIncludePath(resources []*entity.Resource, path string) (bool, error) {
Label:
	for _, resource := range resources {
		resourceParts, e := GetResourceParts(resource.ParentID)
		if e != nil {
			continue
		}
		if len(resource.Type) > 0 {
			resourceParts = fmt.Sprintf("%s/%s", resourceParts, resource.Type)
		}
		if len(resource.Name) > 0 {
			resourceParts = fmt.Sprintf("%s/%s", resourceParts, resource.Name)
		}
		list := strings.Split(resourceParts, "/")
		if len(list) <= 1 {
			return false, nil
		}
		for i := 1; i < len(list); i++ {
			if !strings.Contains(path, list[i]) {
				continue Label
			}
		}
		return true, nil
	}
	return false, nil
}
