package domain

import "regexp"

const EnvironmentStage = "stage"
const EnvironmentProduction = "production"

var PhonePattern, _ = regexp.Compile("^[+]?[(]?[0-9]{3}[)]?[-\\s\\.]?[0-9]{3}[-\\s\\.]?[0-9]{4,6}$")
