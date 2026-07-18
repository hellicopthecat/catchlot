package repo

import "context"

func (r OfficialLottoRepo) RExistsLottoRounds(ctx context.Context) (bool, error) {
	var exsist bool
	q := "SELECT EXISTS(SELECT 1 FROM lotto_rounds LIMIT 1)"
	err := r.db.QueryRowContext(ctx, q).Scan(&exsist)
	return exsist, err
}
