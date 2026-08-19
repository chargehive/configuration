package policy

import "testing"

func boolPtr(v bool) *bool { return &v }

// Every flag on ScaPolicy is a pointer so an omitted field can be told apart from an
// explicit false, which makes what each getter does with nil part of its contract.
func TestScaPolicyFlagDefaults(t *testing.T) {
	tests := []struct {
		name     string
		get      func(ScaPolicy) bool
		set      func(*ScaPolicy, *bool)
		expectNo bool // value returned when the field is unset
	}{
		{
			"RequireSca", ScaPolicy.GetScaRequired,
			func(p *ScaPolicy, v *bool) { p.RequireSca = v }, false,
		},
		{
			"ShouldIdentify", ScaPolicy.GetShouldIdentify,
			func(p *ScaPolicy, v *bool) { p.ShouldIdentify = v }, false,
		},
		{
			"ShouldChallengeOptional", ScaPolicy.GetShouldChallengeOptional,
			func(p *ScaPolicy, v *bool) { p.ShouldChallengeOptional = v }, false,
		},
		{
			"ShouldAuthOnError", ScaPolicy.GetShouldAuthOnError,
			func(p *ScaPolicy, v *bool) { p.ShouldAuthOnError = v }, true,
		},
		{
			"ShouldAuthOnN", ScaPolicy.GetShouldAuthOnN,
			func(p *ScaPolicy, v *bool) { p.ShouldAuthOnN = v }, true,
		},
		{
			// R is the issuer asking that authorization not be attempted, so an unset
			// field must not authorize. This is the one that differs from its siblings.
			"ShouldAuthOnR", ScaPolicy.GetShouldAuthOnR,
			func(p *ScaPolicy, v *bool) { p.ShouldAuthOnR = v }, false,
		},
		{
			"ShouldAuthOnU", ScaPolicy.GetShouldAuthOnU,
			func(p *ScaPolicy, v *bool) { p.ShouldAuthOnU = v }, true,
		},
	}

	for _, test := range tests {
		t.Run(test.name+"/unset", func(t *testing.T) {
			p := ScaPolicy{}
			test.set(&p, nil)
			if got := test.get(p); got != test.expectNo {
				t.Errorf("unset %s: expected %v, got %v", test.name, test.expectNo, got)
			}
		})

		for _, explicit := range []bool{true, false} {
			t.Run(test.name+"/explicit", func(t *testing.T) {
				p := ScaPolicy{}
				test.set(&p, boolPtr(explicit))
				if got := test.get(p); got != explicit {
					t.Errorf("%s set to %v: got %v", test.name, explicit, got)
				}
			})
		}
	}
}
