def word_count(phrase):
    phraseList = phrase.split()
    phraseSet = set(phraseList)
    dict = {}
    for word in phraseSet:
        dict[word] = phraseList.count(word)
    return dict
