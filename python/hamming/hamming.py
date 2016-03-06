def distance(dna, copy):
    distance = 0
    for i in range(len(dna)):
        if dna[i] != copy[i]:
            distance += 1
    return distance

