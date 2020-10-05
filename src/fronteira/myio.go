package fronteira

import (
	"fmt"
	"strconv"
	"strings"
)

// ReadString le uma String da entrada padrao (teclado).
// Retorna uma String com os dados lidos.
func ReadString() string {
	var input string

	var repetir = true
	for repetir {
		_, erro := fmt.Scan(&input)

		if erro != nil {
			fmt.Println(erro)
		} else {
			repetir = false
		}
	}

	return input
}

// ReadUint8 le um numero inteiro unsigned de 8 bits da entrada padrao.
// Retorna um numero inteiro unsigned de 8 bits.
func ReadUint8() uint8 {
	var unsignedInteger uint8

	var repetir = true
	for repetir {

		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para uint64
		var erro error
		var unsignedInteger64 uint64
		unsignedInteger64, erro = strconv.ParseUint(input, 10, 8)

		if erro != nil {
			fmt.Println(erro)
		} else {

			//Convertendo de uint64 para uint8
			unsignedInteger = uint8(unsignedInteger64)
			repetir = false
		}

	}

	return unsignedInteger
}

// ReadUint16 le um numero inteiro unsigned de 16 bits da entrada padrao.
// Retorna um numero inteiro unsigned de 16 bits.
func ReadUint16() uint16 {
	var unsignedInteger uint16

	var repetir = true
	for repetir {

		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para uint64
		var erro error
		var unsignedInteger64 uint64
		unsignedInteger64, erro = strconv.ParseUint(input, 10, 16)

		if erro != nil {
			fmt.Println(erro)
		} else {

			//Convertendo de uint64 para uint16
			unsignedInteger = uint16(unsignedInteger64)
			repetir = false
		}

	}

	return unsignedInteger

}

// ReadUint32 le um numero inteiro unsigned de 32 bits da entrada padrao.
// Retorna um numero inteiro unsigned de 32 bits.
func ReadUint32() uint32 {
	var unsignedInteger uint32

	var repetir = true
	for repetir {

		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para uint64
		var erro error
		var unsignedInteger64 uint64
		unsignedInteger64, erro = strconv.ParseUint(input, 10, 32)

		if erro != nil {
			fmt.Println(erro)
		} else {

			//Convertendo de uint64 para uint32
			unsignedInteger = uint32(unsignedInteger64)
			repetir = false
		}

	}

	return unsignedInteger
}

// ReadUint64 le um numero inteiro unsigned de 64 bits da entrada padrao.
// Retorna um numero inteiro unsigned de 64 bits.
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

// ReadInt8 le um numero inteiro de 8 bits da entrada padrao.
// Retorna um numero inteiro de 8 bits.
func ReadInt8() int8 {
	var integer int8

	var repetir = true
	for repetir {

		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para int64
		var erro error
		var integer64 int64
		integer64, erro = strconv.ParseInt(input, 10, 8)

		if erro != nil {
			fmt.Println(erro)
		} else {

			//Convertendo de int64 para int8
			integer = int8(integer64)
			repetir = false
		}

	}

	return integer
}

// ReadInt16 le um numero inteiro de 16 bits da entrada padrao.
// Retorna um numero inteiro de 16 bits.
func ReadInt16() int16 {
	var integer int16

	var repetir = true
	for repetir {

		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para int64
		var erro error
		var integer64 int64
		integer64, erro = strconv.ParseInt(input, 10, 16)

		if erro != nil {
			fmt.Println(erro)
		} else {

			//Convertendo de int64 para int16
			integer = int16(integer64)
			repetir = false
		}

	}

	return integer
}

// ReadInt32 le um numero inteiro de 32 bits da entrada padrao.
// Retorna um numero inteiro de 32 bits.
func ReadInt32() int32 {
	var integer int32

	var repetir = true
	for repetir {

		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para int64
		var erro error
		var integer64 int64
		integer64, erro = strconv.ParseInt(input, 10, 32)

		if erro != nil {
			fmt.Println(erro)
		} else {

			//Convertendo de int64 para int32
			integer = int32(integer64)
			repetir = false
		}

	}

	return integer

}

// ReadInt64 le um numero inteiro de 64 bits da entrada padrao.
// Retorna um numero inteiro de 64 bits.
func ReadInt64() int64 {
	var integer int64

	var repetir = true
	for repetir {

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
			repetir = false
		}

	}

	return integer
}

// ReadFloat32 le um numero real de 32 bits da entrada padrao.
// Retorna um numero real de 32 bits.
func ReadFloat32() float32 {
	var real float32

	var repetir = true
	for repetir {

		//Lendo entrada
		var input = ReadString()

		//Removendo espacos
		input = strings.TrimSpace(input)

		//Convertendo para float64
		var erro error
		var real64 float64
		real64, erro = strconv.ParseFloat(input, 64)

		if erro != nil {
			fmt.Println(erro)
		} else {

			//Convertendo de float64 para float32
			real = float32(real64)
			repetir = false
		}

	}

	return real
}

// ReadFloat64 le um numero real de 64 bits da entrada padrao.
// Retorna um numero real de 64 bits.
func ReadFloat64() float64 {
	var real float64

	var repetir = true
	for repetir {

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
			repetir = false
		}

	}

	return real
}

// ReadBool le um valor booleano da entrada padrao.
// Retorna um valor booleano da entrada padrao.
func ReadBool() bool {
	var boolean bool

	var repetir = true
	for repetir {

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
			repetir = false
		}
	}

	return boolean

}

// ReadChar le um caractere de entrada padrao.
// Retorna o caractere lido.
func ReadChar() rune {
	var character rune

	//Lendo entrada
	var input = ReadString()

	//Convertendo para char
	character = rune(input[0])

	return character
}
