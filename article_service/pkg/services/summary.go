package services

import (
	"article_service/pkg/models"
	"article_service/pkg/pb"
	"context"
	"net/http"
)

func (s *Server) CreateSummary(ctx context.Context, req *pb.CreateSummaryRequest) (*pb.CreateSummaryResponse, error) {
	var summary models.Summary

	if result := s.ServerDB.DB.Where("name = ? AND type = ? AND article_id = ?", req.Name, req.Type, req.ArticleId).First(&summary); result.Error == nil {
		return &pb.CreateSummaryResponse{
			Status: http.StatusConflict,
			Error:  "Summary already exists",
		}, nil
	}

	summary.Name = req.Name
	summary.FileName = req.NameFile
	summary.Type = req.Type
	summary.Note = req.Note
	summary.ArticleId = req.ArticleId

	s.ServerDB.DB.Create(&summary)

	return &pb.CreateSummaryResponse{
		Status: http.StatusCreated,
		Id:     summary.Id,
	}, nil
}

func (s *Server) FindAllSummary(ctx context.Context, req *pb.FindAllSummaryRequest) (*pb.FindAllSummaryResponse, error) {
	summary := []*pb.Summary{}

	if result := s.ServerDB.DB.Where("article_id = ?", req.ArticleId).Find(&summary); result.Error != nil {
		return &pb.FindAllSummaryResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.FindAllSummaryResponse{
		Status: http.StatusOK,
		Data:   summary,
	}, nil
}

func (s *Server) FindOneSummary(ctx context.Context, req *pb.FindOneSummaryRequest) (*pb.FindOneSummaryResponse, error) {
	summary := []*pb.Summary{}

	if result := s.ServerDB.DB.Where("name = ? AND article_id = ?", req.Name, req.ArticleId).First(&summary); result.Error != nil {
		return &pb.FindOneSummaryResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.FindOneSummaryResponse{
		Status: http.StatusOK,
		Data:   summary,
	}, nil
}

func (s *Server) UpdateSummary(ctx context.Context, req *pb.UpdateSummaryRequest) (*pb.UpdateSummaryResponse, error) {
	var summary models.Summary

	if result := s.ServerDB.DB.Where("id = ?", req.Id).First(&summary); result.Error != nil {
		return &pb.UpdateSummaryResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	summary.Name = req.Name
	summary.FileName = req.NameFile
	summary.Type = req.Type
	summary.Note = req.Note

	s.ServerDB.DB.Save(&summary)

	return &pb.UpdateSummaryResponse{
		Status: http.StatusOK,
		Id:     summary.Id,
	}, nil
}

func (s *Server) DeleteSummary(ctx context.Context, req *pb.DeleteSummaryRequest) (*pb.DeleteSummaryResponse, error) {
	var summary models.Summary

	if result := s.ServerDB.DB.Where("id = ?", req.Id).First(&summary); result.Error != nil {
		return &pb.DeleteSummaryResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	summary.Id = req.Id

	if result := s.ServerDB.DB.Delete(&summary); result.Error != nil {
		return &pb.DeleteSummaryResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.DeleteSummaryResponse{
		Status:    http.StatusOK,
		ArticleId: summary.ArticleId,
	}, nil
}
