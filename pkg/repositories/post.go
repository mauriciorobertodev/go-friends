package repositories

import (
	"database/sql"
	"errors"
	"go-friends/pkg/models"
)

type posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (r posts) Store(post models.Post) (uint64, error) {
	stm, err := r.db.Prepare("INSERT INTO posts (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, nil
	}

	defer stm.Close()

	result, err := stm.Exec(post.Title, post.Content, post.AuthorId)

	if err != nil {
		return 0, nil
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, nil
	}

	return uint64(lastId), nil
}

func (r posts) FindPostById(id uint64) (models.Post, error) {
	rows, err := r.db.Query(
		"SELECT post.*, user.Nick FROM posts post INNER JOIN users user ON user.id = post.author_id WHERE post.id = ?",
		id,
	)

	if err != nil {
		return models.Post{}, err
	}

	defer rows.Close()

	var post models.Post

	if rows.Next() {
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorId,
			&post.AuthorNick,
		)

		if err != nil {
			return models.Post{}, err
		}

		return post, nil
	}

	return models.Post{}, errors.New("post does not exists")
}

func (r posts) Search(userId uint64) ([]models.Post, error) {
	rows, err := r.db.Query(
		`
			SELECT DISTINCT post.*, user.nick FROM posts post
			INNER JOIN users user ON user.id = post.author_id
			INNER JOIN followers follow ON post.author_id = follow.user_id 
			WHERE user.id = ? OR follow.follower_id = ?
			ORDER BY 5 desc
		`,
		userId,
		userId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorId,
			&post.AuthorNick,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r posts) UpdatePost(id uint64, post models.Post) error {
	stm, err := r.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(post.Title, post.Content, id)

	if err != nil {
		return err
	}

	return nil
}

func (r posts) DeletePost(id uint64) error {
	stm, err := r.db.Prepare("DELETE FROM posts WHERE id = ?")

	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (r posts) GetPostsOfUser(userId uint64) ([]models.Post, error) {
	rows, err := r.db.Query(
		`
			SELECT DISTINCT post.*, user.nick FROM posts post
			INNER JOIN users user ON user.id = post.author_id
			WHERE post.author_id = ?
			ORDER BY 5 desc
		`,
		userId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorId,
			&post.AuthorNick,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r posts) LikePost(id uint64) error {
	stm, err := r.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
