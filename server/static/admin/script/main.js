var $$ = mdui.JQ;
var inst = new mdui.Drawer('#main-drawer');

// 图片占位符
Holder.addTheme('gray', {
  bg: '#BCBEC0',
  fg: 'rgba(255, 255, 255, 1)',
  size: 12,
  fontweight: 'normal',
});

// 自动选中menu
let links = Array.from(document.querySelectorAll('a'));
let href = location.href.replace(/\/$/g, '');
for (let i = 0; i < links.length; i++) {
  const element = links[i];
  if (href == element.href.replace(/\/$/g, '')) {
    if ($$(element).hasClass('mdui-list-item')) {
      $$(element).toggleClass('mdui-list-item-active');
      let parent =
        (element.parentElement && element.parentElement.parentElement) || null;
      if (parent) {
        $$(parent).toggleClass('mdui-collapse-item-open');
      }
    }
  }
}
