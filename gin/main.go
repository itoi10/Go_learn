package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "./assets")

	dbInit()

	// 一覧
	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})
	// 作成
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})
	// 詳細
	router.GET("/detail/:id", func(ctx *gin.Context) {
		sId := ctx.Param("id")
		id, err := strconv.Atoi(sId)
		if err != nil {
			// 数値変換失敗
			panic("Error")
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})
	// 更新
	router.POST("/update/:id", func(ctx *gin.Context) {
		sId := ctx.Param("id")
		id, err := strconv.Atoi(sId)
		if err != nil {
			panic("Error")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbUpdate(id, text, status)
		ctx.Redirect(302, "/")
	})
	// 削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		sId := ctx.Param("id")
		id, err := strconv.Atoi(sId)
		if err != nil {
			panic("Error")
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})
	// 削除
	router.POST("/delete/:id", func(ctx *gin.Context) {
		sId := ctx.Param("id")
		id, err := strconv.Atoi(sId)
		if err != nil {
			panic("Error")
		}
		dbDelete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
