package validity

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecial(t *testing.T) {

	// IBAN

	Convey("Given the validation rule \"IBAN\"", t, func() {

		rule := []string{"iban"}

		isValidIban := func(iban string) bool {
			data := map[string]interface{}{
				"IBAN": iban,
			}
			rules := map[string]Field{
				"IBAN": Field{
					Type:  "Special",
					Name:  "IBAN " + iban,
					Rules: rule,
				},
			}
			return ValidateMap(data, rules).IsValid
		}

		Convey("Given a list of correct IBANs", func() {

			list := []string{
				"RO14CECEGR0201RON0269171",
				"RO74RZBR0000060001635742",
				"RO20RNCB0146009062650001",
				"RO51BRDE190SV09770651900",
				"RO19RZBR0000060014507080",
				"RO14CECEGR0201RON0269171",
				"RO89RNCB0146009060300001",
				"RO69RNCB0146009053910001",
				"RO32RNCB0148009659280001",
				"RO32RZBR0000060008527863",
				"RO35BRDE190SV17599841900",
				"RO35BRDE190SV17599841900",
				"RO79RZBR0000060013901108",
				"RO61RZBR0000060005489107",
				"RO41CECEGR0230RON0350215",
				"RO25RZBR0000060008522848",
				"RO35RNCB0146127215140001",
				"RO50RZBR0000060006249397",

				"RO74rzbr0000060001635742",
			}

			Convey("The result should be valid", func() {

				for _, iban := range list {
					isValid := isValidIban(iban)
					if !isValid {
						log.Println("The IBAN " + iban + " should PASS the test ")
					}
					So(isValid, ShouldBeTrue)
				}

			})

		})

		Convey("Given a list of wrong IBANs", func() {

			list := []string{
				"",
				"RO",
				"RO14CECEGR",
				"RO14CECEGR0201RON0269170",
				// "RO14CECEGR0201RON0269171", -- it is good
				"RO14CECEGR0201RON0269172",
				"RO14CECEGR0201RON0269173",
				"RO14CECEGR0201RON0269174",
				"RO14CECEGR0201RON0269175",
				"RO14CECEGR0201RON0269176",
				"RO14CECEGR0201RON0269177",
				"RO14CECEGR0201RON0269178",
				"RO14CECEGR0201RON0269179",
			}

			Convey("The result should not be valid", func() {

				for _, iban := range list {
					isValid := isValidIban(iban)
					if isValid {
						log.Println("The IBAN " + iban + " should NOT pass the test ")
					}
					So(isValid, ShouldBeFalse)
				}
			})

		})

	})

	// cif

	Convey("Given the validation rule \"CIF\"", t, func() {

		rule := []string{"cif"}

		isValidCif := func(cif string) bool {
			data := map[string]interface{}{
				"cif": cif,
			}
			rules := map[string]Field{
				"cif": Field{
					Type:  "Special",
					Name:  "CIF " + cif,
					Rules: rule,
				},
			}
			return ValidateMap(data, rules).IsValid
		}

		Convey("Given a list of correct CIFs", func() {
			list := []string{

				// too short < 6
				"00",

				// too long > 10
				"00000000001",

				// contains other things
				"aaaaaaaaaa",

				"1450123",

				"13880060",
				"13880061",
				"13880062",
				// "13880063", // <-- this is good
				"13880064",
				"13880065",
				"13880066",
				"13880067",
				"13880068",
				"13880069",
			}

			Convey("The result should be valid", func() {
				for _, cif := range list {
					isValid := isValidCif(cif)
					if isValid {
						log.Println("The CIF " + cif + " should NOT pass the test ")
					}
					So(isValid, ShouldBeFalse)
				}
			})

		})

		Convey("Given a list of wrong CIFs", func() {

			list := []string{

				"4446651",
				"3678190",
				"4352719",
				"3627676",
				"7525881",
				"4347666",
				"3372840",
				"5217524",
				"4317819",
				"4192600",
				"4266928",
				"4233815",
				"4288004",
				"4300787",
				"4300574",
				"4332134",
				"3519380",
				"4375178",
				"4297924",
				"3127336",
				"11078781",
				"3897033",
				"4541874",
				"4605609",
				"4245518",
				"4358002",
				"4230436",
				"4266898",
				"2613486",
				"4318113",
				"2844154",
				"4222310",
				"2540830",
				"4287378",
				"3897343",
				"4540062",
				"5397247",
				"4556140",
				"4231881",
				"4244393",
				"4280213",
				"4323179",
				"4567890",
				"6412248",
				"4269304",
				"4321658",
				"4494721",
				"4404613",

				"13880063", // <-- this is valid
			}

			Convey("The result should not be valid", func() {
				for _, cif := range list {
					isValid := isValidCif(cif)
					if !isValid {
						log.Println("The CIF " + cif + " should PASS the test ")
					}
					So(isValid, ShouldBeTrue)
				}
			})

		})

	})

	// CNP

	Convey("Given the validation rule \"CNP\"", t, func() {

		rule := []string{"cnp"}

		isValidCNP := func(cnp string) bool {
			data := map[string]interface{}{
				"NID": cnp,
			}
			rules := map[string]Field{
				"NID": Field{
					Type:  "Special",
					Name:  "NID " + cnp,
					Rules: rule,
				},
			}
			results := ValidateMap(data, rules)
			return results.IsValid
		}

		Convey("Given a list of correct CNPs", func() {

			list := []string{

				// less than 13 characters
				"00",

				// more than 13 characters
				"00000000000000",

				// not all digits
				"             ",
				"a234567890123",
				"1x34567890123",
				"12y4567890123",
				"123(567890123",
				"1234*67890123",
				"12345-7890123",
				"123456+890123",
				"1234567`90123",
				"12345678!0123",
				"123456789@123",
				"12345678901&3",
				"123456789012=",

				// Digit 1 not 0
				"0000000000000",

				// Wrong month
				// Too small
				"1230000000000",
				// Too big
				"1231300000000",
				"1234500000000",
				"1239900000000",

				// Wrong date
				"1000100000000",
				"1000100000000",
				"1231200000000",
				"3431250000000",
				"8430100000000",
				"6431132000000",
				"5160230000000", // this one does not exist

				// Wrong County
				// 0
				"1930426000000",

				// Too big
				"1930426530000",
				"1930426990000",

				// Number - allowed only 001 --> 999
				"1930426010000",

				// Check controll all posible conbinations
				"1930426450030",
				"1930426450031",
				"1930426450032",
				"1930426450033",
				"1930426450034",
				// "1930426450035", // --> this is good
				"1930426450036",
				"1930426450037",
				"1930426450038",
				"1930426450039",

				"1730513635451",
			}

			Convey("The result should be valid", func() {
				for _, cnp := range list {
					isValid := isValidCNP(cnp)
					if isValid {
						log.Println("The CNP " + cnp + " should NOT pass the test ")
					}
					So(isValid, ShouldBeFalse)
				}
			})

		})

		Convey("Given a list of wrong CNPs", func() {

			list := []string{
				// list from http://cnp-orange-young.blogspot.ro/
				"1851021345131",
				"1920617149053",
				"1870505168646",
				"1870619152998",
				"1921204325416",
				"1931011351347",
				"1860601214764",
				"1930114152084",
				"1930729213031",
				"1931110257176",
				"1900806182888",
				"1920201181713",
				"1911202393786",
				"1860114313017",
				"1900129321797",
				"1910319357474",
				"1881110136241",
				"1910313389932",
				"1930324128027",
				"1860713267955",
				"1920110059921",
				"1920323137670",
				"1910524257786",
				"1910219353588",
				"1921113049026",
				"1910510171180",
				"1921219189118",
				"1930614296921",
				"1930408137659",
				"1860525171600",
				"1871008113318",
				"1900809164551",
				"1870418305339",
				"1920622093590",
				"1860228422041",
				"1881024199710",
				"1901119289825",
				"1870623151354",
				"1880717053385",
				"1860331281263",
				"1921127239892",
				"1860726084246",
				"1891209392109",
				"1870610429939",
				"1891101155151",
				"1900515188610",
				"1900817347850",
				"1850603113108",
				"1920219031621",
				"1880716319353",
				"1851123072213",
				"1890503074079",
				"1920929194685",
				"1930604396824",
				"1920721211943",
				"1911125017410",
				"1920328331453",
				"1891012055602",
				"1850412422567",
				"1900830012030",
				"1901109372453",
				"1901119063169",
				"1910801238953",
				"1850126336642",
				"1870619334654",
				"1880930146221",
				"1861019351849",
				"1930502103999",
				"1930425361934",
				"1871108269030",
				"1860622273209",
				"1920723362204",
				"1880919016077",
				"1860627124609",
				"1930210191367",
				"1920306379899",
				"1860621221169",
				"1890514289320",
				"1920421395061",
				"1930426272907",
				"1861224076805",
				"1850204207530",
				"1930726241832",
				"1920229184105",
				"1910318245177",
				"1871225154992",
				"1890425099446",

				"1920918187490",
				"1911017023863",
				"1900514401426",
				"1860117328230",
				"1861019339785",
				"1870119334824",
				"1880316079025",
				"1910216018631",
				"1920425405811",
				"1920904145205",
				"1881201016278",
				"1920731125084",
				"1901003395568",

				"1590512521591",
				"1660418230010",
				"1891119450065",
				"1920310430037",
				"1920504521696",
				"1930426450035",
				"2770926521590",
				"2830111410090",
				"2880330521699",
				"2900117521690",
				"2910825522143",
			}

			Convey("The result should not be valid", func() {
				for _, cnp := range list {
					isValid := isValidCNP(cnp)
					if !isValid {
						log.Println("The CNP " + cnp + " should PASS the test ")
					}
					So(isValid, ShouldBeTrue)
				}
			})

		})

	})

	// shortDate

	Convey("Given the validation rule \"shortDate\"", t, func() {

		rule := Rules{"Foo": Field{
			Type:  "Special",
			Name:  "Foo",
			Rules: []string{"shortDate"},
		},
		}

		Convey("Given the value \"\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": ""}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"04.10.2015\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "04.10.2015"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"28.01.2016T14:03:15\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "28.01.2016T14:03:15"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// longDate

	Convey("Given the validation rule \"longDate\"", t, func() {

		rule := Rules{"Foo": Field{
			Type:  "Special",
			Name:  "Foo",
			Rules: []string{"longDate"},
		}}

		Convey("Given the value \"\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": ""}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"04.10.2015\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "04.10.2015"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"28.01.2016T14:03:15\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "28.01.2016T14:03:15"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

	})

	// email

	Convey("Given the validation rule \"email\"", t, func() {

		rule := Rules{"Foo": Field{
			Type:  "Special",
			Name:  "Foo",
			Rules: []string{"email"},
		}}

		Convey("Given the value \"\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": ""}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"@.\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "@."}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"popescu\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "popescu"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"popescu@\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "popescu@"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"popescu@vlad\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "popescu@vlad"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"popescu@vlad.\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "popescu@vlad."}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"popescu@vlad.ro\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "popescu@vlad.ro"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given a good value (less or eq than 40 characters)", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "pocupoăescupopescupopescupopescu@vlad.ro"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given a good value, but too long (more than 40 characters)", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "epescuoăescupopescupopescupopescu@vlad.ro"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})
}
