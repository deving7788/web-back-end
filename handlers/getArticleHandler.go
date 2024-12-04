package handlers

import (
    "net/http"
    "io"
    "errors"
    "database/sql"
    "encoding/json"
    "strconv"
    "web-back-end/custypes"
    "web-back-end/database"
)

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
    //get article id from request url
    articleIdStr := r.URL.Query()["articleId"][0]
    articleId, err := strconv.Atoi(articleIdStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //get article from database
    var article custypes.Article
    err = database.GetArticleByArticleId(articleId, &article, database.Blogdb)
    if err != nil {
        switch {
        case errors.Is(err, sql.ErrNoRows):
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        default:
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
    //json retrieved article data
    resJsonBytes, err := json.Marshal(article)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    resJsonStr := string(resJsonBytes)
    _, err = io.WriteString(w, resJsonStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

