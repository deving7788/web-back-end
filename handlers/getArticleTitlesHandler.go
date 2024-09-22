package handlers

import (
    "net/http"
    "io"
    "strings"
    "encoding/json"
    "web-back-end/midware"
    "web-back-end/custypes"
    "web-back-end/database"
)

func GetArticleTitlesHandler(w http.ResponseWriter, r *http.Request) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "application/json")

    //hanlder preflight request
    if strings.ToLower(r.Method) == "options" {
        return
    }

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
