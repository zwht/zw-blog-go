
// 添加window.onload方法
function addLoadEvent(func) {
  var oldonload = window.onload;
  if (typeof window.onload != 'function') {
    window.onload = func;
  } else {
    window.onload = function () {
      oldonload();
      func();
    }
  }
}

// 统计rem单位计算
addLoadEvent(function () {
  document.documentElement.style.fontSize = document.documentElement.clientWidth > 1200 ? '100px' : document.documentElement.clientWidth / 1200 * 100 + 'px';
})
