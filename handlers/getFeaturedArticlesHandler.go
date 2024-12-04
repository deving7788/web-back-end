package handlers

import (
    "net/http"
    "encoding/json"
    "io"
    "web-back-end/custypes"
    "web-back-end/database"
)

func GetFeaturedArticlesHandler(w http.ResponseWriter, r *http.Request) {
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
