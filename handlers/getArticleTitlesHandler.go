package handlers

import (
    "net/http"
    "io"
    "encoding/json"
    "web-back-end/custypes"
    "web-back-end/database"
)

func GetArticleTitlesHandler(w http.ResponseWriter, r *http.Request) {

    //get article titles from database
    var articles []custypes.Article
    err := database.GetArticleTitles(&articles, database.Blogdb)    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    //form and send json response
    resJsonBytes, err := json.Marshal(articles)
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
