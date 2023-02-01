import grpc
from concurrent import futures
import logging
import PyPDF2

from protos import summary_pb2_grpc as pb2_grpc
from protos import summary_pb2 as pb2
from protos import file_pb2_grpc as file_client_pb2_grpc
from protos import file_pb2 as file_client_pb2
from protos import article_pb2_grpc as article_client_pb2_grpc
from protos import article_pb2 as article_client_pb2

from reduction import text_rank as tr
from reduction import lsa
from reduction import t5

import os

ip = "192.168.99.100"

class SummaryTextService(pb2_grpc.SummaryTextService):
    def __init__(self, *args, **kwargs):
        pass

    def SummaryTextRank(self, request, context):
        name = request.info.name
        nameFile = name + "TR.txt"
        type = request.info.type
        note = request.info.note
        article_id = request.info.articleId

        countWord = request.info.countWord

        DownloadFileTXT(request.info.nameFile, article_id)

        tr.SummaryTextRank(name, request.info.nameFile, countWord)

        UploadFile(nameFile, article_id)

        info = article_client_pb2.CreateSummaryRequest(name=name, nameFile=nameFile, type=type, note=note, articleId=article_id)

        id = CreateSummaryDB(info)

        result = {'id': id}

        RemoveFile(request.info.nameFile, nameFile)

        return pb2.SummaryTextResponse(**result)

    def SummaryLSA(self, request, context):
        name = request.info.name
        nameFile = name + "LSA.txt"
        type = request.info.type
        note = request.info.note
        article_id = request.info.articleId

        countSentense = request.info.countSentense

        DownloadFileTXT(request.info.nameFile, article_id)

        lsa.SummaryLSA(name, request.info.nameFile, countSentense)

        UploadFile(nameFile, article_id)

        info = article_client_pb2.CreateSummaryRequest(name=name, nameFile=nameFile, type=type, note=note, articleId=article_id)

        id = CreateSummaryDB(info)

        result = {'id': id}

        return pb2.SummaryTextResponse(**result)

    def SummaryT5(self, request, context):
        name = request.info.name
        nameFile = name + "T5.txt"
        type = request.info.type
        note = request.info.note
        article_id = request.info.articleId

        minCount = request.info.minCountSentense
        maxCount = request.info.maxCountSentense

        DownloadFile(request.info.nameFile, article_id)

        ConvertPDF(request.info.nameFile)

        t5.SummaryT5(name, request.info.nameFile, minCount, maxCount)

        UploadFile(nameFile, article_id)

        info = article_client_pb2.CreateSummaryRequest(name=name, nameFile=nameFile, type=type, note=note, articleId=article_id)

        id = CreateSummaryDB(info)

        result = {'id': id}

        return pb2.SummaryTextResponse(**result)


def CreateSummaryDB(info):
    with grpc.insecure_channel(ip + ':9002') as channel:
        stub = article_client_pb2_grpc.ArticleServiceStub(channel)
        res = stub.CreateSummary(info)
        return res.id
    channel.close

    
def DownloadFileTXT(filename, id):
    with grpc.insecure_channel(ip + ':9003') as channel:
        stub = file_client_pb2_grpc.FileServiceStub(channel)
        filepath = filename
        req = file_client_pb2.DownloadFileRequest(id=id, fileName=filename)
        for entry_response in stub.DownloadFileTXT(req):
            with open(filepath + ".txt", mode="ab") as f:
                f.write(entry_response.fileData)

    channel.close

def DownloadFile(filename, id):
    with grpc.insecure_channel(ip + ':9003') as channel:
        stub = file_client_pb2_grpc.FileServiceStub(channel)
        filepath = filename
        req = file_client_pb2.DownloadFileRequest(id=id, fileName=filename)
        for entry_response in stub.DownloadFile(req):
            with open(filepath, mode="ab") as f:
                f.write(entry_response.fileData)

    channel.close

def read_iterfile(id, fileName):
    file_info = file_client_pb2.FileInfo(id=id, fileName=fileName, fileType=".pdf")
    yield file_client_pb2.UploadFileRequest(info=file_info)
    with open(fileName, mode="rb") as f:
        while True:
            chunk = f.read(860160)
            if chunk:
                entry_request = file_client_pb2.UploadFileRequest(info=file_info, fileData=chunk)
                yield entry_request
            else:
                return

def UploadFile(filename, id):
    with grpc.insecure_channel(ip + ':9003') as channel:
        stub = file_client_pb2_grpc.FileServiceStub(channel)

        res = stub.UploadFile(read_iterfile(id=id, fileName=filename))
        print(res)

    channel.close

def ConvertPDF(fileName):
    pdfFileObject = open(r"" + fileName, 'rb')
    pdfReader = PyPDF2.PdfFileReader(pdfFileObject)
    print(" No. Of Pages :", pdfReader.numPages)
    pageCount = pdfReader.numPages
    i = 0
    while i < pageCount:
        pageObject = pdfReader.getPage(i)
        print(pageObject.extractText())
        with open(fileName + ".txt", "a", encoding="utf-8") as myFile:
            myFile.write(pageObject.extractText())

        myFile.close()
        i += 1
    
    pdfFileObject.close()

def RemoveFile(fileName, fileNameSumm):
    os.remove(fileName + ".txt")
    os.remove(fileNameSumm)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_SummaryTextServiceServicer_to_server(SummaryTextService(), server)
    server.add_insecure_port('[::]:9004')
    server.start()
    print("Start summary server on port 9004")
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    serve()