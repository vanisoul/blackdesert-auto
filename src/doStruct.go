package main

type ConfigDo struct {
	Status    bool     `mapstructure:"status"`
	Arms      []string `mapstructure:"arms"`
	PearlArms []string `mapstructure:"pearlArms"`
	Method    []method `mapstructure:"method"`
}

type formula struct {
	Name string `mapstructure:"name"`
	Lov  int    `mapstructure:"LoV"`
}

type method struct {
	Formula []formula `mapstructure:"formula"`
	Recycle []string  `mapstructure:"recycle"`
}
