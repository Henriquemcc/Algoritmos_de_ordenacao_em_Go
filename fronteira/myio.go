package fronteira

import (
	"fmt"
	"strconv"
	"strings"
)

func ReadString() string {
	//Lendo entrada
	var input string
	for {
		_, erro := fmt.Scan(&input)

		if erro != nil {
			fmt.Println(erro)
		} else {
			break
		}
	}

	return input
}

func ReadUint64() uint64 {
	var unsignedInteger uint64
	for {
		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para uint64
		var erro error
		unsignedInteger, erro = strconv.ParseUint(input, 10, 64)

		if erro != nil {
			fmt.Println(erro)
		} else {
			break
		}

	}

	return unsignedInteger
}

func ReadInt64() int64 {
	var integer int64
	for {
		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para int64
		var erro error
		integer, erro = strconv.ParseInt(input, 10, 64)

		if erro != nil {
			fmt.Println(erro)
		} else {
			break
		}

	}

	return integer
}

func ReadFloat64() float64 {
	var real float64
	for {
		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para float64
		var erro error
		real, erro = strconv.ParseFloat(input, 64)

		if erro != nil {
			fmt.Println(erro)
		} else {
			break
		}

	}

	return real
}

func ReadBool() bool {
	var boolean bool
	for {
		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para float64
		var erro error
		boolean, erro = strconv.ParseBool(input)

		if erro != nil {
			fmt.Println(erro)
		} else {
			break
		}
	}

	return boolean

}
