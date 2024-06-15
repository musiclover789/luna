package script

/***

function moveMouse(startX, startY, endX, endY, duration) {
  const path = document.createElementNS("http://www.w3.org/2000/svg", "path");
  const d = `M${startX},${startY} Q${startX},${endY} ${endX},${endY}`;
  path.setAttribute("d", d);
  const pathLength = path.getTotalLength();
  const speed = pathLength / duration;
  let startTime = null;
  let progress = 0;

  document.body.style.cursor = "none"; // 隐藏原有鼠标样式
  const mouse = document.createElement("div");
  mouse.style.width = "20px";
  mouse.style.height = "20px";
  mouse.style.zIndex = 9999;
  mouse.style.borderRadius = "50%";
  mouse.style.backgroundColor = "red";
  mouse.style.position = "absolute";
  document.body.appendChild(mouse);

  function animate(currentTime) {
    if (startTime === null) startTime = currentTime;
    const elapsedTime = currentTime - startTime;
    progress = speed * elapsedTime;
    const point = path.getPointAtLength(progress);

    mouse.style.left = `${point.x}px`;
    mouse.style.top = `${point.y}px`;

    if (progress < pathLength) {
      requestAnimationFrame(animate);
    } else {
      document.body.removeChild(mouse);
      document.body.style.cursor = "default"; // 恢复原有鼠标样式
    }
  }

  requestAnimationFrame(animate);
}

moveMouse(100, 100, 800, 401, 1500);
---下面这个作为备份吧
function moveMouse(startX, startY, endX, endY, targetSize, targetDistance) {
  const path = document.createElementNS("http://www.w3.org/2000/svg", "path");
  const d = `M${startX},${startY} Q${startX},${endY} ${endX},${endY}`;
  path.setAttribute("d", d);
  const pathLength = path.getTotalLength();
  const distance = Math.sqrt(Math.pow(endX - startX, 2) + Math.pow(endY - startY, 2));
  const a = 0.5 * targetSize;
  const b = targetDistance;
  const movementTime = a + b * Math.log2(1 + distance / a);
  const speed = pathLength / movementTime;
  let startTime = null;
  let progress = 0;

  document.body.style.cursor = "none"; // 隐藏原有鼠标样式
  const mouse = document.createElement("div");
  mouse.style.width = "20px";
  mouse.style.height = "20px";
  mouse.style.zIndex = 9999;
  mouse.style.borderRadius = "50%";
  mouse.style.backgroundColor = "red";
  mouse.style.position = "absolute";
  document.body.appendChild(mouse);

  function animate(currentTime) {
    if (startTime === null) startTime = currentTime;
    const elapsedTime = currentTime - startTime;
    progress = speed * elapsedTime;
    const point = path.getPointAtLength(progress);

    mouse.style.left = `${point.x}px`;
    mouse.style.top = `${point.y}px`;

    if (progress < pathLength) {
      requestAnimationFrame(animate);
    } else {
      document.body.removeChild(mouse);
      document.body.style.cursor = "default"; // 恢复原有鼠标样式
    }
  }

  requestAnimationFrame(animate);
}

moveMouse(100, 100, 800, 100, 50, 3000);
---加速减速的也作为备份吧
function moveMouse(startX, startY, endX, endY, targetSize, targetDistance) {
  const path = document.createElementNS("http://www.w3.org/2000/svg", "path");
  const d = `M${startX},${startY} Q${startX},${endY} ${endX},${endY}`;
  path.setAttribute("d", d);
  const pathLength = path.getTotalLength();
  const distance = Math.sqrt(Math.pow(endX - startX, 2) + Math.pow(endY - startY, 2));
  const a = 0.5 * targetSize;
  const b = targetDistance;
  const movementTime = a + b * Math.log2(1 + distance / a);
  let startTime = null;
  let progress = 0;
  let speed = 0;
  let acceleration = 0.2;
  let deceleration = 0.1;
  let velocity = 0;
  let maxSpeed = 200;
  let direction = 1;
  let intervalId = null;

  document.body.style.cursor = "none"; // 隐藏原有鼠标样式
  const mouse = document.createElement("div");
  mouse.style.width = "20px";
  mouse.style.height = "20px";
  mouse.style.zIndex = 9999;
  mouse.style.borderRadius = "50%";
  mouse.style.backgroundColor = "red";
  mouse.style.position = "absolute";
  document.body.appendChild(mouse);

  function updateSpeed(currentProgress, timeElapsed) {
    if (currentProgress < 0.5 * pathLength) {
      // 加速
      velocity += acceleration * timeElapsed;
      speed = Math.min(maxSpeed, velocity);
    } else if (currentProgress > 0.5 * pathLength && currentProgress < pathLength) {
      // 减速
      velocity -= deceleration * timeElapsed;
      speed = Math.max(0, velocity);
    } else {
      clearInterval(intervalId);
      document.body.removeChild(mouse);
      document.body.style.cursor = "default"; // 恢复原有鼠标样式
    }
  }

  function animate() {
    if (startTime === null) startTime = performance.now();
    const currentTime = performance.now();
    const timeElapsed = (currentTime - startTime) / 1000;
    progress += speed * timeElapsed * direction;
    const point = path.getPointAtLength(progress);

    mouse.style.left = `${point.x}px`;
    mouse.style.top = `${point.y}px`;

    updateSpeed(progress, timeElapsed);

    if (progress < 0 || progress > pathLength) {
      direction = -direction;
      velocity = 0;
    }

    intervalId = setTimeout(animate, 10);
  }

  intervalId = setTimeout(animate, 10);
}

moveMouse(100, 100, 800, 100, 50, 3000);
---多阶的
function createPath(points) {
  let d = `M${points[0].x},${points[0].y}`;
  for (let i = 1; i < points.length - 2; i++) {
    const x1 = points[i].x;
    const y1 = points[i].y;
    const x2 = points[i+1].x;
    const y2 = points[i+1].y;
    const xc = (x1 + x2) / 2;
    const yc = (y1 + y2) / 2;
    d += ` Q ${x1},${y1} ${xc},${yc}`;
  }
  const x1 = points[points.length - 2].x;
  const y1 = points[points.length - 2].y;
  const x2 = points[points.length - 1].x;
  const y2 = points[points.length - 1].y;
  d += ` Q ${x1},${y1} ${x2},${y2}`;
  return d;
}

function moveMouse(points, targetSize, targetDistance) {
  const path = document.createElementNS("http://www.w3.org/2000/svg", "path");
  const d = createPath(points);
  path.setAttribute("d", d);
  const pathLength = path.getTotalLength();
  const distance = calculateDistance(points);
  const a = 0.5 * targetSize;
  const b = targetDistance;
  const movementTime = a + b * Math.log2(1 + distance / a);
  const speed = pathLength / movementTime;
  let startTime = null;
  let progress = 0;

  document.body.style.cursor = "none"; // 隐藏原有鼠标样式
  const mouse = document.createElement("div");
  mouse.style.width = "20px";
  mouse.style.height = "20px";
  mouse.style.zIndex = 9999;
  mouse.style.borderRadius = "50%";
  mouse.style.backgroundColor = "red";
  mouse.style.position = "absolute";
  document.body.appendChild(mouse);

  function animate(currentTime) {
    if (startTime === null) startTime = currentTime;
    const elapsedTime = currentTime - startTime;
    progress = speed * elapsedTime;
    const point = path.getPointAtLength(progress);

    mouse.style.left = `${point.x}px`;
    mouse.style.top = `${point.y}px`;

    if (progress < pathLength) {
      requestAnimationFrame(animate);
    } else {
      document.body.removeChild(mouse);
      document.body.style.cursor = "default"; // 恢复原有鼠标样式
    }
  }

  requestAnimationFrame(animate);
}

function calculateDistance(points) {
  let distance = 0;
  for (let i = 0; i < points.length - 1; i++) {
    const dx = points[i + 1].x - points[i].x;
    const dy = points[i + 1].y - points[i].y;
    distance += Math.sqrt(dx * dx + dy * dy);
  }
  return distance;
}

const points = [
  {x: 100, y: 100},
  {x: 250, y: 200},
  {x: 550, y: 50},
  {x: 800, y: 100},
];

moveMouse(points, 50, 3000);

 */

