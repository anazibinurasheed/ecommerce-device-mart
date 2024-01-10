package repo

import (
	"time"

	interfaces "github.com/anazibinurasheed/project-device-mart/pkg/repo/interface"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/request"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"gorm.io/gorm"
)

type orderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{
		DB: DB,
	}
}

func (od *orderDatabase) InsertOrder(order request.NewOrder) (response.OrderLine, error) {
	var NewOrderLine response.OrderLine
	query := `INSERT INTO order_lines (user_id,product_id,addresses_id,qty,price,payment_method_id,order_status_id,coupon_id,created_at,updated_at)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING * ;`
	err := od.DB.Raw(query, order.UserID, order.ProductID, order.AddressID, order.Qty, order.Price, order.PaymentMethodID, order.OrderStatusID, order.CouponID, order.CreatedAt, order.UpdatedAt).Scan(&NewOrderLine).Error
	return NewOrderLine, err
}

func (od *orderDatabase) GetStatusPending() (response.OrderStatus, error) {
	var OrderStatus response.OrderStatus
	query := `SELECT * FROM order_statuses WHERE status = 'Pending';`
	err := od.DB.Raw(query).Scan(&OrderStatus).Error
	return OrderStatus, err
}

func (od *orderDatabase) GetStatusCancelled() (response.OrderStatus, error) {
	var OrderStatus response.OrderStatus
	query := `SELECT * FROM order_statuses WHERE status = 'Cancelled';`
	err := od.DB.Raw(query).Scan(&OrderStatus).Error
	return OrderStatus, err
}

func (od *orderDatabase) GetStatusReturned() (response.OrderStatus, error) {
	var OrderStatus response.OrderStatus
	query := `SELECT * FROM order_statuses WHERE status = 'Returned';`
	err := od.DB.Raw(query).Scan(&OrderStatus).Error
	return OrderStatus, err
}

func (od *orderDatabase) GetUserOrderHistory(userID, startIndex, endIndex int) ([]response.Orders, error) {
	var OrderHistory = make([]response.Orders, 0)
	query := `SELECT
    o.id AS order_id,
    o.product_id,
    p.product_image,
    p.product_name,
    p.price AS product_price,
    o.order_status_id,
    s.status AS order_status,
    o.payment_method_id,
    m.method_name AS payment_method,
    o.addresses_id,
    CONCAT(a.name, ', ', a.locality, ', ', a.address_line, ', ', a.district, ', ', states.name, ', ', a.landmark, ', ', a.pincode, ', ', a.phone_number, ', ', a.alternative_phone) AS delivery_address
FROM
    order_lines o
INNER JOIN products p ON o.product_id = p.id
INNER JOIN order_statuses s ON o.order_status_id = s.id
INNER JOIN payment_methods m ON o.payment_method_id = m.id
INNER JOIN addresses a ON o.addresses_id = a.id
INNER JOIN states states ON a.state_id = states.id
WHERE o.user_id = $1
ORDER BY o.id DESC OFFSET $2 FETCH NEXT $3 ROW ONLY;`
	err := od.DB.Raw(query, userID, startIndex, endIndex).Scan(&OrderHistory).Error
	return OrderHistory, err
}

func (od *orderDatabase) GetInvoiceDataByID(orderID int) (response.Orders, error) {
	var OrderData response.Orders
	query := `SELECT
    o.id AS order_id,
    o.product_id,
    p.product_image,
    p.product_name,
    p.price AS product_price,
    o.order_status_id,
    s.status AS order_status,
    o.payment_method_id,
    m.method_name AS payment_method,
    o.addresses_id,
    CONCAT(a.name, ', ', a.locality, ', ', a.address_line, ', ', a.district, ', ', states.name, ', ', a.landmark, ', ', a.pincode, ', ', a.phone_number, ', ', a.alternative_phone) AS delivery_address
FROM
    order_lines o
INNER JOIN products p ON o.product_id = p.id
INNER JOIN order_statuses s ON o.order_status_id = s.id
INNER JOIN payment_methods m ON o.payment_method_id = m.id
INNER JOIN addresses a ON o.addresses_id = a.id
INNER JOIN states states ON a.state_id = states.id
WHERE  o.id = $1;`

	err := od.DB.Raw(query, orderID).Scan(&OrderData).Error

	return OrderData, err
}

func (od *orderDatabase) GetAllOrderData(startIndex, endIndex int) ([]response.Orders, error) {
	var OrderHistory = make([]response.Orders, 0)
	query := `SELECT
    o.id AS order_id,
    o.product_id,
    p.images,
    p.product_name,
    p.price AS product_price,
    o.order_status_id,
    s.status AS order_status,
    o.payment_method_id,
    m.method_name AS payment_method,
    o.addresses_id,
    CONCAT(a.name, ', ', a.locality, ', ', a.address_line, ', ', a.district, ', ', states.name, ', ', a.landmark, ', ', a.pincode, ', ', a.phone_number, ', ', a.alternative_phone) AS delivery_address
FROM
    order_lines o
INNER JOIN products p ON o.product_id = p.id
INNER JOIN order_statuses s ON o.order_status_id = s.id
INNER JOIN payment_methods m ON o.payment_method_id = m.id
INNER JOIN addresses a ON o.addresses_id = a.id
INNER JOIN states states ON a.state_id = states.id
ORDER BY o.id DESC OFFSET $1 FETCH NEXT $2 ROW ONLY;
`
	err := od.DB.Raw(query, startIndex, endIndex).Scan(&OrderHistory).Error

	return OrderHistory, err
}

