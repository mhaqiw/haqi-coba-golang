package domain

type Warehouse struct {
	ID                 string `json:"id"`
	WarehouseAddress   string `json:"warehouse_address"`
	WarehouseCity      string `json:"warehouse_city"`
	WarehouseLatitude  string `json:"warehouse_latitide"`
	WarehouseLongitude string `json:"warehouse_longitude"`
}


type WarehouseRepository interface {
	CheckWarehouseByID(id int) error
}