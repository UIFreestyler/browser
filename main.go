package main

import (
	"github.com/webview/webview_go"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Eco+ Browser")
	w.SetSize(1200, 720, webview.HintNone)

	html := `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Eco+</title>

<style>

body{
    margin:0;
    font-family: system-ui;
    background:#e9ebe3;
    display:flex;
    justify-content:center;
    align-items:center;
    height:100vh;
    overflow:hidden;
}

/* Center container */

.container{
    text-align:center;
    width:100%;
}

/* Logo */

.logo{
    color:#1c8a3a;
    font-size:80px;
    font-weight:800;
    letter-spacing:2px;
}

/* Search box */

.search-box{
    width:520px;
    max-width:90%;
    margin:25px auto;
}

.search-box input{
    width:100%;
    padding:16px 22px;
    border-radius:30px;
    border:none;
    outline:none;
    font-size:17px;
    box-shadow:0 4px 18px rgba(0,0,0,.08);
}

/* Icon row */

.icons{
    display:flex;
    justify-content:center;
    gap:28px;
    margin-top:28px;
}

.icon-circle{
    width:42px;
    height:42px;
    border-radius:50%;
    background:white;
    box-shadow:0 3px 12px rgba(0,0,0,.08);
}

/* Tree counter */

.counter{
    margin-top:35px;
    font-size:28px;
    font-weight:600;
    color:#2d2d2d;
}

/* Decorative bottom blobs */

.blobs{
    position:absolute;
    bottom:-120px;
    left:0;
    width:100%;
    height:300px;
    background:
    radial-gradient(circle at 20% 40%, #1c8a3a 0, transparent 40%),
    radial-gradient(circle at 40% 60%, #f4b400 0, transparent 45%),
    radial-gradient(circle at 70% 30%, #6fbf73 0, transparent 50%);
    opacity:0.8;
    pointer-events:none;
}

</style>
</head>

<body>

<div class="container">

<div class="logo">ECO+</div>

<div class="search-box">
<input placeholder="Search the web to plant trees..." />
</div>

<div class="icons">
<div class="icon-circle"></div>
<div class="icon-circle"></div>
<div class="icon-circle"></div>
<div class="icon-circle"></div>
<div class="icon-circle"></div>
</div>

<div class="counter">
🌱 156,583,192<br>
<span style="font-size:15px;font-weight:400;">
trees planted by the Eco+ community
</span>
</div>

</div>

<div class="blobs"></div>

</body>
</html>
`

	w.Navigate("data:text/html," + html)
	w.Run()
}