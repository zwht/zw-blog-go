function rep(str) {
  str = "" + str;
  return str.replace(/&lt;/g, "<").replace(/&gt;/g, ">").replace(/&amp;/g, "&").replace(/&quot;/g, '"').replace(/&apos;/g, "'");
}
function baseNode(node) {
  debugger
  var fuck = document.querySelectorAll(node);
  debugger
  for (var i = 0; i < fuck.length; i++) {
    debugger
    fuck[i].innerHTML = rep(fuck[i].innerHTML)
  }
}
