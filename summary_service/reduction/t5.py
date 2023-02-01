import torch
import json
from transformers import T5Tokenizer, T5Config, T5ForConditionalGeneration

def SummaryT5(name, fileName, minCount, maxCount):
    with open(fileName + '.txt', 'r', encoding="utf-8") as file:
        data = file.read().rstrip()

    device = torch.device('cpu')

    summarizer = T5ForConditionalGeneration.from_pretrained('cointegrated/rut5-base-multitask')
    tokenizer = T5Tokenizer.from_pretrained('cointegrated/rut5-base-multitask')

    updated_material = "summarize:" + data

    tokenized_material = tokenizer.encode(updated_material, return_tensors="pt").to(device)

    print("1")
    print(tokenized_material)

    tokenized_summary = summarizer.generate(tokenized_material,
                                    num_beams=5,
                                    no_repeat_ngram_size=2,
                                    min_length=minCount,
                                    max_length=maxCount,
                                    early_stopping=True)
    print("1.2")

    print(tokenized_summary)
    print("2")

    summary = tokenizer.decode(tokenized_summary[0], skip_special_tokens=True)

    print(summary)

    with open(name + "T5.txt", "w", encoding="cp1251") as myFile:
        myFile.write(summary)

