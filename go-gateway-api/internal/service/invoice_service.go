package service

import (
	"github.com/romeolacerda/payment-gateway/go-gateway-api/internal/domain"
	"github.com/romeolacerda/payment-gateway/go-gateway-api/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountService:    accountService,
	}
}

func (s *InvoiceService) Create(input dto.CreateInvoiceInput) (*dto.InvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(&input, accountOutput.ID)
	if err != nil {
		return nil, err
	}

	if err = invoice.Proccess(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(input.APIKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err = s.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}

	output := dto.FromInvoice(invoice)
	return &output, nil
}

func (s *InvoiceService) GetById(id, apiKey string) (*dto.InvoiceOutput, error) {
	invoice, err := s.invoiceRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	account, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != account.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	output := dto.FromInvoice(invoice)
	return &output, nil
}

func (s *InvoiceService) ListByAccount(accountID string) ([]*dto.InvoiceOutput, error) {
	invoice, err := s.invoiceRepository.FindByAccountId(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceOutput, len(invoice))
	for i, invoice := range invoice {
		invoiceOutput := dto.FromInvoice(invoice)
		output[i] = &invoiceOutput
	}

	return output, nil
}

func (s *InvoiceService) ListByAccountAPIKey(apiKey string) ([]*dto.InvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output, err := s.ListByAccount(accountOutput.ID)
	if err != nil {
		return nil, err
	}
	return output, nil

}
