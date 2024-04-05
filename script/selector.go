package script

func GetElementPositionByCss(selector string) string {
	//return `
	//	function getElementPosition(selector) {
	//	  var element = document.querySelector(selector);
	//	  if (element) {
	//		var rect = element.getBoundingClientRect();
	//		return {
	//		  ok: true,
	//		  top: rect.top + window.scrollY,
	//		  left: rect.left + window.scrollX,
	//		  width: rect.width,
	//		  height: rect.height
	//		};
	//	  } else {
	//		return { ok: false};
	//	  }
	//	}
	//	getElementPosition("` + selector + `");
	//`

	return `
	function getElementViewportCoordinates(selector) {
    const element = document.querySelector(selector);
    
    if (!element) {
        return { ok: false};
    }
    const rect = element.getBoundingClientRect();
    return {
 		ok: true,
        top: rect.top,
        left: rect.left,
        width: rect.width,
        height: rect.height
    };
   }

   getElementViewportCoordinates("` + selector + `");

 `
}

func GetElementPositionByXpath(selector string) string {
	//return `
	//	function getElementPositionByXPath(xpath) {
	//	  var element = document.evaluate(xpath, document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
	//	  if (element) {
	//		var rect = element.getBoundingClientRect();
	//		return {
	//	      ok: true,
	//		  top: rect.top + window.scrollY,
	//		  left: rect.left + window.scrollX,
	//		  width: rect.width,
	//		  height: rect.height
	//		};
	//	  } else {
	//		return { ok: false};
	//	  }
	//	}
	//	getElementPositionByXPath(` + "`" + selector + "`" + `);
	//
	//`
	return `
	function getElementByXpath(path) {
		return document.evaluate(path, document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
	}
	
	function getElementViewportCoordinatesByXpath(xpath) {
		const element = getElementByXpath(xpath);
		
		if (!element) {
			return { ok: false};
		}
	
		const rect = element.getBoundingClientRect();
	
		return {
 			ok: true,
			top: rect.top,
			left: rect.left,
			width: rect.width,
			height: rect.height
		};
	}
	getElementViewportCoordinatesByXpath(` + "`" + selector + "`" + `);
	`
}

func JSGetFirstChildElementByCss(selector string) string {
	return `
function getCssSelector(element) {
    if (!(element instanceof Element)) return;
    const selectorList = [];
    while (element.parentElement) {
        let selector = element.tagName.toLowerCase();
        let siblings = Array.from(element.parentElement.children);
            let index = siblings.findIndex(sibling => sibling === element);
            selector += ':nth-child(' + (index + 1) + ')';
            selectorList.unshift(selector);
            element = element.parentElement;
    }
    return selectorList.join(' > ');
}

function getXpathSelector(element) {
    const idx = (sibling, name) => sibling ? idx(sibling.previousElementSibling, name || sibling.tagName) + (sibling.tagName == name) : 1;
    return element.tagName.toLowerCase() + '[' + idx(element) + ']';
}

function getFirstChildNodeInfo(selector) {
    const parentElement = document.querySelector(selector);
    if (parentElement) {
        const firstChildNode = parentElement.firstElementChild;
        if (firstChildNode) {
            const nodeInfo = {
                nodeType: firstChildNode.nodeType,
                nodeName: firstChildNode.nodeName,
                nodeValue: firstChildNode.nodeValue,
                textContent: firstChildNode.textContent,
                htmlContent: firstChildNode.outerHTML,
                cssSelector: getCssSelector(firstChildNode),
                xpathSelector: getXpathSelector(firstChildNode)
            };
            return JSON.stringify(nodeInfo, null, 2);
        } else {
            console.log('该元素没有子节点');
            return null;
        }
    } else {
        console.log('找不到匹配的父元素');
        return null;
    }
}

getFirstChildNodeInfo("` + selector + `")
`
}

func JSGetLastChildElementByCss(selector string) string {
	return `
function getCssSelector(element) {
    if (!(element instanceof Element)) return;
    const selectorList = [];
    while (element.parentElement) {
        let selector = element.tagName.toLowerCase();
         let siblings = Array.from(element.parentElement.children);
            let index = siblings.findIndex(sibling => sibling === element);
            selector += ':nth-child(' + (index + 1) + ')';
            selectorList.unshift(selector);
            element = element.parentElement;
    }
    return selectorList.join(' > ');
}

function getXpathSelector(element) {
    const idx = (sibling, name) => sibling ? idx(sibling.previousElementSibling, name || sibling.tagName) + (sibling.tagName == name) : 1;
    return element.tagName.toLowerCase() + '[' + idx(element) + ']';
}

function getLastChildNodeInfo(selector) {
    const parentElement = document.querySelector(selector);
    if (parentElement) {
        const lastChildNode = parentElement.lastElementChild;
        if (lastChildNode) {
            const nodeInfo = {
                nodeType: lastChildNode.nodeType,
                nodeName: lastChildNode.nodeName,
                nodeValue: lastChildNode.nodeValue,
                textContent: lastChildNode.textContent,
                htmlContent: lastChildNode.outerHTML,
                cssSelector: getCssSelector(lastChildNode),
                xpathSelector: getXpathSelector(lastChildNode)
            };
            return JSON.stringify(nodeInfo, null, 2);
        } else {
            console.log('该元素没有子节点');
            return null;
        }
    } else {
        console.log('找不到匹配的父元素');
        return null;
    }
}
getLastChildNodeInfo("` + selector + `");
`
}

