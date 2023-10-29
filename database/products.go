package database

import (
	"database/sql"
	"errors"
)

const (
	selectProduct = `
	/*name:select_product*/
	SELECT p.id FROM products p
	WHERE p.id = ?;
	`

	selectProductPackSize = `
	/*name:select_product_pack_size*/
	SELECT ps.pid, ps.size FROM packsizes ps
	WHERE ps.pid = ? AND ps.size = ?;
	`

	selectProductPackSizes = `
	/*name:select_product_pack_sizes*/
	SELECT ps.size FROM packsizes ps
	WHERE ps.pid = ?
	ORDER BY ps.size;
	`

	insertProductPackSizes = `
	/*name:insert_product_pack_sizes*/
	INSERT INTO packsizes
	VALUES (?,?);
	`

	deleteProductPackSizes = `
	/*name:delete_product_pack_sizes*/
	DELETE FROM packsizes
	WHERE pid = ? AND size = ?;
	`
)

var (
	ErrProductNotFound       = errors.New("product not found")
	ErrPackSizeAlreadyExists = errors.New("product pack size already exists")
	ErrPackSizeNotFound      = errors.New("product pack size not found")
)

type ProductsTable struct {
	db *sql.DB
}

func NewProductsTable(db *sql.DB) ProductsTable {
	return ProductsTable{
		db: db,
	}
}

func (p ProductsTable) checkProduct(productID int) error {
	row := p.db.QueryRow(selectProduct, productID)
	var pid int
	return row.Scan(&pid)
}

func (p ProductsTable) checkProductPackSize(productID int, size int) error {
	row := p.db.QueryRow(selectProductPackSize, productID, size)
	var pid int
	var psize int
	return row.Scan(&pid, &psize)
}

func (p ProductsTable) GetPackSizes(productID int) ([]int, error) {
	err := p.checkProduct(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrProductNotFound
		}

		return nil, err
	}

	rows, err := p.db.Query(selectProductPackSizes, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]int, 0)
	for rows.Next() {
		var size int
		err = rows.Scan(&size)
		if err != nil {
			return nil, err
		}
		result = append(result, size)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p ProductsTable) AddPackSize(productID int, size int) error {
	err := p.checkProduct(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrProductNotFound
		}

		return err
	}

	err = p.checkProductPackSize(productID, size)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == nil {
		return ErrPackSizeAlreadyExists
	}

	_, err = p.db.Exec(insertProductPackSizes, productID, size)
	return err
}

func (p ProductsTable) DeletePackSize(productID int, size int) error {
	err := p.checkProduct(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrProductNotFound
		}

		return err
	}

	err = p.checkProductPackSize(productID, size)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrPackSizeNotFound
		}

		return err
	}

	_, err = p.db.Exec(deleteProductPackSizes, productID, size)
	return err
}
