package user

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func (db *DbService) CreateInventory(ctx context.Context, ni NewShirtInventory, userId string, now time.Time) (ShirtInventory, error) {
	inv := ShirtInventory{
		UserId:      userId,
		ItemName:    ni.ItemName,
		Quantity:    ni.Quantity,
		DateCreated: now,
		DateUpdated: now,
	}
	const q = `INSERT INTO inventory
	(user_id, item_name, quantity, date_created, date_updated)
	VALUES ( $1, $2, $3, $4, $5)
	Returning id`

	var id int
	row := db.QueryRowContext(ctx, q, userId, inv.ItemName, inv.Quantity, inv.DateCreated, inv.DateUpdated)
	err := row.Scan(&id)
	if err != nil {
		return ShirtInventory{}, fmt.Errorf("inserting inventory %w", err)
	}
	inv.ID = strconv.Itoa(id)
	return inv, nil

}
func (db *DbService) ViewAll(ctx context.Context, userId string) ([]ShirtInventory, error) {
	var inv []ShirtInventory
	rows, err := db.QueryContext(ctx, "Select Id,user_id,item_name,quantity,date_created,date_updated FROM inventory where user_id =$1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var newInv ShirtInventory
		err = rows.Scan(&newInv.ID, &newInv.UserId, &newInv.ItemName, &newInv.Quantity, &newInv.DateCreated, &newInv.DateUpdated)
		if err != nil {
			return nil, err
		}
		inv = append(inv, newInv)
	}
	return inv, nil
}
