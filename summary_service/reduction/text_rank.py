import gensim
from gensim.summarization import summarize

def SummaryTextRank(name, fileName, countWord):
    with open(fileName + '.txt', 'r', encoding="utf8") as file:
        data = file.read().rstrip()

    source_material = "Consciousness, at its simplest, is sentience or awareness of internal and external existence. Despite millennia of analyses, definitions, explanations and debates by philosophers and scientists, consciousness remains puzzling and controversial, being 'at once, the most familiar and also the most mysterious aspect of our lives'. Perhaps the only widely agreed notion about the topic is the intuition that consciousness exists. Opinions differ about what exactly needs to be studied and explained as consciousness. Sometimes, it is synonymous with the mind, and at other times, an aspect of mind. In the past, it was one's 'inner life', the world of introspection, of private thought, imagination and volition. Today, it often includes any kind of cognition, experience, feeling or perception. It may be awareness, awareness of awareness, or self-awareness either continuously changing or not. There might be different levels or orders of consciousness, or different kinds of consciousness, or just one kind with different features. Other questions include whether only humans are conscious, all animals, or even the whole universe. The disparate range of research, notions and speculations raises doubts about whether the right questions are being asked. Examples of the range of descriptions, definitions or explanations are: simple wakefulness, one's sense of selfhood or soul explored by 'looking within'; being a metaphorical 'stream' of contents, or being a mental state, mental event or mental process of the brain; having phanera or qualia and subjectivity; being the 'something that it is like' to 'have' or 'be' it; being the 'inner theatre' or the executive control system of the mind."

    source_material2 = "Оптимальное сочетание инструментов и технологий для построения системы безопасности в компаниях с небольшим отделом ИБ. Решение обеспечивает автоматическое обнаружение и оперативное реагирование на скрытые угрозы и отражает принципы ступенчатого подхода «Лаборатории Касперского» к построению стратегии кибербезопасности. Если вы предпочитаете облачную консоль1, создайте учетную запись для доступа к продуктам «Лаборатории Касперского». Для локальной установки консоли управления загрузите и установите Kaspersky Security Center, используя ссылки ниже. В ходе развертывания защиты для рабочих мест вы сможете скачать все необходимые дистрибутивы непосредственно из консоли управления. Процесс развертывания описан в кратком руководстве по началу работы. При желании вы можете установить ПО непосредственно на рабочие станции, серверы и мобильные устройства на базе любой из перечисленных ниже платформ. Просто загрузите приложения и воспользуйтесь мастером установки. Процесс установки интуитивно понятен. Кроме того, вы всегда можете воспользоваться нашей онлайн-справкой."
    summary = summarize(data,word_count=countWord)

    print(summary)

    with open(name + "TR.txt", "w", encoding="cp1251") as myFile:
        myFile.write(summary)

    #from rouge import Rouge

    #reference = "Оптимальное сочетание инструментов и технологий для построения системы безопасности в компаниях с небольшим отделом ИБ. Решение обеспечивает автоматическое обнаружение и оперативное реагирование на скрытые угрозы и отражает принципы ступенчатого подхода «Лаборатории Касперского» к построению стратегии кибербезопасности. Если вы предпочитаете облачную консоль1, создайте учетную запись для доступа к продуктам «Лаборатории Касперского». Для локальной установки консоли управления загрузите и установите Kaspersky Security Center, используя ссылки ниже. В ходе развертывания защиты для рабочих мест вы сможете скачать все необходимые дистрибутивы непосредственно из консоли управления. Процесс развертывания описан в кратком руководстве по началу работы. При желании вы можете установить ПО непосредственно на рабочие станции, серверы и мобильные устройства на базе любой из перечисленных ниже платформ. Просто загрузите приложения и воспользуйтесь мастером установки. Процесс установки интуитивно понятен. Кроме того, вы всегда можете воспользоваться нашей онлайн-справкой."

    #Rouge = Rouge()

    #print(Rouge.get_scores(summary, reference))
