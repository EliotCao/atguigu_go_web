package dao

import (
	"atguigu_go_web/src/04/util"
	"atguigu_go_web/src/07-bookstore/model"
	"net/http"
)

func AddSession(s *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
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

func GetSession(sid string) (*model.Session, error) {
	sqlStr := "select seesion_id,username,userID from sessions where session_id = ?"
	inStmt, err := util.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	session := &model.Session{}
	row := inStmt.QueryRow(sid)
	row.Scan(&session.SessionID, &session.Username, &session.UserID)
	if err != nil {
		return nil, err
	}
	return session,nil
}

func IsLogin(r *http.Request) (bool, string) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			return true, session.Username
		}
	}
	return false, ""
}