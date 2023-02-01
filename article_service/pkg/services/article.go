package services

import (
	"article_service/pkg/db"
	"article_service/pkg/models"
	"article_service/pkg/pb"
	"context"
	"net/http"
)

type Server struct {
	ServerDB db.Database
}

func (s *Server) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	var article models.Article

	article.Name = req.Name
	article.FileName = req.NameFile
	article.Note = req.Note
	article.DateCreate = req.DateCreate
	article.UserId = req.UserId

	if result := s.ServerDB.DB.Create(&article); result.Error != nil {
		return &pb.CreateArticleResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateArticleResponse{
		Status: http.StatusCreated,
		Id:     article.Id,
	}, nil
}

func (s *Server) FindAllArticle(ctx context.Context, req *pb.FindAllArticleRequest) (*pb.FindAllArticleResponse, error) {
	articles := []*pb.Article{}

	if result := s.ServerDB.DB.Where("user_id = ?", req.UserId).Find(&articles); result.Error != nil {
		return &pb.FindAllArticleResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.FindAllArticleResponse{
		Status: http.StatusOK,
		Data:   articles,
	}, nil
}

func (s *Server) FindOneArticle(ctx context.Context, req *pb.FindOneArticleRequest) (*pb.FindOneArticleResponse, error) {
	var article models.Article

	if result := s.ServerDB.DB.Where("id = ? AND user_id = ?", req.Id, req.UserId).Find(&article); result.Error != nil {
		return &pb.FindOneArticleResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.Article{
		Id:         article.Id,
		Name:       article.Name,
		FileName:   article.FileName,
		Note:       article.Note,
		DateCreate: article.DateCreate,
		UserId:     article.UserId,
	}

	return &pb.FindOneArticleResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) FindArticleByName(ctx context.Context, req *pb.FindArticleByNameRequest) (*pb.FindArticleByNameResponse, error) {
	articles := []*pb.Article{}

	if result := s.ServerDB.DB.Where("name = ? AND user_id = ?", req.Name, req.UserId).Find(&articles); result.Error != nil {
		return &pb.FindArticleByNameResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.FindArticleByNameResponse{
		Status: http.StatusOK,
		Data:   articles,
	}, nil
}

func (s *Server) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	var article models.Article

	if result := s.ServerDB.DB.Where("id = ? AND user_id = ?", req.Id, req.UserId).First(&article); result.Error != nil {
		return &pb.UpdateArticleResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	article.Name = req.Name
	article.Note = req.Note

	result := s.ServerDB.DB.Save(&article)

	if result.RowsAffected == 0 {
		return &pb.UpdateArticleResponse{
			Status: http.StatusConflict,
			Error:  "Failed update row database",
		}, nil
	}

	return &pb.UpdateArticleResponse{
		Status: http.StatusOK,
		Id:     article.Id,
	}, nil
}

func (s *Server) UpdateArticleFileName(ctx context.Context, req *pb.UpdateArticleFileNameRequest) (*pb.UpdateArticleFileNameResponse, error) {
	var article models.Article

	if result := s.ServerDB.DB.Where("id = ?", req.Id).First(&article); result.Error != nil {
		return &pb.UpdateArticleFileNameResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	article.FileName = req.NameFile

	result := s.ServerDB.DB.Save(&article)

	if result.RowsAffected == 0 {
		return &pb.UpdateArticleFileNameResponse{
			Status: http.StatusConflict,
			Error:  "Failed update row database",
		}, nil
	}

	return &pb.UpdateArticleFileNameResponse{
		Status: http.StatusOK,
		Id:     article.Id,
	}, nil
}

func (s *Server) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	var article models.Article

	article.Id = req.Id

	if result := s.ServerDB.DB.Delete(&article); result.Error != nil {
		return &pb.DeleteArticleResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.DeleteArticleResponse{
		Status: http.StatusOK,
	}, nil
}

/*func (s *Server) CreateArticle(stream pb.ArticleService_CreateArticleServer) error {
var article models.Article
var response pb.CreateArticleResponse

req, err := stream.Recv()

if err != nil {
	response = pb.CreateArticleResponse{
		Status: http.StatusConflict,
		Error:  "no data",
	}
}

article.Name = req.Info.Name
article.FileName = req.Info.NameFile
article.Note = req.Info.Note
article.DateCreate = req.Info.DateCreate
article.UserId = req.Info.UserId

if result := s.ServerDB.DB.Create(&article); result.Error != nil {
	response = pb.CreateArticleResponse{
		Status: http.StatusConflict,
		Error:  result.Error.Error(),
	}
}

reqFile := req.GetFile()

info := &pb.FileInfo{
	FileName: article.FileName,
	FileType: ".pdf",
}

_, err = s.FileServiceClient.UploadFile(reqFile, info, stream.Context())

if err != nil {
	response = pb.CreateArticleResponse{
		Status: http.StatusConflict,
		Error:  err.Error(),
	}
}

response = pb.CreateArticleResponse{
	Status: http.StatusCreated,
	Id:     article.Id,
}

stream.SendAndClose(&response)

/*return &pb.CreateArticleResponse{
	Status: http.StatusCreated,
	Id:     article.Id,
}, nil*/
//return nil*/
