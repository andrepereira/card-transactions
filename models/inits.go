package models

// InitOperationTypesIDsTable Receive no argument
// Return nothing
// Used to initialize a in memory table  with Operation Types IDs and your descriptions
func InitOperationTypesIDsTable() {
	//Create Operation Type "Table" as a slice of Operation Types
	var compraAVista OperationType
	compraAVista.OperationTypeID = 1
	compraAVista.Description = "COMPRA A VISTA"

	var compraParcelada OperationType
	compraParcelada.OperationTypeID = 2
	compraParcelada.Description = "COMPRA PARCELADA"

	var saque OperationType
	saque.OperationTypeID = 3
	saque.Description = "SAQUE"

	var pagamento OperationType
	pagamento.OperationTypeID = 4
	pagamento.Description = "PAGAMENTO"

	OperationTypes = append(OperationTypes, compraAVista,
		compraParcelada, saque, pagamento)

}
