package q5

import (
	"sort"
	"strings"
)

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	s = strings.ToLower(s)

	s = strings.ReplaceAll(s, "a", "")
	s = strings.ReplaceAll(s, "e", "")
	s = strings.ReplaceAll(s, "i", "")
	s = strings.ReplaceAll(s, "o", "")
	s = strings.ReplaceAll(s, "u", "")

	var consoantes = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}

	var t string = ""

	for _, letra := range s {
		var l = string(letra)

		if contem(consoantes, l) {
			t = t + "." + l
		} else {
			t = t + l
		}
	}

	return t
}

func contem(s []string, termo string) bool {
	i := sort.SearchStrings(s, termo)
	return i < len(s) && s[i] == termo
}