/**
nextSibling: 下一个兄弟节点
*/

func JSGetNextSiblingElementByCss(selector string) string {
	return `
function getCssSelector(element) {
    if (!(element instanceof Element)) return;
    const selectorList = [];
    while (element.parentElement) {
        let selector = element.tagName.toLowerCase();
       let siblings = Array.from(element.parentElement.children);
            let index = siblings.findIndex(sibling => sibling === element);
            selector += ':nth-child(' + (index + 1) + ')';
            selectorList.unshift(selector);
            element = element.parentElement;
    }
    return selectorList.join(' > ');
}

function getXpathSelector(element) {
    const idx = (sibling, name) => sibling ? idx(sibling.previousElementSibling, name || sibling.tagName) + (sibling.tagName == name) : 1;
    return element.tagName.toLowerCase() + '[' + idx(element) + ']';
}



function getNextSiblingNodeInfo(selector) {
    const currentNode = document.querySelector(selector);
    if (currentNode) {
        const nextSiblingNode = currentNode.nextElementSibling;
        if (nextSiblingNode) {
            const nodeInfo = {
                nodeType: nextSiblingNode.nodeType,
                nodeName: nextSiblingNode.nodeName,
                nodeValue: nextSiblingNode.nodeValue,
                textContent: nextSiblingNode.textContent,
                htmlContent: nextSiblingNode.outerHTML,
                cssSelector: getCssSelector(nextSiblingNode),
                xpathSelector: getXpathSelector(nextSiblingNode)
            };
            return JSON.stringify(nodeInfo, null, 2);
        } else {
            console.log('该节点没有下一个兄弟节点');
            return null;
        }
    } else {
        console.log('找不到匹配的节点');
        return null;
    }
}

getNextSiblingNodeInfo("` + selector + `");
`
}

/**
nextSibling: 上一个兄弟节点
*/

func JSPreviousSiblingElementByCss(selector string) string {
	return `
function getCssSelector(element) {
    if (!(element instanceof Element)) return;
    const selectorList = [];
    while (element.parentElement) {
        let selector = element.tagName.toLowerCase();
       let siblings = Array.from(element.parentElement.children);
            let index = siblings.findIndex(sibling => sibling === element);
            selector += ':nth-child(' + (index + 1) + ')';
            selectorList.unshift(selector);
            element = element.parentElement;
    }
    return selectorList.join(' > ');
}

function getXpathSelector(element) {
    const idx = (sibling, name) => sibling ? idx(sibling.previousElementSibling, name || sibling.tagName) + (sibling.tagName == name) : 1;
    return element.tagName.toLowerCase() + '[' + idx(element) + ']';
}



function getPreviousSiblingNodeInfo(selector) {
    const currentNode = document.querySelector(selector);
    if (currentNode) {
        const previousSiblingNode = currentNode.previousElementSibling;
        if (previousSiblingNode) {
            const nodeInfo = {
                nodeType: previousSiblingNode.nodeType,
                nodeName: previousSiblingNode.nodeName,
                nodeValue: previousSiblingNode.nodeValue,
                textContent: previousSiblingNode.textContent,
                htmlContent: previousSiblingNode.outerHTML,
                cssSelector: getCssSelector(previousSiblingNode),
                xpathSelector: getXpathSelector(previousSiblingNode)
            };
            return JSON.stringify(nodeInfo, null, 2);
        } else {
            console.log('该节点没有上一个兄弟节点');
            return null;
        }
    } else {
        console.log('找不到匹配的节点');
        return null;
    }
}

getPreviousSiblingNodeInfo("` + selector + `");
`
}

/**
nextSibling: 上一个兄弟节点
*/

func JSParentElementByCss(selector string) string {
	return `
	function getCssSelector(element) {
    if (!(element instanceof Element)) return;
    const selectorList = [];
    while (element.parentElement) {
        let selector = element.tagName.toLowerCase();
         let siblings = Array.from(element.parentElement.children);
            let index = siblings.findIndex(sibling => sibling === element);
            selector += ':nth-child(' + (index + 1) + ')';
            selectorList.unshift(selector);
            element = element.parentElement;
    }
    return selectorList.join(' > ');
}

function getXpathSelector(element) {
    const idx = (sibling, name) => sibling ? idx(sibling.previousElementSibling, name || sibling.tagName) + (sibling.tagName == name) : 1;
    return element.tagName.toLowerCase() + '[' + idx(element) + ']';
}



function getParentNodeInfo(selector) {
    const currentNode = document.querySelector(selector);
    if (currentNode) {
        const parentNode = currentNode.parentNode;
        if (parentNode) {
            const nodeInfo = {
                nodeType: parentNode.nodeType,
                nodeName: parentNode.nodeName,
                nodeValue: parentNode.nodeValue,
                textContent: parentNode.textContent,
                htmlContent: parentNode.outerHTML,
                cssSelector: getCssSelector(parentNode),
                xpathSelector: getXpathSelector(parentNode)
            };
            return JSON.stringify(nodeInfo, null, 2);
        } else {
            console.log('该节点没有父节点');
            return null;
        }
    } else {
        console.log('找不到匹配的节点');
        return null;
    }
}

getParentNodeInfo("` + selector + `");

`
}

