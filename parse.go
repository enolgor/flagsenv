package flagsenv

import (
	"flag"
	"os"
	"strconv"
	"time"
)

type FlagsEnv struct {
	flagSet *flag.FlagSet

	boolVar     func(p *bool, name string, value bool, usage string)
	durationVar func(p *time.Duration, name string, value time.Duration, usage string)
	float64Var  func(p *float64, name string, value float64, usage string)
	int64Var    func(p *int64, name string, value int64, usage string)
	intVar      func(p *int, name string, value int, usage string)
	stringVar   func(p *string, name string, value string, usage string)
	uint64Var   func(p *uint64, name string, value uint64, usage string)
	uintVar     func(p *uint, name string, value uint, usage string)
}

func NewFlagsEnv(flagSet *flag.FlagSet) *FlagsEnv {
	fe := &FlagsEnv{}
	if flagSet != nil {
		fe.boolVar = flagSet.BoolVar
		fe.durationVar = flagSet.DurationVar
		fe.float64Var = flagSet.Float64Var
		fe.int64Var = flagSet.Int64Var
		fe.intVar = flagSet.IntVar
		fe.stringVar = flagSet.StringVar
		fe.uint64Var = flagSet.Uint64Var
		fe.uintVar = flagSet.UintVar
	} else {
		fe.boolVar = flag.BoolVar
		fe.durationVar = flag.DurationVar
		fe.float64Var = flag.Float64Var
		fe.int64Var = flag.Int64Var
		fe.intVar = flag.IntVar
		fe.stringVar = flag.StringVar
		fe.uint64Var = flag.Uint64Var
		fe.uintVar = flag.UintVar
	}
	return fe
}

func (fe *FlagsEnv) Env(envKey string) *Env {
	return &Env{fe, os.Getenv(envKey)}
}

type Env struct {
	flagsEnv *FlagsEnv
	val      string
}

func (e Env) BoolVar(p *bool, name string, value bool, usage string) {
	var defValue bool
	var err error
	if e.val != "" {
		if defValue, err = strconv.ParseBool(e.val); err != nil {
			defValue = value
		}
	} else {
		defValue = value
	}
	e.flagsEnv.boolVar(p, name, defValue, usage)
}

func (e Env) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	var defValue time.Duration
	var err error
	if e.val != "" {
		if defValue, err = time.ParseDuration(e.val); err != nil {
			defValue = value
		}
	} else {
		defValue = value
	}
	e.flagsEnv.durationVar(p, name, defValue, usage)
}

func (e Env) Float64Var(p *float64, name string, value float64, usage string) {
	var defValue float64
	var err error
	if e.val != "" {
		if defValue, err = strconv.ParseFloat(e.val, 64); err != nil {
			defValue = value
		}
	} else {
		defValue = value
	}
	e.flagsEnv.float64Var(p, name, defValue, usage)
}

func (e Env) Int64Var(p *int64, name string, value int64, usage string) {
	var defValue int64
	var err error
	if e.val != "" {
		if defValue, err = strconv.ParseInt(e.val, 0, 64); err != nil {
			defValue = value
		}
	} else {
		defValue = value
	}
	e.flagsEnv.int64Var(p, name, defValue, usage)
}

func (e Env) IntVar(p *int, name string, value int, usage string) {
	var defValue int
	var err error
	if e.val != "" {
		if defValue, err = strconv.Atoi(e.val); err != nil {
			defValue = value
		}
	} else {
		defValue = value
	}
	e.flagsEnv.intVar(p, name, defValue, usage)
}

func (e Env) StringVar(p *string, name string, value string, usage string) {
	var defValue string
	if e.val != "" {
		defValue = e.val
	} else {
		defValue = value
	}
	e.flagsEnv.stringVar(p, name, defValue, usage)
}

func (e Env) Uint64Var(p *uint64, name string, value uint64, usage string) {
	var defValue uint64
	var err error
	if e.val != "" {
		if defValue, err = strconv.ParseUint(e.val, 0, 64); err != nil {
			defValue = value
		}
	} else {
		defValue = value
	}
	e.flagsEnv.uint64Var(p, name, defValue, usage)
}

func (e Env) UintVar(p *uint, name string, value uint, usage string) {
	var defValue uint64
	var err error
	if e.val != "" {
		if defValue, err = strconv.ParseUint(e.val, 0, strconv.IntSize); err != nil {
			defValue = uint64(value)
		}
	} else {
		defValue = uint64(value)
	}
	e.flagsEnv.uintVar(p, name, uint(defValue), usage)
}
