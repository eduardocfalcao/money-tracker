package transactions

import (
	"context"
	"errors"
	"io"

	"github.com/aclindsa/ofxgo"
	"github.com/eduardocfalcao/money-tracker/internal/api"
	"github.com/eduardocfalcao/money-tracker/internal/transactions/models"
	"github.com/eduardocfalcao/money-tracker/internal/transactions/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) ImportOFXFile(ctx context.Context, file io.Reader) error {
	user, _ := api.GetContextUser(ctx)
	ofx, err := ofxgo.ParseResponse(file)
	if err != nil {
		return err
	}

	stmt, ok := ofx.Bank[0].(*ofxgo.StatementResponse)
	if !ok {
		return errors.New("not possible to cast 'resp.Bank[0]' to '*ofxgo.StatementResponse)'")
	}

	for _, ofx := range stmt.BankTranList.Transactions {
		amount, _ := ofx.TrnAmt.Float64()
		t := models.RawTransaction{
			TransactionAmount: amount,
			Type:              pgtype.Text{String: ofx.TrnType.String()},
			DatePosted:        ofx.DtPosted.Time,
			AccountID:         pgtype.Text{String: stmt.BankAcctFrom.AcctID.String()},
			Memo:              pgtype.Text{String: ofx.Memo.String()},
			FitID:             pgtype.Text{String: ofx.FiTID.String()},
			Checknum:          pgtype.Text{String: ofx.CheckNum.String()},
			UserID:            user.UserID,
		}

		if err := s.repository.InsertRawTransaction(ctx, t); err != nil {
			return err
		}
	}

	return nil
}
