package abstract_factory



type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDetailDAO interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDao() OrderMainDAO
	CreateOrderDetailDao() OrderDetailDAO
}

type RDBMainDAO struct {}
func (*RDBMainDAO) SaveOrderMain() {}
type RDBDetailDAO struct{}
func (*RDBDetailDAO) SaveOrderDetail() {}
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDao() OrderMainDAO {
	return &RDBMainDAO{}
}
func (*RDBDAOFactory) CreateOrderDetailDao() OrderDetailDAO {
	return &RDBDetailDAO{}
}


type XMLMainDAO struct {}
func (*XMLMainDAO) SaveOrderMain() {}
type XMLDetailDAO struct {}
func (*XMLDetailDAO) SaveOrderDetail() {}

type XMLDAOFactory struct {}
func (*XMLDAOFactory) CreateOrderMainDao() OrderMainDAO {
	return &XMLMainDAO{}
}
func (*XMLDAOFactory) CreateOrderDetailDao() OrderDetailDAO {
	return &XMLDetailDAO{}
}
