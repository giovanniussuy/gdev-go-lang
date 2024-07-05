package chainresponsability

//This function will perform all the steps to simulate medical care
func chainResponsabilityService(p *Patient) {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	//Patient visiting
	reception.execute(p)
}
