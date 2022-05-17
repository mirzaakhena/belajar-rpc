package model

import "fmt"

type Input struct {
	Pertama int
	Kedua   int
}

type Output struct {
	Hasil int
}

type Kalkulator struct {
}

func (r Kalkulator) Jumlah(input *Input, output *Output) error {

	if input.Pertama == 0 || input.Kedua == 0 {
		return fmt.Errorf("input gak boleh nol")
	}

	output.Hasil = input.Pertama + input.Kedua

	fmt.Printf("calculate %d=%d+%d", output.Hasil, input.Pertama, input.Kedua)

	return nil
}
