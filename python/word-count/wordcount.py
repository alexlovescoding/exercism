import re
from string import punctuation

def word_count(phrase):
    phrase = phrase.decode("utf-8").lower()
    r = re.compile(r'[\s{}]+'.format(re.escape(punctuation)))
    phraseList = r.split(phrase)
    print(phraseList)
    return {word: phraseList.count(word) for word in set(phraseList) if word != ''}