/***
**非常完美的代码
<html>
<head>
  <title>Mouse Movement Tracker</title>
  <script>
    let startBtn = null;
    let stopBtn = null;
    let mouseMoveData = [];
    let startTime = 0;
    let startPosition = null;

    function getRandomPosition() {
      const maxX = window.innerWidth - 50;
      const maxY = window.innerHeight - 50;
      const randomX = Math.floor(Math.random() * maxX);
      const randomY = Math.floor(Math.random() * maxY);
      return {x: randomX, y: randomY};
    }

    function startTracking() {
      mouseMoveData = [];
      startTime = Date.now();

      document.addEventListener('mousemove', handleMouseMove);
    }

    function stopTracking() {
      startPosition = getRandomPosition();
      startBtn.style.left = startPosition.x + 'px';
      startBtn.style.top = startPosition.y + 'px';
      //--
      document.removeEventListener('mousemove', handleMouseMove);
      const endTime = Date.now();
      const elapsedTime = endTime - startTime;
      const result = [];
      for (const data of mouseMoveData) {
        const delay = data.time - startTime;
        const x = data.x;
        const y = data.y;
        result.push({x, y, delay});
      }
      console.log(result);
      const outputElem = document.getElementById('output');
      outputElem.innerText = JSON.stringify(result, null, 2);

      const divs = document.querySelectorAll('div');
      divs.forEach(div => {
        div.parentNode.removeChild(div);
      });
    }

    function handleMouseMove(event) {
      const time = Date.now();
      const x = event.pageX;
      const y = event.pageY;
      mouseMoveData.push({x, y, time});

      // 绘制鼠标移动的轨迹
      const div = document.createElement('div');
      div.style.position = 'absolute';
      div.style.width = '2px';
      div.style.height = '2px';
      div.style.background = '#000';
      div.style.left = x + 'px';
      div.style.top = y + 'px';
      document.body.insertBefore(div, document.body.firstChild);

     // document.body.appendChild(div);
    }



    window.onload = function() {
      startBtn = document.getElementById('start');
      stopBtn = document.getElementById('stop');
      stopBtn.disabled = true;

      startBtn.onclick = function() {
        startTracking();
        startBtn.disabled = true;
        stopBtn.disabled = false;
      };

      stopBtn.onclick = function() {
        stopTracking();
        startBtn.disabled = false;
        stopBtn.disabled = true;
      };
    }
  </script>
  <style>
    button {
      position: absolute;
      width: 50px;
      height: 50px;
      background-color: #007aff;
      color: #fff;
      border-radius: 50%;
      border: none;
      outline: none;
      cursor: pointer;
      font-size: 18px;
    }
    #start {
      left: 0;
      top: 0;
    }
    #stop {
      right: 0;
      top: 0;
    }
    #output {
      margin-top: 50px;
      white-space: pre-wrap;
      font-family: monospace;
    }
  </style>
</head>
<body>
  <button id="start">Start</button>
  <button id="stop">Stop</button>
  <pre id="output"></pre>
</body>
</html>

------随时可以看到鼠标的移动的坐标;


const coordinates = document.createElement('div');
coordinates.style.position = 'fixed';
coordinates.style.top = '10px';
coordinates.style.right = '10px';
coordinates.style.background = '#fff';
coordinates.style.border = '1px solid #000';
coordinates.style.padding = '5px';
coordinates.style.userSelect = 'none'; // 禁止选中文字

// 监听鼠标移动事件，更新坐标值
document.addEventListener('mousemove', e => {
  const x = e.clientX;
  const y = e.clientY;
  coordinates.innerHTML = `X: ${x}, Y: ${y}`;
});

// 监听鼠标按下事件，记录初始位置
let isDragging = false;
let initialX;
let initialY;
let currentX;
let currentY;
let xOffset = 0;
let yOffset = 0;

coordinates.addEventListener("mousedown", dragStart);
coordinates.addEventListener("mouseup", dragEnd);
coordinates.addEventListener("mousemove", drag);

function dragStart(e) {
  initialX = e.clientX - xOffset;
  initialY = e.clientY - yOffset;
  if (e.target === coordinates) {
    isDragging = true;
  }
}

function dragEnd(e) {
  initialX = currentX;
  initialY = currentY;
  isDragging = false;
}

function drag(e) {
  if (isDragging) {
    e.preventDefault();

    currentX = e.clientX - initialX;
    currentY = e.clientY - initialY;

    xOffset = currentX;
    yOffset = currentY;

    setTranslate(currentX, currentY, coordinates);
  }
}

function setTranslate(xPos, yPos, el) {
  el.style.transform = "translate3d(" + xPos + "px, " + yPos + "px, 0)";
}

document.body.appendChild(coordinates);

**/


