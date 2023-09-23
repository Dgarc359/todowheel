package util

import (
	"net/http"
    db "todowheel-backend/database"
)


func GetTasks(w http.ResponseWriter, r *http.Request, conn *db.SqliteDatabase) {
    conn.GetTasks()
}
