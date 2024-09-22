package database

import (
    "database/sql"
    "web-back-end/custypes"
)

func GetArticleByArticleId(articleId int, article *custypes.Article, db *sql.DB) error {
    queryStr := "SELECT * FROM blog.articles WHERE article_id = $1"
    err := db.QueryRow(queryStr, articleId).Scan(&article.ArticleId, &article.Author, 
           &article.Title, &article.Content, &article.CreatedAt, &article.ModifiedAt,
           &article.Featured, &article.Category) 
    return err
}
