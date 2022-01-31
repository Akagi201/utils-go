package jobber_test

import (
	"testing"

	"github.com/Akagi201/utils-go/jobber"
	"github.com/stretchr/testify/require"
)

func TestParseFullTimeSpec(t *testing.T) {
	evens := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22}
	threes := []int{1, 4, 7, 10, 13, 16, 19, 22}
	cases := []struct {
		str  string
		spec jobber.FullTimeSpec
	}{
		{
			"0 0 14",
			jobber.FullTimeSpec{
				jobber.OneValTimeSpec{0},
				jobber.OneValTimeSpec{0},
				jobber.OneValTimeSpec{14},
				jobber.WildcardTimeSpec{},
				jobber.WildcardTimeSpec{},
				jobber.WildcardTimeSpec{},
			},
		},
		{
			"0 0 14 * * 1",
			jobber.FullTimeSpec{
				jobber.OneValTimeSpec{0},
				jobber.OneValTimeSpec{0},
				jobber.OneValTimeSpec{14},
				jobber.WildcardTimeSpec{},
				jobber.WildcardTimeSpec{},
				jobber.OneValTimeSpec{1},
			},
		},
		{
			"0 0 */2 * * 1",
			jobber.FullTimeSpec{
				jobber.OneValTimeSpec{0},
				jobber.OneValTimeSpec{0},
				jobber.SetTimeSpec{"*/2", evens},
				jobber.WildcardTimeSpec{},
				jobber.WildcardTimeSpec{},
				jobber.OneValTimeSpec{1},
			},
		},
		{
			"0 0 1,4,7,10,13,16,19,22 * * 1",
			jobber.FullTimeSpec{
				jobber.OneValTimeSpec{0},
				jobber.OneValTimeSpec{0},
				jobber.SetTimeSpec{"1,4,7,10,13,16,19,22", threes},
				jobber.WildcardTimeSpec{},
				jobber.WildcardTimeSpec{},
				jobber.OneValTimeSpec{1},
			},
		},
		{
			"10,20 0 14 1 8 0-5",
			jobber.FullTimeSpec{
				jobber.SetTimeSpec{"10,20", []int{10, 20}},
				jobber.OneValTimeSpec{0},
				jobber.OneValTimeSpec{14},
				jobber.OneValTimeSpec{1},
				jobber.OneValTimeSpec{8},
				jobber.SetTimeSpec{"0-5", []int{0, 1, 2, 3, 4, 5}},
			},
		},
	}

	for _, c := range cases {
		/*
		 * Call
		 */
		var result *jobber.FullTimeSpec
		var err error
		result, err = jobber.ParseFullTimeSpec(c.str)

		/*
		 * Test
		 */
		if err != nil {
			t.Fatalf("Got error: %v", err)
		}
		require.Equal(t, c.spec, *result)
	}
}
