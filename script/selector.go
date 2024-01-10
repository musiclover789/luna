package script

func GetElementPositionByCss(selector string) string {
	return `
		function getElementPosition(selector) {
		  var element = document.querySelector(selector);
		  if (element) {
			var rect = element.getBoundingClientRect();
			return {
			  ok: true,
			  top: rect.top + window.scrollY,
			  left: rect.left + window.scrollX,
			  width: rect.width,
			  height: rect.height
			};
		  } else {
			return { ok: false};
		  }
		}	
		getElementPosition("` + selector + `");
	`
}

func GetElementPositionByXpath(selector string) string {
	return `
		function getElementPositionByXPath(xpath) {
		  var element = document.evaluate(xpath, document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
		  if (element) {
			var rect = element.getBoundingClientRect();
			return {
		      ok: true,
			  top: rect.top + window.scrollY,
			  left: rect.left + window.scrollX,
			  width: rect.width,
			  height: rect.height
			};
		  } else {
			return { ok: false};
		  }
		}
		getElementPositionByXPath(` + "`" + selector + "`" + `);

	`
}
