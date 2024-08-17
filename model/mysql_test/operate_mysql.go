package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type User struct {
	Id       int64          `db:"id"`
	Name     sql.NullString `db:"name"` // The username
	Type     int64          `db:"type"` // The user type, 0:normal,1:vip, for test golang keyword
	CreateAt sql.NullTime   `db:"create_at"`
	UpdateAt time.Time      `db:"update_at"`
}

func queryUserById(Id int64, conn sqlx.SqlConn) (*User, error) {
	query := "SELECT id, name, type, create_at, update_at FROM user WHERE id=?"
	var u User
	err := conn.QueryRow(&u, query, Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No rows found
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}
	return &u, nil
}

// _, err := conn.ExecCtx(context.Background(), "update user set type = ? where name = ?", 2, "test")
func updateUserById(Id int64, userName string, conn sqlx.SqlConn) error {
	update := "UPDATE user SET name=? , type=? where id=?"
	_, err := conn.Exec(update, userName, 22, Id)
	if err != nil {
		return fmt.Errorf("failed to update user name: %w", err)
	}
	return nil
}

// _, err := conn.ExecCtx(context.Background(), "delete from user where `id` = ?", 1)
func DeleteUserById(Id int64, conn sqlx.SqlConn) error {
	delete := "DELETE FROM user where id=?"
	_, err := conn.Exec(delete, Id)
	if err != nil {
		return fmt.Errorf("failed to delete user name: %w", err)
	}
	return nil
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 需要自行确保 dsn 中的 host 账号 密码都配置正确
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewMysql(dsn)

	r, err := conn.ExecCtx(context.Background(), "insert into user (id,type,mobile, name) values (?,?,?,?)", 1, 1, 222, "user2")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.RowsAffected())

	user, err := queryUserById(1, conn)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if user != nil {
		fmt.Printf("User: %+v\n", user)
	} else {
		fmt.Println("User not found")
	}

	// 修改
	/*
		err = updateUserById(1, "张三", conn)
		err = updateUserById(1, "张三", conn)
		if err != nil {
			fmt.Println("Error:", err)
		}
	*/
	err = DeleteUserById(1, conn)

}
