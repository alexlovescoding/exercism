def is_pangram(string):
    ALPHABETS = "abcdefghijklmnopqrstuvwxyz"
    alph_list = list(ALPHABETS)
    string = string.lower()
    for letter in alph_list:
        if string.find(letter) == -1:
            return False
    return True
