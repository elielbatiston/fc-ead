package gateway

import "github.com.br/elielbatiston/fc-ms-wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
