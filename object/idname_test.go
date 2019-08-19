package object

import "testing"

func TestCleanse(t *testing.T) {
	badId := "This)Us=A_£$-Bad__id  to forMa.t,"
	expect := "thisusa-badidtoformat"
	produced := CleanName(badId)
	if expect != produced {
		t.Errorf("expected %s, produced %s", expect, produced)
	}

	badId = "this-id-is-fine-but-too-long-fwkjgfwkejgfkwjegwlekhfwehfwekjgwekjgkwjegwekjghfwejqlqwkhgkwqehg;kwjhgwkjghwjhgwejh348349743"
	expect = "this-id-is-fine-but-too-long-fwkjgfwkejg"
	produced = CleanName(badId)
	if expect != produced {
		t.Errorf("expected %s, produced %s", expect, produced)
	}
}
