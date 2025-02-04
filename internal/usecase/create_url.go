package usecase

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"url/internal/domain/model"
)

type CreateUrlReq struct {
	LongURL string
}

func (u *UrlUseCase) CreateUrl(req CreateUrlReq) (*model.URL, error) {
	// Шаг 1. Вычисляем SHA-256 хеш строки.
	hasher := sha256.Sum256([]byte(req.LongURL))

	// Преобразуем хеш в большое число.
	hashInt := new(big.Int).SetBytes(hasher[:])

	// Шаг 2. Определяем основание и длину кода.
	base := int64(63)
	length := 10

	// Вычисляем modValue = 63^10.
	modValue := new(big.Int).Exp(big.NewInt(base), big.NewInt(int64(length)), nil)
	// Берём остаток от деления.
	hashInt.Mod(hashInt, modValue)

	// Алфавит допустимых символов.
	charSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

	// Шаг 3. Преобразовываем число в систему счисления с основанием 63.
	// Результат будем заполнять с конца, чтобы гарантировать фиксированную длину.
	result := make([]byte, length)
	bigBase := big.NewInt(base)

	for i := length - 1; i >= 0; i-- {
		remainder := new(big.Int)
		// hashInt.DivMod делит число hashInt на bigBase,
		// записывая остаток в remainder.
		hashInt.DivMod(hashInt, bigBase, remainder)
		// Выбираем соответствующий символ из алфавита.
		result[i] = charSet[remainder.Int64()]
	}
	hash := string(result)

	url := model.NewURL(hash, req.LongURL)
	url, err := u.urlRepo.CreateUrl(url)
	if err != nil {
		return nil, fmt.Errorf("urlRepo.CreateUrl: %w", err)
	}

	return url, nil
}
