package repo

import (
	"context"
	"log"
)

type GakSooCache struct {
	GakSooMap map[int]string
}

func (r *GakSooRepo) RFindGakSooAllID(ctx context.Context) (*GakSooCache, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, gak_num FROM gak_soo")
	if err != nil {
		log.Printf("❌ FAILED INITIALIZE GAK SOO :: %s", err)
		return nil, err
	}
	defer rows.Close()

	cache := &GakSooCache{
		GakSooMap: make(map[int]string),
	}

	for rows.Next() {
		var (
			id     string
			number int
		)
		rows.Scan(&id, &number)
		cache.GakSooMap[number] = id
	}
	log.Println("✅ SUCCESS INITIALIZE GAK SOO")
	return cache, nil
}
