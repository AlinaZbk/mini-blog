package service

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/AlinaZbk/mini-blog.git/model"
)

var (
	posts        = make([]*model.Post, 0)
	nextID int64 = 1
)

func CreatePost(req model.CreatePostRequest) (*model.Post, error) {
	title := strings.TrimSpace(req.Title)
	content := strings.TrimSpace(req.Content)
	if title == "" || utf8.RuneCountInString(title) > 200 {
		return nil, errors.New("title is required and must be <= 200 chars")
	}
	if content == "" {
		return nil, errors.New("content is required")
	}
	now := time.Now()
	p := &model.Post{
		ID:        nextID,
		Title:     title,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}
	nextID++
	posts = append(posts, p)
	return p, nil
}

func GetListPosts() []*model.Post { return posts }

func GetPost(id int64) (*model.Post, error) {
	for _, p := range posts {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, errors.New("post not found")
}

func DeletePost(id int64) error {
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			return nil
		}
	}
	return errors.New("post not found")
}

func UpdatePost(id int64, req model.UpdatePostRequest) (*model.Post, error) {
	title := strings.TrimSpace(req.Title)
	content := strings.TrimSpace(req.Content)
	if title == "" || utf8.RuneCountInString(title) > 200 {
		return nil, errors.New("title is required and must be <= 200 chars")
	}
	if content == "" {
		return nil, errors.New("content is required")
	}

	for _, post := range posts {
		if post.ID == id {
			post.Title = req.Title
			post.Content = req.Content
			post.UpdatedAt = time.Now()
			return post, nil
		}
	}
	return nil, errors.New("post not found")
}
