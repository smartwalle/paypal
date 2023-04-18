package paypal_test

import (
	"github.com/smartwalle/paypal"
	"testing"
)

func TestPayPal_CreateWebExperienceProfile(t *testing.T) {
	var p = &paypal.WebProfiles{}
	p.Name = "Test Name"
	p.Presentation = &paypal.WebProfilesPresentation{}
	p.Presentation.LogoImage = ""
	p.Presentation.BrandName = ""

	var result, err = client.CreateWebExperienceProfile(p)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.Id, result.Presentation.LogoImage)
}

func TestPayPal_GetWebExperienceProfileList(t *testing.T) {
	var result, err = client.GetWebExperienceProfileList()
	if err != nil {
		t.Fatal(err)
	}

	for _, profile := range result {
		t.Logf(profile.Id, profile.Name)
	}
}

func TestPayPal_DeleteWebExperienceProfile(t *testing.T) {
	//var err = client.DeleteWebExperienceProfile("XP-6MSR-MERU-75MJ-SCXL")
	//t.Log(err)
}

func TestPayPal_GetWebhookDetails(t *testing.T) {
	var webhook, err = client.GetWebExperienceProfileDetails("XP-BEFQ-A67P-RNXQ-LJLM")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(webhook.Id, webhook.Name)
}
