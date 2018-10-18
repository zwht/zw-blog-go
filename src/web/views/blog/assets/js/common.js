function rep(str) {
  str = "" + str;
  return str.replace(/&lt;/g, "<").replace(/&gt;/g, ">").replace(/&amp;/g, "&").replace(/&quot;/g, '"').replace(/&apos;/g, "'");
}
function baseNode(node) {
  var fuck = document.querySelectorAll(node);
  for (var i = 0; i < fuck.length; i++) {
    fuck[i].innerHTML = rep(fuck[i].innerHTML)
  }
}
