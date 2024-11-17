package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/submitContact", submitContactHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>校园导航</title>
    <style>
        body {
            background-color: #e6f7ff;
            font-family: Arial, sans-serif;
            text-align: center;
            margin: 0;
            padding: 50px;
        }
        h1 {
            color: #333;
        }
        .link-container {
            display: flex;
            justify-content: center;
            align-items: center;
            flex-direction: column;
        }
        .link {
            margin: 10px 0;
            padding: 10px 20px;
            background-color: #007BFF;
            color: white;
            text-decoration: none;
            border-radius: 5px;
            transition: background-color 0.3s ease;
        }
        .link:hover {
            background-color: #0056b3;
        }
        .time {
            margin: 20px 0;
        }
        .contact-form {
            margin: 20px 0;
        }
        .form-group {
            margin: 10px 0;
        }
        .form-control {
            padding: 10px;
            width: 300px;
        }
        .submit-button {
            padding: 10px 20px;
            background-color: #007BFF;
            color: white;
            border: none;
            border-radius: 5px;
            cursor: pointer;
        }
        .submit-button:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <h1>SUT校园导航</h1>
    <div class="time">现在是工大时间: %s</div>
    <div class="link-container" id="linkContainer">
        <a href="https://www.sut.edu.cn/" class="link">SUT</a>
        <a href="https://jwzx.sut.edu.cn/" class="link">教务处</a>
        <a href="https://library.sut.edu.cn/" class="link">图书馆</a>
        <a href="http://sygy.bysjy.com.cn/" class="link">招生就业网</a>
        <a href="https://ss.sut.edu.cn/index.htm" class="link">软件学院</a>
        <a href="https://pd.qq.com/s/2djxwptbi /index.htm" class="link">校园论坛</a>
        <a href="https://www.csdn.net/index.htm" class="link">CSDN</a>

    </div> 
    <div class="contact-form">
        <h2>请联系我们</h2>
        <form action="/submitContact" method="POST">
            <div class="form-group">
                <label for="name">姓名:</label>
                <input type="text" id="name" name="name" class="form-control" required>
            </div>
            <div class="form-group">
                <label for="email">邮箱:</label>
                <input type="email" id="email" name="email" class="form-control" required>
            </div>
            <div class="form-group">
                <label for="message">留言:</label>
                <textarea id="message" name="message" class="form-control" rows="4" required></textarea>
            </div>
            <button type="submit" class="submit-button">提交</button>
        </form>
    </div>
</body>
</html>
`, currentTime)
	fmt.Fprint(w, html)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "contact.html")
}

func submitContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		content := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s\n", name, email, message)
		err := ioutil.WriteFile("contact.txt", []byte(content), 0644)
		if err != nil {
			http.Error(w, "提交失败", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "无效方法", http.StatusMethodNotAllowed)
	}
}