/*
根据css选择器,返回节点信息
*/
func JSGetElementBycss(selector string) string {
	return `
function getNodeInfo(selector) {
    const element = document.querySelector(selector);
    if (element) {
        const nodeInfo = {
            nodeType: element.nodeType,
            nodeName: element.nodeName,
            nodeValue: element.nodeValue,
            textContent: element.textContent,
            htmlContent: element.outerHTML,
            cssSelector: getCssSelector(element),
            xpathSelector: getXpathSelector(element)
        };
        return JSON.stringify(nodeInfo, null, 2);
    } else {
        console.log('找不到匹配的元素');
        return null;
    }
}

function getCssSelector(element) {
    if (!(element instanceof Element)) return;
    const selectorList = [];
    while (element.parentElement) {
        let selector = element.tagName.toLowerCase();
         let siblings = Array.from(element.parentElement.children);
            let index = siblings.findIndex(sibling => sibling === element);
            selector += ':nth-child(' + (index + 1) + ')';
            selectorList.unshift(selector);
            element = element.parentElement;
    }
    return selectorList.join(' > ');
}

function getXpathSelector(element) {
    const idx = (sibling, name) => sibling ? idx(sibling.previousElementSibling, name || sibling.tagName) + (sibling.tagName == name) : 1;
    return element.tagName.toLowerCase() + '[' + idx(element) + ']';
}

getNodeInfo("` + selector + `"); // 选择的CSS选择器
`
}

func JSGetAllChildElementByCss(selector string) string {
	return `
	function getAllChildNodesInfo(selector) {
    const element = document.querySelector(selector);
    if (element) {
        const childNodes = Array.from(element.children).map(child => getNodeInfo(child));
        return JSON.stringify(childNodes, null, 2);
    } else {
        console.log('找不到匹配的元素');
        return null;
    }
}

function getNodeInfo(node) {
    const nodeInfo = {
        nodeType: node.nodeType,
        nodeName: node.nodeName,
        nodeValue: node.nodeValue,
        textContent: node.textContent,
        htmlContent: node.outerHTML,
        cssSelector: getCssSelector(node),
        xpathSelector: getXpathSelector(node)
    };
    return nodeInfo;
}

function getCssSelector(element) {
    if (!(element instanceof Element)) return;
    const selectorList = [];
    while (element.parentElement) {
        let selector = element.tagName.toLowerCase();
        let siblings = Array.from(element.parentElement.children);
            let index = siblings.findIndex(sibling => sibling === element);
            selector += ':nth-child(' + (index + 1) + ')';
            selectorList.unshift(selector);
            element = element.parentElement;
    }
    return selectorList.join(' > ');
}

function getXpathSelector(element) {
    const idx = (sibling, name) => sibling ? idx(sibling.previousElementSibling, name || sibling.tagName) + (sibling.tagName == name) : 1;
    return element.tagName.toLowerCase() + '[' + idx(element) + ']';
}

getAllChildNodesInfo("` + selector + `");
`
}

func JSGetElementPositionAndWindowViewportByCss(selector string) string {

	return `
function getElementPositionAndWindowViewport(cssSelector) {
    var element = document.querySelector(cssSelector);
    var rect = element.getBoundingClientRect();

    // 获取元素在文档中的位置信息
    var elementPosition = {
        x: rect.left + window.scrollX,
        y: rect.top + window.scrollY
    };

    // 获取视窗相对于整个文档的上下位置信息
    var viewportPosition = {
        top: window.scrollY,
        bottom: window.scrollY + window.innerHeight
    };

    return JSON.stringify({
        elementPosition: elementPosition,
        viewportPosition: viewportPosition
    });
}


getElementPositionAndWindowViewport("` + selector + `");
`
}

func JSGetRandomCoordinates() string {
	return `
	function getRandomCoordinates() {
		var windowWidth = window.innerWidth;
		var windowHeight = window.innerHeight;
		
		var randomX = Math.floor(Math.random() * windowWidth);
		var randomY = Math.floor(Math.random() * windowHeight);
    
		//return { x: randomX, y: randomY };
	   return JSON.stringify({
			x: randomX,
			y: randomY
		});
	}

	getRandomCoordinates();

`
}
