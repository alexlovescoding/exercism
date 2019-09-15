const CODON_TO_PROTEIN = {
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
    "UAA": "STOP",
    "UAG": "STOP",
    "UGA": "STOP"
}

export const translate = (rna) => {
    if (!rna) {
        return [];
    }
    let codons = rna.match(/.{1,3}/g);
    let proteins = [];
    for (let i = 0; i < codons.length; i++) {
        let c = codons[i];
        let p = CODON_TO_PROTEIN[c];
        if (!p) {
            throw new Error("Invalid codon");
        } else if (p === 'STOP') {
            return proteins;
        }
        proteins.push(p);
    }
    return proteins;
};