var ShowMousePosition= func() string{
	return `
(function(){ 

function luna_createMouseTrail() {
	  const trail = document.createElement('div');
	  trail.style.position = 'fixed';
	  trail.style.top = '0';
	  trail.style.left = '0';
	  trail.style.pointerEvents = 'none';
	  trail.style.zIndex = '9999';
	  trail.style.width = '10px';
	  trail.style.height = '10px';
	  trail.style.borderRadius = '50%';
	  trail.style.backgroundColor = 'black';
	  trail.style.opacity = '0.5';
	  trail.style.transition = 'transform 0.1s ease-out';
	
	  document.body.appendChild(trail);
	
	  const coordinates = document.createElement('div');
	  coordinates.style.position = 'fixed';
	  coordinates.style.top = '10px';
	  coordinates.style.right = '10px';
	  coordinates.style.background = '#fff';
	  coordinates.style.border = '1px solid #000';
	  coordinates.style.padding = '5px';
	  coordinates.style.userSelect = 'none'; // 禁止选中文字
	  coordinates.style.zIndex = '9999';
	
	  document.body.appendChild(coordinates);
	
	  document.addEventListener('mousemove', (event) => {
		const { clientX, clientY } = event;
		const x = clientX - 5;
		const y = clientY - 5;
		trail.style.transform = 'translate(' + x + 'px, ' + y + 'px)';
		const initialX = clientX;
		const initialY = clientY;
		const pageWidth = window.innerWidth;
		const offsetX = (clientX > window.innerWidth / 2) ? 20 : -80; // 偏移量
		coordinates.style.right = (window.innerWidth - initialX + offsetX) + 'px';
		coordinates.style.top = (initialY + 10) + 'px';
		coordinates.innerHTML = 'X: ' + clientX + ', Y: ' + clientY + '';
	  });
};
//
luna_createMouseTrail();

})();
`
}