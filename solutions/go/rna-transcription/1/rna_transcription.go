package rnatranscription

func ToRNA(dna string) string {
    var rna []rune
    for _, sign := range dna {
        var mapped rune
        switch {
            case sign == 'G':
                mapped = 'C'
            case sign == 'C':
            	mapped = 'G'
            case sign == 'T':
            	mapped = 'A'
            case sign == 'A':
            	mapped = 'U'
        }
        rna = append(rna, mapped)
    }
    return string(rna)
}
