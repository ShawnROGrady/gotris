package game

import "testing"

var schemeFromNameTests = []struct {
	input               string
	expectedDescription string
	expectErr           bool
}{
	{
		input:               HomeRowName,
		expectedDescription: HomeRow().Description(),
	},
	{
		input:               ArrowKeysName,
		expectedDescription: ArrowKeys().Description(),
	},
	{
		input:               StandardName,
		expectedDescription: Standard().Description(),
	},
	{
		input:     "Fake input",
		expectErr: true,
	},
}

func TestSchemeFromName(t *testing.T) {
	for _, test := range schemeFromNameTests {
		scheme, err := SchemeFromName(test.input)
		if err != nil {
			if !test.expectErr {
				t.Errorf("Unexpected error creating scheme from name '%s': %s", test.input, err)
			}
			continue
		}

		if scheme.Description() != test.expectedDescription {
			t.Errorf("Unexpected scheme description from name '%s' [expected = %s, actual = %s]", test.input, test.expectedDescription, scheme.Description())
		}
	}
}

func TestAvailableSchemes(t *testing.T) {
	// check that all available schemes can be constructed using name
	for _, available := range AvailableSchemes() {
		name := available.String()

		scheme, err := SchemeFromName(name)
		if err != nil {
			t.Errorf("Unexpected err constructing scheme from available scheme name '%s': %s", name, err)
			continue
		}

		if scheme.Description() != available.Description() {
			t.Errorf("Unexpected scheme description from available scheme '%s' [expected = %s, actual = %s]", name, available.Description(), scheme.Description())
		}
	}
}

var schemeDescriptionTests = []struct {
	scheme              ControlScheme
	expectedDescription string
	expectedName        string
}{
	{
		scheme:              HomeRow(),
		expectedDescription: "move left: h\nmove right: l\nmove down: j\nmove up: k\nrotate left: a\nrotate right: d",
		expectedName:        "home-row",
	},
	{
		scheme:              ArrowKeys(),
		expectedDescription: "move left: ←\nmove right: →\nmove down: ↓\nmove up: ↑\nrotate left: z\nrotate right: x",
		expectedName:        "arrow-keys",
	},
	{
		scheme:              Standard(),
		expectedDescription: "move left: ←\nmove right: →\nmove down: ↓\nmove up: SPACE\nrotate left: ↑",
		expectedName:        "standard",
	},
	{
		scheme:              ControlSchemes([]ControlScheme{HomeRow(), ArrowKeys()}),
		expectedDescription: "move left: h, ←\nmove right: l, →\nmove down: j, ↓\nmove up: k, ↑\nrotate left: a, z\nrotate right: d, x",
		expectedName:        "home-row, arrow-keys",
	},
}

func TestSchemeDescription(t *testing.T) {
	for _, test := range schemeDescriptionTests {
		if test.scheme.Description() != test.expectedDescription {
			t.Errorf("Unexpected description for scheme '%s' [expected = %s, actual = %s]", test.scheme.String(), test.expectedDescription, test.scheme.Description())
		}

		if test.scheme.String() != test.expectedName {
			t.Errorf("Unexpected string for scheme '%s' [expected = %s, actual = %s]", test.scheme.String(), test.expectedName, test.scheme.String())
		}
	}
}
