package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Writer.Write([]byte(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dym树洞空间，不开心就来这吧!</title>
</head>
<body>
    <div style="
           background-color:lightskyblue;
            text-align: center;
           padding:40px;
    ">
     <img href = "https://wx1.sinaimg.cn/mw690/006ynZgoly1hs0pmkgffcj30t80t8gqz.jpg"><img>
    <p>把每天发生的快乐的事情和难过的事情都写下来吧，总有人会倾听你的故事</p>
    </div>
    <div style="
         max-width: 700px;
         margin: 30px auto;
         padding: 15px;
         line-height:1.7">
        <label for="username">用户名：</label>
        <input type="text" id="username" name="username" required><br>

        <label for="email">邮箱：</label>
        <input type="email" id="email" name="email" required><br>

        <label for="password">密码：</label>
        <input type="password" id="password" name="password" required><br>

        <input type="submit" value="提交">
    </form>
        <p>在线聊天</p >
        <p>视频通话</p >
        <p>信件传递</p >
    </div>
    <div style="
           background-color:darkgray;
           text-align:center;
           padding:40px;
           font-size:12px;">
        <p>我的网站</p>
            <a href="https://www.sut.edu.cn/">友情链接</a>
    </div>

    <
</body>
</html>
`))
	})
	router.POST("/submit", func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"email":    email,
			"password": password,
		})
	})
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
