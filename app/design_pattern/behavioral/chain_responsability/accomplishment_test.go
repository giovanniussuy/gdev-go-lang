package chainresponsability

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test(t *testing.T) {

	patient := &Patient{name: "abc"}

	chainResponsabilityService(patient)

	assert.Equal(t, patient.doctorCheckUpDone, true)
	assert.Equal(t, patient.medicineDone, true)
	assert.Equal(t, patient.paymentDone, true)
	assert.Equal(t, patient.registrationDone, true)
	assert.Equal(t, patient.name, "abc")
}
