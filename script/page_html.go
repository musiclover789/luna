package script
//document.documentElement.outerHTML;

func OuterHTML() string {
	return `
	function getScreenInfo() {
	  return document.documentElement.outerHTML;
	}
    getScreenInfo();
`
}
