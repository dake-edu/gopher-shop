package catalog

// priceAfterDiscount принимает цену в минимальных денежных единицах.
// Допустимая цена: от 0 до 100000000 включительно; скидка: от 0 до 100.
// Дробную минимальную единицу в скидке отбрасываем в пользу продавца.
func priceAfterDiscount(priceMinor int64, percent int64) (int64, bool) {
	if priceMinor < 0 || priceMinor > 100000000 {
		return 0, false
	}
	if percent < 0 || percent > 100 {
		return 0, false
	}
	discountMinor := priceMinor * percent / 100
	return priceMinor - discountMinor, true
}
