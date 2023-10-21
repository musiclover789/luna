package script

/***
计算屏幕的缩放因子
*/
func GetScalingFactor() string {
	return `
	function getScalingFactor() {
	  var devicePixelRatio = window.devicePixelRatio || 1;
	  var cssPixelRatio = window.innerWidth / document.documentElement.clientWidth;
	  var scalingFactor = devicePixelRatio / cssPixelRatio;
	  return scalingFactor;
	}
 	getScalingFactor();
`
}

func ScreenInfo() string {
	return `
	function getScreenInfo() {
	  var screenInfo = {
		width: window.screen.width,
		height: window.screen.height,
		availWidth: window.screen.availWidth,
		availHeight: window.screen.availHeight,
		devicePixelRatio: window.devicePixelRatio || 1,
	  };
	  return screenInfo;
	}
    getScreenInfo();
`
}
