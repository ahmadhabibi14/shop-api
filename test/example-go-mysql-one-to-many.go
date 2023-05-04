package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dropTable(db *sql.DB, table string) {
	_, err := db.Exec("DROP TABLE IF EXISTS " + table)
	if err != nil {
		panic(err)
	}
}

type Comment struct {
	Id      int64
	Comment string
}
type PostView struct {
	Id       int64
	Post     string
	Comments []Comment
}
type PostViewRepo struct {
	*sql.DB
}

func SetPostViewRepo(db *sql.DB) (repo *PostViewRepo) {
	repo = &PostViewRepo{db}
	return
}
func (r *PostViewRepo) GetAll() (posts []PostView) {
	rows, err := r.Query(`
SELECT A.id, A.post, B.id, B.comment FROM post A LEFT JOIN comment B ON A.id = B.post_id
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	lastPostId := int64(-1)
	lastPostIndex := -1
	var (
		post    PostView
		comment Comment
	)
	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Post, &comment.Id, &comment.Comment)
		if err != nil {
			panic(err)
		}
		if post.Id != lastPostId {
			posts = append(posts, post)
			lastPostId = post.Id
			lastPostIndex++
		}
		posts[lastPostIndex].Comments = append(posts[lastPostIndex].Comments, comment)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return
}

type PostRepo struct {
	*sql.DB
}

func SetPostRepo(db *sql.DB) (repo *PostRepo) {
	repo = &PostRepo{db}
	return
}
func (r *PostRepo) Insert(post string) (postId int64) {
	result, err := r.Exec("INSERT INTO post(post) VALUES (?)", post)
	if err != nil {
		panic(err)
	}
	postId, _ = result.LastInsertId()
	return
}

type CommentRepo struct {
	*sql.DB
}

func SetCommentRepo(db *sql.DB) (repo *CommentRepo) {
	repo = &CommentRepo{db}
	return
}
func (r *CommentRepo) Insert(postId int64, comment string) {
	_, err := r.Exec("INSERT INTO comment(comment, post_id) VALUES (?, ?)", comment, postId)
	if err != nil {
		panic(err)
	}
}

func createTablePost(db *sql.DB) {
	_, err := db.Exec(`
CREATE TABLE post (
	id INT AUTO_INCREMENT,
	post VARCHAR(999),
	PRIMARY KEY (id)
);
	`)
	if err != nil {
		panic(err)
	}
}
func createTableComment(db *sql.DB) {
	_, err := db.Exec(`
CREATE TABLE comment (
	id INT AUTO_INCREMENT,
	comment VARCHAR(999),
	post_id INT,
	PRIMARY KEY (id),
	INDEX FK_comment_post (post_id),
	CONSTRAINT FK_comment_post FOREIGN KEY (post_id) REFERENCES post (id)
);
	`)
	if err != nil {
		panic(err)
	}
}

func main() {
	username, password, database := "root", "p4ssw0rd", "test"
	dsn := fmt.Sprintf("%s:%s@/%s", username, password, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dropTable(db, "comment")
	dropTable(db, "post")
	createTablePost(db)
	createTableComment(db)

	postRepo := SetPostRepo(db)
	commentRepo := SetCommentRepo(db)
	postViewRepo := SetPostViewRepo(db)

	postId1 := postRepo.Insert("post1")
	commentRepo.Insert(postId1, "comment1 of post1")
	commentRepo.Insert(postId1, "comment2 of post1")
	postId2 := postRepo.Insert("post2")
	commentRepo.Insert(postId2, "comment1 of post2")
	commentRepo.Insert(postId2, "comment2 of post2")

	allPostsView := postViewRepo.GetAll()
	b, _ := json.MarshalIndent(allPostsView, "", " ")
	fmt.Printf("%s\n", b)
}
