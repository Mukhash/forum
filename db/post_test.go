package db

import "testing"

func TestGetTags(t *testing.T) {
	tests := map[string][]string{
		"Simple #tag in the middle":             {"tag"},            //0
		"Simple #tags #in the middle":           {"tags", "in"},     //1
		"Simple tag in the #end":                {"end"},            //2
		"Simple tags in #the #end":              {"the", "end"},     //3
		"#Simple tag at the start":              {"Simple"},         //4
		"#Simple #tags at the start":            {"Simple", "tags"}, //5
		"Simple #tag! closed by non-space byte": {"tag"},            //6
		"Intricate #tag#, no, this #one":        {"one"},            //7
		"Intricate tag #in# the #end":           {"end"},            //8
		"No tag #":                              {},                 //9
		"No tag in the #end#":                   {},                 //10
		"Intricate #tag## with double #hash":    {"hash"},           //11
		"Intricate #tag#3 with #digit":          {"digit"},          //12
		"Intricate #tag######with many #hashes": {"with", "hashes"}, //13
		"Intricate #tag###3#with no tag":        {},                 //14
		"#Intricate# at the #start3":            {"start3"},         //15
		"#######":                               {},                 //16
		"###########3":                          {},                 //17
		"#########tag3":                         {"tag3"},           //18
		"":                                      {},                 //19
		"No tag no bug":                         {},                 //20
		"Теперь на #кириллице в центре": {"кириллице"}, //21
	}

	count := 0
	for k, v := range tests {

		res := GetTags(k)

		if len(v) != len(res) {
			t.Errorf("Lens are different: should be %v\ngot %v\n", v, res)
			return
		}
		for i, tag := range v {
			if res[i] != tag {
				t.Errorf("Test #%d: invalid tag: should be %s\ngot %s\n", count, v, res)
				return
			}
		}
		count++
	}
}
