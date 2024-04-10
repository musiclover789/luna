package script

/**
***
 */
var TestWebDriver = func() string {
	return `
	var luna_result="";
	function luna_test_js(){
    	console.log("appCodeName: " + navigator.appCodeName);
		console.log("appName: " + navigator.appName);
		console.log("appVersion: " + navigator.appVersion);
		console.log("cookieEnabled: " + navigator.cookieEnabled);
		console.log("language: " + navigator.language);
		
		const locale = new Intl.DateTimeFormat().resolvedOptions().locale;
		console.log("Intl Locale: " + locale);
		
		console.log("platform: " + navigator.platform);
		console.log("userAgent: " + navigator.userAgent);
		
		console.log("availHeight: " + screen.availHeight);
		console.log("availWidth: " + screen.availWidth);
		console.log("availLeft: " + screen.availLeft);
		console.log("availTop: " + screen.availTop);
		console.log("height: " + screen.height);
		console.log("width: " + screen.width);
		console.log("colorDepth: " + screen.colorDepth);
		console.log("pixelDepth: " + screen.pixelDepth);
		
		console.log("Device Memory: " + navigator.deviceMemory);
		console.log("Hardware Concurrency: " + navigator.hardwareConcurrency);
		
		const canvas = document.createElement("canvas");
		const gl = canvas.getContext("webgl");
		const glExtensions = gl.getExtension("WEBGL_debug_renderer_info");
		const glVendor = gl.getParameter(glExtensions.UNMASKED_VENDOR_WEBGL);
		const glRenderer = gl.getParameter(glExtensions.UNMASKED_RENDERER_WEBGL);
		const glVersion = gl.getParameter(gl.VERSION);
        const shadingLanguageVersion = gl.getParameter(gl.SHADING_LANGUAGE_VERSION);
		const extensions = gl.getSupportedExtensions();
		console.log("UNMASKED_VENDOR_WEBGL: " + glVendor);
		console.log("UNMASKED_RENDERER_WEBGL: " + glRenderer);
		console.log("GL_VERSION: " + glVersion);
		console.log("GL_SHADING_LANGUAGE_VERSION:", shadingLanguageVersion);
		console.log("extensions"+extensions);
		const date = new Date();
		const timeZone = date.getTimezoneOffset();
		const timeZoneOffset = -timeZone / 60;
		console.log("Time Zone: " + Intl.DateTimeFormat().resolvedOptions().timeZone);
		console.log("Time Zone Offset: " + timeZoneOffset);
		
		if (typeof webdriver !== 'undefined') {
		  console.log("在 WebDriver 环境下执行的代码");
		} else {
		  console.log("不在WebDriver环境下执行的代码");
		}
		if (typeof window.domAutomation !== 'undefined') {
		  console.log("domAutomation 的环境下执行的代码");
		} else {
		  console.log("不在domAutomation环境下执行的代码");
		}
		
		const isPuppeteer = typeof puppeteer !== 'undefined';
		if (isPuppeteer) {
		  console.log("这是在 Puppeteer 环境下执行的代码");
		} else {
		  console.log("这不是在 Puppeteer 环境下执行的代码");
		}
		
		const isPlaywright = typeof playwright !== 'undefined';
		if (isPlaywright) {
		  console.log("这是在 Playwright 环境下执行的代码");
		} else {
		  console.log("这不是在 Playwright 环境下执行的代码");
		}
	}
	luna_test_js();
`
}
