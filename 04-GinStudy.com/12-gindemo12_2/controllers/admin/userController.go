package admin

import (
	"fmt"
	"gindemo12/models"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "用户列表--")
	// con.success(c)
}
func (con UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

/*
1、获取上传的文件
2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
3、创建图片保存目录  static/upload/20200623
4、生成文件名称和文件保存的目录
5、执行上传
*/
func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	// 1、获取上传的文件
	file, err := c.FormFile("face")
	if err == nil {
		// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
		extName := path.Ext(file.Filename)

		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			c.String(200, "上传的文件类型不合法")
			return
		}
		// 3、创建图片保存目录  static/upload/20210624
		day := models.GetDay()
		dir := "./static/upload/" + day
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			fmt.Println(err)
			c.String(200, "MkdirAll失败")
			return
		}
		// 4、生成文件名称和文件保存的目录   111111111111.jpeg
		fileName := strconv.FormatInt(models.GetUnix(), 10) + extName
		// 5、执行上传
		dst := path.Join(dir, fileName)
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	})
}
