package internal

type User struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Session    bool
	ErrorL     bool
	ErrorE     bool
	ErrorEm    bool
	ErrorEmpty bool
}

type LikeNdis struct {
	Like int
	Dis  int
	User User
	Post Post
}

type Forum struct {
	User     User
	Category []Category
	Post     Post
	Comment  Comment
	LikeNdis LikeNdis
}
type Category struct {
	ID   int
	Name string
	Rows []Category
}
type Post struct {
	ID           int
	Name         string
	Body         string
	CategoryNull bool
	TitBodNull   bool
	User         User
	Cat          []Category
	Likes        int
	Dislikes     int
	Image        string
	Comm         []Comment
	Rows         []Post
}
type Comment struct {
	ID       int
	Body     string
	User     User
	Post     Post
	Likes    int
	Dislikes int
	Rows     []Comment
}
