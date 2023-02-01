import sumy
import nltk
nltk.download('punkt')
from sumy.summarizers.lsa import LsaSummarizer
from sumy.nlp.tokenizers import Tokenizer
from sumy.parsers.plaintext import PlaintextParser

def SummaryLSA(name, fileName, countSentense):
    with open(fileName + '.txt', 'r', encoding="utf8") as file:
        data = file.read().rstrip()

    LANGUAGE = "ru"
    parser = PlaintextParser.from_string(data,Tokenizer(LANGUAGE))

    summarizer = LsaSummarizer()

    testsummary = summarizer(parser.document, sentences_count=countSentense)

    summary = ""
    for sentence in testsummary:
        summary+=str(sentence)
    print(summary)

    with open(name + "LSA.txt", "w", encoding="cp1251") as myFile:
        myFile.write(summary)
