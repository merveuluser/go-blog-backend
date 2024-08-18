package helpers

import "blog-backend/models"

func ValidateJSONPost(action string, post ...*models.Post) bool {
	switch action {
	case "create":
		for _, field := range post {
			if field.Title == "" || field.Content == "" {
				return false
			}
		}
	case "update":
		for _, field := range post {
			if field.ID == 0 || field.Title == "" || field.Content == "" {
				return false
			}
		}
	case "delete":
		for _, field := range post {
			if field.ID == 0 {
				return false
			}
		}
	}
	return true
}

func ValidateJSONComment(action string, comment ...*models.Comment) bool {
	switch action {
	case "add":
		for _, field := range comment {
			if field.Content == "" || field.PostID == 0 {
				return false
			}
		}
	case "update":
		for _, field := range comment {
			if field.ID == 0 || field.Content == "" {
				return false
			}
		}
	case "delete":
		for _, field := range comment {
			if field.ID == 0 {
				return false
			}
		}
	}
	return true
}

func ValidateJSONCategory(action string, category ...*models.Category) bool {
	switch action {
	case "create":
		for _, field := range category {
			if field.Name == "" {
				return false
			}
		}
	case "update":
		for _, field := range category {
			if field.ID == 0 || field.Name == "" {
				return false
			}
		}
	case "delete":
		for _, field := range category {
			if field.Name == "" {
				return false
			}
		}
	}
	return true
}
