package services

import (
	"bytes"
	"file_service/pkg/client"
	"file_service/pkg/pb"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"github.com/ledongthuc/pdf"
)

type Server struct {
	ArticleClient client.ArticleClient
}

func (s *Server) UploadFile(stream pb.FileService_UploadFileServer) error {
	imageData := bytes.Buffer{}
	imageSize := 0

	var path string
	var nameFile string

	for {
		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}

		fmt.Println(req.Info)

		chunk := req.GetFileData()
		size := len(chunk)

		imageSize += size

		_, err = imageData.Write(chunk)

		if err != nil {
			log.Print("no more write")
		}

		idString := strconv.Itoa(int(req.Info.Id))

		pathFile := fmt.Sprintf("file/" + idString + "/")

		err = os.MkdirAll(pathFile, 0777)
		if err != nil {
			log.Print("no more panic!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			res := &pb.UploadFileResponse{
				Status: http.StatusConflict,
				Error:  "Failed create file",
			}

			stream.SendAndClose(res)
		}

		dst, err := os.Create(filepath.Join(pathFile, req.Info.FileName))

		if err != nil {
			log.Print("no more dst")
		}

		defer dst.Close()

		io.Copy(dst, &imageData)

		path = pathFile + req.Info.FileName
		nameFile = req.Info.FileName
	}

	fmt.Println(path)
	fmt.Println(nameFile)
	//convertTXTtoPDF(nameFile, path)

	res := &pb.UploadFileResponse{
		Status: http.StatusOK,
	}

	stream.SendAndClose(res)

	return nil
}

func convertTXTtoPDF(nameFile, path string) {
	fmt.Println(path)
	text, err := ioutil.ReadFile(path)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.SetFont("Helvetica", "B", 16)
	pdf.MoveTo(0, 10)
	pdf.Cell(1, 1, "арвоарова")
	pdf.MoveTo(0, 20)
	pdf.SetFont("Arial", "", 14)
	width, _ := pdf.GetPageSize()
	pdf.MultiCell(width, 10, string(text), "", "", false)
	fmt.Println(nameFile)
	err = pdf.OutputFileAndClose(nameFile + ".pdf")
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) UploadFileArticle(stream pb.FileService_UploadFileArticleServer) error {
	imageData := bytes.Buffer{}
	imageSize := 0

	var id int64
	var nameFile string

	for {
		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}

		fmt.Println(req.Info)

		chunk := req.GetFileData()
		size := len(chunk)

		imageSize += size

		_, err = imageData.Write(chunk)

		if err != nil {
			log.Print("no more write")
		}

		idString := strconv.Itoa(int(req.Info.Id))

		pathFile := fmt.Sprintf("file/" + idString + "/")

		err = os.MkdirAll(pathFile, 0777)
		if err != nil {
			log.Print("no more panic!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			res := &pb.UploadFileResponse{
				Status: http.StatusConflict,
				Error:  "Failed create file",
			}

			stream.SendAndClose(res)
		}

		dst, err := os.Create(filepath.Join(pathFile, req.Info.FileName))

		if err != nil {
			log.Print("no more dst")
		}

		defer dst.Close()

		io.Copy(dst, &imageData)

		id = req.Info.Id
		nameFile = req.Info.FileName
	}

	res := &pb.UploadFileResponse{
		Status: http.StatusOK,
	}

	stream.SendAndClose(res)

	s.ArticleClient.UpdateArticleFileName(id, nameFile)

	return nil
}

func (s *Server) DownloadFile(req *pb.DownloadFileRequest, stream pb.FileService_DownloadFileServer) error {
	idString := strconv.Itoa(int(req.Id))

	pathFile := fmt.Sprintf("file/" + idString + "/")

	bufferSize := 64 * 1024 //64KiB, tweak this as desired
	file, err := os.Open(pathFile + req.GetFileName())
	if err != nil {
		fmt.Println("Error open file file server", err)
		return err
	}
	defer file.Close()
	buff := make([]byte, bufferSize)
	for {
		bytesRead, err := file.Read(buff)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Read file file server", err)
			}
			break
		}
		resp := &pb.DownloadFileResponse{
			FileData: buff[:bytesRead],
		}
		err = stream.Send(resp)
		if err != nil {
			log.Println("error while sending chunk:", err)
			return err
		}
	}
	return nil
}

func convertPDFtoTXT(fileName string) (string, error) {
	f, r, err := pdf.Open(fileName)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)

	if err := os.WriteFile(fileName+".txt", buf.Bytes(), 0777); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (s *Server) DownloadFileTXT(req *pb.DownloadFileTXTRequest, stream pb.FileService_DownloadFileTXTServer) error {
	idString := strconv.Itoa(int(req.Id))

	pathFile := fmt.Sprintf("file/" + idString + "/")

	_, err := convertPDFtoTXT(pathFile + req.GetFileName())

	if err != nil {
		fmt.Println("err")
		return err
	}

	bufferSize := 64 * 1024 //64KiB, tweak this as desired
	file, err := os.Open(pathFile + req.GetFileName() + ".txt")

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	buff := make([]byte, bufferSize)
	for {
		bytesRead, err := file.Read(buff)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		resp := &pb.DownloadFileTXTResponse{
			FileData: buff[:bytesRead],
		}
		err = stream.Send(resp)
		if err != nil {
			log.Println("error while sending chunk:", err)
			return err
		}
	}
	return nil
}
