package service

/*
	service -- the service layer provides a boundary to the backend, exposed through a set of interfaces

	This provides the regions of the world, as a set of static values, that will not change as shown below:

	GetListOfRegions -- List of countries in the world, along with their emoji flag, id and iso values
	GetEmojiForCountry -- Provides, for a given country, the emoji flag
*/

import (
	"github.com/colinSchofield/go-covid/model"
)

type RegionService interface {
	GetListOfRegions() []model.Region
	GetEmojiForCountry(country string) string
	GetIsoForCountry(country string) string
}

type regionService struct {
	regions []model.Region
}

func (rs regionService) GetListOfRegions() []model.Region {
	return rs.regions
}

func (rs regionService) GetEmojiForCountry(country string) string {

	for _, region := range rs.GetListOfRegions() {
		if region.Key == country {
			return region.Flag
		}
	}
	return ""
}

func (rs regionService) GetIsoForCountry(country string) string {

	for _, region := range rs.GetListOfRegions() {
		if region.Key == country {
			return region.Iso
		}
	}
	return ""
}

func NewRegionService() RegionService {

	return regionService{
		regions: []model.Region{
			{Key: "Afghanistan", Location: "Afghanistan", CountryCode: "AF", Iso: "afg", Flag: "🇦🇫"},
			{Key: "Albania", Location: "Albania", CountryCode: "AL", Iso: "alb", Flag: "🇦🇱"},
			{Key: "Algeria", Location: "Algeria", CountryCode: "DZ", Iso: "dza", Flag: "🇩🇿"},
			{Key: "Andorra", Location: "Andorra", CountryCode: "AD", Iso: "and", Flag: "🇦🇩"},
			{Key: "Angola", Location: "Angola", CountryCode: "AO", Iso: "ago", Flag: "🇦🇴"},
			{Key: "Anguilla", Location: "Anguilla", CountryCode: "AI", Iso: "aia", Flag: "🇦🇮"},
			{Key: "Antigua-and-Barbuda", Location: "Antigua And Barbuda", CountryCode: "AG", Iso: "atg", Flag: "🇦🇬"},
			{Key: "Argentina", Location: "Argentina", CountryCode: "AR", Iso: "arg", Flag: "🇦🇷"},
			{Key: "Armenia", Location: "Armenia", CountryCode: "AM", Iso: "arm", Flag: "🇦🇲"},
			{Key: "Aruba", Location: "Aruba", CountryCode: "AW", Iso: "abw", Flag: "🇦🇼"},
			{Key: "Australia", Location: "Australia", CountryCode: "AU", Iso: "aus", Flag: "🇦🇺"},
			{Key: "Austria", Location: "Austria", CountryCode: "AT", Iso: "aut", Flag: "🇦🇹"},
			{Key: "Azerbaijan", Location: "Azerbaijan", CountryCode: "AZ", Iso: "aze", Flag: "🇦🇿"},
			{Key: "Bahamas", Location: "Bahamas", CountryCode: "BS", Iso: "bhs", Flag: "🇧🇸"},
			{Key: "Bahrain", Location: "Bahrain", CountryCode: "BH", Iso: "bhr", Flag: "🇧🇭"},
			{Key: "Bangladesh", Location: "Bangladesh", CountryCode: "BD", Iso: "bgd", Flag: "🇧🇩"},
			{Key: "Barbados", Location: "Barbados", CountryCode: "BB", Iso: "brb", Flag: "🇧🇧"},
			{Key: "Belarus", Location: "Belarus", CountryCode: "BY", Iso: "blr", Flag: "🇧🇾"},
			{Key: "Belgium", Location: "Belgium", CountryCode: "BE", Iso: "bel", Flag: "🇧🇪"},
			{Key: "Belize", Location: "Belize", CountryCode: "BZ", Iso: "blz", Flag: "🇧🇿"},
			{Key: "Benin", Location: "Benin", CountryCode: "BJ", Iso: "ben", Flag: "🇧🇯"},
			{Key: "Bermuda", Location: "Bermuda", CountryCode: "BM", Iso: "bmu", Flag: "🇧🇲"},
			{Key: "Bhutan", Location: "Bhutan", CountryCode: "BT", Iso: "btn", Flag: "🇧🇹"},
			{Key: "Bolivia", Location: "Bolivia", CountryCode: "BO", Iso: "bol", Flag: "🇧🇴"},
			{Key: "Bosnia-and-Herzegovina", Location: "Bosnia And Herzegovina", CountryCode: "BA", Iso: "bih", Flag: "🇧🇦"},
			{Key: "Botswana", Location: "Botswana", CountryCode: "BW", Iso: "bwa", Flag: "🇧🇼"},
			{Key: "Brazil", Location: "Brazil", CountryCode: "BR", Iso: "bra", Flag: "🇧🇷"},
			{Key: "US-Virgin-Islands", Location: "US Virgin Islands", CountryCode: "VG", Iso: "vgb", Flag: "🇻🇬"},
			{Key: "Brunei", Location: "Brunei", CountryCode: "BN", Iso: "brn", Flag: "🇧🇳"},
			{Key: "Bulgaria", Location: "Bulgaria", CountryCode: "BG", Iso: "bgr", Flag: "🇧🇬"},
			{Key: "Burkina-Faso", Location: "Burkina Faso", CountryCode: "BF", Iso: "bfa", Flag: "🇧🇫"},
			{Key: "Burundi", Location: "Burundi", CountryCode: "BI", Iso: "bdi", Flag: "🇧🇮"},
			{Key: "CAR", Location: "CAR", CountryCode: "CF", Iso: "caf", Flag: "🇨🇫"},
			{Key: "Cabo-Verde", Location: "Cabo Verde", CountryCode: "CV", Iso: "cpv", Flag: "🇨🇻"},
			{Key: "Cambodia", Location: "Cambodia", CountryCode: "KH", Iso: "khm", Flag: "🇰🇭"},
			{Key: "Cameroon", Location: "Cameroon", CountryCode: "CM", Iso: "cmr", Flag: "🇨🇲"},
			{Key: "Canada", Location: "Canada", CountryCode: "CA", Iso: "can", Flag: "🇨🇦"},
			{Key: "Caribbean-Netherlands", Location: "Caribbean Netherlands", CountryCode: "", Iso: "bes", Flag: "🇧🇶"},
			{Key: "Cayman-Islands", Location: "Cayman Islands", CountryCode: "KY", Iso: "cym", Flag: "🇰🇾"},
			{Key: "Chad", Location: "Chad", CountryCode: "TD", Iso: "tcd", Flag: "🇹🇩"},
			{Key: "Channel-Islands", Location: "Channel Islands", CountryCode: "JE", Iso: "usa", Flag: "🇯🇪"},
			{Key: "Chile", Location: "Chile", CountryCode: "CL", Iso: "chl", Flag: "🇨🇱"},
			{Key: "China", Location: "China", CountryCode: "CN", Iso: "chn", Flag: "🇨🇳"},
			{Key: "Colombia", Location: "Colombia", CountryCode: "CO", Iso: "col", Flag: "🇨🇴"},
			{Key: "Comoros", Location: "Comoros", CountryCode: "KM", Iso: "com", Flag: "🇰🇲"},
			{Key: "Congo", Location: "Congo", CountryCode: "CG", Iso: "cog", Flag: "🇨🇩"},
			{Key: "Costa-Rica", Location: "Costa Rica", CountryCode: "CR", Iso: "cri", Flag: "🇨🇷"},
			{Key: "Croatia", Location: "Croatia", CountryCode: "HR", Iso: "hrv", Flag: "🇭🇷"},
			{Key: "Cuba", Location: "Cuba", CountryCode: "CU", Iso: "cub", Flag: "🇨🇺"},
			{Key: "Curaçao", Location: "Curaçao", CountryCode: "CW", Iso: "cuw", Flag: "🇨🇼"},
			{Key: "Cyprus", Location: "Cyprus", CountryCode: "CY", Iso: "cyp", Flag: "🇨🇾"},
			{Key: "Czechia", Location: "Czechia", CountryCode: "CZ", Iso: "cze", Flag: "🇨🇿"},
			{Key: "DRC", Location: "DRC", CountryCode: "CD", Iso: "cod", Flag: "🇨🇬"},
			{Key: "Denmark", Location: "Denmark", CountryCode: "DK", Iso: "dnk", Flag: "🇩🇰"},
			{Key: "Djibouti", Location: "Djibouti", CountryCode: "DJ", Iso: "dji", Flag: "🇩🇯"},
			{Key: "Dominica", Location: "Dominica", CountryCode: "DM", Iso: "dma", Flag: "🇩🇲"},
			{Key: "Dominican-Republic", Location: "Dominican Republic", CountryCode: "DO", Iso: "dom", Flag: "🇩🇴"},
			{Key: "Ecuador", Location: "Ecuador", CountryCode: "EC", Iso: "ecu", Flag: "🇪🇨"},
			{Key: "Egypt", Location: "Egypt", CountryCode: "EG", Iso: "egy", Flag: "🇪🇬"},
			{Key: "El-Salvador", Location: "El Salvador", CountryCode: "SV", Iso: "slv", Flag: "🇸🇻"},
			{Key: "Equatorial-Guinea", Location: "Equatorial Guinea", CountryCode: "", Iso: "gnq", Flag: "🇬🇶"},
			{Key: "Eritrea", Location: "Eritrea", CountryCode: "ER", Iso: "eri", Flag: "🇪🇷"},
			{Key: "Estonia", Location: "Estonia", CountryCode: "EE", Iso: "est", Flag: "🇪🇪"},
			{Key: "Eswatini", Location: "Eswatini", CountryCode: "SZ", Iso: "swz", Flag: "🇸🇿"},
			{Key: "Ethiopia", Location: "Ethiopia", CountryCode: "ET", Iso: "eth", Flag: "🇪🇹"},
			{Key: "Faeroe-Islands", Location: "Faeroe Islands", CountryCode: "FO", Iso: "fro", Flag: "🇫🇴"},
			{Key: "Falkland-Islands", Location: "Falkland Islands", CountryCode: "FK", Iso: "flk", Flag: "🇫🇰"},
			{Key: "Fiji", Location: "Fiji", CountryCode: "FJ", Iso: "fji", Flag: "🇫🇯"},
			{Key: "Finland", Location: "Finland", CountryCode: "FI", Iso: "fin", Flag: "🇫🇮"},
			{Key: "France", Location: "France", CountryCode: "FR", Iso: "fra", Flag: "🇫🇷"},
			{Key: "French-Guiana", Location: "French Guiana", CountryCode: "GF", Iso: "guf", Flag: "🇬🇫"},
			{Key: "French-Polynesia", Location: "French Polynesia", CountryCode: "PF", Iso: "pyf", Flag: "🇵🇫"},
			{Key: "Gabon", Location: "Gabon", CountryCode: "GA", Iso: "gab", Flag: "🇬🇦"},
			{Key: "Gambia", Location: "Gambia", CountryCode: "GM", Iso: "gmb", Flag: "🇬🇲"},
			{Key: "Georgia", Location: "Georgia", CountryCode: "GE", Iso: "geo", Flag: "🇬🇪"},
			{Key: "Germany", Location: "Germany", CountryCode: "DE", Iso: "deu", Flag: "🇩🇪"},
			{Key: "Ghana", Location: "Ghana", CountryCode: "GH", Iso: "gha", Flag: "🇬🇭"},
			{Key: "Gibraltar", Location: "Gibraltar", CountryCode: "GI", Iso: "gib", Flag: "🇬🇮"},
			{Key: "Greece", Location: "Greece", CountryCode: "GR", Iso: "grc", Flag: "🇬🇷"},
			{Key: "Greenland", Location: "Greenland", CountryCode: "GL", Iso: "grl", Flag: "🇬🇱"},
			{Key: "Grenada", Location: "Grenada", CountryCode: "GD", Iso: "grd", Flag: "🇬🇩"},
			{Key: "Guadeloupe", Location: "Guadeloupe", CountryCode: "GP", Iso: "glp", Flag: "🇬🇵"},
			{Key: "Guatemala", Location: "Guatemala", CountryCode: "GT", Iso: "gtm", Flag: "🇬🇹"},
			{Key: "Guinea", Location: "Guinea", CountryCode: "GN", Iso: "gin", Flag: "🇬🇳"},
			{Key: "Guinea-Bissau", Location: "Guinea Bissau", CountryCode: "GW", Iso: "gnb", Flag: "🇬🇼"},
			{Key: "Guyana", Location: "Guyana", CountryCode: "GY", Iso: "guy", Flag: "🇬🇾"},
			{Key: "Haiti", Location: "Haiti", CountryCode: "HT", Iso: "hti", Flag: "🇭🇹"},
			{Key: "Honduras", Location: "Honduras", CountryCode: "HN", Iso: "hnd", Flag: "🇭🇳"},
			{Key: "Hong-Kong", Location: "Hong Kong", CountryCode: "HK", Iso: "hkg", Flag: "🇭🇰"},
			{Key: "Hungary", Location: "Hungary", CountryCode: "HU", Iso: "hun", Flag: "🇭🇺"},
			{Key: "Iceland", Location: "Iceland", CountryCode: "IS", Iso: "isl", Flag: "🇮🇸"},
			{Key: "India", Location: "India", CountryCode: "IN", Iso: "ind", Flag: "🇮🇳"},
			{Key: "Indonesia", Location: "Indonesia", CountryCode: "ID", Iso: "idn", Flag: "🇮🇩"},
			{Key: "Iran", Location: "Iran", CountryCode: "IR", Iso: "irn", Flag: "🇮🇷"},
			{Key: "Iraq", Location: "Iraq", CountryCode: "IQ", Iso: "irq", Flag: "🇮🇶"},
			{Key: "Ireland", Location: "Ireland", CountryCode: "IE", Iso: "irl", Flag: "🇮🇪"},
			{Key: "Isle-of-Man", Location: "Isle Of Man", CountryCode: "IM", Iso: "imn", Flag: "🇮🇲"},
			{Key: "Israel", Location: "Israel", CountryCode: "IL", Iso: "isr", Flag: "🇮🇱"},
			{Key: "Italy", Location: "Italy", CountryCode: "IT", Iso: "ita", Flag: "🇮🇹"},
			{Key: "Ivory-Coast", Location: "Ivory Coast", CountryCode: "CI", Iso: "civ", Flag: "🇨🇮"},
			{Key: "Jamaica", Location: "Jamaica", CountryCode: "JM", Iso: "jam", Flag: "🇯🇲"},
			{Key: "Japan", Location: "Japan", CountryCode: "JP", Iso: "jpn", Flag: "🇯🇵"},
			{Key: "Jordan", Location: "Jordan", CountryCode: "JO", Iso: "jor", Flag: "🇯🇴"},
			{Key: "Kazakhstan", Location: "Kazakhstan", CountryCode: "KZ", Iso: "kaz", Flag: "🇰🇿"},
			{Key: "Kenya", Location: "Kenya", CountryCode: "KE", Iso: "ken", Flag: "🇰🇪"},
			{Key: "Kuwait", Location: "Kuwait", CountryCode: "KW", Iso: "kwt", Flag: "🇰🇼"},
			{Key: "Kyrgyzstan", Location: "Kyrgyzstan", CountryCode: "KG", Iso: "kgz", Flag: "🇰🇬"},
			{Key: "Laos", Location: "Laos", CountryCode: "LA", Iso: "lao", Flag: "🇱🇦"},
			{Key: "Latvia", Location: "Latvia", CountryCode: "LV", Iso: "lva", Flag: "🇱🇻"},
			{Key: "Lebanon", Location: "Lebanon", CountryCode: "LB", Iso: "lbn", Flag: "🇱🇧"},
			{Key: "Lesotho", Location: "Lesotho", CountryCode: "LS", Iso: "lso", Flag: "🇱🇸"},
			{Key: "Liberia", Location: "Liberia", CountryCode: "LR", Iso: "lbr", Flag: "🇱🇷"},
			{Key: "Libya", Location: "Libya", CountryCode: "LY", Iso: "lby", Flag: "🇱🇾"},
			{Key: "Liechtenstein", Location: "Liechtenstein", CountryCode: "LI", Iso: "lie", Flag: "🇱🇮"},
			{Key: "Lithuania", Location: "Lithuania", CountryCode: "LT", Iso: "ltu", Flag: "🇱🇹"},
			{Key: "Luxembourg", Location: "Luxembourg", CountryCode: "LU", Iso: "lux", Flag: "🇱🇺"},
			{Key: "Macao", Location: "Macao", CountryCode: "MO", Iso: "mac", Flag: "🇲🇴"},
			{Key: "Madagascar", Location: "Madagascar", CountryCode: "MG", Iso: "mdg", Flag: "🇲🇬"},
			{Key: "Malawi", Location: "Malawi", CountryCode: "MW", Iso: "mwi", Flag: "🇲🇼"},
			{Key: "Malaysia", Location: "Malaysia", CountryCode: "MY", Iso: "mys", Flag: "🇲🇾"},
			{Key: "Maldives", Location: "Maldives", CountryCode: "MV", Iso: "mdv", Flag: "🇲🇻"},
			{Key: "Mali", Location: "Mali", CountryCode: "ML", Iso: "mli", Flag: "🇲🇱"},
			{Key: "Malta", Location: "Malta", CountryCode: "MT", Iso: "mlt", Flag: "🇲🇹"},
			{Key: "Marshall-Islands", Location: "Marshall Islands", CountryCode: "MH", Iso: "mhl", Flag: "🇲🇭"},
			{Key: "Martinique", Location: "Martinique", CountryCode: "MQ", Iso: "mtq", Flag: "🇲🇶"},
			{Key: "Mauritania", Location: "Mauritania", CountryCode: "MR", Iso: "mrt", Flag: "🇲🇷"},
			{Key: "Mauritius", Location: "Mauritius", CountryCode: "MU", Iso: "mus", Flag: "🇲🇺"},
			{Key: "Mayotte", Location: "Mayotte", CountryCode: "YT", Iso: "myt", Flag: "🇾🇹"},
			{Key: "Mexico", Location: "Mexico", CountryCode: "MX", Iso: "mex", Flag: "🇲🇽"},
			{Key: "Moldova", Location: "Moldova", CountryCode: "MD", Iso: "mda", Flag: "🇲🇩"},
			{Key: "Monaco", Location: "Monaco", CountryCode: "MC", Iso: "mco", Flag: "🇲🇨"},
			{Key: "Mongolia", Location: "Mongolia", CountryCode: "MN", Iso: "mng", Flag: "🇲🇳"},
			{Key: "Montenegro", Location: "Montenegro", CountryCode: "ME", Iso: "mne", Flag: "🇲🇪"},
			{Key: "Montserrat", Location: "Montserrat", CountryCode: "MS", Iso: "msr", Flag: "🇲🇸"},
			{Key: "Morocco", Location: "Morocco", CountryCode: "MA", Iso: "mar", Flag: "🇲🇦"},
			{Key: "Mozambique", Location: "Mozambique", CountryCode: "MZ", Iso: "moz", Flag: "🇲🇿"},
			{Key: "Myanmar", Location: "Myanmar", CountryCode: "MM", Iso: "mmr", Flag: "🇲🇲"},
			{Key: "Namibia", Location: "Namibia", CountryCode: "NA", Iso: "nam", Flag: "🇳🇦"},
			{Key: "Nepal", Location: "Nepal", CountryCode: "NP", Iso: "npl", Flag: "🇳🇵"},
			{Key: "Netherlands", Location: "Netherlands", CountryCode: "NL", Iso: "nld", Flag: "🇳🇱"},
			{Key: "New-Caledonia", Location: "New Caledonia", CountryCode: "NC", Iso: "ncl", Flag: "🇳🇨"},
			{Key: "New-Zealand", Location: "New Zealand", CountryCode: "NZ", Iso: "nzl", Flag: "🇳🇿"},
			{Key: "Nicaragua", Location: "Nicaragua", CountryCode: "NI", Iso: "nic", Flag: "🇳🇮"},
			{Key: "Niger", Location: "Niger", CountryCode: "NE", Iso: "ner", Flag: "🇳🇪"},
			{Key: "Nigeria", Location: "Nigeria", CountryCode: "NG", Iso: "nga", Flag: "🇳🇬"},
			{Key: "North-Macedonia", Location: "North Macedonia", CountryCode: "MK", Iso: "mkd", Flag: "🇲🇰"},
			{Key: "Norway", Location: "Norway", CountryCode: "NO", Iso: "nor", Flag: "🇳🇴"},
			{Key: "Oman", Location: "Oman", CountryCode: "OM", Iso: "omn", Flag: "🇴🇲"},
			{Key: "Pakistan", Location: "Pakistan", CountryCode: "PK", Iso: "pak", Flag: "🇵🇰"},
			{Key: "Palestine", Location: "Palestine", CountryCode: "PS", Iso: "pse", Flag: "🇵🇸"},
			{Key: "Panama", Location: "Panama", CountryCode: "PA", Iso: "pan", Flag: "🇵🇦"},
			{Key: "Papua-New-Guinea", Location: "Papua New Guinea", CountryCode: "PG", Iso: "png", Flag: "🇵🇬"},
			{Key: "Paraguay", Location: "Paraguay", CountryCode: "PY", Iso: "pry", Flag: "🇵🇾"},
			{Key: "Peru", Location: "Peru", CountryCode: "PE", Iso: "per", Flag: "🇵🇪"},
			{Key: "Philippines", Location: "Philippines", CountryCode: "PH", Iso: "phl", Flag: "🇵🇭"},
			{Key: "Poland", Location: "Poland", CountryCode: "PL", Iso: "pol", Flag: "🇵🇱"},
			{Key: "Portugal", Location: "Portugal", CountryCode: "PT", Iso: "prt", Flag: "🇵🇹"},
			{Key: "Qatar", Location: "Qatar", CountryCode: "QA", Iso: "qat", Flag: "🇶🇦"},
			{Key: "Romania", Location: "Romania", CountryCode: "RO", Iso: "rou", Flag: "🇷🇴"},
			{Key: "Russia", Location: "Russia", CountryCode: "RU", Iso: "rus", Flag: "🇷🇺"},
			{Key: "Rwanda", Location: "Rwanda", CountryCode: "RW", Iso: "rwa", Flag: "🇷🇼"},
			{Key: "Réunion", Location: "Réunion", CountryCode: "RE", Iso: "reu", Flag: "🇷🇪"},
			{Key: "S-Korea", Location: "S Korea", CountryCode: "KR", Iso: "kor", Flag: "🇰🇷"},
			{Key: "Saint-Kitts-and-Nevis", Location: "Saint Kitts And Nevis", CountryCode: "KN", Iso: "kna", Flag: "🇰🇳"},
			{Key: "Saint-Lucia", Location: "Saint Lucia", CountryCode: "LC", Iso: "lca", Flag: "🇱🇨"},
			{Key: "Saint-Martin", Location: "Saint Martin", CountryCode: "MF", Iso: "maf", Flag: "🇲🇫"},
			{Key: "Saint-Pierre-Miquelon", Location: "Saint Pierre Miquelon", CountryCode: "PM", Iso: "spm", Flag: "🇵🇲"},
			{Key: "Samoa", Location: "Samoa", CountryCode: "WS", Iso: "wsm", Flag: "🇼🇸"},
			{Key: "San-Marino", Location: "San Marino", CountryCode: "SM", Iso: "smr", Flag: "🇸🇲"},
			{Key: "Sao-Tome-and-Principe", Location: "Sao Tome And Principe", CountryCode: "ST", Iso: "stp", Flag: "🇸🇹"},
			{Key: "Saudi-Arabia", Location: "Saudi Arabia", CountryCode: "sau", Iso: "sau", Flag: "🇸🇦"},
			{Key: "Senegal", Location: "Senegal", CountryCode: "SN", Iso: "sen", Flag: "🇸🇳"},
			{Key: "Serbia", Location: "Serbia", CountryCode: "RS", Iso: "srb", Flag: "🇷🇸"},
			{Key: "Seychelles", Location: "Seychelles", CountryCode: "SC", Iso: "syc", Flag: "🇸🇨"},
			{Key: "Sierra-Leone", Location: "Sierra Leone", CountryCode: "SL", Iso: "sle", Flag: "🇸🇱"},
			{Key: "Singapore", Location: "Singapore", CountryCode: "SG", Iso: "sgp", Flag: "🇸🇬"},
			{Key: "Sint-Maarten", Location: "Sint Maarten", CountryCode: "", Iso: "sxm", Flag: "🇸🇽"},
			{Key: "Slovakia", Location: "Slovakia", CountryCode: "SK", Iso: "svk", Flag: "🇸🇰"},
			{Key: "Slovenia", Location: "Slovenia", CountryCode: "SI", Iso: "svn", Flag: "🇸🇮"},
			{Key: "Solomon-Islands", Location: "Solomon Islands", CountryCode: "SB", Iso: "slb", Flag: "🇸🇧"},
			{Key: "Somalia", Location: "Somalia", CountryCode: "SO", Iso: "som", Flag: "🇸🇴"},
			{Key: "South-Africa", Location: "South Africa", CountryCode: "ZA", Iso: "zaf", Flag: "🇿🇦"},
			{Key: "South-Sudan", Location: "South Sudan", CountryCode: "SS", Iso: "ssd", Flag: "🇸🇸"},
			{Key: "Spain", Location: "Spain", CountryCode: "ES", Iso: "esp", Flag: "🇪🇸"},
			{Key: "Sri-Lanka", Location: "Sri Lanka", CountryCode: "LK", Iso: "lka", Flag: "🇱🇰"},
			{Key: "St-Barth", Location: "St Barth", CountryCode: "BL", Iso: "blm", Flag: "🇧🇱"},
			{Key: "St-Vincent-Grenadines", Location: "St Vincent Grenadines", CountryCode: "VC", Iso: "vct", Flag: "🇻🇨"},
			{Key: "Sudan", Location: "Sudan", CountryCode: "SD", Iso: "sdn", Flag: "🇸🇩"},
			{Key: "Suriname", Location: "Suriname", CountryCode: "SR", Iso: "sur", Flag: "🇸🇷"},
			{Key: "Sweden", Location: "Sweden", CountryCode: "SE", Iso: "swe", Flag: "🇸🇪"},
			{Key: "Switzerland", Location: "Switzerland", CountryCode: "CH", Iso: "che", Flag: "🇨🇭"},
			{Key: "Syria", Location: "Syria", CountryCode: "SY", Iso: "syr", Flag: "🇸🇾"},
			{Key: "Taiwan", Location: "Taiwan", CountryCode: "TW", Iso: "twn", Flag: "🇹🇼"},
			{Key: "Tajikistan", Location: "Tajikistan", CountryCode: "TJ", Iso: "tjk", Flag: "🇹🇯"},
			{Key: "Tanzania", Location: "Tanzania", CountryCode: "TZ", Iso: "tza", Flag: "🇹🇿"},
			{Key: "Thailand", Location: "Thailand", CountryCode: "TH", Iso: "tha", Flag: "🇹🇭"},
			{Key: "Timor-Leste", Location: "Timor Leste", CountryCode: "TL", Iso: "tls", Flag: "🇹🇱"},
			{Key: "Togo", Location: "Togo", CountryCode: "TG", Iso: "tgo", Flag: "🇹🇬"},
			{Key: "Trinidad-and-Tobago", Location: "Trinidad And Tobago", CountryCode: "TT", Iso: "tto", Flag: "🇹🇹"},
			{Key: "Tunisia", Location: "Tunisia", CountryCode: "TN", Iso: "tun", Flag: "🇹🇳"},
			{Key: "Turkey", Location: "Turkey", CountryCode: "TR", Iso: "tur", Flag: "🇹🇷"},
			{Key: "Turks-and-Caicos", Location: "Turks And Caicos", CountryCode: "TC", Iso: "tca", Flag: "🇹🇨"},
			{Key: "UAE", Location: "UAE", CountryCode: "AE", Iso: "are", Flag: "🇦🇪"},
			{Key: "UK", Location: "UK", CountryCode: "UK", Iso: "gbr", Flag: "🇬🇧"},
			{Key: "USA", Location: "USA", CountryCode: "US", Iso: "usa", Flag: "🇺🇸"},
			{Key: "Uganda", Location: "Uganda", CountryCode: "UG", Iso: "uga", Flag: "🇺🇬"},
			{Key: "Ukraine", Location: "Ukraine", CountryCode: "UA", Iso: "ukr", Flag: "🇺🇦"},
			{Key: "Uruguay", Location: "Uruguay", CountryCode: "UY", Iso: "ury", Flag: "🇺🇾"},
			{Key: "Uzbekistan", Location: "Uzbekistan", CountryCode: "UZ", Iso: "uzb", Flag: "🇺🇿"},
			{Key: "Vanuatu", Location: "Vanuatu", CountryCode: "VU", Iso: "vut", Flag: "🇻🇺"},
			{Key: "Vatican-City", Location: "Vatican City", CountryCode: "VA", Iso: "vat", Flag: "🇻🇦"},
			{Key: "Venezuela", Location: "Venezuela", CountryCode: "VE", Iso: "ven", Flag: "🇻🇪"},
			{Key: "Vietnam", Location: "Vietnam", CountryCode: "VN", Iso: "vnm", Flag: "🇻🇳"},
			{Key: "Wallis-and-Futuna", Location: "Wallis and Futuna", CountryCode: "WF", Iso: "wlf", Flag: "🇼🇫"},
			{Key: "Western-Sahara", Location: "Western Sahara", CountryCode: "EH", Iso: "esh", Flag: "🇪🇭"},
			{Key: "Yemen", Location: "Yemen", CountryCode: "YE", Iso: "yem", Flag: "🇾🇪"},
			{Key: "Zambia", Location: "Zambia", CountryCode: "ZM", Iso: "zmb", Flag: "🇿🇲"},
			{Key: "Zimbabwe", Location: "Zimbabwe", CountryCode: "ZW", Iso: "zwe", Flag: "🇿🇼"},
		},
	}
}
