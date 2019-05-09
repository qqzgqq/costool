package template

// StringUPSET upset string
func StringUPSET(s string) string {
	var str []string
	str = StringCHANGE2(StringTOSHUZU(s))
	str = StringCHANGE(str)
	return SHUZUtoSTring(str)
}

// StringRESTORE restore string
func StringRESTORE(s string) string {
	var str []string
	str = StringCHANGE(StringTOSHUZU(s))
	str = StringCHANGE2(str)
	return SHUZUtoSTring(str)
}