func (od *orderDatabase) GetOrderStatuses() ([]response.OrderStatus, error) {
	var OrderStatus = make([]response.OrderStatus, 0)
	query := `SELECT * FROM order_statuses ORDER BY id;`
	err := od.DB.Raw(query).Scan(&OrderStatus).Error
	return OrderStatus, err

}

func (od *orderDatabase) ChangeOrderStatusByID(statusID int, orderID int) (response.OrderLine, error) {
	var UpdatedOrder response.OrderLine
	query := `UPDATE order_lines SET order_status_id = $1 WHERE id = $2 RETURNING * ;`
	err := od.DB.Raw(query, statusID, orderID).Scan(&UpdatedOrder).Error
	return UpdatedOrder, err
}

func (od *orderDatabase) FindOrderByUserIDAndProductID(userID, productID int) (response.OrderLine, error) {
	var OrderData response.OrderLine

	query := `SELECT * FROM order_lines WHERE user_id= $1 AND product_id=$2 ;`
	err := od.DB.Raw(query, userID, productID).Scan(&OrderData).Error
	return OrderData, err
}

func (od *orderDatabase) FindOrderStatusByID(statusID int) (string, error) {
	var status string
	query := `SELECT status FROM order_statuses WHERE id = $1 ;`
	err := od.DB.Raw(query, statusID).Scan(&status).Error

	return status, err
}

func (od *orderDatabase) FindOrderByID(orderID int) (response.OrderLine, error) {
	var Order response.OrderLine
	query := `SELECT * FROM order_lines WHERE id = $1 ;`
	err := od.DB.Raw(query, orderID).Scan(&Order).Error
	return Order, err
}

func (od *orderDatabase) FindOrdersBoughtUsingCoupon(couponID int) ([]response.OrderLine, error) {
	var Orders = make([]response.OrderLine, 0)
	query := `SELECT * FROM order_lines WHERE coupon_id = $1 ; `
	err := od.DB.Raw(query, couponID).Scan(&Orders).Error
	return Orders, err
}

func (od *orderDatabase) InitializeNewUserWallet(userID int) (response.Wallet, error) {
	var NewWallet response.Wallet

	query := `INSERT INTO wallets (user_id,amount)VALUES($1,$2) RETURNING *;`
	err := od.DB.Raw(query, userID, 0).Scan(&NewWallet).Error
	return NewWallet, err
}

func (od *orderDatabase) FindUserWalletByID(userID int) (response.Wallet, error) {
	var Wallet response.Wallet

	query := `SELECT * FROM wallets WHERE user_id = $1 ;`
	err := od.DB.Raw(query, userID).Scan(&Wallet).Error
	return Wallet, err
}

func (od *orderDatabase) UpdateUserWalletBalance(userID int, amount float32) (response.Wallet, error) {
	var UpdatedWallet response.Wallet

	query := `UPDATE wallets SET amount = $1 WHERE user_id = $2 RETURNING *;`
	err := od.DB.Raw(query, amount, userID).Scan(&UpdatedWallet).Error
	return UpdatedWallet, err
}
func (od *orderDatabase) UpdateWalletTransactionHistory(update request.WalletTransactionHistory) (response.WalletTransactionHistory, error) {

	var updatedHistory response.WalletTransactionHistory

	query := `INSERT INTO wallet_transaction_histories(transaction_time,user_id,amount,transaction_type) 
	VALUES($1,$2,$3,$4) RETURNING *;`

	err := od.DB.Raw(query, update.TransactionTime, update.UserID, update.Amount, updatedHistory.TransactionType).Scan(&updatedHistory).Error

	return updatedHistory, err
}

func (od *orderDatabase) GetWalletHistoryByUserID(userID int) ([]response.WalletTransactionHistory, error) {

	var walletHistory = make([]response.WalletTransactionHistory, 0)
	query := `SELECT * FROM wallet_transaction_histories WHERE user_id = $1 ;`
	err := od.DB.Raw(query, userID).Scan(&walletHistory).Error
	return walletHistory, err
}

func (o *orderDatabase) TopSellingProduct(startDate, endDate time.Time) (response.TopSelling, error) {

	var data response.TopSelling
	query := ` SELECT product_id, SUM(qty) AS quantity
	FROM order_lines	WHERE created_at >= ? AND created_at <= ?
	GROUP BY product_id
	ORDER BY quantity DESC
	LIMIT 1;`

	err := o.DB.Raw(query, startDate, endDate).Scan(&data).Error
	return data, err

}

func (o *orderDatabase) GetTotalSaleCount(startDate, endDate time.Time) (int, error) {

	var count int
	query := `SELECT COUNT(*) AS count
	FROM order_lines
	WHERE created_at >= $1 AND created_at <= $2;`

	err := o.DB.Raw(query, startDate, endDate).Scan(&count).Error
	return count, err

}

func (o *orderDatabase) GetAverageOrderValue(startDate, endDate time.Time) (float32, error) {

	var avg float32
	query := `SELECT AVG(qty * price) AS average_order_value
	FROM order_lines WHERE created_at >= $1 AND created_at <= $2 ;`

	err := o.DB.Raw(query, startDate, endDate).Scan(&avg).Error
	return avg, err
}

func (o *orderDatabase) GetTotalRevenue(status int, startDate, endDate time.Time) ([]response.OrderLine, error) {
	var data = make([]response.OrderLine, 0)
	query := `SELECT * from order_lines WHERE created_at >= $1 AND created_at <= $2 AND order_status_id != $3 ;`

	err := o.DB.Raw(query, startDate, endDate, status).Scan(&data).Error
	return data, err
}
