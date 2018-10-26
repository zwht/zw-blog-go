// 解码html字符串
function rep(str) {
  str = "" + str;
  return str.replace(/&lt;/g, "<").replace(/&gt;/g, ">").replace(/&amp;/g, "&").replace(/&quot;/g, '"').replace(/&apos;/g, "'");
}
// 根据node解码html字符串
function baseNode(node) {
  var fuck = document.querySelectorAll(node);
  for (var i = 0; i < fuck.length; i++) {
    fuck[i].innerHTML = rep(fuck[i].innerHTML)
  }
}
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
