package product

import (
	"fmt"
	"product-management/models"
	mysql "product-management/sql"
)

func Register(product models.Product) error {
	fmt.Println(product)
	query := "INSERT INTO product (manager_id, category, price, `name`, description, `size`, expired_date) values (?, ?, ?, ?, ?, ?, ?)"
	_, err := mysql.DBConn.Exec(query, product.ManagerID, product.Category, product.Price, product.Name, product.Description, product.Size, product.ExpiredDate)
	if err != nil {
		fmt.Println("error!!!!!!!!", err)
		return err
	}
	return nil
	//db.ExecContext();
	//db.QueryContext();
	//db.QueryRowContext();
	//db.PrepareContext();
}

