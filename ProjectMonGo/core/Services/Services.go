package Services

import (
	"ProjectMonGo/Middwares/HashPassword"
	"ProjectMonGo/api/Request"
	"ProjectMonGo/core/entities"
	"ProjectMonGo/core/posts"
	"context"
	"log"
	"time"
)

type PostService struct {
	PostRepository posts.PostService
}

func NewUserService(userRepositoryPort posts.PostService) *PostService {
	return &PostService{
		PostRepository: userRepositoryPort,
	}
}

func (u *PostService) Create_Post(ctx context.Context, req *Request.CreatePostRequest) (*entities.DBPost, error) {
	post, err := u.PostRepository.CreatePost(&Request.CreatePostRequest{
		Title:     req.Title,
		Email:     req.Email,
		Token:     req.Token,
		Content:   req.Content,
		Image:     req.Image,
		User:      req.User,
		CreateAt:  time.Time{},
		UpdatedAt: time.Time{},
	})
	if err != nil {
		log.Fatalf("err : ", err)
		return nil, err
	}
	return post, nil
}
func (u *PostService)Create_PostHashToken(ctx context.Context,req *Request.CreatePostRequest)(*entities.DBPost,error){
	hashedToken ,_:= HashPassword.HashPassword(req.Token)
	req.Token=hashedToken
	post, err := u.PostRepository.CreatePost(&Request.CreatePostRequest{
		Title:     req.Title,
		Email:     req.Email,
		Token:     req.Token,
		Content:   req.Content,
		Image:     req.Image,
		User:      req.User,
		CreateAt:  time.Time{},
		UpdatedAt: time.Time{},
	})
	if err != nil {
		log.Fatalf("err : ", err)
		return nil, err
	}
	return post, nil
}

func (u *PostService)CreateToken(request *Request.TokenRequest) (string,error) {
	post,err := u.PostRepository.GenratorToken(request)
	if err != nil{
		return "", err
	}
	return post,nil
}

func (u *PostService) Update_Post(id string, req *Request.UpdatePost) (*Request.UpdatePost, error) {
	post, err := u.PostRepository.UpdatePost(id, req)
	if err != nil {
		log.Fatalf("err ", err)
		return nil, err
	}
	return post, nil
}
func (u *PostService) Delete_Post(id string) error {
	err := u.PostRepository.DeletePosts(id)
	if err != nil {
		log.Fatalf("err ", err)
		return err
	}
	return nil
}

func (u *PostService) Find_Id(id string) (*entities.DBPost, error) {
	//var post *entities.DBPost
	post, err := u.PostRepository.FindPostById(id)
	if err != nil {
		log.Fatalf("err ", err)
		return nil, err
	}
	return post, nil
}

func (u *PostService) FindPost(page int, limit int) ([]*entities.DBPost, error) {
	post, err := u.PostRepository.FindPosts(page, limit)
	if err != nil {
		log.Fatalf("err ", err)
		return nil, err
	}
	return post, nil
}

func (u *PostService) SearchForm(req *Request.FromReq) ([]*entities.DBPost,error) {
	post,err := u.PostRepository.FindPostByIdForm(req)
	if err != nil{
		log.Fatalf("err",err)
		return nil,err
	}
	return post,nil
}

























