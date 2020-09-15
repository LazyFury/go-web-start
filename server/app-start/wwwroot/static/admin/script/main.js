// var inst = new mdui.Drawer('#main-drawer');
mdui.mutation();

// 图片占位符
Holder.addTheme('gray', {
  bg: '#BCBEC0',
  fg: 'rgba(255, 255, 255, 1)',
  size: 12,
  fontweight: 'normal',
});

// 自动选中menu
let links = Array.from(document.querySelectorAll('a'));
let href = (location.origin + location.pathname).replace(/\/$/g, '');
for (let i = 0; i < links.length; i++) {
  const element = links[i];
  const eHref = element.href.split('?')[0].replace(/\/$/g, '');
  const isMatch = href == eHref;
  // console.log({ eHref, reg, isMatch, href });
  if (isMatch) {
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

function params2JSON(params) {
  params = decodeURIComponent(params);
  let arr = params.split('&');
  let obj = {};
  for (let i = 0; i < arr.length; i++) {
    let ele = arr[i].split('=');
    obj[ele[0]] = ele[1];
  }
  return obj;
}
