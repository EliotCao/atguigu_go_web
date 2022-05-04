package dao

import (
	"atguigu_go_web/src/04/util"
	"atguigu_go_web/src/07-bookstore/model"
)

func AddSession(s *model.Session) error {
	sqlStr := "insert into session values(?,?,?)"
	_, err := util.Db.Exec(sqlStr, s.SessionID, s.Username, s.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sid string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := util.Db.Exec(sqlStr, sid)
	if err != nil {
		 return err
	}
	return nil
}
