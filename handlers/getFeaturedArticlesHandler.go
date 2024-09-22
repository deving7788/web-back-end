package handlers

import (
    "net/http"
    "strings"
    "encoding/json"
    "io"
    "web-back-end/midware"
    "web-back-end/custypes"
    "web-back-end/database"
)

func GetFeaturedArticlesHandler(w http.ResponseWriter, r *http.Request) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "application/json")
    //handle pre-flight request
    if strings.ToLower(r.Method) == "options" {
        return
    }
    //get featured articles from database
    var articles []custypes.Article
    err := database.GetFeaturedArticles(&articles, database.Blogdb)
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
