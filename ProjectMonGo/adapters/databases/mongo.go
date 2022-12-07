package databases

import (
	"ProjectMonGo/Middwares/auth"
	"ProjectMonGo/adapters/databases/mapper"
	"ProjectMonGo/api/Request"
	"ProjectMonGo/core/entities"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type PostServiceImpl struct {
	postCollection *mongo.Collection
	ctx            context.Context
}

func NewMonGoDb(collection *mongo.Collection) *PostServiceImpl {
	return &PostServiceImpl{
		postCollection: collection,
	}
}

func (p *PostServiceImpl) CreatePost(post *Request.CreatePostRequest) (*entities.DBPost, error) {
	post.CreateAt = time.Now()
	post.UpdatedAt = post.CreateAt

	userModel := &Request.CreatePostRequest{
		Title:     post.Title,
		Email:     post.Email,
		Token:     post.Token,
		Content:   post.Content,
		Image:     post.Image,
		User:      post.User,
		CreateAt:  time.Time{},
		UpdatedAt: time.Time{},
	}
	_, err := p.postCollection.InsertOne(p.ctx, userModel)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (p *PostServiceImpl)CreatePostHashToken(post *Request.CreatePostRequest)(*entities.DBPost,error){
	post.CreateAt = time.Now()
	post.UpdatedAt = post.CreateAt

	userModel := &Request.CreatePostRequest{
		Title:     post.Title,
		Email:     post.Email,
		Token:     post.Token,
		Content:   post.Content,
		Image:     post.Image,
		User:      post.User,
		CreateAt:  time.Time{},
		UpdatedAt: time.Time{},
	}
	_, err := p.postCollection.InsertOne(p.ctx, userModel)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *PostServiceImpl)GenratorToken(request *Request.TokenRequest)(string,error){
	var admin *entities.DBPost
	err := p.postCollection.FindOne(p.ctx,bson.M{
		"email":request.Email,
	}).Decode(&admin)
	if err != nil{
		log.Fatal(err)
		return "",err
	}
	err2 := bcrypt.CompareHashAndPassword([]byte(admin.Token),[]byte(request.Token))
	if err2 != nil{
		log.Fatal(err2)
		return "",nil
	}
	Token,err3 := auth.GenerateJWT(admin.Email,admin.Token)
	if err3 != nil{
		log.Fatal(err3)
		return "",nil
	}
	return Token,nil
}


func (p *PostServiceImpl) UpdatePost(id string, data *Request.UpdatePost) (*Request.UpdatePost, error) {
	doc, err := mapper.ToDoc(data)
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "title", Value: id}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.postCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatePost *Request.UpdatePost

	if err := res.Decode(&updatePost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatePost, nil
}
func (p *PostServiceImpl) FindPostById(id string) (*entities.DBPost, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}
	var post *entities.DBPost
	if err := p.postCollection.FindOne(p.ctx, query).Decode(&post); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}
		return nil, err
	}
	return post, nil
}

func (p *PostServiceImpl)FindPostByIdForm(req *Request.FromReq)([]*entities.DBPost,error){
	filter := bson.M{}
	var post []*entities.DBPost
	if req.Email != ""{
		filter["email"]=req.Email
	}
	if req.User != ""{
		filter["user"]=req.User
	}
	if req.Title != ""{
		filter["title"]=req.Title
	}
	if req.Content != ""{
		filter["content"]=req.Content
	}
	cursor,err := p.postCollection.Find(p.ctx,filter,nil)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	if err = cursor.All(context.TODO(), &post); err != nil {
		log.Fatal(err)
		return nil,err
	}
	return post,nil
}

func (p *PostServiceImpl) FindPosts(page int, limit int) ([]*entities.DBPost, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	skip := (page - 1) * limit
	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	query := bson.M{}
	cursor, err := p.postCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	var posts []*entities.DBPost

	for cursor.Next(p.ctx) {
		post := &entities.DBPost{}
		err := cursor.Decode(post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return []*entities.DBPost{}, nil
	}
	return posts, nil
}
func (p *PostServiceImpl) DeletePosts(title string) error {

	filter := bson.M{"title": title}
	_, err := p.postCollection.DeleteOne(p.ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
