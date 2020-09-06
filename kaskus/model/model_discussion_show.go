package model

type DiscussionShow struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Date         string `json:"date"`
	Message      string `json:"message"`
	ProfileImage string `json:"profile_image"`
	Total        int    `json:"total"`
}

type DiscussionPost struct {
	DiscussionID uint     `json:"discussion_id"`
	UserID       uint     `json:"user_id"`
	CatagoryID   uint     `json:"catagory_id"`
	Title        string   `json:"title"`
	Date         string   `json:"date"`
	Message      string   `json:"message"`
	Images       []string `json:"images"`
	Files        []string `json:"files"`
}

type DiscussionDetailShow struct {
	ID             uint                        `json:"id"`
	Name           string                      `json:"name"`
	Date           string                      `json:"date"`
	Message        string                      `json:"message"`
	ProfileImage   string                      `json:"profile_image"`
	Images         []string                    `json:"images"`
	Files          []string                    `json:"files"`
	FirsDiscussion []DiscussionFirstDetailShow `json:"answer_first"`
}

type DiscussionFirstDetailShow struct {
	ID               uint                         `json:"id"`
	Name             string                       `json:"name"`
	Date             string                       `json:"date"`
	Message          string                       `json:"message"`
	ProfileImage     string                       `json:"profile_image"`
	Image            string                       `json:"image"`
	File             string                       `json:"file"`
	SecondDiscussion []DiscussionSecondDetailShow `json:"answer_second"`
}

type DiscussionSecondDetailShow struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Date         string `json:"date"`
	Message      string `json:"message"`
	ProfileImage string `json:"profile_image"`
	Image        string `json:"image"`
	File         string `json:"file"`
}
