import re


def hey(what):
    if what == "" or re.compile(r'^\s*\t$').match(what) is not None:
        return "Fine. Be that way!"
    if what.isupper():
        return "Whoa, chill out!"
    if re.compile(r'^.*\?\s*$').match(what) is not None:
        return "Sure."
    return "Whatever."
