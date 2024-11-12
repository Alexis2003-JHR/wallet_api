package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"golang.org/x/exp/rand"
)

type TransactionDTO struct {
	TransactionType    string
	Amount             float64
	SourceAccount      string
	DestinationAccount string
	Status             string
}

func (s *WalletHandler) GenerateExcel(g *gin.Context) {
	rand.Seed(uint64(time.Now().UnixNano()))
	f := excelize.NewFile()

	sheet := "Sheet1"
	f.NewSheet(sheet)

	headers := []string{"TransactionsType", "Amount", "SourceAccount", "DestinationAccount", "Status"}
	for i, header := range headers {
		f.SetCellValue(sheet, fmt.Sprintf("%s1", string(rune(65+i))), header)
	}

	for row := 2; row <= 10001; row++ {
		transaction := generateTransactionData()
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), transaction.TransactionType)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), transaction.Amount)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), transaction.SourceAccount)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), transaction.DestinationAccount)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), transaction.Status)
	}

	if err := f.SaveAs("transactions.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func generateTransactionData() TransactionDTO {
	transactionTypes := []string{"Deposit", "Withdrawal", "Transfer"}
	statusTypes := []string{"Pending", "Completed", "Failed"}

	transaction := TransactionDTO{
		TransactionType:    transactionTypes[rand.Intn(len(transactionTypes))],
		Amount:             rand.Float64() * 1000,
		SourceAccount:      fmt.Sprintf("ACCT%05d", rand.Intn(10000)),
		DestinationAccount: fmt.Sprintf("ACCT%05d", rand.Intn(10000)),
		Status:             statusTypes[rand.Intn(len(statusTypes))],
	}

	return transaction
}
