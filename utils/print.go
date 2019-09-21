package utils

var count = 1

func PrintSite(site *Site) {
	count = count + 1
	//a := strings.Join([]string{strings.Repeat("    ", indent), (*page).URL.String()}, "")
	//fmt.Println(a)
	if len((*site).Links) > 0 {
		//d := strings.Join([]string{strings.Repeat("    ", indent+1), "Links:"}, "")
		//fmt.Println(d)
		for _, subpage := range (*site).Links {
			PrintSite(subpage)
		}
	}
}
