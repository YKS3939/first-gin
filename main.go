package main

import (
	"database/sql" // Go 语言自带的数据库标准库
	"fmt"
	"log"

	// 引入我们刚刚下载的纯 Go 版 SQLite 驱动
	// 前面的下划线 "_" 是一种魔法：意思是“我只要它在底层偷偷注册，但我代码里不直接调用它的名字”
	_ "modernc.org/sqlite"
)

func main() {
	// ==========================================
	// 1. 创世：连接（或创建）数据库文件
	// ==========================================
	// 这行代码执行后，你的代码文件夹里会神奇地多出一个叫 "my_first_db.db" 的文件！
	db, err := sql.Open("sqlite", "my_first_db.db")
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}
	defer db.Close() // 礼貌退场：程序结束时关掉数据库连接

	// ==========================================
	// 2. 开天辟地：用 SQL 语句建立一张 "users" 表
	// ==========================================
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,  -- 让数据库自动给用户发 ID（1,2,3...）
		name TEXT,                             -- 名字是文本
		skill TEXT                             -- 技能也是文本
	);`
	
	_, err = db.Exec(createTableSQL) // Exec 就是让数据库执行这句 SQL
	if err != nil {
		log.Fatal("建表失败: ", err)
	}
	fmt.Println("✅ 数据库文件创建成功，[users] 表已准备就绪！")

	// ==========================================
	// 3. 存入数据 (Insert)
	// ==========================================
	// SQL 里的问号 "?" 是占位符，专门用来防止黑客的 SQL 注入攻击
	insertSQL := `INSERT INTO users (name, skill) VALUES (?, ?)`
	
	// 把 "Kobe" 和他的技能填进问号里
	_, err = db.Exec(insertSQL, "Kobe", "Go与数据库魔法")
	if err != nil {
		log.Fatal("插入数据失败: ", err)
	}
	// 我们再偷偷插入一个人
	db.Exec(insertSQL, "Alice", "前端架构师")
	fmt.Println("✅ 成功将 Kobe 和 Alice 写入硬盘，实现数据永生！")

	// ==========================================
	// 4. 提取数据 (Select)
	// ==========================================
	// 告诉数据库：把 users 表里所有的 id, name, skill 都交出来！
	rows, err := db.Query("SELECT id, name, skill FROM users")
	if err != nil {
		log.Fatal("查询失败: ", err)
	}
	defer rows.Close()

	fmt.Println("\n--- 🔍 开始扫描数据库里的数据 ---")
	
	// Rows 就像一个游标，用 Next() 一行一行地往下读
	for rows.Next() {
		// 准备三个容器，用来接住数据库扔出来的数据
		var dbID int
		var dbName string
		var dbSkill string

		// Scan：把这一行的数据，按顺序装进我们准备好的容器里
		rows.Scan(&dbID, &dbName, &dbSkill)
		
		fmt.Printf("从硬盘读出 -> ID: %d | 姓名: %s | 技能: %s\n", dbID, dbName, dbSkill)
	}
	fmt.Println("---------------------------------")
}