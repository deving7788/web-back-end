package database

import (
    "database/sql"
    "web-back-end/custypes"
)

func GetAllArticles(articles *[]custypes.Article, db *sql.DB) (error) {
    queryStr := "SELECT * FROM blog.articles"
    rows, err := db.Query(queryStr)
    if err != nil {
        return err
    }

    for rows.Next() {
        var article custypes.Article
        err = rows.Scan(&article.ArticleId, &article.Author, &article.Title, &article.Content, &article.CreatedAt, &article.ModifiedAt, &article.Featured)
        if err != nil {
            return err
        }
        *articles = append(*articles, article)
    }

    return nil

}
