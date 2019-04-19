package btcutil_test

import (
	"fmt"
	"math"

	"github.com/btcsuite/btcutil"
)

func ExampleAmount() {

	a := btcutil.Amount(0)
	fmt.Println("Zero gro:", a)

	a = btcutil.Amount(1e8)
	fmt.Println("100,000,000 gros:", a)

	a = btcutil.Amount(1e5)
	fmt.Println("100,000 gros:", a)
	// Output:
	// Zero gro: 0 GRS
	// 100,000,000 gros: 1 GRS
	// 100,000 gros: 0.001 GRS
}

func ExampleNewAmount() {
	amountOne, err := btcutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := btcutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := btcutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := btcutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 GRS
	// 0.01234567 GRS
	// 0 GRS
	// invalid groestlcoin amount
}

func ExampleAmount_unitConversions() {
	amount := btcutil.Amount(44433322211100)

	fmt.Println("gro to kGRS:", amount.Format(btcutil.AmountKiloBTC))
	fmt.Println("gro to GRS:", amount)
	fmt.Println("gro to MilliGRS:", amount.Format(btcutil.AmountMilliBTC))
	fmt.Println("gro to MicroGRS:", amount.Format(btcutil.AmountMicroBTC))
	fmt.Println("gro to gro:", amount.Format(btcutil.AmountSatoshi))

	// Output:
	// gro to kGRS: 444.333222111 kGRS
	// gro to GRS: 444333.222111 GRS
	// gro to MilliGRS: 444333222.111 mGRS
	// gro to MicroGRS: 444333222111 μGRS
	// gro to gro: 44433322211100 gro
}
