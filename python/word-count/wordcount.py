import re
from string import punctuation

def word_count(phrase):
    phrase = phrase.decode("utf-8").lower()
    r = re.compile(ur'[\s{}\U0001f596]+'.format(re.escape(punctuation)))
    phraseList = r.split(phrase)
    return {word: phraseList.count(word) for word in set(phraseList) if word != ''}
