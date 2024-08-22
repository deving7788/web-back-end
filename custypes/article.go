package custypes

import (
    "time"
)

type Article struct {
    ArticleId int `json:"articleId,omitempty"`
    Author string `json:"author,omitempty"`
    Title string `json:"title,omitempty"`
    Content string `json:"content,omitempty"`
    CreatedAt time.Time `json:"createdAt,omitempty"`
    ModifiedAt time.Time `json:"modifiedAt,omitempty"`
    Featured bool `json:"featured"`
}
