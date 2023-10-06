package config

// Framework-less "dependency injection" for code paths that are difficult to pass these Arguments to

var arguments *Arguments

func configure(args *Arguments) {
	arguments = args
}

func GetArguments() *Arguments {
	return arguments
}
