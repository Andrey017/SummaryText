FROM python:3.8

RUN python -m pip install grpcio
RUN python -m pip install grpcio-tools

RUN python -m pip install PyPDF2

RUN python -m pip install transformers
RUN python -m pip install sentencepiece
RUN python -m pip --no-cache-dir install torch

RUN python -m pip install sumy

RUN python -m pip install gensim==3.6.0

RUN python -m pip install rouge

RUN python -c "import nltk; nltk.download('punkt')"

COPY . /summary/

WORKDIR /summary/