def to_rna(dna):
    dict = {"G":"C", "C":"G", "T":"A", "A":"U"}
    dna = list(dna)
    rna = []
    for c in dna:
        rna.append(dict[c])
    return "".join(rna)
