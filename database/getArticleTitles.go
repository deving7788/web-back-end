package database

import (
    "database/sql"
    "web-back-end/custypes"
)

func GetArticleTitles(articles *[]custypes.Article, db *sql.DB) (error) {
    queryStr := "SELECT article_id, title, created_at, modified_at, featured FROM blog.articles"
    rows, err := db.Query(queryStr)
    if err != nil {
        return err
    }

    for rows.Next() {
        var article custypes.Article
        err = rows.Scan(&article.ArticleId, &article.Title, &article.CreatedAt, &article.ModifiedAt, &article.Featured)
        if err != nil {
            return err
        }
        *articles = append(*articles, article)
    }

    return nil

}
