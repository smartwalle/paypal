package paypal

import (
	"testing"
)

func TestPayPal_CreateWebExperienceProfile(t *testing.T) {
	var p = &WebProfiles{}
	p.Name = "Test Name"
	p.Presentation = &WebProfilesPresentation{}
	p.Presentation.LogoImage = ""
	p.Presentation.BrandName = ""

	//var result, err = getPayPal().CreateWebExperienceProfile(p)
	//fmt.Println(err)
	//fmt.Println(result.Id, result.Presentation.LogoImage)
}