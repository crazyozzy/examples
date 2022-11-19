package math

import "testing"

func TestAddPositive(t *testing.T) {
    sum, err := Add(1, 2)
    if err != nil {
        t.Error("unexpected error")
    }
    if sum != 3 {
        t.Errorf("sum expected to be 3; got %d", sum)
    }
}


func TestAddNegative(t *testing.T) {
    _, err := Add(-1, 2)
    if err == nil {
        t.Error("first arg negative - expected error not be nil" )
    }
    _, err = Add(1, -2)
    if err == nil {
        t.Error("second arg negative - expected error not be nil" )
    }
    _, err = Add(-1, -2)
    if err == nil {
        t.Error("all arg negative - expected error not be nil" )
    }
}

func TestAddNull(t *testing.T) {
    _, err := Add(0, 1)
    if err == nil {
        t.Error("first arg is 0 - expected error not be nil")
    }

    _, err = Add(1, 0)
    if err == nil {
        t.Error("second arg is 0 - expected error not be nil")
    }

    _, err = Add(0, 0)
    if err == nil {
        t.Error("all arg is 0 - expected error not be nil")
    }
}

func TestEstimateValue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
            name: "small",
            args: args{
                value: 5,
            },
            want: "small",
        },
        {
            name: "medium",
            args: args{
                value: 23,
            },
            want: "medium",
        },
        {
            name: "big",
            args: args{
                value: 537,
            },
            want: "big",
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EstimateValue(tt.args.value); got != tt.want {
				t.Errorf("EstimateValue() = %v, want %v", got, tt.want)
			}
		})
	}
}